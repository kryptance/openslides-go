// Package http implements the HTTP server and routing for the backend service.
package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/presenter"
	"github.com/OpenSlides/openslides-go/environment"
)

var envBackendPort = environment.NewVariable("BACKEND_PORT", "9002", "Port on which the service listens.")

// Server is the HTTP server for the backend service.
type Server struct {
	Addr string
	lst  net.Listener
}

// New initializes a new Server.
func New(lookup environment.Environmenter) Server {
	return Server{
		Addr: ":" + envBackendPort.Value(lookup),
	}
}

// Run starts the HTTP service.
func (s *Server) Run(ctx context.Context, auth authenticater) error {
	mux := registerHandlers(auth)

	srv := &http.Server{
		Handler:     mux,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	wait := make(chan error)
	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			wait <- fmt.Errorf("HTTP server shutdown: %w", err)
			return
		}
		wait <- nil
	}()

	if s.lst == nil {
		var err error
		s.lst, err = net.Listen("tcp", s.Addr)
		if err != nil {
			return fmt.Errorf("open %s: %w", s.Addr, err)
		}
	}

	s.Addr = s.lst.Addr().String()
	log.Printf("Listen on %s\n", s.Addr)
	if err := srv.Serve(s.lst); err != http.ErrServerClosed {
		return fmt.Errorf("HTTP Server failed: %v", err)
	}

	return <-wait
}

type authenticater interface {
	Authenticate(http.ResponseWriter, *http.Request) (context.Context, error)
	FromContext(context.Context) int
}

func registerHandlers(auth authenticater) *http.ServeMux {
	mux := http.NewServeMux()

	// External endpoints (require authentication).
	mux.Handle("POST /system/action/handle_request", handleExternal(handleActionRequest(auth)))
	mux.Handle("POST /system/action/handle_separately", handleExternal(handleActionSeparately(auth)))
	mux.Handle("POST /system/presenter/handle_request", handleExternal(handlePresenterRequest(auth)))

	// Internal endpoints (no auth required).
	mux.Handle("POST /internal/action/handle_request", handleInternal(handleInternalAction()))

	// Health check.
	mux.Handle("GET /system/action/health", handleExternal(handleHealth()))

	return mux
}

func handleActionRequest(auth authenticater) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx, err := auth.Authenticate(w, r)
		if err != nil {
			return err
		}

		userID := auth.FromContext(ctx)

		var payload []action.ActionPayloadElement
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			return fmt.Errorf("decoding action payload: %w", err)
		}

		responses, err := action.HandleRequest(ctx, payload, userID, false)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		results := make([][]map[string]any, len(responses))
		for i, resp := range responses {
			results[i] = resp.Results
		}
		return json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"results": results,
		})
	}
}

func handleActionSeparately(auth authenticater) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx, err := auth.Authenticate(w, r)
		if err != nil {
			return err
		}

		userID := auth.FromContext(ctx)

		var payload []action.ActionPayloadElement
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			return fmt.Errorf("decoding action payload: %w", err)
		}

		responses, err := action.HandleSeparately(ctx, payload, userID)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		results := make([][]map[string]any, len(responses))
		for i, resp := range responses {
			results[i] = resp.Results
		}
		return json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"results": results,
		})
	}
}

func handlePresenterRequest(auth authenticater) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx, err := auth.Authenticate(w, r)
		if err != nil {
			return err
		}

		userID := auth.FromContext(ctx)

		var payload []json.RawMessage
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			return fmt.Errorf("decoding presenter payload: %w", err)
		}

		results, err := presenter.Handle(ctx, userID, payload)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(results)
	}
}

func handleInternalAction() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var payload []action.ActionPayloadElement
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			return fmt.Errorf("decoding action payload: %w", err)
		}

		responses, err := action.HandleRequest(r.Context(), payload, -1, true)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		results := make([][]map[string]any, len(responses))
		for i, resp := range responses {
			results[i] = resp.Results
		}
		return json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"results": results,
		})
	}
}

func handleHealth() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"healthy":true}`)
		return nil
	}
}

// Handler is like http.Handler but returns an error.
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request) error
}

// HandlerFunc is like http.HandlerFunc but returns an error.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return f(w, r)
}

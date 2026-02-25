package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

func handleInternal(handler Handler) http.Handler {
	return resolveError(handler, true)
}

func handleExternal(handler Handler) http.Handler {
	return resolveError(handler, false)
}

func resolveError(handler Handler, internalRoute bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler.ServeHTTP(w, r)
		if err == nil || errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return
		}

		writeStatusCode(w, err)
		writeFormattedError(w, err, internalRoute)
	}
}

func writeStatusCode(w http.ResponseWriter, err error) {
	code := 400

	var statusCoder backenderr.StatusCoder
	if errors.As(err, &statusCoder) {
		code = statusCoder.StatusCode()
	}

	var typed backenderr.Typed
	if !errors.As(err, &typed) {
		code = 500
	}

	w.WriteHeader(code)
}

func writeFormattedError(w io.Writer, err error, internalRoute bool) {
	errType := "internal"
	var typed backenderr.Typed
	if errors.As(err, &typed) {
		errType = typed.Type()
	}

	msg := err.Error()
	if errType == "internal" {
		log.Printf("Internal error: %s", msg)
		if !internalRoute {
			msg = "Internal server error"
		}
	}

	out := struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{
		errType,
		msg,
	}

	if encErr := json.NewEncoder(w).Encode(out); encErr != nil {
		log.Printf("Error encoding error message: %v", encErr)
		fmt.Fprint(w, `{"error":"internal","message":"Something went wrong encoding the error message"}`)
	}
}

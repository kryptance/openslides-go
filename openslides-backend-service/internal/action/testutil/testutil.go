// Package testutil provides test utilities for action testing,
// equivalent to Python's BaseActionTestCase.
package testutil

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/datastore"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"
)

// ActionResponse wraps the result of an action request.
type ActionResponse struct {
	Results []map[string]any
	Events  []event.Event
	Err     error
}

// ActionTestCase provides utilities for testing actions.
type ActionTestCase struct {
	T         *testing.T
	Datastore *datastore.Extended
	UserID    int
}

// New creates a new ActionTestCase with a fresh in-memory datastore.
func New(t *testing.T) *ActionTestCase {
	t.Helper()
	return &ActionTestCase{
		T:         t,
		Datastore: datastore.NewExtended(),
		UserID:    1,
	}
}

// SetModels batch-sets models in the test datastore.
// models is a map of FQID string ("collection/id") to field data.
func (tc *ActionTestCase) SetModels(models map[string]map[string]any) {
	tc.T.Helper()
	for fqid, fields := range models {
		collection, id := parseFQID(fqid)
		if id == 0 {
			tc.T.Fatalf("invalid FQID: %s", fqid)
		}
		// Set the id field.
		fields["id"] = id
		tc.Datastore.ApplyChangedModel(collection, id, fields)
		// Update the next ID counter so ReserveIDs won't conflict with pre-set models.
		tc.Datastore.SetNextID(collection, id)
	}
}

// Request executes an action and returns the response.
func (tc *ActionTestCase) Request(actionName string, data map[string]any) (*ActionResponse, error) {
	tc.T.Helper()
	return tc.performAction(actionName, data, false)
}

// RequestMulti executes an action with multiple data instances.
func (tc *ActionTestCase) RequestMulti(actionName string, dataList []map[string]any) (*ActionResponse, error) {
	tc.T.Helper()

	a, err := action.Lookup(actionName)
	if err != nil {
		return nil, err
	}

	rm := relation.NewManager()
	params := &action.ActionParams{
		UserID:          tc.UserID,
		Internal:        false,
		Datastore:       tc.Datastore,
		RelationManager: rm,
	}

	instances := make([]action.Instance, len(dataList))
	for i, d := range dataList {
		instances[i] = action.Instance(d)
	}

	events, results, err := action.Perform(context.Background(), a, params, instances)
	if err != nil {
		return &ActionResponse{Err: err}, err
	}

	return &ActionResponse{Results: results, Events: events}, nil
}

// RequestInternal executes an action as internal (no permission check).
func (tc *ActionTestCase) RequestInternal(actionName string, data map[string]any) (*ActionResponse, error) {
	tc.T.Helper()
	return tc.performAction(actionName, data, true)
}

func (tc *ActionTestCase) performAction(actionName string, data map[string]any, internal bool) (*ActionResponse, error) {
	tc.T.Helper()

	a, err := action.Lookup(actionName)
	if err != nil {
		return nil, err
	}

	rm := relation.NewManager()
	params := &action.ActionParams{
		UserID:          tc.UserID,
		Internal:        internal,
		Datastore:       tc.Datastore,
		RelationManager: rm,
	}

	instances := []action.Instance{action.Instance(data)}
	events, results, err := action.Perform(context.Background(), a, params, instances)
	if err != nil {
		return &ActionResponse{Err: err}, err
	}

	return &ActionResponse{Results: results, Events: events}, nil
}

// AssertSuccess asserts the response has no error.
func (tc *ActionTestCase) AssertSuccess(resp *ActionResponse, err error) {
	tc.T.Helper()
	if err != nil {
		tc.T.Fatalf("expected success, got error: %v", err)
	}
}

// AssertError asserts the response has an error.
func (tc *ActionTestCase) AssertError(resp *ActionResponse, err error) {
	tc.T.Helper()
	if err == nil {
		tc.T.Fatal("expected error, got success")
	}
}

// AssertErrorContains asserts the error message contains the given substring.
func (tc *ActionTestCase) AssertErrorContains(err error, substring string) {
	tc.T.Helper()
	if err == nil {
		tc.T.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), substring) {
		tc.T.Fatalf("expected error containing %q, got: %v", substring, err)
	}
}

// AssertModelExists checks that a model exists in the datastore with the expected fields.
func (tc *ActionTestCase) AssertModelExists(fqid string, expectedFields ...map[string]any) {
	tc.T.Helper()
	collection, id := parseFQID(fqid)
	model, found := tc.Datastore.GetChangedModel(collection, id)
	if !found {
		tc.T.Fatalf("model %s does not exist", fqid)
	}

	if len(expectedFields) > 0 {
		for key, expected := range expectedFields[0] {
			actual, ok := model[key]
			if !ok {
				tc.T.Errorf("model %s missing field %q", fqid, key)
				continue
			}
			if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected) {
				tc.T.Errorf("model %s field %q: expected %v, got %v", fqid, key, expected, actual)
			}
		}
	}
}

// AssertModelDeleted checks that a model is marked as deleted.
func (tc *ActionTestCase) AssertModelDeleted(fqid string) {
	tc.T.Helper()
	collection, id := parseFQID(fqid)
	if !tc.Datastore.IsDeleted(collection, id) {
		tc.T.Fatalf("model %s should be deleted but is not", fqid)
	}
}

// AssertModelNotExists checks that a model doesn't exist.
func (tc *ActionTestCase) AssertModelNotExists(fqid string) {
	tc.T.Helper()
	collection, id := parseFQID(fqid)
	_, found := tc.Datastore.GetChangedModel(collection, id)
	if found {
		tc.T.Fatalf("model %s should not exist but does", fqid)
	}
}

// AssertEventCount checks the number of events returned by the action.
func (tc *ActionTestCase) AssertEventCount(resp *ActionResponse, expected int) {
	tc.T.Helper()
	if resp == nil {
		tc.T.Fatal("response is nil")
	}
	if len(resp.Events) != expected {
		tc.T.Fatalf("expected %d events, got %d", expected, len(resp.Events))
	}
}

// AssertHasEvent checks that an event of the given type and FQID exists.
func (tc *ActionTestCase) AssertHasEvent(resp *ActionResponse, eventType event.Type, fqid string) {
	tc.T.Helper()
	if resp == nil {
		tc.T.Fatal("response is nil")
	}
	for _, e := range resp.Events {
		if e.Type == eventType && e.FQID == fqid {
			return
		}
	}
	tc.T.Fatalf("expected event type=%s fqid=%s not found in %d events", eventType, fqid, len(resp.Events))
}

// GetModel returns the model data from the test datastore.
func (tc *ActionTestCase) GetModel(fqid string) map[string]any {
	tc.T.Helper()
	collection, id := parseFQID(fqid)
	model, found := tc.Datastore.GetChangedModel(collection, id)
	if !found {
		tc.T.Fatalf("model %s does not exist", fqid)
	}
	return model
}

// CreateMeeting creates a basic meeting with standard groups for testing.
func (tc *ActionTestCase) CreateMeeting() {
	tc.T.Helper()
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name":               "Test Organization",
			"active_meeting_ids": []any{1},
			"committee_ids":      []any{1},
		},
		"committee/1": {
			"name":            "Test Committee",
			"organization_id": 1,
			"meeting_ids":     []any{1},
		},
		"meeting/1": {
			"name":                         "Test Meeting",
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
			"group_ids":                    []any{1, 2, 3},
			"default_group_id":             1,
			"admin_group_id":               2,
		},
		"group/1": {
			"name":                         "Default",
			"meeting_id":                   1,
			"default_group_for_meeting_id": 1,
		},
		"group/2": {
			"name":                       "Admin",
			"meeting_id":                 1,
			"admin_group_for_meeting_id": 1,
		},
		"group/3": {
			"name":       "Delegates",
			"meeting_id": 1,
		},
		"user/1": {
			"username":         "admin",
			"is_active":        true,
			"organization_id":  1,
			"meeting_user_ids": []any{1},
		},
		"meeting_user/1": {
			"user_id":    1,
			"meeting_id": 1,
			"group_ids":  []any{2},
		},
	})
}

// parseFQID splits "collection/id" into collection and id.
func parseFQID(fqid string) (string, int) {
	parts := strings.SplitN(fqid, "/", 2)
	if len(parts) != 2 {
		return "", 0
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return parts[0], 0
	}
	return parts[0], id
}

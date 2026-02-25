package mixin_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/datastore"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"
)

func TestWithWeight(t *testing.T) {
	a := action.NewCreateAction("test.weight", model.MustGet("topic"))
	mixin.WithWeight(a, "weight")

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	// Instance without explicit weight.
	instance := action.Instance{
		"meeting_id": 1,
		"title":      "Test Topic",
	}

	_, results, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) == 0 {
		t.Fatal("expected results")
	}

	id := results[0]["id"]
	var idInt int
	switch v := id.(type) {
	case float64:
		idInt = int(v)
	case int:
		idInt = v
	}

	pm, found := ds.GetChangedModel("topic", idInt)
	if !found {
		t.Fatal("expected topic model to exist in datastore")
	}

	// Weight should be auto-assigned to 1 (default behavior).
	if pm["weight"] != 1 {
		t.Errorf("expected weight 1, got %v", pm["weight"])
	}
}

func TestWithWeightExplicit(t *testing.T) {
	a := action.NewCreateAction("test.weight.explicit", model.MustGet("topic"))
	mixin.WithWeight(a, "weight")

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	// Instance with explicit weight.
	instance := action.Instance{
		"meeting_id": 1,
		"title":      "Test Topic",
		"weight":     42,
	}

	_, results, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) == 0 {
		t.Fatal("expected results")
	}

	id := results[0]["id"]
	var idInt int
	switch v := id.(type) {
	case float64:
		idInt = int(v)
	case int:
		idInt = v
	}

	pm, found := ds.GetChangedModel("topic", idInt)
	if !found {
		t.Fatal("expected topic model to exist in datastore")
	}

	// Explicit weight should be preserved.
	if pm["weight"] != 42 {
		t.Errorf("expected weight 42, got %v", pm["weight"])
	}
}

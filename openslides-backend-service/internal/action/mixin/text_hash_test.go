package mixin_test

import (
	"context"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/datastore"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"
)

func TestWithTextHash(t *testing.T) {
	a := action.NewCreateAction("test.texthash", model.MustGet("topic"))
	mixin.WithTextHash(a, "text", "text_hash")

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{
		"meeting_id": 1,
		"title":      "Test",
		"text":       "Hello World",
	}

	_, results, err := action.Perform(context.Background(), a, params, []action.Instance{instance})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) == 0 {
		t.Fatal("expected results")
	}

	// Verify the model in the datastore has text_hash set.
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

	expectedHash := fmt.Sprintf("%x", sha256.Sum256([]byte("Hello World")))
	if pm["text_hash"] != expectedHash {
		t.Errorf("expected text_hash %q, got %v", expectedHash, pm["text_hash"])
	}
}

func TestWithTextHashEmptyString(t *testing.T) {
	a := action.NewCreateAction("test.texthash.empty", model.MustGet("topic"))
	mixin.WithTextHash(a, "text", "text_hash")

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	instance := action.Instance{
		"meeting_id": 1,
		"title":      "Test",
		"text":       "",
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

	// Empty string should result in nil text_hash.
	if pm["text_hash"] != nil {
		t.Errorf("expected text_hash to be nil for empty text, got %v", pm["text_hash"])
	}
}

func TestWithTextHashNoTextField(t *testing.T) {
	a := action.NewCreateAction("test.texthash.nofield", model.MustGet("topic"))
	mixin.WithTextHash(a, "text", "text_hash")

	ds := datastore.NewExtended()
	rm := relation.NewManager()

	params := &action.ActionParams{
		UserID:          1,
		Internal:        true,
		Datastore:       ds,
		RelationManager: rm,
	}

	// Instance without the text field.
	instance := action.Instance{
		"meeting_id": 1,
		"title":      "Test",
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

	// text_hash should not be set if text field is not in the instance.
	if _, ok := pm["text_hash"]; ok {
		t.Error("expected text_hash to not be set when text is not in instance")
	}
}

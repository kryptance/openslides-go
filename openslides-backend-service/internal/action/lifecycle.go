package action

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

// Perform executes the full action lifecycle for a list of data instances.
//
// Lifecycle:
//  1. Prefetch(data) - batch fetch needed data
//  2. Per instance: ValidateSchema → ValidateFields → CheckPermissions
//  3. Per instance: UpdateInstance → HandleRelationUpdates → CreateEvents
//  4. Append relation manager events
//  5. Return all events
func Perform(ctx context.Context, action *BaseAction, params *ActionParams, data []Instance) ([]event.Event, []map[string]any, error) {
	// 1. Prefetch.
	if err := action.Prefetch(ctx, params, data); err != nil {
		return nil, nil, fmt.Errorf("prefetch %s: %w", action.Name, err)
	}

	var allEvents []event.Event
	results := make([]map[string]any, 0, len(data))

	for _, instance := range data {
		// 2. Validate schema (if schema is defined).
		if action.Schema != nil {
			if err := ValidateSchema(action.Schema, instance); err != nil {
				return nil, nil, fmt.Errorf("schema validation %s: %w", action.Name, err)
			}
		}

		// 3. Validate fields.
		if err := action.ValidateFields(ctx, params, instance); err != nil {
			return nil, nil, fmt.Errorf("validate fields %s: %w", action.Name, err)
		}

		// 4. Check permissions (skip for internal calls).
		if !params.Internal {
			if err := action.CheckPermissions(ctx, params, instance); err != nil {
				return nil, nil, fmt.Errorf("check permissions %s: %w", action.Name, err)
			}
		}

		// 5. Update instance (apply defaults, transformations, mixins).
		updated, err := action.UpdateInstance(ctx, params, instance)
		if err != nil {
			return nil, nil, fmt.Errorf("update instance %s: %w", action.Name, err)
		}

		// 6. Handle relation updates via the relation manager.
		if params.RelationManager != nil && action.Model != nil {
			idVal, _ := extractID(updated)
			if idVal > 0 {
				params.RelationManager.HandleRelationUpdates(action.Model, idVal, map[string]any(updated))
			}
		}

		// 7. Create events.
		events, err := action.CreateEvents(ctx, params, updated)
		if err != nil {
			return nil, nil, fmt.Errorf("create events %s: %w", action.Name, err)
		}

		allEvents = append(allEvents, events...)

		// Build result element.
		result := make(map[string]any)
		if id, ok := updated["id"]; ok {
			result["id"] = id
		}
		results = append(results, result)
	}

	// 8. Append relation manager events.
	if params.RelationManager != nil {
		relEvents := params.RelationManager.GetEvents()
		allEvents = append(allEvents, relEvents...)
	}

	return allEvents, results, nil
}

// extractID extracts the integer ID from an instance.
func extractID(instance Instance) (int, bool) {
	id, ok := instance["id"]
	if !ok {
		return 0, false
	}
	switch v := id.(type) {
	case float64:
		return int(v), true
	case int:
		return v, true
	}
	return 0, false
}

package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

// TreeNode represents a node in a tree sort payload.
type TreeNode struct {
	ID       int        `json:"id"`
	Children []TreeNode `json:"children,omitempty"`
}

// WithTreeSort provides tree-sorting functionality for an action.
// The instance should contain a "tree" field with the tree structure.
//
// This replaces Python's TreeSortMixin.
func WithTreeSort(a *action.BaseAction, weightField, parentField string) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		result, err := origUpdate(ctx, params, instance)
		if err != nil {
			return nil, err
		}

		treeData, ok := result["tree"]
		if !ok {
			return result, nil
		}

		// Parse tree nodes and compute weights + parent assignments.
		nodes, ok := treeData.([]any)
		if !ok {
			return nil, fmt.Errorf("tree field must be an array")
		}

		// Flatten tree and assign weights.
		weight := 1
		if err := flattenTree(params, a.Model.Collection, nodes, 0, &weight, weightField, parentField); err != nil {
			return nil, fmt.Errorf("flatten tree: %w", err)
		}

		delete(result, "tree")
		return result, nil
	}
}

func flattenTree(params *action.ActionParams, collection string, nodes []any, parentID int, weight *int, weightField, parentField string) error {
	for _, node := range nodes {
		nodeMap, ok := node.(map[string]any)
		if !ok {
			continue
		}

		var id int
		switch v := nodeMap["id"].(type) {
		case float64:
			id = int(v)
		case int:
			id = v
		default:
			continue
		}

		fields := map[string]any{
			weightField: *weight,
		}
		if parentID > 0 {
			fields[parentField] = parentID
		} else {
			fields[parentField] = nil
		}
		params.Datastore.ApplyChangedModel(collection, id, fields)
		*weight++

		if children, ok := nodeMap["children"].([]any); ok {
			if err := flattenTree(params, collection, children, id, weight, weightField, parentField); err != nil {
				return err
			}
		}
	}
	return nil
}

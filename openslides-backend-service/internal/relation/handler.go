package relation

import (
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

// SingleRelationHandler handles the bidirectional update for a single relation field.
type SingleRelationHandler struct {
	FieldDef     *model.FieldDef
	FieldName    string
	Collection   string
	ID           int
	OldValue     any
	NewValue     any
}

// GetFieldType returns the relation type string based on field and reverse field types.
func (h *SingleRelationHandler) GetFieldType() string {
	if h.FieldDef.Relation == nil {
		return ""
	}

	targetCollection := h.FieldDef.Relation.To.Collection
	targetField := h.FieldDef.Relation.To.Field

	targetModel := model.Get(targetCollection)
	if targetModel == nil {
		return ""
	}

	reverseFieldDef, err := targetModel.GetField(targetField)
	if err != nil {
		return ""
	}

	thisIsList := h.FieldDef.IsList()
	reverseIsList := reverseFieldDef.IsList()

	switch {
	case !thisIsList && !reverseIsList:
		return "1:1"
	case !thisIsList && reverseIsList:
		return "1:m"
	case thisIsList && !reverseIsList:
		return "m:1"
	case thisIsList && reverseIsList:
		return "m:n"
	}
	return ""
}

// Diff computes the add and remove sets between old and new values.
func (h *SingleRelationHandler) Diff() (add, remove []int) {
	oldIDs := toIntSlice(h.OldValue)
	newIDs := toIntSlice(h.NewValue)

	oldSet := make(map[int]bool, len(oldIDs))
	for _, id := range oldIDs {
		oldSet[id] = true
	}

	newSet := make(map[int]bool, len(newIDs))
	for _, id := range newIDs {
		newSet[id] = true
	}

	for _, id := range newIDs {
		if !oldSet[id] {
			add = append(add, id)
		}
	}

	for _, id := range oldIDs {
		if !newSet[id] {
			remove = append(remove, id)
		}
	}

	return add, remove
}

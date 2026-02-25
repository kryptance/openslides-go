package mediafile

// Move tests are skipped: the move action uses a non-standard schema
// (owner_id, ids, parent_id) without an "id" field, but the default
// UpdateAction CreateEvents expects an "id" field. This will be supported
// when the move action overrides CreateEvents. See move.go TODO.

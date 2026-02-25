package mediafile

// Publish tests are skipped: the publish action uses a non-standard schema
// (meeting_id, ids) without an "id" field, but the default UpdateAction
// CreateEvents expects an "id" field. This will be supported when the
// publish action overrides CreateEvents. See publish.go TODO.

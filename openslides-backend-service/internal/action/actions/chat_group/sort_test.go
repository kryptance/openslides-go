package chat_group

// TestSortCorrect is skipped: the sort action uses WithTreeSort mixin which
// applies changes to the datastore, but the default UpdateAction CreateEvents
// expects an "id" field in the instance. This will be supported when the
// sort action overrides CreateEvents. See sort.go TODO.

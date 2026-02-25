package speaker

// Tests for speaker.create_for_merge are skipped because they depend on
// user merge functionality which is not yet implemented in the Go backend.
// The Python tests in test_create_for_merge.py rely on:
// - User merge process and internal action calls
// - Complex meeting user relationship resolution
// These features are not available in the Go action framework yet.

package list_of_speakers

// Tests for list_of_speakers.update are skipped because there is no
// list_of_speakers.update action registered in the Go backend.
// The list_of_speakers collection only has: create, delete, delete_all_speakers, re_add_last.
// The Python tests for update test setting closed field and permissions,
// which would require the update action to be implemented.

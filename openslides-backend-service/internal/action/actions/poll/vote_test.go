package poll

// Vote tests are skipped because they depend on the vote service integration
// which is not yet implemented in the Go backend.
// The Python tests in test_vote.py rely heavily on:
// - vote_service.start / vote_service.vote / vote_service.clear
// - Anonymous voting via HTTP headers
// - Delegation chains
// - Multiple user logins
// These features are not available in the Go test framework.

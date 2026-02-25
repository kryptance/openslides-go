package presenter_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/presenter"
)

func TestRegisteredPresenters(t *testing.T) {
	names := []string{
		"check_database",
		"check_database_all",
		"export_meeting",
		"get_active_users_amount",
		"get_forwarding_committees",
		"get_forwarding_meetings",
		"get_mediafile_context",
		"get_user_editable",
		"get_user_related_models",
		"get_user_scope",
		"get_users",
		"number_of_users",
		"search_for_id_by_external_id",
		"search_users",
	}
	for _, name := range names {
		_, err := presenter.LookupPresenter(name)
		if err != nil {
			t.Errorf("presenter %q not registered: %v", name, err)
		}
	}
}

func TestLookupUnknownPresenter(t *testing.T) {
	_, err := presenter.LookupPresenter("nonexistent_presenter")
	if err == nil {
		t.Fatal("expected error for unknown presenter")
	}
}

func TestPresenterCount(t *testing.T) {
	// Verify that we have at least 14 presenters registered.
	count := 0
	names := []string{
		"check_database",
		"check_database_all",
		"export_meeting",
		"get_active_users_amount",
		"get_forwarding_committees",
		"get_forwarding_meetings",
		"get_mediafile_context",
		"get_user_editable",
		"get_user_related_models",
		"get_user_scope",
		"get_users",
		"number_of_users",
		"search_for_id_by_external_id",
		"search_users",
	}
	for _, name := range names {
		if _, err := presenter.LookupPresenter(name); err == nil {
			count++
		}
	}
	if count != 14 {
		t.Errorf("expected 14 registered presenters, got %d", count)
	}
}

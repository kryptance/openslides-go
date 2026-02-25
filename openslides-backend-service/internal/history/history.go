// Package history computes history information for write requests.
package history

import (
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

// Information represents a history entry for one FQID.
type Information struct {
	FQID string
	Info []string
}

// BuildInformation creates history information strings for a single event.
func BuildInformation(e event.Event) []string {
	switch e.Type {
	case event.TypeCreate:
		return []string{"Object created"}
	case event.TypeUpdate:
		return []string{"Object updated"}
	case event.TypeDelete:
		return []string{"Object deleted"}
	default:
		return nil
	}
}

// BuildAllInformation creates history information entries for a list of events.
func BuildAllInformation(events []event.Event) []Information {
	var infos []Information

	for _, e := range events {
		info := BuildInformation(e)
		if len(info) > 0 {
			infos = append(infos, Information{
				FQID: e.FQID,
				Info: info,
			})
		}
	}

	return infos
}

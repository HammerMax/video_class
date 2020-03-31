package console

import (
	"net/http"
	"server/api/console/station"
	"server/api/console/storage"
)

var Router = map[string]http.HandlerFunc {
	"/video/create": station.Create,
	"/video/update": station.Update,
	"/video/list":   station.List,
	"/video/detail": station.Detail,
	"/video/apply_vid": station.ApplyVid,

	"/storage/upload": storage.Upload,
	"/storage/transcode": storage.Transcode,
}
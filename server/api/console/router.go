package console

import "net/http"

var Router = map[string]http.HandlerFunc {
	"/video/create": Create,
	"/video/update": Update,
	"/video/list": List,
	"/video/detail": Detail,
}
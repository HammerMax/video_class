package main

import (
	"net/http"
	"server/api/console"
	"server/modules/video"
)

func main() {
	video.InitVideo()

	for router, handlerFunc := range console.Router {
		http.HandleFunc(router, handlerFunc)
	}

	http.ListenAndServe(":9999", nil)
}

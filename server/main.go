package main

import (
	"net/http"
	"server/api/console"
	"server/modules/file_metadata"
	"server/modules/transcode"
	"server/modules/video"
)

func main() {
	// Register all formats and codecs
	transcode.Transcode()


	return
	video.InitVideo()
	file_metadata.InitEngine()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(file_metadata.DataDir))))

	for router, handlerFunc := range console.Router {
		http.HandleFunc(router, midWare(handlerFunc))
	}

	http.ListenAndServe(":9999", nil)
}

func midWare(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		if request.Method == http.MethodOptions {
			return
		}
		f(writer, request)
	}
}

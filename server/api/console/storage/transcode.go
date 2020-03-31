package storage

import (
	"log"
	"net/http"
	"server/utils"
)

func Transcode(w http.ResponseWriter, r *http.Request) {
	newFile, fh, err := r.FormFile("file")
	if err != nil {
		log.Printf("api console storage transcode ParseFile err:%v", err)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	key, err := writeFile(newFile, fh)
	if err != nil {
		log.Printf("api console storage transcode writeFile err:%v", err)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, map[string]string{
		"key": key,
	})
}

package storage

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"server/modules/file_metadata"
	"server/utils"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	newFile, fh, err := r.FormFile("file")
	if err != nil {
		log.Printf("api console storage upload ParseFile err:%v", err)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	key, err := writeFile(newFile, fh)
	if err != nil {
		log.Printf("api console storage upload writeFile err:%v", err)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, map[string]string{
		"key": key,
	})
}

func writeFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	var metadata = file_metadata.Metadata{
		Key:        header.Filename,
		Name:       header.Filename,
		Size:       header.Size,
	}
	err := file_metadata.Create(&metadata)
	if err != nil {
		return "", err
	}

	f, err := os.Create(file_metadata.DataDir + "/" + metadata.Key)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(f, file)
	return metadata.Key, err
}

package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJson(w http.ResponseWriter, code int, message string, data interface{}) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response{
		Code:    code,
		Message: message,
		Data:    data,
	})
	if err != nil {
		log.Printf("utils response err:%v", err)
	}
}

func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	ResponseJson(w, 0, "success", data)
}
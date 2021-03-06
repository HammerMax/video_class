package station

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/modules/video"
	"server/utils"
)

// 创建视频
func Create(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("api console create receive body:%s", body)

	var attr video.Attr
	err := json.Unmarshal(body, &attr)
	if err != nil {
		log.Printf("api console create json unmarshal err:%v", err)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	err = video.Create(&attr)
	if err != nil {
		log.Printf("api console create video create err:%v", err)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, nil)
}


package station

import (
	"log"
	"net/http"
	"server/modules/video"
	"server/utils"
)

func Update(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	vid := params.Get("vid")
	if vid == "" {
		utils.ResponseJson(w, -1, "vid required", nil)
		return
	}

	updateFields := make(map[string]string)
	for key := range params {
		updateFields[key] = params.Get(key)
	}

	err := video.Update(vid, updateFields)
	if err != nil {
		log.Printf("api console update err:%v, vid:%s, params:%v", err, vid, updateFields)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}
	utils.ResponseSuccess(w, nil)
}
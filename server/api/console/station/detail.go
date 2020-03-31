package station

import (
	"log"
	"net/http"
	"server/modules/video"
	"server/utils"
)

func Detail(w http.ResponseWriter, r *http.Request) {
	vid := r.URL.Query().Get("vid")
	if vid == "" {
		utils.ResponseJson(w, -1, "vid required", nil)
		return
	}

	videoAttr, err := video.Find(vid)
	if err != nil {
		log.Printf("api console detail err:%v, vid:%s", err, vid)
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, videoAttr)
}

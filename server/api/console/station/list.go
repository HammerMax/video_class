package station

import (
	"net/http"
	"server/modules/video"
	"server/utils"
)

func List(w http.ResponseWriter, r *http.Request) {
	videos, err := video.List()
	if err != nil {
		utils.ResponseJson(w, -1, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, videos)
}

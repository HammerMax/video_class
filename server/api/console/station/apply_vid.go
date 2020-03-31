package station

import (
	"fmt"
	"net/http"
	"server/utils"
	"time"
)

func ApplyVid(w http.ResponseWriter, r *http.Request) {
	utils.ResponseSuccess(w, map[string]string{
		"vid": generateVid(),
	})
}

func generateVid() string {
	return fmt.Sprintf("%x", time.Now().Unix() & 0xFFFFFF)
}

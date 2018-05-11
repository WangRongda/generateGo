
import "net/http"

/* ------------
 * 申请自定义技能
 * ------------*/
func VoiceSkillApply(w http.ResponseWriter, r *http.Request) {
	var resData struct {
		Skillid int `json:"skillid"`
		Status  int `json:"status"`
		Skill   struct {
			IntLinst []int `json:"intLinst"`
			List     []struct {
				ID      int  `json:"id"`
				Publish bool `json:"publish"`
			} `json:"list"`
			Status bool `json:"status"`
		} `json:"skill"`
	}

}

/* ------------
 * 申请自定义技能
 * ------------*/
func VoiceSkillApply(w http.ResponseWriter, r *http.Request) {
	var resData struct {
		Skillid int `json:"skillid"`
		Status int `json:"status"`
		Skill struct {
			IntLinst []int `json:"intLinst"`
			List []struct {
				ID int `json:"id"`
				Publish bool `json:"publish"`
			} `json:"list"`
			Status bool `json:"status"`
		} `json:"skill"`
	}

}

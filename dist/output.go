
/** 申请自定义技能 **/
func VoiceSkillApply(w http.ResponseWriter, r *http.Request) {
	var resData struct {
		Skillid int `json:"skillid"`
		Status int `json:"status"`
		Skill struct {
			Status bool `json:"status"`
			IntLinst []int `json:"intLinst"`
			List []struct {
				ID int `json:"id"`
				Publish bool `json:"publish"`
			} `json:"list"`
		} `json:"skill"`
	}
}/** 申请自定义技能 VoiceSkillApply()  **/

/** 王蓉打 **/
func Wangrongda(w http.ResponseWriter, r *http.Request) {
	var resData KitResData}/** 王蓉打 Wangrongda()  **/

/** 申请自定义技能 **/
func VoiceSkillApply(w http.ResponseWriter, r *http.Request) {
struct {
		Name string `json:"name"`
		Platform int `json:"platform"`
		Appid int `json:"appid"`
		Oauthid int `json:"oauthid"`
	}
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
}/** 申请自定义技能 VoiceSkillApply()  **/

/** 王蓉打 **/
func Wangrongda(w http.ResponseWriter, r *http.Request) {
struct {
		Appid int `json:"appid"`
		Oauthid int `json:"oauthid"`
		Name string `json:"name"`
		Platform int `json:"platform"`
	}
	var resData KitResData}/** 王蓉打 Wangrongda()  **/

/** 申请自定义技能 **/
func VoiceSkillApply(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Oauthid int `json:"oauthid"`
		Name string `json:"name"`
		Platform int `json:"platform"`
		Appid int `json:"appid"`
	}
	var resData struct {
		Status int `json:"status"`
		Skill struct {
			IntLinst []int `json:"intLinst"`
			List []struct {
				ID int `json:"id"`
				Publish bool `json:"publish"`
			} `json:"list"`
			Status bool `json:"status"`
		} `json:"skill"`
		Skillid int `json:"skillid"`
	}
}/** 申请自定义技能 VoiceSkillApply()  **/

/** 王蓉打 **/
func Wangrongda(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Oauthid int `json:"oauthid"`
		Name string `json:"name"`
		Platform int `json:"platform"`
		Appid int `json:"appid"`
	}
	var resData KitResData}/** 王蓉打 Wangrongda()  **/

/** 申请自定义技能 **/
func VoiceSkillApply(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Name string `json:"name"`
		Platform int `json:"platform"`
		Appid int `json:"appid"`
		Oauthid int `json:"oauthid"`
	}
	var resData struct {
		Status int `json:"status"`
		Skill struct {
			IntLinst []int `json:"intLinst"`
			List []struct {
				Publish bool `json:"publish"`
				ID int `json:"id"`
			} `json:"list"`
			Status bool `json:"status"`
		} `json:"skill"`
		Skillid int `json:"skillid"`
	}
	
}/** 申请自定义技能 VoiceSkillApply()  **/

/** 王蓉打 **/
func Wangrongda(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Name string `json:"name"`
		Platform int `json:"platform"`
		Appid int `json:"appid"`
		Oauthid int `json:"oauthid"`
	}
	var resData KitResData
	defer func() {
		if "" == resData.Msg {
			resData.Msg = codemsg[resData.Code]
		}
		resBody, _ := json.Marshal(resData)
		w.Head().Set("Content-Type", "application/json")
		w.Write(resBod)
	}()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resData.Code = RequestParse
		return
	}
}/** 王蓉打 Wangrongda()  **/

/** 申请自定义技能 **/
func VoiceSkillApply(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Name string `json:"name"`
		Platform int `json:"platform"`
		Appid int `json:"appid"`
		Oauthid int `json:"oauthid"`
	}
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

	defer func() {
		if "" == resData.Msg {
			resData.Msg = codemsg[resData.Code]
		}
		resBody, _ := json.Marshal(resData)
		w.Head().Set("Content-Type", "application/json")
		w.Write(resBod)
	}()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resData.Code = RequestParse
		return
	}
	if err := json.Unmarshal(reqBody, &reqData); nil != err {
		resData.Code = RequestParse
		return
	}}/** 申请自定义技能 VoiceSkillApply()  **/

/** 王蓉打 **/
func Wangrongda(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Oauthid int `json:"oauthid"`
		Name string `json:"name"`
		Platform int `json:"platform"`
		Appid int `json:"appid"`
	}
	var resData KitResData

	defer func() {
		if "" == resData.Msg {
			resData.Msg = codemsg[resData.Code]
		}
		resBody, _ := json.Marshal(resData)
		w.Head().Set("Content-Type", "application/json")
		w.Write(resBod)
	}()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resData.Code = RequestParse
		return
	}
	if err := json.Unmarshal(reqBody, &reqData); nil != err {
		resData.Code = RequestParse
		return
	}}/** 王蓉打 Wangrongda()  **/

/** 申请自定义技能 **/
func VoiceSkillApply(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Oauthid int `json:"oauthid"`
		Name string `json:"name"`
		Platform int `json:"platform"`
		Appid int `json:"appid"`
	}
	var resData struct {
		Skill struct {
			IntLinst []int `json:"intLinst"`
			List []struct {
				Publish bool `json:"publish"`
				ID int `json:"id"`
			} `json:"list"`
			Status bool `json:"status"`
		} `json:"skill"`
		Skillid int `json:"skillid"`
		Status int `json:"status"`
	}

	defer func() {
		if "" == resData.Msg {
			resData.Msg = codemsg[resData.Code]
		}
		resBody, _ := json.Marshal(resData)
		w.Head().Set("Content-Type", "application/json")
		w.Write(resBod)
	}()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resData.Code = RequestParse
		return
	}
	if err := json.Unmarshal(reqBody, &reqData); nil != err {
		resData.Code = RequestParse
		return
	}
}/** 申请自定义技能 VoiceSkillApply()  **/

/** 王蓉打 **/
func Wangrongda(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Name string `json:"name"`
		Platform int `json:"platform"`
		Appid int `json:"appid"`
		Oauthid int `json:"oauthid"`
	}
	var resData KitResData

	defer func() {
		if "" == resData.Msg {
			resData.Msg = codemsg[resData.Code]
		}
		resBody, _ := json.Marshal(resData)
		w.Head().Set("Content-Type", "application/json")
		w.Write(resBod)
	}()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resData.Code = RequestParse
		return
	}
	if err := json.Unmarshal(reqBody, &reqData); nil != err {
		resData.Code = RequestParse
		return
	}
}/** 王蓉打 Wangrongda()  **/

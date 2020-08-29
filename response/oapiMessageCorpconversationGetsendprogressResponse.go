package response

type OapiMessageCorpconversationGetsendprogressResponse struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	Progress struct {
		ProgressInPercent int `json:"progress_in_percent"`
		Status            int `json:"status"`
	} `json:"progress"`
}

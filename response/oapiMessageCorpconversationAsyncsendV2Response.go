package response

type OapiMessageCorpconversationAsyncsendV2Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	TaskId  int64  `json:"task_id"`
}

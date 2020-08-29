package response

type OapiFileUploadSingleResponse struct {
	ErrCode int    `json:"errcode"`
	EerrMsg string `json:"errmsg"`
	MediaId string `json:"media_id"`
}

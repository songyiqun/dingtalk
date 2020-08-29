package response

type OapiFileUploadChunkResponse struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}

package response

type OapiFileUploadTransactionResponse struct {
	ErrMsg   string `json:"errmsg"`
	ErrCode  int    `json:"errcode"`
	UploadId string `json:"upload_id"`
	MediaId  string `json:"media_id"`
}

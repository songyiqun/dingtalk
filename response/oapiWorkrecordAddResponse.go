package response

type OapiWorkrecordAddResponse struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	RecordId string `json:"record_id"`
}

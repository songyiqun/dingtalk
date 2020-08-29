package response

type OapiWorkrecordUpdateResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Result  bool   `json:"result"`
}

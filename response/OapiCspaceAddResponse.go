package response


type OapiCspaceAddResponse struct {
	ErrCode int    `json:"errcode"`
	EerrMsg string `json:"errmsg"`
	Dentry  string `json:"dentry"`
}
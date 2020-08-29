package response

type OapiUserGetByMobileResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	UserId  string `json:"userid"`
}

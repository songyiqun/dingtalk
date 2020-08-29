package response

type OapiServiceGetCorpTokenResponse struct {
	ErrCode     int    `json:"errcode"`
	EerrMsg     string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

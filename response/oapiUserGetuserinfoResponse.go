package response

type OapiUserGetuserinfoResponse struct {
	DeviceId string `json:"deviceId"`
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	IsSys    bool   `json:"is_sys"`
	SysLevel int    `json:"sys_level"`
	UserId   string `json:"userid"`
}

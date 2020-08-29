package response

type OapiCspaceGetCustomSpaceResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	SpaceId int64  `json:"spaceid"`
}

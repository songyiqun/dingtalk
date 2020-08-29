package request

type OapiUserGetuserinfoRequest struct {
	BaseDingTalkRequest
	code string
}

func (this *OapiUserGetuserinfoRequest) GetApiName() string {
	return "dingtalk.oapi.user.getuserinfo"
}

func (this *OapiUserGetuserinfoRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiUserGetuserinfoRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["code"] = this.code
	return parameters
}

func (this *OapiUserGetuserinfoRequest) Validate() {}

func (this *OapiUserGetuserinfoRequest) SetCode(code string) {
	this.code = code
}

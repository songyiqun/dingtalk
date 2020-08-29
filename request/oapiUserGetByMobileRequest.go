package request

type OapiUserGetByMobileRequest struct {
	BaseDingTalkRequest
	mobile string
}

func (this *OapiUserGetByMobileRequest) GetApiName() string {
	return "dingtalk.oapi.user.get_by_mobile"
}

func (this *OapiUserGetByMobileRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiUserGetByMobileRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["mobile"] = this.mobile
	return parameters
}

func (this *OapiUserGetByMobileRequest) Validate() {
	ValidateRequired("mobile", this.mobile)
}

func (this *OapiUserGetByMobileRequest) SetMobile(mobile string) {
	this.mobile = mobile
}

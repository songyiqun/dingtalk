package request

type OapiUserGetRequest struct {
	BaseDingTalkRequest
	userid string
	lang   string
}

func (this *OapiUserGetRequest) GetApiName() string {
	return "dingtalk.oapi.user.get"
}

func (this *OapiUserGetRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiUserGetRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["userid"] = this.userid
	if this.lang == "" {
		parameters["lang"] = "zh_CN"
	}
	return parameters
}

func (this *OapiUserGetRequest) Validate() {
	ValidateRequired("userid", this.userid)
}

func (this *OapiUserGetRequest) SetUserId(userid string) {
	this.userid = userid
}

func (this *OapiUserGetRequest) SetLang(lang string) {
	this.lang = lang
}

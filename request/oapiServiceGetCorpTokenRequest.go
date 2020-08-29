package request

type OapiServiceGetCorpTokenRequest struct {
	BaseDingTalkRequest
	accessKey     string
	timestamp  string
	suiteTicket     string
	signature string
	authCorpid string
}

func (this *OapiServiceGetCorpTokenRequest) GetApiName() string {
	return "dingtalk.oapi.service.get_corp_token"
}

func (this *OapiServiceGetCorpTokenRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiServiceGetCorpTokenRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["auth_corpid"] = this.authCorpid
	return parameters
}

func (this *OapiServiceGetCorpTokenRequest) GetUrlParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["accessKey"] = this.accessKey
	parameters["timestamp"] = this.timestamp
	parameters["suiteTicket"] = this.suiteTicket
	parameters["signature"] = this.signature
	return parameters
}

//发送前检查数据
func (this *OapiServiceGetCorpTokenRequest) Validate() {
	ValidateRequired("accessKey", this.accessKey)
}

func (this *OapiServiceGetCorpTokenRequest) SetAccessKey(accessKey string) {
	this.accessKey = accessKey
}

func (this *OapiServiceGetCorpTokenRequest) SetTimestamp(timestamp string) {
	this.timestamp = timestamp
}

func (this *OapiServiceGetCorpTokenRequest) SetSuiteTicket(suiteTicket string) {
	this.suiteTicket = suiteTicket
}

func (this *OapiServiceGetCorpTokenRequest) SetSignature(signature string) {
	this.signature = signature
}

func (this *OapiServiceGetCorpTokenRequest) SetAuthCorpid(authCorpid string) {
	this.authCorpid = authCorpid
}
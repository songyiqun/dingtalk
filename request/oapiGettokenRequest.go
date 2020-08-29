package request

type OapiGettokenRequest struct {
	BaseDingTalkRequest
	appkey     string
	appsecret  string
	corpid     string
	corpsecret string
}

func (this *OapiGettokenRequest) GetApiName() string {
	return "dingtalk.oapi.gettoken"
}

func (this *OapiGettokenRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiGettokenRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["appkey"] = this.appkey
	parameters["appsecret"] = this.appsecret
	if this.corpid != "" {
		parameters["corpid"] = this.corpid
	}
	if this.corpsecret != "" {
		parameters["corpsecret"] = this.corpsecret
	}
	return parameters
}

//发送前检查数据
func (this *OapiGettokenRequest) Validate() {
	ValidateRequired("appkey", this.appkey)
	ValidateRequired("appsecret", this.appsecret)
}

func (this *OapiGettokenRequest) SetAppKey(appKey string) {
	this.appkey = appKey
}

func (this *OapiGettokenRequest) SetAppsecret(appsecret string) {
	this.appsecret = appsecret
}

func (this *OapiGettokenRequest) SetCorpId(corpId string) {
	this.corpid = corpId
}

func (this *OapiGettokenRequest) SetCorpSecret(corpSecret string) {
	this.corpsecret = corpSecret
}

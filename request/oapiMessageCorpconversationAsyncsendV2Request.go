package request

type OapiMessageCorpconversationAsyncsendV2Request struct {
	BaseDingTalkRequest
	agentId    string
	useridList string
	deptIdList string
	toAllUser  bool
	msg        interface{}
}

func (this *OapiMessageCorpconversationAsyncsendV2Request) GetApiName() string {
	return "dingtalk.oapi.message.corpconversation.asyncsend_v2"
}

func (this *OapiMessageCorpconversationAsyncsendV2Request) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiMessageCorpconversationAsyncsendV2Request) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["agent_id"] = this.agentId
	if this.useridList != "" {
		parameters["userid_list"] = this.useridList
	}
	if this.deptIdList != "" {
		parameters["dept_id_list"] = this.deptIdList
	}
	if this.toAllUser == true{
		parameters["to_all_user"] = this.toAllUser
	}
	parameters["msg"] = this.msg
	return parameters
}

//发送前检查数据
func (this *OapiMessageCorpconversationAsyncsendV2Request) Validate() {
	ValidateRequired("agent_id", this.agentId)
	//ValidateRequired("msg", this.msg)
}

func (this *OapiMessageCorpconversationAsyncsendV2Request) SetAgentId(agentId string) {
	this.agentId = agentId
}

func (this *OapiMessageCorpconversationAsyncsendV2Request) SetUseridList(useridList string) {
	this.useridList = useridList
}

func (this *OapiMessageCorpconversationAsyncsendV2Request) SetDeptIdList(deptIdList string) {
	this.deptIdList = deptIdList
}

func (this *OapiMessageCorpconversationAsyncsendV2Request) SetToAllUser(toAllUser bool) {
	this.toAllUser = toAllUser
}

func (this *OapiMessageCorpconversationAsyncsendV2Request) SetMsg(msg interface{}) {
	this.msg = msg
}

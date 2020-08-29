package request

type OapiMessageCorpconversationRecallRequest struct {
	BaseDingTalkRequest
	agentId   string
	msgTaskId int64
}

func (this *OapiMessageCorpconversationRecallRequest) GetApiName() string {
	return "dingtalk.oapi.message.corpconversation.recall"
}

func (this *OapiMessageCorpconversationRecallRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiMessageCorpconversationRecallRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["agent_id"] = this.agentId
	parameters["msg_task_id"] = this.msgTaskId
	return parameters
}

func (this *OapiMessageCorpconversationRecallRequest) SetAgentId(agentId string) {
	this.agentId = agentId
}

func (this *OapiMessageCorpconversationRecallRequest) SetMsgTaskId(msgTaskId int64) {
	this.msgTaskId = msgTaskId
}

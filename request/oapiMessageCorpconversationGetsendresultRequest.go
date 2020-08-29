package request

type OapiMessageCorpconversationGetsendresultRequest struct {
	BaseDingTalkRequest
	agentId string
	taskId  int64
}

func (this *OapiMessageCorpconversationGetsendresultRequest) GetApiName() string {
	return "dingtalk.oapi.message.corpconversation.getsendresult"
}

func (this *OapiMessageCorpconversationGetsendresultRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiMessageCorpconversationGetsendresultRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["agent_id"] = this.agentId
	parameters["task_id"] = this.taskId
	return parameters
}

func (this *OapiMessageCorpconversationGetsendresultRequest) SetAgentId(agentId string) {
	this.agentId = agentId
}

func (this *OapiMessageCorpconversationGetsendresultRequest) SetTaskId(taskId int64) {
	this.taskId = taskId
}

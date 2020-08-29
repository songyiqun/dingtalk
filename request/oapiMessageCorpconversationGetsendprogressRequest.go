package request

type OapiMessageCorpconversationGetsendprogressRequest struct {
	BaseDingTalkRequest
	agentId string
	taskId  int64
}

func (this *OapiMessageCorpconversationGetsendprogressRequest) GetApiName() string {
	return "dingtalk.oapi.message.corpconversation.getsendprogress"
}

func (this *OapiMessageCorpconversationGetsendprogressRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiMessageCorpconversationGetsendprogressRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["agent_id"] = this.agentId
	parameters["task_id"] = this.taskId
	return parameters
}

func (this *OapiMessageCorpconversationGetsendprogressRequest) SetAgentId(agentId string) {
	this.agentId = agentId
}

func (this *OapiMessageCorpconversationGetsendprogressRequest) SetTaskId(taskId int64) {
	this.taskId = taskId
}

package request

//OapiCspaceGrantCustomSpaceRequest

type OapiCspaceGrantCustomSpaceRequest struct {
	BaseDingTalkRequest
	agentId    string
	domain     string
	actionType string
	userId     string
	path       string
	fileids    string
	duration   string
}

func (this *OapiCspaceGrantCustomSpaceRequest) GetApiName() string {
	return "dingtalk.oapi.file.upload.chunk"
}

func (this *OapiCspaceGrantCustomSpaceRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiCspaceGrantCustomSpaceRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["agent_id"] = this.agentId
	parameters["domain"] = this.domain
	parameters["type"] = this.actionType
	parameters["userid"] = this.userId
	if this.path != ""{
		parameters["path"] = this.path
	}
	parameters["fileids"] = this.fileids
	parameters["duration"] = this.duration
	return parameters
}

func (this *OapiCspaceGrantCustomSpaceRequest) Validate() {}

func (this *OapiCspaceGrantCustomSpaceRequest) SetAgentId(agentId string) {
	this.agentId = agentId
}

func (this *OapiCspaceGrantCustomSpaceRequest) SetDomain(domain string) {
	this.domain = domain
}

func (this *OapiCspaceGrantCustomSpaceRequest) SetActionType(actionType string) {
	this.actionType = actionType
}

func (this *OapiCspaceGrantCustomSpaceRequest) SetUserId(userId string) {
	this.userId = userId
}

func (this *OapiCspaceGrantCustomSpaceRequest) SetPath(path string) {
	this.path = path
}

func (this *OapiCspaceGrantCustomSpaceRequest) SetFileids(fileids string) {
	this.fileids = fileids
}

func (this *OapiCspaceGrantCustomSpaceRequest) SetDuration(duration string) {
	this.duration = duration
}

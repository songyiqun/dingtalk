package request

type OapiCspaceGetCustomSpaceRequest struct {
	BaseDingTalkRequest
	domain  string
	agentId string
}

func (this *OapiCspaceGetCustomSpaceRequest) GetApiName() string {
	return "dingtalk.oapi.cspace.get_custom_space"
}

func (this *OapiCspaceGetCustomSpaceRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiCspaceGetCustomSpaceRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["domain"] = this.domain
	parameters["agent_id"] = this.agentId
	return parameters
}

func (this *OapiCspaceGetCustomSpaceRequest) Validate() {}

func (this *OapiCspaceGetCustomSpaceRequest) SetDomain(domain string) {
	this.domain = domain
}

func (this *OapiCspaceGetCustomSpaceRequest) SetAgentId(agentId string) {
	this.agentId = agentId
}

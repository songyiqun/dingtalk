package request

type OapiCspaceAddRequest struct {
	BaseDingTalkRequest
	agentId   string
	code      string
	mediaId   string
	spaceId   string
	folderId  string
	name      string
	overWrite string
	path      string
}

func (this *OapiCspaceAddRequest) GetApiName() string {
	return "dingtalk.oapi.cspace.add"
}

func (this *OapiCspaceAddRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiCspaceAddRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["agent_id"] = this.agentId
	parameters["code"] = this.code
	parameters["media_id"] = this.mediaId
	parameters["space_id"] = this.spaceId
	if this.folderId != "" {
		parameters["folder_id"] = this.folderId
	}
	parameters["name"] = this.name
	parameters["overwrite"] = this.overWrite
	if this.path != "" {
		parameters["path"] = this.path
	}
	return parameters
}

func (this *OapiCspaceAddRequest) Validate() {}

func (this *OapiCspaceAddRequest) SetAgentId(agentId string) {
	this.agentId = agentId
}

func (this *OapiCspaceAddRequest) SetCode(code string) {
	this.code = code
}

func (this *OapiCspaceAddRequest) SetMediaId(mediaId string) {
	this.mediaId = mediaId
}

func (this *OapiCspaceAddRequest) SetSpaceId(spaceId string) {
	this.spaceId = spaceId
}

func (this *OapiCspaceAddRequest) SetFolderId(folderId string) {
	this.folderId = folderId
}

func (this *OapiCspaceAddRequest) SetName(name string) {
	this.name = name
}

func (this *OapiCspaceAddRequest) SetOverWrite(overWrite string) {
	this.overWrite = overWrite
}

func (this *OapiCspaceAddRequest) SetPath(path string) {
	this.path = path
}
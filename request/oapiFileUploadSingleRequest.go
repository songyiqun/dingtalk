package request

type OapiFileUploadSingleRequest struct {
	BaseDingTalkRequest
	agentId  string
	fileSize string
	filePath string
}

func (this *OapiFileUploadSingleRequest) GetApiName() string {
	return "dingtalk.oapi.cspace.get_custom_space"
}

func (this *OapiFileUploadSingleRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiFileUploadSingleRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["agent_id"] = this.agentId
	parameters["file_size"] = this.fileSize
	return parameters
}

func (this *OapiFileUploadSingleRequest) Validate() {}

func (this *OapiFileUploadSingleRequest) SetAgentId(agentId string) {
	this.agentId = agentId
}

func (this *OapiFileUploadSingleRequest) SetFileSize(fileSize string) {
	this.fileSize = fileSize
}

func (this *OapiFileUploadSingleRequest) SetFilePath(filePath string) {
	this.filePath = filePath
}

func (this *OapiFileUploadSingleRequest) GetFilePath() string {
	return this.filePath
}

package request

type OapiFileUploadChunkRequest struct {
	BaseDingTalkRequest
	agentId       string
	uploadId      string
	chunkSequence string
	filePath      string
}

func (this *OapiFileUploadChunkRequest) GetApiName() string {
	return "dingtalk.oapi.file.upload.chunk"
}

func (this *OapiFileUploadChunkRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiFileUploadChunkRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["agent_id"] = this.agentId
	parameters["upload_id"] = this.uploadId
	parameters["chunk_sequence"] = this.chunkSequence
	return parameters
}

func (this *OapiFileUploadChunkRequest) Validate() {}

func (this *OapiFileUploadChunkRequest) SetAgentId(agentId string) {
	this.agentId = agentId
}

func (this *OapiFileUploadChunkRequest) SetUploadId(uploadId string) {
	this.uploadId = uploadId
}

func (this *OapiFileUploadChunkRequest) SetChunkSequence(chunkSequence string) {
	this.chunkSequence = chunkSequence
}

func (this *OapiFileUploadChunkRequest) SetFilePath(filePath string) {
	this.filePath = filePath
}

func (this *OapiFileUploadChunkRequest) GetFilePath() string {
	return this.filePath
}

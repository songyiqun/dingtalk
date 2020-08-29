package request

type OapiFileUploadTransactionRequest struct {
	BaseDingTalkRequest
	agentId      string
	fileSize     string
	chunkNumbers string
	uploadId     string
}

func (this *OapiFileUploadTransactionRequest) GetApiName() string {
	return "dingtalk.oapi.file.upload.transaction"
}

func (this *OapiFileUploadTransactionRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiFileUploadTransactionRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["agent_id"] = this.agentId
	parameters["file_size"] = this.fileSize
	parameters["chunk_numbers"] = this.chunkNumbers
	if this.uploadId != "" {
		parameters["upload_id"] = this.uploadId
	}
	return parameters
}

func (this *OapiFileUploadTransactionRequest) Validate() {}

func (this *OapiFileUploadTransactionRequest) SetAgentId(agentId string) {
	this.agentId = agentId
}

func (this *OapiFileUploadTransactionRequest) SetFileSize(fileSize string) {
	this.fileSize = fileSize
}

func (this *OapiFileUploadTransactionRequest) SetChunkNumbers(chunkNumbers string) {
	this.chunkNumbers = chunkNumbers
}

func (this *OapiFileUploadTransactionRequest) SetUploadId(uploadId string) {
	this.uploadId = uploadId
}

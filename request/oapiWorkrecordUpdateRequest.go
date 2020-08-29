package request

type OapiWorkrecordUpdateRequest struct {
	BaseDingTalkRequest
	userId   string
	recordId string
}

func (this *OapiWorkrecordUpdateRequest) GetApiName() string {
	return "dingtalk.oapi.workrecord.update"
}

func (this *OapiWorkrecordUpdateRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiWorkrecordUpdateRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["userid"] = this.userId
	parameters["record_id"] = this.recordId
	return parameters
}

func (this *OapiWorkrecordUpdateRequest) Validate() {

}

func (this *OapiWorkrecordUpdateRequest) SetUserId(userId string) {
	this.userId = userId
}

func (this *OapiWorkrecordUpdateRequest) SetRecordId(recordId string) {
	this.recordId = recordId
}

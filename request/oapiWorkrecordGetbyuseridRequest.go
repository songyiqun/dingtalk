package request

import "fmt"

type OapiWorkrecordGetbyuseridRequest struct {
	BaseDingTalkRequest
	userId string
	offset int
	limit  int
	status int
}

func (this *OapiWorkrecordGetbyuseridRequest) GetApiName() string {
	return "dingtalk.oapi.workrecord.getbyuserid"
}

func (this *OapiWorkrecordGetbyuseridRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiWorkrecordGetbyuseridRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["userid"] = this.userId
	parameters["offset"] = this.offset
	parameters["limit"] = this.limit
	parameters["status"] = this.status
	return parameters
}

//检查数据
func (this *OapiWorkrecordGetbyuseridRequest) Validate() {
	if this.userId == "" {
		fmt.Println(this.GetApiName() + "Validate userid")
		return
	}
	if this.offset == 0 {
		fmt.Println(this.GetApiName() + "Validate offset")
		return
	}
	if this.limit == 0 {
		fmt.Println(this.GetApiName() + "Validate limit")
		return
	}
	if this.status == 0 {
		fmt.Println(this.GetApiName() + "Validate status")
		return
	}
}

func (this *OapiWorkrecordGetbyuseridRequest) SetUserId(userId string) {
	this.userId = userId
}

func (this *OapiWorkrecordGetbyuseridRequest) SetOffset(offset int) {
	this.offset = offset
}

func (this *OapiWorkrecordGetbyuseridRequest) SetLimit(limit int) {
	this.limit = limit
}

func (this *OapiWorkrecordGetbyuseridRequest) SetStatus(status int) {
	this.status = status
}

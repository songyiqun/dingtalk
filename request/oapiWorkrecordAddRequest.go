package request

import (
	"github.com/songyiqun/dingtalk/utils"
)

type OapiWorkrecordAddRequest struct {
	BaseDingTalkRequest
	userid       string
	createTime   int64
	title        string
	url          string
	pcUrl        string //pc端跳转url，不传则使用url参数 (可空)
	formItemList []FormItemList
	sourceName   string //待办来源名称 (可空)
	pcOpenType   string    //待办的pc打开方式。2表示在pc端打开，4表示在浏览器打开 (可空)
	bizId        string
}

func (this *OapiWorkrecordAddRequest) GetApiName() string {
	return "dingtalk.oapi.workrecord.add"
}

func (this *OapiWorkrecordAddRequest) GetApiCallType() string {
	return CALL_TYPE_OAPI
}

func (this *OapiWorkrecordAddRequest) GetParameters() map[string]interface{} {
	parameters := make(map[string]interface{})
	parameters["userid"] = this.userid
	parameters["create_time"] = utils.IntToString(this.createTime)
	parameters["title"] = this.title
	parameters["url"] = this.url
	if this.pcUrl != ""{
		parameters["pcUrl"] = this.pcUrl
	}

	parameters["formItemList"] = this.formItemList
	if this.sourceName != ""{
		parameters["source_name"] = this.sourceName
	}
	if this.pcOpenType != ""{
		parameters["pc_open_type"] = this.pcOpenType
	}
	if this.bizId != "" {
		parameters["biz_id"] = this.bizId
	}
	return parameters
}

func (this *OapiWorkrecordAddRequest) Validate() {}

func (this *OapiWorkrecordAddRequest) SetUserId(code string) {
	this.userid = code
}

func (this *OapiWorkrecordAddRequest) SetCreateTime(createTime int64) {
	this.createTime = createTime
}

func (this *OapiWorkrecordAddRequest) SetTitle(title string) {
	this.title = title
}

func (this *OapiWorkrecordAddRequest) SetUrl(url string) {
	this.url = url
}

func (this *OapiWorkrecordAddRequest) SetPcUrl(pcUrl string) {
	this.pcUrl = pcUrl
}

func (this *OapiWorkrecordAddRequest) SetFormItemList(formItemList []FormItemList) {
	this.formItemList = formItemList
}

func (this *OapiWorkrecordAddRequest) SetSourceName(sourceName string) {
	this.sourceName = sourceName
}

func (this *OapiWorkrecordAddRequest) SetPcOpenType(pcOpenType string) {
	this.pcOpenType = pcOpenType
}

func (this *OapiWorkrecordAddRequest) SetBizId(bizId string) {
	this.bizId = bizId
}

type FormItemList struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
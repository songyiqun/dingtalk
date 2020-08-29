package dingtalk

import (
	"encoding/json"
	"fmt"
	"github.com/songyiqun/dingtalk/request"
	"github.com/songyiqun/dingtalk/response"
	"github.com/songyiqun/dingtalk/utils"
)

type DefaultDingTalkClient struct {
	IDingTalkClient
	serverUrl        string
	httpMethod       string
	systemParameters map[string]string
}

func NewDefaultDingTalkClient(serverUrl string) *DefaultDingTalkClient {
	return &DefaultDingTalkClient{
		serverUrl: serverUrl,
	}
}

func (this *DefaultDingTalkClient) OapiGettokenExecute(request request.OapiGettokenRequest) (response response.OapiGettokenResponse, err error) {
	var result []byte
	request.GetApiName()
	request.Validate()
	textParams := request.GetParameters()
	result, err = utils.DoGet(this.serverUrl, textParams, nil)
	if err != nil {
		return response, err
	}
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiUserGetuserinfoExecute(request request.OapiUserGetuserinfoRequest, token string) (response response.OapiUserGetuserinfoResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	textParams[ACCESS_TOKEN] = token
	result, err = utils.DoGet(this.serverUrl, textParams, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiCspaceGetCustomSpaceExecute(request request.OapiCspaceGetCustomSpaceRequest, token string) (response response.OapiCspaceGetCustomSpaceResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	textParams[ACCESS_TOKEN] = token
	result, err = utils.DoGet(this.serverUrl, textParams, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

//发起待办
func (this *DefaultDingTalkClient) OapiWorkrecordAddExecute(request request.OapiWorkrecordAddRequest, token string) (response response.OapiWorkrecordAddResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	otherParams := make(map[string]interface{})
	otherParams[ACCESS_TOKEN] = token
	result, err = utils.DoPost(this.serverUrl, textParams, nil, otherParams)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

//更新待办
func (this *DefaultDingTalkClient) OapiWorkrecordUpdateExecute(request request.OapiWorkrecordUpdateRequest, token string) (response response.OapiWorkrecordUpdateResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	otherParams := make(map[string]interface{})
	otherParams[ACCESS_TOKEN] = token
	result, err = utils.DoPost(this.serverUrl, textParams, nil, otherParams)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

//获取用户待办事项
func (this *DefaultDingTalkClient) OapiWorkrecordGetbyuseridExecute(request request.OapiWorkrecordGetbyuseridRequest, token string) (response response.OapiWorkrecordGetbyuseridResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	otherParams := make(map[string]interface{})
	otherParams[ACCESS_TOKEN] = token
	result, err = utils.DoPost(this.serverUrl, textParams, nil, otherParams)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

//工作消息
func (this *DefaultDingTalkClient) OapiMessageCorpconversationAsyncsendV2Execute(request request.OapiMessageCorpconversationAsyncsendV2Request, token string) (response response.OapiMessageCorpconversationAsyncsendV2Response, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	otherParams := make(map[string]interface{})
	otherParams[ACCESS_TOKEN] = token
	result, err = utils.DoPost(this.serverUrl, textParams, nil, otherParams)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

//发送工作消息进度
func (this *DefaultDingTalkClient) OapiMessageCorpconversationGetsendprogressExecute(request request.OapiMessageCorpconversationGetsendprogressRequest, token string) (response response.OapiMessageCorpconversationGetsendprogressResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	otherParams := make(map[string]interface{})
	otherParams[ACCESS_TOKEN] = token
	result, err = utils.DoPost(this.serverUrl, textParams, nil, otherParams)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiMessageCorpconversationGetsendresultExecute(request request.OapiMessageCorpconversationGetsendresultRequest, token string) (response response.OapiMessageCorpconversationGetsendresultResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	otherParams := make(map[string]interface{})
	otherParams[ACCESS_TOKEN] = token
	result, err = utils.DoPost(this.serverUrl, textParams, nil, otherParams)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiMessageCorpconversationRecallExecute(request request.OapiMessageCorpconversationRecallRequest, token string) (response response.OapiMessageCorpconversationRecallResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	otherParams := make(map[string]interface{})
	otherParams[ACCESS_TOKEN] = token
	result, err = utils.DoPost(this.serverUrl, textParams, nil, otherParams)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiUserGetByMobileExecute(request request.OapiUserGetByMobileRequest, token string) (response response.OapiUserGetByMobileResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	textParams[ACCESS_TOKEN] = token
	result, err = utils.DoGet(this.serverUrl, textParams, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiUserGetExecute(request request.OapiUserGetRequest, token string) (response response.OapiUserGetResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	textParams[ACCESS_TOKEN] = token
	result, err = utils.DoGet(this.serverUrl, textParams, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiFileUploadSingleExecute(request request.OapiFileUploadSingleRequest, token string) (response response.OapiFileUploadSingleResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	textParams[ACCESS_TOKEN] = token
	result, err = utils.DoFilePost(this.serverUrl, textParams, "file", request.GetFilePath())
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiFileUploadTransactionExecute(request request.OapiFileUploadTransactionRequest, token string) (response response.OapiFileUploadTransactionResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	textParams[ACCESS_TOKEN] = token
	result, err = utils.DoGet(this.serverUrl, textParams, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiFileUploadChunkExecute(request request.OapiFileUploadChunkRequest, token string) (response response.OapiFileUploadChunkResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	textParams[ACCESS_TOKEN] = token
	result, err = utils.DoFilePost(this.serverUrl, textParams, "file", request.GetFilePath())
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiCspaceGrantCustomSpaceExecute(request request.OapiCspaceGrantCustomSpaceRequest, token string) (response response.OapiCspaceGrantCustomSpaceResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	textParams[ACCESS_TOKEN] = token
	result, err = utils.DoGet(this.serverUrl, textParams, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

func (this *DefaultDingTalkClient) OapiCspaceAddExecute(request request.OapiCspaceAddRequest, token string) (response response.OapiCspaceAddResponse, err error) {
	var result []byte
	request.GetApiName()
	textParams := request.GetParameters()
	textParams[ACCESS_TOKEN] = token
	result, err = utils.DoGet(this.serverUrl, textParams, nil)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	return response, err
}


func (this *DefaultDingTalkClient) OapiServiceGetCorpTokenExecute(request request.OapiServiceGetCorpTokenRequest)(response response.OapiServiceGetCorpTokenResponse, err error){
	var result []byte
	request.GetApiName()
	otherParams := request.GetUrlParameters()
	textParams := request.GetParameters()
	result, err = utils.DoPost(this.serverUrl, textParams, nil, otherParams)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(result, &response)
	if err != nil {
		return response, err
	}
	fmt.Println(string(result))
	return response, err
}
package dingtalk

import (
	"github.com/songyiqun/dingtalk/request"
	"github.com/songyiqun/dingtalk/response"
)

func OapiGettokenHandle(params OapiGettokenParams) (response.OapiGettokenResponse, error) {
	var err error
	var request = request.OapiGettokenRequest{}
	var res = response.OapiGettokenResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/gettoken")
	request.SetHttpMethod("GET")
	request.SetAppKey(params.Appkey)
	request.SetAppsecret(params.Appsecret)
	request.SetCorpId(params.Corpid)
	request.SetCorpSecret(params.Corpsecret)
	res, err = client.OapiGettokenExecute(request)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiUserGetByMobileHandle(params OapiUserGetByMobileParams, token string) (response.OapiUserGetByMobileResponse, error) {
	var err error
	var request = request.OapiUserGetByMobileRequest{}
	var res = response.OapiUserGetByMobileResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/user/get_by_mobile")
	request.SetMobile(params.Mobile)
	res, err = client.OapiUserGetByMobileExecute(request, token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiUserGetHandle(params OapiUserGetParams, token string) (response.OapiUserGetResponse, error) {
	var err error
	var request = request.OapiUserGetRequest{}
	var res = response.OapiUserGetResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/user/get")
	request.SetUserId(params.UserId)
	request.SetLang(params.Lang)
	res, err = client.OapiUserGetExecute(request, token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiUserGetuserinfoHandle(params OapiUserGetuserinfoParams, token string) (response.OapiUserGetuserinfoResponse, error) {
	var err error
	var request = request.OapiUserGetuserinfoRequest{}
	var res = response.OapiUserGetuserinfoResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/user/getuserinfo")
	request.SetCode(params.Code)
	res, err = client.OapiUserGetuserinfoExecute(request, token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiCspaceGetCustomSpaceHandle(params OapiCspaceGetCustomSpaceParams, token string) (response.OapiCspaceGetCustomSpaceResponse, error) {
	var err error
	var request = request.OapiCspaceGetCustomSpaceRequest{}
	var res = response.OapiCspaceGetCustomSpaceResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/cspace/get_custom_space")
	request.SetAgentId(params.AgentId)
	request.SetDomain(params.Domain)
	res, err = client.OapiCspaceGetCustomSpaceExecute(request, token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiFileUploadSingleHandle(params OapiFileUploadSingleParams, token string) (response.OapiFileUploadSingleResponse, error) {
	var err error
	var request = request.OapiFileUploadSingleRequest{}
	var res = response.OapiFileUploadSingleResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/file/upload/single")
	request.SetAgentId(params.AgentId)
	request.SetFileSize(params.FileSize)
	request.SetFilePath(params.FilePath)
	res, err = client.OapiFileUploadSingleExecute(request, token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiFileUploadTransactionHandle(params OapiFileUploadTransactionParams, token string) (response.OapiFileUploadTransactionResponse, error) {
	var err error
	var request = request.OapiFileUploadTransactionRequest{}
	var res = response.OapiFileUploadTransactionResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/file/upload/transaction")
	request.SetAgentId(params.AgentId)
	request.SetFileSize(params.FileSize)
	request.SetChunkNumbers(params.ChunkNumbers)
	request.SetUploadId(params.UploadId)
	res, err = client.OapiFileUploadTransactionExecute(request, token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiFileUploadChunkHandle(params OapiFileUploadChunkParams, token string) (response.OapiFileUploadChunkResponse, error) {
	var err error
	var request = request.OapiFileUploadChunkRequest{}
	var res = response.OapiFileUploadChunkResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/file/upload/chunk")
	request.SetAgentId(params.AgentId)
	request.SetUploadId(params.UploadId)
	request.SetChunkSequence(params.ChunkSequence)
	request.SetFilePath(params.FilePath)
	res, err = client.OapiFileUploadChunkExecute(request, token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiCspaceGrantCustomSpaceHandle(params OapiCspaceGrantCustomSpaceParams, token string)(response.OapiCspaceGrantCustomSpaceResponse, error){
	var err error
	var request = request.OapiCspaceGrantCustomSpaceRequest{}
	var res = response.OapiCspaceGrantCustomSpaceResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/cspace/grant_custom_space")
	request.SetAgentId(params.AgentId)
	request.SetDomain(params.Domain)
	request.SetActionType(params.ActionType)
	request.SetUserId(params.UserId)
	request.SetPath(params.Path)
	request.SetFileids(params.Fileids)
	request.SetDuration(params.Duration)
	res, err = client.OapiCspaceGrantCustomSpaceExecute(request, token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiCspaceAddHandle(params OapiCspaceAddParams, token string) (response.OapiCspaceAddResponse, error){
	var err error
	var request = request.OapiCspaceAddRequest{}
	var res = response.OapiCspaceAddResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/cspace/add")
	request.SetAgentId(params.AgentId)
	request.SetCode(params.Code)
	request.SetMediaId(params.MediaId)
	request.SetSpaceId(params.SpaceId)
	request.SetName(params.Name)
	request.SetOverWrite(params.OverWrite)
	res, err = client.OapiCspaceAddExecute(request, token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiWorkrecordAddHandle(params OapiWorkrecordAddParams, token string) (response.OapiWorkrecordAddResponse, error){
	var err error
	var request = request.OapiWorkrecordAddRequest{}
	var res = response.OapiWorkrecordAddResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/topapi/workrecord/add")
	request.SetUserId(params.Userid)
	request.SetCreateTime(params.CreateTime)
	request.SetTitle(params.Title)
	request.SetUrl(params.Url)
	request.SetPcUrl(params.PcUrl)
	request.SetFormItemList(params.FormItemList)
	request.SetSourceName(params.SourceName)
	request.SetPcOpenType(params.PcOpenType)
	request.SetBizId(params.BizId)
	res, err = client.OapiWorkrecordAddExecute(request,token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiMessageCorpconversationAsyncsendV2Handle(params OapiMessageCorpconversationAsyncsendV2Params, token string) (response.OapiMessageCorpconversationAsyncsendV2Response, error){
	var err error
	var request = request.OapiMessageCorpconversationAsyncsendV2Request{}
	var res = response.OapiMessageCorpconversationAsyncsendV2Response{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2")
	request.SetAgentId(params.AgentId)
	request.SetUseridList(params.UseridList)
	request.SetDeptIdList(params.DeptIdList)
	request.SetToAllUser(params.ToAllUser)
	request.SetMsg(params.Msg)
	res, err = client.OapiMessageCorpconversationAsyncsendV2Execute(request,token)
	if err != nil {
		return res, err
	}
	return res, err
}

func OapiServiceGetCorpTokenHandle(params OapiServiceGetCorpTokenParams)(response.OapiServiceGetCorpTokenResponse, error){
	var err error
	var request = request.OapiServiceGetCorpTokenRequest{}
	var res = response.OapiServiceGetCorpTokenResponse{}
	client := NewDefaultDingTalkClient("https://oapi.dingtalk.com/service/get_corp_token")
	request.SetAccessKey(params.AccessKey)
	request.SetSignature(params.Signature)
	request.SetSuiteTicket(params.SuiteTicket)
	request.SetTimestamp(params.Timestamp)
	request.SetAuthCorpid(params.AuthCorpid)
	res, err = client.OapiServiceGetCorpTokenExecute(request)
	if err != nil {
		return res, err
	}
	return res, err
}
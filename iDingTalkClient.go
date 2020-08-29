package dingtalk

import (
	"github.com/songyiqun/dingtalk/request"
	"github.com/songyiqun/dingtalk/response"
)

type IDingTalkClient interface {
	OapiGettokenExecute(request request.OapiGettokenRequest) (response.OapiGettokenResponse, error)
	OapiUserGetuserinfoExecute(request request.OapiUserGetuserinfoRequest, token string) (response.OapiUserGetuserinfoResponse, error)
	OapiCspaceGetCustomSpaceExecute(request request.OapiCspaceGetCustomSpaceRequest, token string) (response.OapiCspaceGetCustomSpaceResponse, error)
	OapiWorkrecordAddExecute(request request.OapiWorkrecordAddRequest, token string) (response.OapiWorkrecordAddResponse, error)
	OapiWorkrecordUpdateExecute(request request.OapiWorkrecordUpdateRequest, token string) (response.OapiWorkrecordUpdateResponse, error)
	OapiWorkrecordGetbyuseridExecute(request request.OapiWorkrecordGetbyuseridRequest, token string) (response.OapiWorkrecordGetbyuseridResponse, error)
	OapiMessageCorpconversationAsyncsendV2Execute(request request.OapiMessageCorpconversationAsyncsendV2Request, token string) (response response.OapiMessageCorpconversationAsyncsendV2Response, err error)
	OapiMessageCorpconversationGetsendprogressExecute(request request.OapiMessageCorpconversationGetsendprogressRequest, token string) (response response.OapiMessageCorpconversationGetsendprogressResponse, err error)
	OapiMessageCorpconversationGetsendresultExecute(request request.OapiMessageCorpconversationGetsendresultRequest, token string) (response response.OapiMessageCorpconversationGetsendresultResponse, err error)
	OapiMessageCorpconversationRecallExecute(request request.OapiMessageCorpconversationRecallRequest, token string) (response response.OapiMessageCorpconversationRecallResponse, err error)
	OapiUserGetByMobileExecute(request request.OapiUserGetByMobileRequest, token string) (response response.OapiUserGetByMobileResponse, err error)
	OapiUserGetExecute(request request.OapiUserGetRequest, token string) (response response.OapiUserGetResponse, err error)

	OapiFileUploadSingleExecute(request request.OapiFileUploadSingleRequest, token string) (response response.OapiFileUploadSingleResponse, err error)

	OapiFileUploadTransactionExecute(request request.OapiFileUploadTransactionRequest, token string) (response response.OapiFileUploadTransactionResponse, err error)
	OapiFileUploadChunkExecute(request request.OapiFileUploadChunkRequest, token string) (response response.OapiFileUploadChunkResponse, err error)
	OapiCspaceGrantCustomSpaceExecute(request request.OapiCspaceGrantCustomSpaceRequest, token string) (response response.OapiCspaceGrantCustomSpaceResponse, err error)
	OapiCspaceAddExecute(request request.OapiCspaceAddRequest, token string) (response response.OapiCspaceAddResponse, err error)
	OapiServiceGetCorpTokenExecute(request request.OapiServiceGetCorpTokenRequest)(response response.OapiServiceGetCorpTokenResponse, err error)
}

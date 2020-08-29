package dingtalk

import "github.com/songyiqun/dingtalk/request"

type OapiGettokenParams struct {
	Appkey     string
	Appsecret  string
	Corpid     string
	Corpsecret string
}

type OapiUserGetByMobileParams struct {
	Mobile string
}

type OapiUserGetParams struct {
	UserId string
	Lang   string
}

type OapiUserGetuserinfoParams struct {
	Code string
}

type OapiCspaceGetCustomSpaceParams struct {
	Domain  string
	AgentId string
}

type OapiFileUploadSingleParams struct {
	AgentId  string
	FileSize string
	FilePath string
}

type OapiFileUploadTransactionParams struct {
	AgentId      string
	FileSize     string
	ChunkNumbers string
	UploadId     string
}

type OapiFileUploadChunkParams struct {
	AgentId       string
	UploadId      string
	ChunkSequence string
	FilePath      string
}

type OapiCspaceGrantCustomSpaceParams struct {
	AgentId    string
	Domain     string
	ActionType string
	UserId     string
	Path       string
	Fileids    string
	Duration   string
}

type OapiCspaceAddParams struct {
	AgentId   string
	Code      string
	MediaId   string
	SpaceId   string
	FolderId  string
	Name      string
	OverWrite string
	Path      string
}

type OapiWorkrecordAddParams struct {
	Userid       string
	CreateTime   int64
	Title        string
	Url          string
	PcUrl        string //pc端跳转url，不传则使用url参数 (可空)
	FormItemList []request.FormItemList
	SourceName   string //待办来源名称 (可空)
	PcOpenType   string    //待办的pc打开方式。2表示在pc端打开，4表示在浏览器打开 (可空)
	BizId        string
}

type OapiMessageCorpconversationAsyncsendV2Params struct {
	AgentId    string
	UseridList string
	DeptIdList string
	ToAllUser  bool
	Msg        interface{}
}

type OapiServiceGetCorpTokenParams struct {
	AccessKey     string
	Timestamp  string
	SuiteTicket     string
	Signature string
	AuthCorpid string
}

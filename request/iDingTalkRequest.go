package request

type IDingTalkRequest interface {
	//获取TOP的API名称
	GetApiName() string

	//获取API请求方式
	GetApiCallType() string

	//获取所有的Key-Value形式的文本请求参数字典
	GetParameters() map[string]interface{}

	//获取自定义HTTP请求头参数
	GetHeaderParameters() map[string]string

	//客户端参数检查，减少服务端无效调用
	Validate()

	//获取http method
	GetHttpMethod() string
}

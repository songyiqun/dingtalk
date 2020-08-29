package request

type BaseDingTalkRequest struct {
	httpMethod string
}

func (this *BaseDingTalkRequest) SetHttpMethod(httpMethod string) {
	this.httpMethod = httpMethod
}

func (this *BaseDingTalkRequest) GetHttpMethod() string {
	return this.httpMethod
}

func (this *BaseDingTalkRequest) GetHeaderParameters() map[string]string {
	return make(map[string]string)
}

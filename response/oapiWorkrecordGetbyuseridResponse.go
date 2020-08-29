package response

type OapiWorkrecordGetbyuseridResponse struct {
	ErrCode int     `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
	Records records `json:"records"`
}

type records struct {
	HasMore bool `json:"has_more"`
	List    list `json:"list"`
}

type list struct {
	RecordId   string  `json:"record_id"`
	CreateTime int64   `json:"create_time"`
	Title      string  `json:"title"`
	Url        string  `json:"url"`
	forms      []forms `json:"forms"`
}

type forms struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

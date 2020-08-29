package response

type OapiMessageCorpconversationGetsendresultResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	SendResult struct {
		InvalidUserIdList   []string        `json:"invalid_user_id_list"`
		ForbiddenUserIdList []string        `json:"forbidden_user_id_list"`
		FailedUserIdList    []string        `json:"failed_user_id_list"`
		ReadUserIdList      []string        `json:"read_user_id_list"`
		UnreadUserIdList    []string        `json:"unread_user_id_list"`
		InvalidDeptIdList   []string        `json:"invalid_dept_id_list"`
		ForbiddenList       []ForbiddenList `json:"forbidden_list"`
	} `json:"send_result"`
}

type ForbiddenList struct {
	Code   int `json:"code"`
	Count  int `json:"count"`
	UserId int `json:"userid"`
}

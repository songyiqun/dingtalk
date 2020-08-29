package response

import "time"

type OapiUserGetResponse struct {
	ErrCode         int               `json:"errcode"`
	UnionId         string            `json:"unionid"`
	Remark          string            `json:"remark"`
	UserId          string            `json:"userid"`
	IsLeaderInDepts string            `json:"isLeaderInDepts"`
	IsBoss          bool              `json:"isBoss"`
	HiredDate       time.Time         `json:"hiredDate"`
	IsSenior        bool              `json:"isSenior"`
	Tel             string            `json:"tel"`
	Department      []int             `json:"department"`
	WorkPlace       string            `json:"workPlace"`
	Email           string            `json:"email"`
	OrderInDepts    string            `json:"orderInDepts"`
	Mobile          string            `json:"mobile"`
	ErrMsg          string            `json:"errmsg"`
	Active          bool              `json:"active"`
	Avatar          string            `json:"avatar"`
	IsAdmin         bool              `json:"isAdmin"`
	IsHide          bool              `json:"isHide"`
	Jobnumber       string            `json:"jobnumber"`
	Name            string            `json:"name"`
	Extattr         map[string]string `json:"extattr"`
	StateCode       string            `json:"stateCode"`
	Position        string            `json:"position"`
	Roles           []Roles           `json:"roles"`
}

type Roles struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
}

package response

type FileDentry struct {
	Attachment  struct{} `json:"attachment"`
	Cid         string   `json:"cid"`
	ContentType string   `json:"contentType"`
	CopyCount   int      `json:"copyCount"`
	CreateTime  int64    `json:"createTime"`
	Creator     struct {
		Account string `json:"account"`
		Site    int    `json:"site"`
	} `json:"creator"`
	DentrySpaceType string `json:"dentrySpaceType"`
	DownloadCount   int    `json:"downloadCount"`
	Encrypt         int    `json:"encrypt"`
	Extension       string `json:"extension"`
	Extra           string `json:"extra"`
	ID              string `json:"id"`
	ModifiedTime    int64  `json:"modifiedTime"`
	Modifier        struct {
		Account string `json:"account"`
		Site    int    `json:"site"`
	} `json:"modifier"`
	Name           string `json:"name"`
	OpenFileEdit   bool   `json:"openFileEdit"`
	OrgId          int64  `json:"orgId"`
	ParentId       int    `json:"parentId"`
	Path           string `json:"path"`
	PreviewCount   int    `json:"previewCount"`
	Size           int64  `json:"size"`
	SpaceId        int64  `json:"spaceId"`
	Status         int    `json:"status"`
	StoreSrc       string `json:"storeSrc"`
	Type           string `json:"type"`
	Version        string `json:"version"`
	WaitingForEdit bool   `json:"waitingForEdit"`
}

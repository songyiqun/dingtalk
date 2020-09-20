package request

type TextMessage struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type ImageMessage struct {
	MsgType string `json:"msgtype"`
	Image   struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

type VoiceMessage struct {
	MsgType string `json:"msgtype"`
	Voice   struct {
		MediaId  string `json:"media_id"`
		Duration string `json:"duration"`
	} `json:"voice"`
}

type FileMessage struct {
	MsgType string `json:"msgtype"`
	File    struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
}

type LinkMessage struct {
	MsgType string `json:"msgtype"`
	Link    struct {
		MessageUrl string `json:"messageUrl"`
		PicUrl     string `JSON:"picUrl"`
		Title      string `json:"title"`
	} `json:"link"`
}

type OaMessageForm struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OaMessage struct {
	MsgType string `json:"msgtype"`
	Oa      struct {
		MessageUrl string `json:"message_url"`
		Head       struct {
			Bgcolor string `json:"bgcolor"`
			Text    string `json:"text"`
		} `json:"head"`
		Body struct {
			Title string          `json:"title"`
			Form  []OaMessageForm `json:"form"`
		} `json:"body"`
		Rich struct {
			Num  string `json:"num"`
			Unit string `json:"unit"`
		} `json:"rich"`
		Content   string `json:"content"`
		Image     string `json:"image"`
		FileCount string `JSON:"file_count"`
		Author    string `json:"author"`
	} `json:"oa"`
}

type MarkdownMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
}

type ActionCardMessage_1 struct {
	MsgType    string                         `json:"msgtype"`
	ActionCard ActionCardMessage_1_ActionCard `json:"action_card"`
}

type ActionCardMessage_1_ActionCard struct {
	Title       string `json:"title"`
	Markdown    string `json:"markdown"`
	SingleTitle string `json:"single_title"`
	SingleUrl   string `json:"single_url"`
}

type ActionCardMessage_2_BtnJsonList struct {
	Title     string `json:"title"`
	ActionUrl string `json:"action_url"`
}

type ActionCardMessage_2 struct {
	MsgType    string `json:"msgtype"`
	ActionCard ActionCardMessage_2_ActionCard  `json:"action_card"`
}

type ActionCardMessage_2_ActionCard struct {
	Title          string                            `json:"title"`
	Markdown       string                            `json:"markdown"`
	BtnOrientation string                            `json:"btn_orientation"`
	BtnJsonList    []ActionCardMessage_2_BtnJsonList `json:"btn_json_list"`
}
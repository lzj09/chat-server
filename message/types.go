package message

const (
	// LoginMsgType 登录消息
	LoginMsgType int64 = 1

	// PersonalDialMsgType 个人对话消息
	PersonalDialMsgType int64 = 2

	// FeedbackMsgType 反馈消息
	FeedbackMsgType int64 = 20

	// ErrorStatus 错误状态
	ErrorStatus int64 = 500

	// SuccessStatus 正常状态
	SuccessStatus int64 = 200
)

// Msg 消息实体
type Msg struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	MsgType int64  `json:"msgType,omitempty"`
	FromID  string `json:"fromId,omitempty"`
	ToID    string `json:"toId,omitempty"`
	Status  int64  `json:"status,omitempty"`
}

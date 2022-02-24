package message

const (
	// 登录消息
	LoginMsgType int64 = 1

	// 反馈消息
	FeedbackMsgType int64 = 20

	// 错误状态
	ErrorStatus int64 = 500
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

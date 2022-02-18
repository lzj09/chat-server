package message

// Msg 消息实体
type Msg struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	MsgType int64  `json:"msgType,omitempty"`
	FromID  string `json:"fromId,omitempty"`
	ToID    string `json:"toId,omitempty"`
	Status  int64  `json:"status,omitempty"`
}

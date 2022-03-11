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

	// MessageDBName 消息库名称
	MessageDBName = "msg_db"

	// MessageTableName 消息表名称
	MessageTableName = "msg"
)

// Msg 消息实体
type Msg struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Content     string `json:"content,omitempty" bson:"content,omitempty"`
	ContentType int64  `json:"contentType,omitempty" bson:"contentType,omitempty"`
	MsgType     int64  `json:"msgType,omitempty" bson:"msgType,omitempty"`
	FromID      string `json:"fromId,omitempty" bson:"fromId,omitempty"`
	ToID        string `json:"toId,omitempty" bson:"toId,omitempty"`
	Status      int64  `json:"status,omitempty" bson:"status,omitempty"`
}

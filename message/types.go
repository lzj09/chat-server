package message

import "time"

const (
	// LogoutMsgType 退出消息
	LogoutMsgType int64 = -1

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

	// UnreadStatus 未读状态
	UnreadStatus int64 = 10

	// DBName 消息库名称
	DBName = "msg_db"

	// TableName 消息表名称
	TableName = "msg"
)

// Msg 消息实体
type Msg struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Content     string    `json:"content,omitempty" bson:"content,omitempty"`
	ContentType int64     `json:"contentType,omitempty" bson:"contentType,omitempty"`
	MsgType     int64     `json:"msgType,omitempty" bson:"msgType,omitempty"`
	FromID      string    `json:"fromId,omitempty" bson:"fromId,omitempty"`
	ToID        string    `json:"toId,omitempty" bson:"toId,omitempty"`
	Status      int64     `json:"status,omitempty" bson:"status,omitempty"`
	CreateTime  time.Time `json:"createTime,omitempty" bson:"createTime,omitempty"`
}

// Conversation 会话实体
type Conversation struct {
	ID       string    `json:"id,omitempty" db:"id"`
	UID      string    `json:"uid,omitempty" db:"uid"`
	ToID     string    `json:"toId,omitempty" db:"to_id"`
	ChatId   string    `json:"chatId,omitempty" db:"chat_id"`
	ChatType int64     `json:"chatType,omitempty" db:"chat_type"`
	Unread   int64     `json:"unread,omitempty" db:"unread"`
	IsDelete int64     `json:"isDelete,omitempty" db:"is_delete"`
	LastTime time.Time `json:"lastTime,omitempty" db:"last_time"`
}

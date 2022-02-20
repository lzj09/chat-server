package user

import "time"

// User 用户实体
type User struct {
	ID         string    `json:"id,omitempty"`
	UserName   string    `json:"userName,omitempty"`
	Password   string    `json:"password,omitempty"`
	NickName   string    `json:"nickName,omitempty"`
	Gender     int64     `json:"gender,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
}

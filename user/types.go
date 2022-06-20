package user

import (
	"time"
)

// User 用户实体
type User struct {
	ID         string    `json:"id,omitempty" db:"id"`
	Username   string    `json:"username,omitempty" db:"username"`
	Password   string    `json:"password,omitempty" db:"password"`
	NickName   string    `json:"nickName,omitempty" db:"nick_name"`
	Gender     int64     `json:"gender,omitempty" db:"gender"`
	Avatar     *string   `json:"avatar,omitempty" db:"avatar"`
	CreateTime time.Time `json:"createTime,omitempty" db:"create_time"`
}

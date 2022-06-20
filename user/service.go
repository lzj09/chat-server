package user

import (
	"errors"
	"github.com/lzj09/chat-server/persistent/mysql"
	"k8s.io/klog/v2"
)

// Service 用户信息服务接口
type Service interface {
	// Login 登录
	Login(username, password string) (*User, error)

	// Get 根据id获取
	Get(id string) (*User, error)
}

// DefaultUserService 默认用户服务实现
type DefaultUserService struct {
}

func (svc *DefaultUserService) Login(username, password string) (*User, error) {
	var user User
	err := mysql.MysqlClient.Get(&user, "select * from chat_user where username = ?", username)
	if err != nil {
		klog.Errorf("get user by username %v error: %v", username, err)
		return nil, err
	}

	// TODO 密码加密，后续完善
	if password != user.Password {
		return nil, errors.New("username or password error")
	}

	return &user, nil
}

func (svc *DefaultUserService) Get(id string) (*User, error) {
	var user User
	if err := mysql.MysqlClient.Get(&user, "select * from chat_user where id = ?", id); err != nil {
		klog.Errorf("get user by id error: %v", err)
		return nil, err
	}

	return &user, nil
}

package user

// Service 用户信息服务接口
type Service interface {
	// Login 登录
	Login(username, password string) (*User, error)
}

// DefaultUserService 默认用户服务实现
type DefaultUserService struct {
}

func (svc *DefaultUserService) Login(username, password string) (*User, error) {
	// TODO 模拟登录，后续完善
	return &User{
		ID:       "123",
		UserName: username,
		Password: password,
	}, nil
}

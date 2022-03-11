package message

// Service 消息服务接口
type Service interface {
	// Save 保存
	Save(msg *Msg) (*Msg, error)
}

// DefaultMessageService 默认消息服务实现
type DefaultMessageService struct {
}

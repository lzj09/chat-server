package message

import (
	"context"
	"github.com/lzj09/chat-server/persistent/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"k8s.io/klog/v2"
)

// Service 消息服务接口
type Service interface {
	// Save 保存
	Save(msg *Msg) (*Msg, error)
}

// DefaultMessageService 默认消息服务实现
type DefaultMessageService struct {
}

func (svc *DefaultMessageService) Save(msg *Msg) (*Msg, error) {
	// 指定要操作的数据集
	collection := mongo.MongoClient.Database(DBName).Collection(TableName)

	result, err := collection.InsertOne(context.Background(), msg)
	if err != nil {
		klog.Errorf("insert msg error: %v", err)
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)
	msg.ID = id.Hex()
	return msg, nil
}

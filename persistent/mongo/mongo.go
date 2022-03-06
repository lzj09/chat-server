package mongo

import (
	"context"
	"github.com/lzj09/chat-server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// Init 初始化mongo
func Init() {
	clientOptions := options.Client().SetHosts([]string{utils.GetEnv("MONGO_HOSTS", "127.0.0.1:27017")}).SetAuth(options.Credential{Username: utils.GetEnv("MONGO_USERNAME", "lzj"), Password: utils.GetEnv("MONGO_PASSWORD", "123")})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	MongoClient = client
}

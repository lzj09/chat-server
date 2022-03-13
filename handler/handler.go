package handler

import (
	"github.com/lzj09/chat-server/message"
	"github.com/lzj09/chat-server/persistent/mongo"
	"github.com/lzj09/chat-server/persistent/mysql"
	"github.com/lzj09/chat-server/user"
	"github.com/lzj09/chat-server/utils"
)

// Init 初始化结构类
func Init() {
	userService := &user.DefaultUserService{}
	utils.Register(userService)

	messageService := &message.DefaultMessageService{}
	utils.Register(messageService)

	mongo.Init()
	mysql.Init()
}

package main

import (
	"encoding/json"
	"github.com/lzj09/chat-server/message"
	"github.com/lzj09/chat-server/user"
	"k8s.io/klog/v2"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9156")
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)
	contents := map[string]string{
		"username": "lzj",
		"password": "123",
	}
	contentBytes, _ := json.Marshal(contents)
	login := message.Msg{
		MsgType: message.LoginMsgType,
		Content: string(contentBytes),
	}
	loginBytes, _ := json.Marshal(login)
	conn.Write(loginBytes)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			klog.Errorf("conn read error: %v", err)
			break
		}

		var msgObj message.Msg
		json.Unmarshal(buffer[:n], &msgObj)
		if msgObj.Status == message.SuccessStatus {
			// 登录成功，获取用户信息
			var userObj user.User
			json.Unmarshal([]byte(msgObj.Content), &userObj)

			toMsg := message.Msg{
				Content: "你好，测试，测试，测试，测试，测试，测试",
				MsgType: message.PersonalDialMsgType,
				FromID:  userObj.ID,
				ToID:    "msg-002",
			}

			toMsgBytes, _ := json.Marshal(toMsg)
			conn.Write(toMsgBytes)

			logoutMsg := message.Msg{
				MsgType: message.LogoutMsgType,
				FromID:  userObj.ID,
			}
			logoutMsgBytes, _ := json.Marshal(logoutMsg)
			conn.Write(logoutMsgBytes)

			break
		}
	}
	conn.Close()
}

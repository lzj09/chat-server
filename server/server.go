package server

import (
	"encoding/json"
	"fmt"
	"github.com/lzj09/chat-server/message"
	"github.com/lzj09/chat-server/user"
	"github.com/lzj09/chat-server/utils"
	"k8s.io/klog/v2"
	"net"
)

var connections = make(map[string]net.Conn)

// ChatServer chat服务
type ChatServer struct {
	IP   string
	Port int64
}

// Run 启动
func (cs *ChatServer) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", cs.IP, cs.Port))
	if err != nil {
		klog.Errorf("chat server run error: %v", err)
		panic(err)
	}

	klog.Infof("chat server start successful!")

	for {
		conn, err := listener.Accept()
		if err != nil {
			klog.Errorf("chat server accept client error: %v", err)
			continue
		}

		klog.Infof("chat server accept client: %v", conn)

		go readConn(conn)
	}
}

// readConn 读取信息
func readConn(conn net.Conn) {
	buffer := make([]byte, 1024)

	loopFlag := false
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			klog.Errorf("conn read error: %v", err)
			continue
		}

		// 反序列化消息
		var msgObj message.Msg
		msgBytes := buffer[0:n]
		if err := json.Unmarshal(msgBytes, &msgObj); err != nil {
			klog.Errorf("json unmarshal msg error: %v", err)

			bytes, err := feedbackMsg("data error", message.ErrorStatus)
			if err != nil {
				klog.Errorf("get feedback msg error: %v", err)
				continue
			}

			conn.Write(bytes)
			continue
		}

		switch msgObj.MsgType {
		case message.LoginMsgType:
			// 处理登录
			var contents map[string]string
			if err := json.Unmarshal([]byte(msgObj.Content), &contents); err != nil {
				klog.Errorf("json unmarshal msg content error: %v", err)

				bytes, err := feedbackMsg("login data error", message.ErrorStatus)
				if err != nil {
					klog.Errorf("get feedback msg error: %v", err)
					break
				}

				conn.Write(bytes)
				break
			}
			userService := utils.Obtain(new(user.DefaultUserService)).(*user.DefaultUserService)
			user, err := userService.Login(contents["username"], contents["password"])
			if err != nil {
				klog.Errorf("user login error: %v", err)

				bytes, err := feedbackMsg("username or password error", message.ErrorStatus)
				if err != nil {
					klog.Errorf("get feedback msg error: %v", err)
					break
				}

				conn.Write(bytes)
				break
			}

			content, err := json.Marshal(user)
			bytes, err := feedbackMsg(string(content), message.SuccessStatus)
			if err != nil {
				klog.Errorf("get feedback msg error: %v", err)
				break
			}

			// 保存连接
			connections[user.ID] = conn
			conn.Write(bytes)

		case message.PersonalDialMsgType:
			// 校验是否登录，暂时校验是否带上FromID
			if msgObj.FromID == "" {
				bytes, err := feedbackMsg("please login", message.ErrorStatus)
				if err != nil {
					klog.Errorf("get feedback msg error: %v", err)
					break
				}

				conn.Write(bytes)
				break
			}

			if msgObj.ToID == "" {
				bytes, err := feedbackMsg("data error", message.ErrorStatus)
				if err != nil {
					klog.Errorf("get feedback msg error: %v", err)
					break
				}

				conn.Write(bytes)
				break
			}

			targetConn, ok := connections[msgObj.ToID]
			if ok {
				// 对方在线，则转发消息
				targetConn.Write(msgBytes)
			}
			// 将消息持久化
			msgObj.Status = message.UnreadStatus
			messageService := utils.Obtain(new(message.DefaultMessageService)).(*message.DefaultMessageService)
			messageService.Save(&msgObj)

		case message.LogoutMsgType:
			// 校验是否登录，暂时校验是否带上FromID
			if msgObj.FromID == "" {
				bytes, err := feedbackMsg("please login", message.ErrorStatus)
				if err != nil {
					klog.Errorf("get feedback msg error: %v", err)
					break
				}

				conn.Write(bytes)
				break
			}

			// 清除连接
			delete(connections, msgObj.FromID)
			loopFlag = true
		}

		if loopFlag {
			break
		}
	}
	conn.Close()
}

// feedbackMsg 封装反馈消息
func feedbackMsg(content string, status int64) ([]byte, error) {
	msg := message.Msg{
		Content: content,
		MsgType: message.FeedbackMsgType,
		Status:  status,
	}

	// 序列化
	return json.Marshal(msg)
}

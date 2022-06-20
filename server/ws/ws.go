package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lzj09/chat-server/message"
	"github.com/lzj09/chat-server/utils"
	"k8s.io/klog/v2"
)

var connections = make(map[string]*websocket.Conn)
var upgrader = websocket.Upgrader{}

func Run(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		klog.Errorf("connection websocket error: %v", err)
		return
	}

	// 解析token获取当前用户id

	go readConn(conn)
}

func readConn(conn *websocket.Conn) {
	loopFlag := false

	for {
		n, msg, err := conn.ReadMessage()
		if err != nil {
			klog.Errorf("websocket conn read error: %v", err)
			continue
		}

		// 反序列化消息
		var msgObj message.Msg
		if err := json.Unmarshal(msg, &msgObj); err != nil {
			klog.Errorf("websocket json unmarshal msg error: %v", err)
			continue
		}

		switch msgObj.MsgType {
		case message.PersonalDialMsgType:
			if msgObj.ToID == "" {
				klog.Errorf("websocket data error")
				break
			}

			targetConn, ok := connections[msgObj.ToID]
			if ok {
				// 对方在线，则转发消息
				if err := targetConn.WriteMessage(n, msg); err != nil {
					klog.Errorf("websocket write message error: %v", err)
				}
			}

			// 将消息持久化
			msgObj.Status = message.UnreadStatus
			messageService := utils.Obtain(new(message.DefaultMessageService)).(*message.DefaultMessageService)
			messageService.Save(&msgObj)
		case message.LogoutMsgType:
			if msgObj.FromID == "" {
				klog.Errorf("websocket data error")
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

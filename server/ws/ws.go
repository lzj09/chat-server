package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lzj09/chat-server/message"
	"github.com/lzj09/chat-server/utils"
	"k8s.io/klog/v2"
	"net/http"
	"time"
)

var Void void

// WebsocketManager 初始化websocket管理器
var WebsocketManager = Manager{
	Conns:      make(map[string]*Client),
	Ids:        make(map[string]map[string]void),
	Register:   make(chan *Client, 128),
	UnRegister: make(chan *Client, 128),
}

var upgrader = websocket.Upgrader{

	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// GetConnId 获取连接的id标志
func GetConnId(id, sign string) string {
	return fmt.Sprintf("%v_%v", id, sign)
}

// Start websocket管理器开始运行
func (manager *Manager) Start() {
	klog.Infoln("websocket管理器启动...")
	for {
		select {
		case client := <-manager.Register:
			// 新连接的客户端
			manager.Lock.Lock()
			connId := GetConnId(client.ID, client.Sign)
			manager.Conns[connId] = client

			ids, ok := manager.Ids[client.ID]
			if !ok {
				ids = make(map[string]void)
			}
			ids[client.Sign] = Void
			manager.Ids[client.ID] = ids

			klog.Infof("客户端【%v】连接成功...", connId)
			manager.Lock.Unlock()
		case client := <-manager.UnRegister:
			// 客户端退出
			manager.Lock.Lock()
			connId := GetConnId(client.ID, client.Sign)

			if ids, ok := manager.Ids[client.ID]; ok {
				delete(ids, client.Sign)
				if len(ids) == 0 {
					delete(manager.Ids, client.ID)
				}
			}

			if _, ok := manager.Conns[connId]; ok {
				// 移除该客户端的连接
				delete(manager.Conns, connId)
				klog.Infof("客户端【%v】退出...", connId)
			}
			manager.Lock.Unlock()
		}
	}
}

// RegisterClient 连接客户端
func (manager *Manager) RegisterClient(client *Client) {
	manager.Register <- client
}

// UnRegisterClient 退出客户端
func (manager *Manager) UnRegisterClient(client *Client) {
	manager.UnRegister <- client
}

func (client *Client) Monitor() {
	defer func() {
		if err := client.Conn.Close(); err != nil {
			klog.Errorf("客户端【%v】断开失败：%v", GetConnId(client.ID, client.Sign), err)
		}

		WebsocketManager.UnRegisterClient(client)
	}()

	loopFlag := false
	for {
		n, msg, err := client.Conn.ReadMessage()
		if n == websocket.CloseMessage {
			klog.Errorf("客户端【%v】读取监控数据失败: %v", GetConnId(client.ID, client.Sign), err)
			break
		}
		if err != nil {
			klog.Errorf("websocket conn read error: %v", err)
			break
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

			if signs, ok := WebsocketManager.Ids[msgObj.ToID]; ok {
				// 对方在线，则转发消息
				for sign, _ := range signs {
					connId := GetConnId(msgObj.ToID, sign)
					if targetClient, ok := WebsocketManager.Conns[connId]; ok {
						if err := targetClient.Conn.WriteMessage(n, msg); err != nil {
							klog.Errorf("websocket write message error: %v", err)
						}
					}
				}
			}

			// 将消息持久化
			msgObj.Status = message.UnreadStatus
			msgObj.CreateTime = time.Now()
			messageService := utils.Obtain(new(message.DefaultMessageService)).(*message.DefaultMessageService)
			messageService.Save(&msgObj)
		case message.LogoutMsgType:
			if msgObj.FromID == "" {
				klog.Errorf("websocket data error")
				break
			}

			loopFlag = true
		}

		if loopFlag {
			break
		}
	}
}

func Run(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		klog.Errorf("connection websocket error: %v", err)
		return
	}

	// 解析token获取当前用户id
	id := utils.UserInfo(c)
	sign := c.Request.FormValue("sign")
	client := &Client{
		ID:   id,
		Conn: conn,
		Sign: sign,
	}
	WebsocketManager.RegisterClient(client)

	go client.Monitor()
}

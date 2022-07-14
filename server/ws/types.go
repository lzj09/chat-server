package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

// Client websocket连接客户端
type Client struct {
	ID   string
	Sign string
	Conn *websocket.Conn
}

// void 定义空类型
type void struct{}

// Manager websocket连接管理器
type Manager struct {
	Conns      map[string]*Client
	Ids        map[string]map[string]void
	Lock       sync.Mutex
	Register   chan *Client
	UnRegister chan *Client
}

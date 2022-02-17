package server

import (
	"fmt"
	"k8s.io/klog/v2"
	"net"
)

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

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			klog.Errorf("conn read error: %v", err)
			continue
		}

		msg := string(buffer[0:n])

		if msg == "bye" {
			conn.Write([]byte(msg))
			break
		}

		conn.Write([]byte(fmt.Sprintf("已读: %v", msg)))
	}
	conn.Close()
}

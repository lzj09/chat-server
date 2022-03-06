package main

import (
	"encoding/json"
	"fmt"
	"github.com/lzj09/chat-server/message"
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

		fmt.Println(string(buffer[:n]))
	}
	conn.Close()
}

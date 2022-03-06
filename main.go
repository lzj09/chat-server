package main

import (
	"flag"
	"github.com/lzj09/chat-server/handler"
	"github.com/lzj09/chat-server/server"
	"github.com/lzj09/chat-server/utils"
	"k8s.io/klog/v2"
	"strconv"
)

func main() {
	klog.InitFlags(flag.CommandLine)
	defer klog.Flush()

	handler.Init()

	port, err := strconv.ParseInt(utils.GetEnv("SERVER_PORT", "9156"), 10, 64)
	if err != nil {
		klog.Errorf("server port parse error: %v", err)
		panic(err)
	}
	svr := server.ChatServer{
		IP:   utils.GetEnv("SERVER_IP", "127.0.0.1"),
		Port: port,
	}
	svr.Run()
}

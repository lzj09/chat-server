package main

import (
	"flag"
	"fmt"
	"github.com/lzj09/chat-server/controller"
	"github.com/lzj09/chat-server/handler"
	"github.com/lzj09/chat-server/server"
	"github.com/lzj09/chat-server/utils"
	"k8s.io/klog/v2"
	"net/http"
	"strconv"
	"time"
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
	go svr.Run()

	errChan := make(chan error)

	go func() {
		errChan <- newServer().ListenAndServe()
	}()

	klog.Fatal(<-errChan)
}

func newServer() *http.Server {
	port := utils.GetEnv("WEB_SERVER_PORT", "8080")
	return &http.Server{
		Addr:              fmt.Sprintf(":%v", port),      //Addr可选地指定服务器要监听的TCP地址，形式为“host:port”。如果为空，则使用“:http”(端口80)
		Handler:           controller.NewServerHandler(), //**调用的处理程序**
		TLSConfig:         nil,                           //可选地为ServeTLS和ListenAndServeTLS提供使用的TLS配置
		ReadTimeout:       600 * time.Second,             //读取整个请求(包括请求体)的最大持续时间。
		ReadHeaderTimeout: 0,                             //ReadHeaderTimeout是允许读取请求头的时间。如果ReadHeaderTimeout为0，则使用ReadTimeout的值。如果两者都为零，则不存在超时。
		WriteTimeout:      600 * time.Second,             //响应写超时之前的最大持续时间
		IdleTimeout:       0,                             //启用keep-alive时等待下一个请求的最大时间。如果IdleTimeout为0，则使用ReadTimeout的值。如果两者都为零，则不存在超时。
		MaxHeaderBytes:    1 << 20,                       //控制服务器解析请求头的键和值(包括请求行)时读取的最大字节数,1 << 20 十进制的值为1048576
		TLSNextProto:      nil,                           //
		ConnState:         nil,                           //一个可选的回调函数，当客户端连接改变状态时被调用
		ErrorLog:          nil,                           //指定一个可选的日志记录器，用于接收连接错误、处理程序的意外行为和底层文件系统错误。如果为nil，日志记录是通过日志包的标准日志记录程序来完成的
		BaseContext:       nil,                           //可选地指定一个函数，该函数返回此服务器上传入请求的基本上下文,如果BaseContext为nil，默认值为context.Background()
		ConnContext:       nil,
	}
}

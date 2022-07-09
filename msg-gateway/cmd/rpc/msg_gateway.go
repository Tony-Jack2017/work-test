package main

import (
	"fmt"
	wsconfig "work-test/msg-gateway/cmd/rpc/internal/websocket/config"
	"work-test/msg-gateway/cmd/rpc/internal/websocket/handler"
	"work-test/msg-gateway/cmd/rpc/internal/websocket/servectx"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"time"
)

var wsConfigFile = flag.String("w", "etc/msg_gateway_ws.yaml", "ws config file")
func ws() {
	var wsConfig wsconfig.Config
	conf.MustLoad(*wsConfigFile, &wsConfig)
	serveCtx := servectx.NewServiceContext(wsConfig)
	server := rest.MustNewServer(wsConfig.RestConf)
	defer server.Stop()

	handler.RegisterHandler(server, serveCtx)
	fmt.Println("Starting websocket server at")
	server.Start()
}

func main() {
	ws()
	time.Sleep(time.Second)
}

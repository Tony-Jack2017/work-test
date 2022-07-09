package handler

import (
	"work-test/msg-gateway/cmd/rpc/internal/websocket/logic"
	"work-test/msg-gateway/cmd/rpc/internal/websocket/servectx"
	"context"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterHandler(server *rest.Server, serveCtx *servectx.ServiceContext) {
	logic.NewMsgGatewayLogic(context.Background(), serveCtx)
	server.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/login",
			Handler: MsgGatewayHandler(serveCtx),
		},
	})
}

package handler

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"work-test/msg-gateway/cmd/rpc/internal/websocket/logic"
	"work-test/msg-gateway/cmd/rpc/internal/websocket/servectx"
	"work-test/msg-gateway/cmd/rpc/internal/websocket/types"
)

func MsgGatewayHandler(serveCtx *servectx.ServiceContext) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req types.Request
		if err := httpx.Parse(request, &req); err != nil {
			httpx.Error(writer, err)
			return
		}
		ws := logic.NewMsgGatewayLogic(context.Background(), serveCtx)

		logx.Info("resp uid is", req.Uid)
		err := ws.WsUpgrade(req.Uid, writer, request, nil)
		if err != nil {
			logx.WithContext(request.Context()).Errorf("ws.WsUpgrade error: %s", err)
			return
		}
		httpx.WriteJson(writer, http.StatusOK, "ws success")

	}
}

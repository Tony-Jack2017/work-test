package logic

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
	"work-test/msg-gateway/cmd/rpc/internal/websocket/servectx"
)

type UserConn struct {
	*websocket.Conn
	w *sync.Mutex
}

type MsgGatewayLogic struct {
	ctx          context.Context
	Svc          *servectx.ServiceContext
	wxMaxConnNum int
	wsUpGrader   *websocket.Upgrader
	WsUserToConn map[string]*UserConn `json:"wsUserToConn"`
}

var MsgGwLogic *MsgGatewayLogic

func NewMsgGatewayLogic(ctx context.Context, serviceContext *servectx.ServiceContext) *MsgGatewayLogic {
	if MsgGwLogic != nil {
		return MsgGwLogic
	}

	ws := &MsgGatewayLogic{
		ctx: ctx,
		Svc: serviceContext,
	}

	//ws.wxMaxConnNum = config.Conf.Websocket.MaxConnNum
	ws.wxMaxConnNum = 10000
	ws.WsUserToConn = make(map[string]*UserConn)
	ws.wsUpGrader = &websocket.Upgrader{
		//HandshakeTimeout: time.Duration(config.Conf.Websocket.TimeOut),
		HandshakeTimeout: time.Duration(10 * time.Second),
		//ReadBufferSize: config.Conf.Websocket.ReadBufferSize,
		ReadBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	MsgGwLogic = ws
	return MsgGwLogic
}

// WsUpgrade
// http upgrade to websocket
func (mgLogic *MsgGatewayLogic) WsUpgrade(uid string, w http.ResponseWriter, r *http.Request, header http.Header) error {
	conn, err := mgLogic.wsUpGrader.Upgrade(w, r, header)
	if err != nil {
		return err
	}
	var wl = new(sync.Mutex)
	newConn := &UserConn{conn, wl}
	err = mgLogic.addUserConn(uid, newConn)
	if err != nil {
		fmt.Println("addUserConn err: ", err)
		return err
	}
	mgLogic.readMsg(newConn, uid)
	return nil
}

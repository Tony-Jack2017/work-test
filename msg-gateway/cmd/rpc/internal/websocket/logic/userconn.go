package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

func (mgLogic *MsgGatewayLogic) addUserConn(uid string, conn *UserConn) error {
	conn.w.Lock()
	defer conn.w.Unlock()

	if oldConn, ok := mgLogic.WsUserToConn[uid]; ok {
		mgLogic.WsUserToConn[uid] = oldConn
	} else {
		mgLogic.WsUserToConn[uid] = conn
	}

	return nil
}

func (mgLogic *MsgGatewayLogic) delUserConn(ctx context.Context, conn *UserConn) {
	conn.w.Lock()
	defer conn.w.Unlock()
	var uid string
	if _, ok := mgLogic.WsUserToConn[uid]; ok {
		delete(mgLogic.WsUserToConn, uid)
	}

	err := conn.Close()
	if err != nil {
		logx.WithContext(mgLogic.ctx).Error("close conn err", "", "uid", uid)
	}
}

package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	routerType "work-test/msg-gateway/cmd/rpc/internal/websocket/types"
)

/*******		read and parse message  		********/

func (mgLogic *MsgGatewayLogic) readMsg(conn *UserConn, uid string) {
	messageType, msg, err := conn.ReadMessage()
	if messageType == websocket.PongMessage {
		mgLogic.sendMsg(conn, &routerType.RespBody{
			Event: "ping",
			Code:  0,
			Msg:   "",
			Data:  "",
		})
	}

	if err != nil {
		logx.Error("WS ReadMsg error ", " userIP ", conn.RemoteAddr().String(), " userUid ", uid, " error ", err.Error())
		mgLogic.delUserConn(context.Background(), conn)
		return
	}
	mgLogic.msgParse(conn, msg, uid)
}

func (mgLogic *MsgGatewayLogic) msgParse(conn *UserConn, binaryMsg []byte, uid string) error {
	msg := &routerType.ReqParams{}
	err := json.Unmarshal(binaryMsg, msg)
	if err != nil {
		logx.Error("parse error", msg)
		err = conn.Close()
		if err != nil {
			fmt.Println("ws conn close error", err.Error())
		}
		return err
	}
	switch msg.ReqId {
	case 1000:
		mgLogic.getSeqReq(conn, msg, uid)
	default:
	}
	return nil
}

/********	write and send message	********/

func (mgLogic *MsgGatewayLogic) writeMsg(conn *UserConn, messageType int, msg []byte) error {
	conn.w.Lock()
	defer conn.w.Unlock()
	err := conn.WriteMessage(messageType, msg)
	if err != nil {
		logx.Error(err)
	} else {
		logx.Info("send success")
	}
	return err
}

func (mgLogic *MsgGatewayLogic) sendMsg(conn *UserConn, resp *routerType.RespBody) {
	data, err := json.Marshal(resp)
	if err != nil {
		logx.Error(resp.Event, " ", resp.Code, " ", resp.Msg)
	}
	err = mgLogic.writeMsg(conn, websocket.TextMessage, data)
	if err != nil {
		logx.Error("Write Message Error")
	}
}


/********	the handler of message type	********/

func (mgLogic *MsgGatewayLogic) getSeqReq(conn *UserConn, reqBody *routerType.ReqParams, uid string) {
	mgLogic.getSeqResp(conn, reqBody, "test")
}

func (mgLogic *MsgGatewayLogic) getSeqResp(conn *UserConn, reqBody *routerType.ReqParams, resp string) error {
	logx.Info(resp)
	msgReply := &routerType.RespBody{
		Event:   reqBody.ReqFunction,
		Code:    0,
		Msg:     "error",
		Success: false,
		Data:    resp,
	}
	mgLogic.sendMsg(conn, msgReply)
	return nil
}

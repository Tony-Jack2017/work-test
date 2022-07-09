package types

/*****		ws build and request		*****/

type Request struct {
	Uid      string `form:"uid"`
}

type Response struct {
	Uid     string `json:"uid"`
}

type ReqParams struct {
	ReqId       int    `json:"reqId"`
	SendId      string `json:"sendId"`
	SenderType  string `json:"senderType"`
	ReqFunction string `json:"reqFunction"`
	Data        string `json:"data"`
}

type RespBody struct {
	Event   string `json:"event"`
	Code    uint32 `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

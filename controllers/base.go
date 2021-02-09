package controllers

type ApiJson struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type BackendListRq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BackendListRsp struct {
	Count int64       `json:"count"`
	List  interface{} `json:"list"`
}

func ApiResource(code int, objects interface{}, msg string) (apijson *ApiJson) {
	apijson = &ApiJson{Code: code, Data: objects, Msg: msg}
	return
}

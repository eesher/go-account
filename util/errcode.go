package util

import (
	"log"
	"runtime"
)

type Result struct {
	Errcode int    `json:"errcode"`
	Msg     string `json:"msg"`
}

type Errcode struct {
	OK            int
	DB_ERROR      int
	INVA_PARAM    int
	INVA_PLATFORM int
	LOGIN_ERROR   int
	NO_USER       int
	INVA_TOKEN    int
}

var error_info = map[int]string{}
var ERRCODE = Errcode{}

func ErrcodeInit() {
	ERRCODE.OK = 0
	error_info[ERRCODE.OK] = "ok"

	ERRCODE.DB_ERROR = 30001
	error_info[ERRCODE.DB_ERROR] = "db error"

	ERRCODE.NO_USER = 30002
	error_info[ERRCODE.NO_USER] = "can not found user"

	ERRCODE.INVA_PARAM = 31001
	error_info[ERRCODE.INVA_PARAM] = "invalid params"

	ERRCODE.INVA_PLATFORM = 31002
	error_info[ERRCODE.INVA_PLATFORM] = "invalid platform"

	ERRCODE.INVA_TOKEN = 31003
	error_info[ERRCODE.INVA_TOKEN] = "invalid token"
	//	error_info[31003] = "invalid passwd"
	//	error_info[31004] = "invalid device_id"
	ERRCODE.LOGIN_ERROR = 32001
	error_info[ERRCODE.LOGIN_ERROR] = "login error"
	//	error_info[32002] = "registered user"
	//	error_info[33001] = "signin error"
	//	error_info[33003] = "passwd not match"
}

func GetResult(errcode int) Result {
	return Result{errcode, error_info[errcode]}
}

func CheckErr(err error) bool {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] [%s:%d] %v", fn, line, err)
		return false
	}
	return true
}

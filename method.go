package main

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"time"
)

var ERRCODE map[int]string
var platform_method = make(map[string]func(map[string]string) (map[string]string, int))

func ErrcodeInit() {
	ERRCODE = make(map[int]string)
	ERRCODE[0] = "ok"
	ERRCODE[30001] = "db error"
	ERRCODE[31001] = "invalid params"
	ERRCODE[31002] = "invalid user"
	ERRCODE[31003] = "invalid passwd"
	ERRCODE[31004] = "invalid device_id"
	ERRCODE[32001] = "signup error"
	ERRCODE[32002] = "registered user"
	ERRCODE[33001] = "signin error"
	ERRCODE[33002] = "can not found user"
	ERRCODE[33003] = "passwd not match"
	ERRCODE[33004] = "invalid token"

	platform_method["guest"] = GuestLogin
}

var dur_time int64 = time.Date(2016, time.January, 0, 0, 0, 0, 0, time.UTC).Unix()

type Result struct {
	Errcode int    `json:"errcode"`
	Msg     string `json:"msg"`
}

type LoginRet struct {
	Result
	Openid           string `json:"openid"`
	Uid              string `json:"uid"`
	ThirdOpenid      string `json:"third_openid"`
	ThirdAccessToken string `json:"third_access_token"`
	AccessToken      string `json:"access_token"`
	//	AdminLevel  int    `json:"admin_level"`
}

type UserInfo struct {
	Uid         string `json:"uid"`
	PortraitUrl string `json:"portrait_url"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	AdminLevel  int    `json:"admin_level"`
	CreateTime  string `json:"create_time"`
}

type AuthRet struct {
	Result
	UserInfo
}

func GetResult(errcode int) Result {
	return Result{errcode, ERRCODE[errcode]}
}

func MakeUser(msg map[string]string) string {
	uid := time.Now().Unix() - dur_time
	uid_str := strconv.FormatInt(uid, 16)
	msg["uid"] = uid_str
	msg["third_openid"] = msg["device_id"]
	msg["name"] = "default name"

	//var json_data []byte
	if NewUserMap(msg) != 0 {
		return ""
	}
	if NewUser(msg) != 0 {
		return ""
	}

	return uid_str
}

func GuestLogin(msg map[string]string) (map[string]string, int) {
	ret := make(map[string]string)
	uid, errcode := GetUid("guest", msg["device_id"])
	if errcode != 0 {
		return ret, errcode
	}
	if len(uid) == 0 {
		uid = MakeUser(msg)
	}
	ret["uid"] = uid
	ret["third_openid"] = msg["device_id"]
	return ret, 0
}

func Login(msg map[string]string, json_data *[]byte) {
	var ret map[string]string
	var errcode int
	if method, ok := platform_method[msg["platform"]]; ok {
		ret, errcode = method(msg)
		if errcode != 0 {
			*json_data, _ = json.Marshal(GetResult(errcode))
			return
		}
	}

	ret["device_id"] = msg["device_id"]
	access_token := base64.URLEncoding.EncodeToString(GenerateToken(ret["uid"]))
	*json_data, _ = json.Marshal(LoginRet{GetResult(0), ret["openid"], ret["uid"], ret["third_openid"], ret["third_access_token"], access_token})
}

func Auth(msg map[string]string, json_data *[]byte) {
	if !AuthToken(msg["uid"], msg["access_token"]) {
		*json_data, _ = json.Marshal(GetResult(33003))
		return
	}
	user_info := UserInfo{}
	if errcode := GetUserInfo(msg["uid"], &user_info); errcode != 0 {
		*json_data, _ = json.Marshal(GetResult(errcode))
		return
	}

	*json_data, _ = json.Marshal(AuthRet{GetResult(0), user_info})
}

package main

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var ERRCODE map[int]string

func ErrcodeInit() {
	ERRCODE = make(map[int]string)
	ERRCODE[0] = "ok"
	ERRCODE[30001] = "db error"
	ERRCODE[31001] = "invalid params"
	ERRCODE[31002] = "invalid platform"
	ERRCODE[31003] = "invalid passwd"
	ERRCODE[31004] = "invalid device_id"
	ERRCODE[32001] = "login error"
	ERRCODE[32002] = "registered user"
	ERRCODE[33001] = "signin error"
	ERRCODE[33002] = "can not found user"
	ERRCODE[33003] = "passwd not match"
	ERRCODE[33004] = "invalid token"
}

var dur_time int64 = time.Date(2016, time.January, 0, 0, 0, 0, 0, time.UTC).Unix()

type Result struct {
	Errcode int    `json:"errcode"`
	Msg     string `json:"msg"`
}

type LoginRet struct {
	Result
	LoginInfo
}

type LoginInfo struct {
	Openid           string `json:"openid"`
	Uid              string `json:"uid"`
	ThirdOpenid      string `json:"third_openid"`
	ThirdAccessToken string `json:"third_access_token"`
	AccessToken      string `json:"access_token"`
	DeviceID         string `json:"device_id"`
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

func (login_info *LoginInfo) MakeUser(msg map[string]string) int {
	uid := time.Now().Unix() - dur_time
	uid_str := strconv.FormatInt(uid, 16)

	msg["uid"] = uid_str
	msg["third_openid"] = msg["device_id"]
	msg["name"] = "default name"

	//var json_data []byte
	errcode := 0
	if errcode = NewUserMap(msg); errcode != 0 {
		return errcode
	}
	if errcode = NewUser(msg); errcode != 0 {
		return errcode
	}

	login_info.Uid = uid_str
	return errcode
}

func (login_info *LoginInfo) GuestLogin(msg map[string]string) int {
	var errcode int
	login_info.Uid, errcode = GetUid("guest", msg["device_id"])
	if errcode != 0 {
		return errcode
	}
	if len(login_info.Uid) == 0 {
		if errcode = login_info.MakeUser(msg); errcode != 0 {
			return errcode
		}
	}
	login_info.ThirdOpenid = msg["device_id"]
	return 0
}

func Login(msg map[string]string, json_data *[]byte) {
	//var errcode int
	login_info := LoginInfo{}
	method := reflect.ValueOf(&login_info).MethodByName(strings.Title(msg["platform"]) + "Login")
	if !method.IsValid() {
		*json_data, _ = json.Marshal(GetResult(31002))
		return
	}
	ret := method.Call([]reflect.Value{reflect.ValueOf(msg)})
	if errcode := ret[0].Int(); errcode != 0 {
		*json_data, _ = json.Marshal(GetResult(int(errcode)))
		return
	}

	login_info.DeviceID = msg["device_id"]
	login_info.AccessToken = base64.URLEncoding.EncodeToString(GenerateToken(login_info.Uid))
	*json_data, _ = json.Marshal(LoginRet{GetResult(0), login_info})
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

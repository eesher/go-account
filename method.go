package main

import (
	"encoding/json"
	"go-account/util"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var dur_time = time.Date(2016, time.January, 0, 0, 0, 0, 0, time.UTC).UnixNano()

type LoginRet struct {
	util.Result
	LoginInfo
}

type LoginInfo struct {
	Openid           string `json:"openid"`
	Uid              string `json:"uid"`
	ThirdOpenid      string `json:"third_openid"`
	ThirdAccessToken string `json:"third_access_token"`
	AccessToken      string `json:"access_token"`
	ExpireTime       int64  `json:"expire_time"`
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
	util.Result
	UserInfo
}

func (login_info *LoginInfo) MakeUser(msg map[string]string) int {
	uid := time.Now().UnixNano() - dur_time
	uid_str := strconv.FormatInt(uid, 36)

	msg["uid"] = uid_str
	msg["third_openid"] = msg["device_id"]
	msg["name"] = "default name"

	//var json_data []byte
	errcode := 0
	if errcode = util.NewUserMap(msg); errcode != util.ERRCODE.OK {
		return errcode
	}
	if errcode = util.NewUser(msg); errcode != util.ERRCODE.OK {
		return errcode
	}

	login_info.Uid = uid_str
	return errcode
}

func (login_info *LoginInfo) GuestLogin(msg map[string]string) int {
	var errcode int
	login_info.Uid, errcode = util.GetUid("guest", msg["device_id"])
	if errcode != util.ERRCODE.OK {
		return errcode
	}
	if len(login_info.Uid) == 0 {
		errcode = login_info.MakeUser(msg)
	}
	return errcode
}

func Login(msg map[string]string, json_data *[]byte) {
	//var errcode int
	login_info := LoginInfo{}
	method := reflect.ValueOf(&login_info).MethodByName(strings.Title(msg["platform"]) + "Login")
	if !method.IsValid() {
		*json_data, _ = json.Marshal(util.GetResult(util.ERRCODE.INVA_PLATFORM))
		return
	}
	ret := method.Call([]reflect.Value{reflect.ValueOf(msg)})
	if errcode := int(ret[0].Int()); errcode != util.ERRCODE.OK {
		*json_data, _ = json.Marshal(util.GetResult(errcode))
		return
	}

	login_info.DeviceID = msg["device_id"]
	login_info.ThirdOpenid = msg["device_id"]
	login_info.ExpireTime = util.TOKEN_DATA.ExpireTime
	login_info.AccessToken = util.GenerateToken(login_info.Uid)
	*json_data, _ = json.Marshal(LoginRet{util.GetResult(util.ERRCODE.OK), login_info})
}

func Auth(msg map[string]string, json_data *[]byte) {
	if !util.AuthToken(msg["uid"], msg["access_token"]) {
		*json_data, _ = json.Marshal(util.GetResult(util.ERRCODE.INVA_TOKEN))
		return
	}
	user_info := UserInfo{}
	var db_info map[string]interface{}
	var errcode int
	if db_info, errcode = util.GetUserInfo(msg["uid"]); errcode != util.ERRCODE.OK {
		*json_data, _ = json.Marshal(util.GetResult(errcode))
		return
	}
	user_info.Uid = string(db_info["uid"].([]byte))
	user_info.PortraitUrl = string(db_info["portrait_url"].([]byte))
	user_info.Name = string(db_info["name"].([]byte))
	user_info.Gender = int(db_info["gender"].(int64))
	user_info.Email = string(db_info["email"].([]byte))
	user_info.Phone = string(db_info["phone"].([]byte))
	user_info.AdminLevel = int(db_info["admin_level"].(int64))
	user_info.CreateTime = string(db_info["create_time"].([]byte))

	*json_data, _ = json.Marshal(AuthRet{util.GetResult(util.ERRCODE.OK), user_info})
}

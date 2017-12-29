package main

import (
	"bytes"
	//"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	//"fmt"
	"strconv"
	"time"
)

type Token struct {
	salt        string
	expire_time int64
}

var token_data Token

func TokenInit(salt string, expire_time int64) {
	token_data = Token{salt: salt, expire_time: expire_time}
}

func GenerateToken(data string) []byte {
	var buff bytes.Buffer
	buff.WriteString(data)
	buff.WriteString(token_data.salt)

	now := time.Now()
	epoch := now.Unix() / token_data.expire_time
	tmp := strconv.FormatInt(epoch, 10)
	buff.WriteString(tmp)

	hash := sha512.New384()
	hash.Write(buff.Bytes())
	return hash.Sum(nil)
}

/*
func CheckPasswd(user string, passwd string, device_id string, is_guest bool) (string, int) {
	uid, origin_passwd, origin_device, admin_level := GetUserInfo("mu77", user)
	if len(uid) != 0 {
		if is_guest {
			if origin_passwd != passwd {
				return "-1", 0
			}
		} else {
			var buff bytes.Buffer
			buff.WriteString(origin_passwd)
			buff.WriteString(device_id)

			hash := md5.New()
			hash.Write(buff.Bytes())
			if fmt.Sprintf("%x", hash.Sum(nil)) != passwd {
				return "-1", 0
			}

			if device_id != origin_device {
				UpdateDevice(uid, device_id)
			}
		}
	}

	return uid, admin_level
}
*/

func AuthToken(data string, token string) bool {
	tmp, _ := base64.URLEncoding.DecodeString(token)
	return bytes.Equal(GenerateToken(data), tmp)
}

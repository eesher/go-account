package util

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"strconv"
	"time"
)

type Token struct {
	Salt       string
	ExpireTime int64
}

var TOKEN_DATA = Token{}

func TokenInit(config TokenConfig) {
	TOKEN_DATA.Salt = config.Salt
	TOKEN_DATA.ExpireTime = config.Expire
}

func sum(data string) []byte {
	var buff bytes.Buffer
	buff.WriteString(data)
	buff.WriteString(TOKEN_DATA.Salt)

	now := time.Now()
	epoch := now.Unix() / TOKEN_DATA.ExpireTime
	tmp := strconv.FormatInt(epoch, 10)
	buff.WriteString(tmp)

	hash := sha1.New()
	hash.Write(buff.Bytes())
	return hash.Sum(nil)
}

func GenerateToken(data string) string {
	return base64.URLEncoding.EncodeToString(sum(data))
}

func AuthToken(data string, token string) bool {
	tmp, _ := base64.URLEncoding.DecodeString(token)
	return bytes.Equal(sum(data), tmp)
}

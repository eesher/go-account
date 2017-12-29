package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func CheckErr(err error) bool {
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func MysqlInit() {
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/platform?charset=utf8mb4")
	CheckErr(err)
}

func GetUid(platform string, third_uid string) (string, int) {
	rows, err := db.Query("select uid from third_platform_user_map where platform=? and third_openid=?", platform, third_uid)
	if !CheckErr(err) {
		return "", 30001
	}
	defer rows.Close()

	var uid string

	if rows.Next() && !CheckErr(rows.Scan(&uid)) {
		return "", 30001
	}

	return uid, 0
}

func NewUserMap(info map[string]string) int {
	stmt, err := db.Prepare("INSERT third_platform_user_map SET uid=?,platform=?, channel=?, device_id=?, third_openid=?,third_token=?,third_unionid=?")
	if !CheckErr(err) {
		return 30001
	}

	_, err = stmt.Exec(info["uid"], info["platform"], info["channel"], info["device_id"], info["third_openid"], info["third_token"], info["third_unionid"])
	if !CheckErr(err) {
		return 30001
	}
	return 0
}

func NewUser(info map[string]string) int {
	stmt, err := db.Prepare("INSERT users SET uid=?,portrait_url=?,name=?,gender=?,email=?,phone=?,admin_level=?,create_time=now()")
	if !CheckErr(err) {
		return 30001
	}

	_, err = stmt.Exec(info["uid"], info["portrait_url"], info["name"], info["gender"], info["email"], info["phone"], info["admin_level"])
	if !CheckErr(err) {
		return 30001
	}
	return 0
}

func GetUserInfo(uid string, user_info *UserInfo) int {
	rows, err := db.Query("select * from users where uid=?", uid)
	if !CheckErr(err) {
		return 30001
	}

	defer rows.Close()

	if rows.Next() && !CheckErr(rows.Scan(&user_info.Uid, &user_info.PortraitUrl, &user_info.Name, &user_info.Gender, &user_info.Email, &user_info.Phone, &user_info.AdminLevel, &user_info.CreateTime)) {
		return 30001
	}

	return 0
}

/*
func CheckUser(platform string, user string, device_id string) string {
	rows, err := db.Query("select uid from users where platform=? and (device_id=? or user=?)", platform, device_id, user)
	if !CheckErr(err) {
		return "-1"
	}

	defer rows.Close()
	var uid string
	rows.Next()
	rows.Scan(&uid)

	return uid
}

func UpdateDevice(uid string, device_id string) {
	db.Query("update users set device_id = ? where uid = ?", device_id, uid)
}

func NewUser(uid string, platform string, channel string, user string, passwd string, device_id string) int {
	has_user := CheckUser(platform, user, device_id)

	if len(has_user) != 0 {
		return 32002
	}

	stmt, err := db.Prepare("INSERT users SET uid=?,platform=?,channel=?,user=?,passwd=?,device_id=?,date=now()")
	if !CheckErr(err) {
		return 32001
	}

	_, err = stmt.Exec(uid, platform, channel, user, passwd, device_id)
	if !CheckErr(err) {
		return 32001
	}

	return 0
}

func NewGuest(uid string, device_id string) bool {
	stmt, err := db.Prepare("INSERT users SET uid=?,platform='mu77',channel='guest',device_id=?,create_time=now()")
	if !CheckErr(err) {
		return false
	}

	_, err = stmt.Exec(uid, device_id)
	if !CheckErr(err) {
		return false
	}
	return true
}
*/

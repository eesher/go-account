package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MysqlInit(config DbConfig) {
	var err error
	db, err = sql.Open(config.Driver, config.Connection)
	CheckErr(err)
}

func ScanSingleRowToMap(rows *sql.Rows) (map[string]interface{}, error) {
	cols, _ := rows.Columns()
	// Create a slice of interface{}'s to represent each column,
	// and a second slice to contain pointers to each item in the columns slice.
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i, _ := range columns {
		columnPointers[i] = &columns[i]
	}

	// Scan the result into the column pointers...
	if err := rows.Scan(columnPointers...); err != nil {
		return nil, err
	}

	// Create our map, and retrieve the value for each column from the pointers slice,
	// storing it in the map with the name of the column as the key.
	m := make(map[string]interface{})
	for i, colName := range cols {
		val := columnPointers[i].(*interface{})
		m[colName] = *val
	}
	return m, nil
}

func GetUid(platform string, third_uid string) (string, int) {
	rows, err := db.Query("select uid from third_platform_user_map where platform=? and third_openid=?", platform, third_uid)
	if !CheckErr(err) {
		return "", ERRCODE.DB_ERROR
	}
	defer rows.Close()

	var uid string

	if rows.Next() && !CheckErr(rows.Scan(&uid)) {
		return "", ERRCODE.DB_ERROR
	}

	return uid, ERRCODE.OK
}

func NewUserMap(info map[string]string) int {
	stmt, err := db.Prepare("INSERT third_platform_user_map SET uid=?,platform=?, channel=?, device_id=?, third_openid=?,third_token=?,third_unionid=?")
	if !CheckErr(err) {
		return ERRCODE.DB_ERROR
	}

	_, err = stmt.Exec(info["uid"], info["platform"], info["channel"], info["device_id"], info["third_openid"], info["third_token"], info["third_unionid"])
	if !CheckErr(err) {
		return ERRCODE.DB_ERROR
	}
	return ERRCODE.OK
}

func NewUser(info map[string]string) int {
	stmt, err := db.Prepare("INSERT users SET uid=?,portrait_url=?,name=?,gender=?,email=?,phone=?,admin_level=?,create_time=now()")
	if !CheckErr(err) {
		return ERRCODE.DB_ERROR
	}

	_, err = stmt.Exec(info["uid"], info["portrait_url"], info["name"], info["gender"], info["email"], info["phone"], info["admin_level"])
	if !CheckErr(err) {
		return ERRCODE.DB_ERROR
	}
	return ERRCODE.OK
}

func GetUserInfo(uid string) (map[string]interface{}, int) {
	rows, err := db.Query("select * from users where uid=?", uid)
	if !CheckErr(err) {
		return nil, ERRCODE.DB_ERROR
	}

	defer rows.Close()

	rows.Next()
	user_info, err := ScanSingleRowToMap(rows)
	if !CheckErr(err) {
		return nil, ERRCODE.DB_ERROR
	}

	return user_info, ERRCODE.OK
}

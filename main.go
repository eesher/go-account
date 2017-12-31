package main

import (
	"go-account/util"
	"log"
	"net/http"
)

func main() {

	util.MysqlInit()
	util.TokenInit("wewillfuckyou", 3600)
	ErrcodeInit()
	RoutesInit()

	http.HandleFunc("/login", Handler(Login))
	http.HandleFunc("/auth", Handler(Auth))

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

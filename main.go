package main

import (
    "log"
    "net/http"
)

func main() {

	MysqlInit()
	TokenInit("wewillfuckyou", 3600)
	ErrcodeInit()
	RoutesInit()

    http.HandleFunc("/login", Handler(Login))
    http.HandleFunc("/auth", Handler(Auth))

    log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

package main

import (
	"go-account/util"
	"log"
	"net/http"
)

func main() {

	config := util.LoadConfiguration("config.cfg")
	if config == nil {
		return
	}
	util.MysqlInit(config.Db)
	util.TokenInit(config.Token)
	util.ErrcodeInit()
	//RoutesInit()

	http.HandleFunc("/login", Handler(Login))
	http.HandleFunc("/auth", Handler(Auth))

	log.Fatal(http.ListenAndServe(config.Server.Listen, nil))
}

package util

import "gopkg.in/gcfg.v1"

type ServerConfig struct {
	Listen    string
	Verbose   bool
	AccessLog string
	ErrorLog  string
}

type DbConfig struct {
	Driver     string
	Connection string
	Prefix     string
}

type TokenConfig struct {
	Salt   string
	Expire int64
}

type ConfigFile struct {
	Server ServerConfig
	Db     DbConfig
	Token  TokenConfig
}

const defaultConfig = `
    [server]
	listen = 0.0.0.0:8080
    verbose = false
    accessLog = -
    errorLog  = -

	[db]
    driver     = mysql
    connection = root@tcp(127.0.0.1:3306)/platform?charset=utf8mb4
    prefix  =

	[token]
	salt = fuckoff
	expire = 3600
`

func LoadConfiguration(cfgFile string) *ConfigFile {
	var err error
	var cfg ConfigFile

	if cfgFile != "" {
		err = gcfg.ReadFileInto(&cfg, cfgFile)
	} else {
		err = gcfg.ReadStringInto(&cfg, defaultConfig)
	}

	if !CheckErr(err) {
		return nil
	}

	return &cfg
}

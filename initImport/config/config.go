package config

import "log"

var (
	LogVer int = 3
)

func init() {
	log.Println("log init..")
}

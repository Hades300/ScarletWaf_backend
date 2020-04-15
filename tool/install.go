package tool

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"scarlet/common"
)

var temp = `
[redis]
addr = "127.0.0.1:6379"
password = "123456"

[mysql]
addr = "127.0.0.1:3306"
database = "scarlet"
password = "123456"
username = "scarlet"`

func CheckInstall() bool {
	if _, err := os.Stat("../scarlet.toml"); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func Install() {
	conf := common.Conf{}
	if _, err := toml.Decode(temp, &conf); err != nil {
		log.Fatal("ERR INSTALL", err)
	}
	fmt.Printf("%+v", conf)
}

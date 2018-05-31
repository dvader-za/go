package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Load ...
func Load(path string, conf interface{}) {
	if !exists(path) {
		panic("The config file is missing at the location " + path)
	}
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &conf)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

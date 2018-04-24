package main

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"encoding/json"
)
func main() {


	config := make(map[string]interface{})
	config["filename"] = "./logcollect.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	err = logs.SetLogger(logs.AdapterFile, string(configStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(configStr))
	logs.Debug("test debug")
	logs.Info("test info")
	logs.Warn("test warning")
	logs.Error("test error")

}

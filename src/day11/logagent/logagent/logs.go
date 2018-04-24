package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
	"day11/logagent/conf"
)

func initLogs() (err error)  {
	config := make(map[string]interface{})
	config["filename"] = conf.AppConf.LogPath
	switch conf.AppConf.LogLevel {
	case "info":
		config["level"] = logs.LevelInfo
	case "debug":
		config["level"] = logs.LevelDebug
	case "error":
		config["error"] = logs.LevelError
	default:
		config["error"] = logs.LevelWarn



	}

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	err = logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

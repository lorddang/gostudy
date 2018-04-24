package conf

import (
	"github.com/astaxie/beego/config"
)

type appConf struct {
	LogPath string
	LogLevel string
}

type collectionConf struct {
	LogFile string
	Topic string
}

type CollConfMgr struct {
	Confs []*collectionConf
}

type KfkConf struct {
	ServerAddr string
}

var (
	conf config.Configer
	AppConf *appConf
	CollConfs *CollConfMgr
	KafkaConf *KfkConf
)

func InitConfig(confType string, filePath string) (err error) {
	conf, err = config.NewConfig(confType, filePath)
	initAppConf()
	initCollConf()
	initKafkaConf()
	return
}

func initAppConf(){
	AppConf = &appConf{
	}
	AppConf.LogLevel = conf.String("log::log_level")
	AppConf.LogPath = conf.String("log::log_path")
}

func initKafkaConf(){
	KafkaConf = &KfkConf{}
	KafkaConf.ServerAddr = conf.String("kafka::server_addr")
}

func initCollConf()  {

	cc := &collectionConf{}
	cc.LogFile = conf.String("collection::log_file_path")
	cc.Topic = conf.String("collection::topic")
	CollConfs = &CollConfMgr{}
	CollConfs.Confs = append(CollConfs.Confs, cc)
}

package tailf

import (
	"github.com/hpcloud/tail"
	"fmt"
	"time"
	"day11/logagent/conf"
)

var (
	lineChannel = make(chan *LineMsg, 10)
)

type LineMsg struct {
	Msg string
	Topic string
}

func InitTail(collconf *conf.CollConfMgr) (err error){
	for _, v := range collconf.Confs {
		var tails *tail.Tail
		tails, err = tail.TailFile(v.LogFile, tail.Config{
			ReOpen:    true,
			Follow:    true,
			Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})
		if err != nil {
			fmt.Println("tail file err:", err)
			return
		}

		go func() {
			for true {
				msg, ok := <- tails.Lines
				if !ok {
					fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
					time.Sleep(100 * time.Millisecond)

					continue
				}
				sendToChannel(msg.Text, v.Topic)
				fmt.Println(msg.Text)
			}
		}()
	}
	return
}

func sendToChannel(msg string, top string)  {
	lineChannel <- &LineMsg{
		Msg:msg,
		Topic:top,
	}
	return
}

func ReadFromChan() *LineMsg {
	lmsg := <- lineChannel
	return lmsg
}
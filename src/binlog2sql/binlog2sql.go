package main
import (
	"github.com/siddontang/go-mysql/replication"
	"github.com/siddontang/go-mysql/mysql"
	"context"
	"os"
)

func main() {
	cfg := replication.BinlogSyncerConfig {
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "",
	}
	syncer := replication.NewBinlogSyncer(cfg)

	// Start sync with sepcified binlog file and position
	binlogFile := "mysql-bin.000051"
	binlogPos := uint32(4)
	binLogPostion := mysql.Position{
		Name:binlogFile,
		Pos:binlogPos,
	}
	streamer, _ := syncer.StartSync(binLogPostion)

	// or you can start a gtid replication like
	// streamer, _ := syncer.StartSyncGTID(gtidSet)
	// the mysql GTID set likes this "de278ad0-2106-11e4-9f8e-6edd0ca20947:1-2"
	// the mariadb GTID set likes this "0-1-100"

	for {
		ev, _ := streamer.GetEvent(context.Background())
		// Dump event
		if ev.Header.EventType == 4 {
			ev.Dump(os.Stdout)
		}


	}

}

package main


type binlogStream struct {
	event chan interface{}
	tables map[uint8]table
}

type table struct {
	schemaName string
	tableName string
}

func (bs *binlogStream)Write(data []byte) (n int, err error)  {
	bs.event <- data
	//fmt.Println("write ", string(data))
	return  len(data), nil
}
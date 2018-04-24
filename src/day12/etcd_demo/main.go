package main

import (
	etcdClient "github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
)


func main() {
	client , err := etcdClient.New(etcdClient.Config{
		Endpoints: []string{"zk1:2379", "zk2:2379", "zk3:2379"},
		DialTimeout: 5 * time.Second,

	})
	if err != nil {
		fmt.Println("connect faild error: ", err)
		return
	}
	defer client.Close()
	fmt.Println("connect success ")
	//getrps, err:= client.Get(context.Background(), "/usr/local/etcd")
	//mblsresp, err :=client.Cluster.MemberList(context.Background())
	//for _,v := range mblsresp.Members {
	//	fmt.Println(v.Name, v.ClientURLs)
	//}
	//for k,v := range getrps.Kvs {
	//	fmt.Println(k,string(v.Key), string(v.Value))
	//}
	//wc := client.Watch(context.Background(), "/usr/local/etcd")
	//for msg :=  range wc {
	//	for _, v := range msg.Events {
	//		fmt.Println(v.Kv, v.PrevKv, v.Type)
	//	}
	//}
	client.Put(context.Background(), "/usr/local/etcd", "ff")

}

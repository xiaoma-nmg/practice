package etcdconnect

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
)

/*
	服务端注册etcd
*/

const schema = "MQ"

var cli *clientv3.Client

func Register(etcdAddr string, name string, addr string, ttl int64) error {
	var err error
	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(etcdAddr, ";"),
			DialTimeout: 15 * time.Second,
		})
		if err != nil {
			fmt.Println("connect to etcd server error: ", err)
			return err
		}
	}

	ticker := time.NewTicker(time.Second * time.Duration(ttl))

	go func() {
		for {
			key := "/" + schema + "/" + name + "/" + addr

			getResp, err := cli.Get(context.Background(), key)
			if err != nil {
				fmt.Printf("get key:[%s], error:%#v", key, err)
			} else if getResp.Count == 0 {
				err = withAlive(name, addr, ttl)
				if err != nil {
					fmt.Printf("withAlive error: %#v\n", err)
				}
			} else {
				//fmt.Printf("get [%s], result is [%#v]\n", key, getResp)
			}

			<-ticker.C
		}
	}()

	return nil
}

func withAlive(name string, addr string, ttl int64) error {
	leaseResp, err := cli.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}

	key := "/" + schema + "/" + name + "/" + addr
	_, err = cli.Put(context.Background(), key, addr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		fmt.Printf("put key:[%s] error: %v\n", key, err)
		return err
	}

	ch, err := cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		fmt.Println("KeepAlive error: ", err)
		return err
	}

	// 清空 KeepAlive 返回的channel 以防channel爆满
	go func() {
		for {
			<-ch
		}
	}()

	return nil
}

func UnRegister(name string, addr string) {
	if cli != nil {
		_, _ = cli.Delete(context.Background(), "/"+schema+"/"+name+"/"+addr)
	}
}

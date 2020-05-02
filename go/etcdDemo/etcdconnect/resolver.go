package etcdconnect

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/resolver"
)

/*
   客户端解析etcd
*/

type etcdResolver struct {
	etcdAddr string
	conn     resolver.ClientConn
}

func NewetcdResolver(etcdAddr string) *etcdResolver {
	return &etcdResolver{
		etcdAddr: etcdAddr,
	}
}

func (r *etcdResolver) Scheme() string {
	return schema
}

func (r *etcdResolver) watch(keyPrefix string) {
	var addrList []resolver.Address

	// 获得初始状态下的地址列表
	getResp, err := cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("get key:[%s], error:[%#v] \n", keyPrefix, err)
		return
	} else {
		for i := range getResp.Kvs {
			addrList = append(addrList, resolver.Address{
				Addr: string(getResp.Kvs[i].Value),
			})
		}
	}

	r.conn.NewAddress(addrList)

	// 使用Watch机制，监控服务器地址的变化情况
	// Watch 是长连接一直在后台工作
	// 监控keyPrefix这个前缀
	rch := cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := string(ev.Kv.Value)
			switch ev.Type {
			case mvccpb.PUT:
				if !exist(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})
					r.conn.NewAddress(addrList)
				}
			case mvccpb.DELETE:
				if newAddrList, ok := remove(addrList, addr); ok {
					addrList = newAddrList
					r.conn.NewAddress(addrList)
				}
			}
		}
	}
}

func exist(addrList []resolver.Address, addr string) bool {
	for _, val := range addrList {
		if val.Addr == addr {
			return true
		}
	}
	return false
}

func remove(addrList []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range addrList {
		if addrList[i].Addr == addr {
			addrList[i] = addrList[len(addrList)-1]
			return addrList[:len(addrList)-1], true
		}
	}
	return nil, false
}

func (r *etcdResolver) Build(target resolver.Target,
	cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var err error

	// 构建etcd cli
	if cli == nil {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(r.etcdAddr, ";"),
			DialTimeout: 15 * time.Second,
		})
		if err != nil {
			fmt.Println("connect to etcd server error: ", err)
			return nil, err
		}
	}

	r.conn = cc

	go r.watch("/" + target.Scheme + "/" + target.Endpoint + "/")
	return r, nil
}

func (r *etcdResolver) ResolveNow(resolver.ResolveNowOptions) {
	//TODO
}

// Close closes the resolver.
func (r *etcdResolver) Close() {
	//TODO
}

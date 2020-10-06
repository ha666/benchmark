package main

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/select/roundrobin"
	"strings"
	"time"
)

const etcdAddr = "127.0.0.1:2379"

func NewClient(ServiceName string) client.Client {
	addrList := strings.Split(etcdAddr, ",")
	return client.NewClient(
		client.RequestTimeout(time.Second*10),
		client.PoolSize(100),
		client.Wrap(roundrobin.NewClientWrapper()),
		client.Registry(etcdv3.NewRegistry(
			registry.Addrs(addrList...),
		),
		),
	)
}

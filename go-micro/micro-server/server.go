package main

import (
	"github.com/ha666/logs"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/select/roundrobin"
	"strconv"
	"strings"
	"time"
)

const etcdAddr = "127.0.0.1:2379"

func NewService(registryHandle func(server server.Server), serviceName, serviceVersion string, servicePort int) micro.Service {
	service := micro.NewService(appOptions(serviceName, serviceVersion, servicePort)...)
	registryHandle(service.Server())
	if err := service.Server().Init(server.Wait(nil)); err != nil {
		logs.Emergency(err.Error())
		return nil
	}
	return service
}

func appOptions(serviceName, serviceVersion string, servicePort int) (options []micro.Option) {
	options = []micro.Option{
		micro.Name(serviceName),
		micro.Version(serviceVersion),
		micro.Metadata(map[string]string{"micro_metric_port": strconv.Itoa(servicePort + 2)}),
		micro.Address(":" + strconv.Itoa(servicePort)),
		micro.Broker(broker.NewBroker(broker.Addrs(":" + strconv.Itoa(servicePort+1)))),
		micro.RegisterTTL(time.Second * 10),
		micro.RegisterInterval(time.Second * 30),
	}
	options = append(options, appRegistry(), appClientOption())
	return
}

func appRegistry() micro.Option {
	addrList := strings.Split(etcdAddr, ",")
	return micro.Registry(etcdv3.NewRegistry(
		registry.Addrs(addrList...),
	))
}

func appClientOption() micro.Option {
	addrList := strings.Split(etcdAddr, ",")
	return micro.Client(client.NewClient(
		client.RequestTimeout(time.Second*10),
		client.PoolSize(100),
		client.Wrap(roundrobin.NewClientWrapper()),
		client.Registry(etcdv3.NewRegistry(
			registry.Addrs(addrList...),
		))))
}

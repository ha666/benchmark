package main

import (
	"context"
	"github.com/ha666/benchmark/rpcx-pb/helloworld"
	"github.com/ha666/logs"
	"github.com/smallnest/rpcx/client"
	"runtime"
)

const (
	basePath = "/rpcx_pb_test"
	etcdAddr = "localhost:2379"
)

func init() {
	initLog()
	initConcurrency()
	initClients()
}

func initLog() {
	_ = logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}

func initConcurrency() {
	concurrency = runtime.NumCPU() * 2
	if concurrency < 1 || concurrency > 512 {
		logs.Emergency("无效的并发数:%d", concurrency)
	}
}

func initClients() {
	clients = make(map[int]*helloworld.GreeterClient, concurrency)
	d := client.NewEtcdV3Discovery(basePath, "Greeter", []string{etcdAddr}, nil)
	for i := 0; i < concurrency; i++ {
		args := &helloworld.HelloRequest{
			A: int32(i),
			B: int32(i),
		}
		clients[i] = helloworld.NewGreeterClient(
			client.NewXClient(
				"Greeter",
				client.Failfast,
				client.RoundRobin,
				d,
				client.DefaultOption))
		reply, err := clients[i].SayHello(context.Background(), args)
		if err != nil {
			logs.Emergency("failed to call: %v", err)
		}
		logs.Info("%d * %d = %d", args.A, args.B, reply.C)
	}
}

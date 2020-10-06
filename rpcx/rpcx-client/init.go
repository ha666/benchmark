package main

import (
	"context"
	"github.com/ha666/benchmark/rpcx/example"
	"github.com/ha666/logs"
	"github.com/smallnest/rpcx/client"
	"runtime"
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
	clients = make(map[int]client.XClient, concurrency)
	d := client.NewEtcdV3Discovery(*basePath, "Arith", []string{*etcdAddr}, nil)
	args := &example.Args{
		A: 10,
		B: 20,
	}
	for i := 0; i < concurrency; i++ {
		clients[i] = client.NewXClient("Arith", client.Failfast, client.RoundRobin, d, client.DefaultOption)
		reply := &example.Reply{}
		err := clients[i].Call(context.Background(), "Mul", args, reply)
		if err != nil {
			logs.Emergency("failed to call: %v", err)
		}
	}
}

func closeClients() {
	for i := 0; i < concurrency; i++ {
		clients[i].Close()
	}
}

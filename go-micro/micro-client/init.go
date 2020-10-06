package main

import (
	"context"
	"fmt"
	"github.com/ha666/benchmark/go-micro/proto/arith"
	"github.com/ha666/logs"
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
	clients = make(map[int]arith.ArithService, concurrency)
	for i := 0; i < concurrency; i++ {
		serverName := fmt.Sprintf("micro-client-%d", i)
		microClientInstance := NewClient(serverName)
		clients[i] = arith.NewArithService("arith-service", microClientInstance)
		req := &arith.Args{
			A: 10,
			B: 20,
		}
		_, err := clients[i].Mul(context.Background(), req)
		if err != nil {
			logs.Emergency("failed to call: %v", err)
		}
	}
}

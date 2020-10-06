package main

import (
	"github.com/ha666/benchmark/go-micro/proto/arith"
	"github.com/ha666/logs"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

func main() {
	var service micro.Service
	service = NewService(func(server server.Server) {
		arith.RegisterArithHandler(server, new(Arith))
	}, "arith-service", "2020.1006.1500", 50801)
	service.Init()
	if err := service.Run(); err != nil {
		logs.Emergency(err)
	}
}

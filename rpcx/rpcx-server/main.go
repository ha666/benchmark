package main

import (
	"flag"
	"github.com/ha666/benchmark/rpcx/example"
	"github.com/ha666/logs"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

var (
	addr     = flag.String("addr", "localhost:8972", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(example.Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		logs.Emergency(err)
	}
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		logs.Emergency(err)
	}
	s.Plugins.Add(r)
}

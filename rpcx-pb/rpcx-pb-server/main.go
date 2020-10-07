package main

import (
	"context"
	"flag"
	"github.com/ha666/benchmark/rpcx-pb/helloworld"
	"github.com/ha666/logs"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

var (
	addr     = flag.String("addr", "localhost:9972", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_pb_test", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Greeter", new(GreeterImpl), "")
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

type GreeterImpl struct{}

// SayHello is server rpc method as defined
func (s *GreeterImpl) SayHello(ctx context.Context, args *helloworld.HelloRequest, reply *helloworld.HelloReply) (err error) {
	*reply = helloworld.HelloReply{
		C: args.A * args.B,
	}
	return nil
}

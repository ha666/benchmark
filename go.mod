module github.com/ha666/benchmark

go 1.14

replace (
	github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/golang/protobuf v1.4.2
	github.com/ha666/golibs v2020.712.900+incompatible // indirect
	github.com/ha666/logs v2019.830.1004+incompatible
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/montanaflynn/stats v0.6.3
	github.com/smallnest/rpcx v0.0.0-20200924044220-f2cdd4dea15a
	google.golang.org/protobuf v1.23.0
)

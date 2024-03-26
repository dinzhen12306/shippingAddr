package main

import (
	"flag"
	"fmt"
	"shippingAddr/internal/config"
	"shippingAddr/internal/server"
	"shippingAddr/internal/svc"
	"shippingAddr/models/mysql"
	"shippingAddr/shippingAddrs"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/shippingaddr.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	err := mysql.Sync(c)
	if err != nil {
		panic(err)
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		shippingAddrs.RegisterShippingAddrServer(grpcServer, server.NewShippingAddrServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

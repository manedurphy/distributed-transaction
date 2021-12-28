package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "dist-tranx/api/orders/v1"
	service "dist-tranx/orders/internal"

	_ "github.com/go-sql-driver/mysql"
)

type opts struct {
	Debug bool  `short:"d" long:"debug" description:"set logs to debug level"`
	Port  int32 `short:"p" long:"port" description:"the port to run the gRPC server"`
}

func main() {
	var (
		ctx        context.Context
		lis        net.Listener
		cmdOpts    opts
		grpcOpts   []grpc.ServerOption
		grpcServer *grpc.Server
		srv        *service.Service
		logger     *logrus.Logger
		formatter  *logrus.JSONFormatter
		address    string
		err        error
	)

	if _, err = flags.Parse(&cmdOpts); err != nil {
		panic(err)
	}

	ctx = context.Background()
	logger = logrus.New()

	formatter = new(logrus.JSONFormatter)
	if cmdOpts.Debug {
		formatter.PrettyPrint = true
		logger.Level = logrus.DebugLevel
	}
	logger.Formatter = formatter

	address = ":8081"
	if cmdOpts.Port > 0 {
		address = fmt.Sprintf(":%d", cmdOpts.Port)
	}
	if lis, err = net.Listen("tcp", address); err != nil {
		panic(err)
	}

	grpcServer = grpc.NewServer(grpcOpts...)
	if srv, err = service.NewService(service.Config{
		Logger: logger,
	}); err != nil {
		panic(err)
	}

	go srv.ListenForPayments(ctx)
	pb.RegisterOrderServiceServer(grpcServer, srv)

	logger.WithFields(logrus.Fields{
		"address":      address,
		"service_name": "orders",
	}).Infoln("gRPC server started")
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

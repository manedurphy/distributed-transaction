package main

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "dist-tranx/api/orders/v1"
	service "dist-tranx/orders/internal"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		ctx        context.Context
		lis        net.Listener
		opts       []grpc.ServerOption
		grpcServer *grpc.Server
		srv        *service.Service
		logger     *logrus.Logger
		err        error
	)

	ctx = context.Background()
	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		PrettyPrint: true,
	}
	if lis, err = net.Listen("tcp", ":8081"); err != nil {
		panic(err)
	}

	grpcServer = grpc.NewServer(opts...)
	if srv, err = service.NewService(service.Config{
		Logger: logger,
	}); err != nil {
		panic(err)
	}
	go srv.ListenForPayments(ctx)

	pb.RegisterOrderServiceServer(grpcServer, srv)

	logger.WithFields(logrus.Fields{
		"port":         8081,
		"service_name": "orders",
	}).Infoln("gRPC server started")
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

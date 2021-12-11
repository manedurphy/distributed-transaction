package main

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "dist-tranx/customers/customer/v1"
	service "dist-tranx/customers/internal"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		lis        net.Listener
		opts       []grpc.ServerOption
		grpcServer *grpc.Server
		srv        *service.Service
		logger     *logrus.Logger
		ctx        context.Context
		err        error
	)

	ctx = context.Background()
	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		PrettyPrint: true,
	}
	if lis, err = net.Listen("tcp", ":8080"); err != nil {
		panic(err)
	}

	grpcServer = grpc.NewServer(opts...)
	if srv, err = service.NewService(service.Config{
		Logger: logger,
	}); err != nil {
		panic(err)
	}

	go srv.ListenForOrders(ctx)
	pb.RegisterCustomerServiceServer(grpcServer, srv)

	logger.WithFields(logrus.Fields{
		"port":         8080,
		"service_name": "customers",
	}).Infoln("gRPC server started")
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "dist-tranx/api/payments/v1"
	service "dist-tranx/payments/internal"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		lis        net.Listener
		opts       []grpc.ServerOption
		grpcServer *grpc.Server
		srv        *service.Service
		logger     *logrus.Logger
		err        error
	)

	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		PrettyPrint: true,
	}
	if lis, err = net.Listen("tcp", ":8082"); err != nil {
		panic(err)
	}

	grpcServer = grpc.NewServer(opts...)
	if srv, err = service.NewService(service.Config{
		Logger: logger,
	}); err != nil {
		panic(err)
	}

	pb.RegisterPaymentServiceServer(grpcServer, srv)

	logger.WithFields(logrus.Fields{
		"port":         8082,
		"service_name": "payments",
	}).Infoln("gRPC server started")
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

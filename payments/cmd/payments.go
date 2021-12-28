package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "dist-tranx/api/payments/v1"
	service "dist-tranx/payments/internal"

	_ "github.com/go-sql-driver/mysql"
)

type opts struct {
	Debug bool  `short:"d" long:"debug" description:"set logs to debug level"`
	Port  int32 `short:"p" long:"port" description:"the port to run the gRPC server"`
}

func main() {
	var (
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

	logger = logrus.New()
	formatter = new(logrus.JSONFormatter)
	if cmdOpts.Debug {
		formatter.PrettyPrint = true
		logger.Level = logrus.DebugLevel
	}
	logger.Formatter = formatter

	address = ":8082"
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
	pb.RegisterPaymentServiceServer(grpcServer, srv)

	logger.WithFields(logrus.Fields{
		"address":      address,
		"service_name": "payments",
	}).Infoln("gRPC server started")
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

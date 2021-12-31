package main

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "dist-tranx/api/customers/v1"
	service "dist-tranx/customers/internal"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultConfigPath = "/app/configs/customers.hcl"
)

type (
	opts struct {
		Debug      bool   `short:"d" long:"debug" description:"set logs to debug level"`
		Port       int32  `short:"p" long:"port" description:"the port to run the gRPC server"`
		ConfigFile string `short:"f" long:"file" description:"the path to the configuration file"`
	}

	Config struct {
		Debug bool  `hcl:"debug" json:"debug"`
		Port  int32 `hcl:"port" json:"port"`
	}
)

func main() {
	var (
		lis        net.Listener
		cmdOpts    opts
		grpcOpts   []grpc.ServerOption
		grpcServer *grpc.Server
		srv        *service.Service
		logger     *logrus.Logger
		formatter  *logrus.JSONFormatter
		ctx        context.Context
		cfg        Config
		debug      bool
		port       int32
		cfgPath    string
		address    string
		err        error
	)

	if _, err = flags.Parse(&cmdOpts); err != nil {
		panic(err)
	}

	ctx = context.Background()
	logger = logrus.New()
	formatter = new(logrus.JSONFormatter)
	cfgPath = defaultConfigPath

	if cmdOpts.ConfigFile != "" {
		cfgPath = defaultConfigPath
	}

	if err = hclsimple.DecodeFile(cfgPath, nil, &cfg); err != nil {
		panic(err)
	}

	debug = cfg.Debug
	if cmdOpts.Debug {
		debug = cmdOpts.Debug
	}

	if debug {
		formatter.PrettyPrint = true
		logger.Level = logrus.DebugLevel
	}
	logger.Formatter = formatter

	logger.WithFields(logrus.Fields{
		"config": cfg,
	}).Infoln("config file loaded")

	port = cfg.Port
	if cmdOpts.Port > 0 {
		port = cmdOpts.Port
	}

	address = fmt.Sprintf(":%d", port)
	if lis, err = net.Listen("tcp", address); err != nil {
		logger.WithFields(logrus.Fields{
			"address": address,
			"err":     err,
		}).Panicln("could not created listener")
	}

	grpcServer = grpc.NewServer(grpcOpts...)
	if srv, err = service.NewService(service.Config{
		Logger: logger,
	}); err != nil {
		logger.WithFields(logrus.Fields{
			"address": address,
			"err":     err,
		}).Panicln("could not created service")
	}

	go srv.ListenForEvents(ctx)
	pb.RegisterCustomerServiceServer(grpcServer, srv)

	logger.WithFields(logrus.Fields{
		"address":      address,
		"service_name": "customers",
	}).Infoln("gRPC server started")

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "dist-tranx/api/payments/v1"
	service "dist-tranx/payments/internal"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultConfigPath = "/app/configs/payments.hcl"
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
		cfg        Config
		port       int32
		debug      bool
		cfgPath    string
		address    string
		err        error
	)

	if _, err = flags.Parse(&cmdOpts); err != nil {
		panic(err)
	}

	logger = logrus.New()
	formatter = new(logrus.JSONFormatter)

	cfgPath = defaultConfigPath
	if cmdOpts.ConfigFile != "" {
		cfgPath = cmdOpts.ConfigFile
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
	pb.RegisterPaymentServiceServer(grpcServer, srv)

	logger.WithFields(logrus.Fields{
		"address":      address,
		"service_name": "payments",
	}).Infoln("gRPC server started")
	if err = grpcServer.Serve(lis); err != nil {
		logger.WithFields(logrus.Fields{
			"address": address,
			"err":     err,
		}).Panicln("could not start gRPC server")
	}
}

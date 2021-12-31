package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/jessevdk/go-flags"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	customerv1 "dist-tranx/api/customers/v1"
	orderv1 "dist-tranx/api/orders/v1"
	paymentv1 "dist-tranx/api/payments/v1"
)

const (
	defaultConfigPath = "/app/configs/apigateway.hcl"
)

type (
	apigateway struct {
		logger *logrus.Logger
	}

	opts struct {
		Debug      bool   `short:"d" long:"debug" description:"set logs to debug level"`
		FsPort     int32  `short:"s" long:"file-port" description:"the port to run the file server"`
		ApiPort    int32  `short:"a" long:"api-port" description:"the port to run the API server"`
		ConfigFile string `short:"f" long:"file" description:"the path to the configuration file"`
	}

	Config struct {
		Ports struct {
			FileServer int32 `hcl:"file_server"`
			Api        int32 `hcl:"api"`
		} `hcl:"ports,block"`
	}
)

func newAPIGateway() *apigateway {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	return &apigateway{
		logger: logger,
	}
}

func main() {
	var (
		ctx        context.Context
		mux        *runtime.ServeMux
		grpcOpts   []grpc.DialOption
		srv        http.Server
		fs         http.Handler
		httpMux    *http.ServeMux
		apigw      *apigateway
		cfg        Config
		cmdOpts    opts
		cfgPath    string
		apiAddress string
		err        error
	)

	if _, err = flags.Parse(&cmdOpts); err != nil {
		panic(err)
	}

	cfgPath = cmdOpts.ConfigFile
	if cmdOpts.ConfigFile == "" {
		cfgPath = defaultConfigPath
	}

	if err = hclsimple.DecodeFile(cfgPath, nil, &cfg); err != nil {
		panic(err)
	}

	ctx = context.Background()
	apigw = newAPIGateway()

	mux = runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
		}),
	)

	grpcOpts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = orderv1.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, os.Getenv("ORDERS_SERVICE_ADDR"), grpcOpts); err != nil {
		apigw.logger.WithFields(logrus.Fields{
			"err": err,
		}).Panicln("could not register apigateway with orders service")
	}

	if err = customerv1.RegisterCustomerServiceHandlerFromEndpoint(ctx, mux, os.Getenv("CUSTOMERS_SERVICE_ADDR"), grpcOpts); err != nil {
		apigw.logger.WithFields(logrus.Fields{
			"err": err,
		}).Panicln("could not register apigateway with customers service")
	}

	if err = paymentv1.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, os.Getenv("PAYMENTS_SERVICE_ADDR"), grpcOpts); err != nil {
		apigw.logger.WithFields(logrus.Fields{
			"err": err,
		}).Panicln("could not register apigateway with payments service")
	}

	fs = http.StripPrefix("/distributed-transaction/", http.FileServer(http.Dir("./static")))
	httpMux = http.NewServeMux()

	httpMux.Handle("/distributed-transaction/", fs)
	httpMux.HandleFunc("/distributed-transaction/healthz", func(w http.ResponseWriter, r *http.Request) {
		apigw.logger.Infoln("healthy!")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	go func() {
		var (
			address string
			httpSrv http.Server
		)

		address = fmt.Sprintf(":%d", cfg.Ports.FileServer)
		httpSrv = http.Server{
			Addr:    address,
			Handler: httpMux,
		}

		apigw.logger.WithFields(logrus.Fields{
			"address": address,
		}).Infoln("starting apigateway static file server")
		if err := httpSrv.ListenAndServe(); err != nil {
			apigw.logger.WithFields(logrus.Fields{
				"err": err,
			}).Panicln("could not start static file server")
		}
	}()

	mux.HandlePath("GET", "/distributed-transaction/sse/*", apigw.getOrderStatus)

	apiAddress = fmt.Sprintf(":%d", cfg.Ports.Api)
	srv = http.Server{
		Addr:    apiAddress,
		Handler: cors.Default().Handler(mux),
	}

	apigw.logger.WithFields(logrus.Fields{
		"address": apiAddress,
	}).Infoln("starting apigateway server")
	if err = srv.ListenAndServe(); err != nil {
		apigw.logger.WithFields(logrus.Fields{
			"err": err,
		}).Panicln("could not start apigateway server")
	}
}

func (ag *apigateway) getOrderStatus(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var (
		id          int
		b           []byte
		ctx         context.Context
		conn        *grpc.ClientConn
		resp        *orderv1.GetOrderStatusResponse
		orderClient orderv1.OrderServiceClient
		err         error
	)
	if id, err = strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/distributed-transaction/sse/")); err != nil {
		ag.logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("trimm URL prefix error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"something went wrong\"}"))
		return
	}

	ctx = context.Background()
	if conn, err = grpc.DialContext(ctx, os.Getenv("ORDERS_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		ag.logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("dialing error occurred")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"something went wrong\"}"))
		return
	}

	orderClient = orderv1.NewOrderServiceClient(conn)
	if resp, err = orderClient.GetOrderStatus(ctx, &orderv1.GetOrderStatusRequest{
		OrderId: int32(id),
	}); err != nil {
		ag.logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("could not get order status")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"could not get order status\"}"))
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	if b, err = protojson.Marshal(resp); err != nil {
		ag.logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("marshaling error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"something went wrong\"}"))
		return
	}

	ag.logger.WithFields(logrus.Fields{
		"data": string(b),
	}).Infoln("streaming data to customer")

	fmt.Fprintf(w, "data: %v\n\n", string(b))
	if f, valid := w.(http.Flusher); valid {
		f.Flush()
	}
}

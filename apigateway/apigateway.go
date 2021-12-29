package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	customerv1 "dist-tranx/api/customers/v1"
	orderv1 "dist-tranx/api/orders/v1"
	paymentv1 "dist-tranx/api/payments/v1"
)

func main() {
	var (
		ctx  context.Context
		mux  *runtime.ServeMux
		opts []grpc.DialOption
		srv  http.Server
		fs   http.Handler
		err  error
	)

	ctx = context.Background()
	mux = runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
		}),
	)

	opts = []grpc.DialOption{grpc.WithInsecure()}
	if err = orderv1.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, os.Getenv("ORDERS_SERVICE_ADDR"), opts); err != nil {
		panic(err)
	}

	if err = customerv1.RegisterCustomerServiceHandlerFromEndpoint(ctx, mux, os.Getenv("CUSTOMERS_SERVICE_ADDR"), opts); err != nil {
		panic(err)
	}

	if err = paymentv1.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, os.Getenv("PAYMENTS_SERVICE_ADDR"), opts); err != nil {
		panic(err)
	}

	fs = http.StripPrefix("/distributed-transaction/", http.FileServer(http.Dir("./static")))
	http.Handle("/distributed-transaction/", fs)
	http.HandleFunc("/distributed-transaction/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	go func() {
		if err := http.ListenAndServe(":3000", nil); err != nil {
			panic(err)
		}
	}()

	mux.HandlePath("GET", "/distributed-transaction/sse/*", getOrderStatus)
	srv = http.Server{
		Addr:    ":8000",
		Handler: cors.Default().Handler(mux),
	}

	fmt.Println("API Gateway started on port 8000")
	if err = srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func getOrderStatus(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
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
		panic(err)
	}

	ctx = context.Background()
	if conn, err = grpc.DialContext(ctx, os.Getenv("ORDERS_SERVICE_ADDR"), grpc.WithInsecure()); err != nil {
		panic(err)
	}

	orderClient = orderv1.NewOrderServiceClient(conn)
	if resp, err = orderClient.GetOrderStatus(ctx, &orderv1.GetOrderStatusRequest{
		OrderId: int32(id),
	}); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	if b, err = protojson.Marshal(resp); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "data: %v\n\n", string(b))
	if f, valid := w.(http.Flusher); valid {
		f.Flush()
	}
}

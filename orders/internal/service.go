package service

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	customerPb "dist-tranx/customers/customer/v1"
	pb "dist-tranx/orders/order/v1"
)

var (
	connString = fmt.Sprintf("%s:%s@tcp(%s)/orders", os.Getenv("ORDERS_MYSQL_USERNAME"), os.Getenv("ORDERS_MYSQL_PASSWORD"), os.Getenv("ORDERS_DB_ADDR"))
)

type (
	Config struct {
		Logger *logrus.Logger
	}

	Service struct {
		pb.UnimplementedOrderServiceServer
		db     *sql.DB
		conn   *redis.Client
		logger *logrus.Logger
	}
)

func NewService(cfg Config) (*Service, error) {
	var (
		conn         *redis.Client
		db           *sql.DB
		sqlStatement string
		err          error
	)
	conn = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	sqlStatement = "TRUNCATE TABLE orders_table"
	db, err = sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	if _, err = db.Exec(sqlStatement); err != nil {
		return nil, err
	}

	return &Service{
		conn:   conn,
		db:     db,
		logger: cfg.Logger,
	}, nil
}

func (s *Service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	var (
		tx                 *sql.Tx
		result             sql.Result
		makePaymentRequest *customerPb.MakePaymentRequest
		b                  []byte
		sqlStatement       string
		orderId            int64
		err                error
	)
	if tx, err = s.db.BeginTx(ctx, nil); err != nil {
		return &pb.CreateOrderResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	sqlStatement = fmt.Sprintf("INSERT INTO orders_table (customer_id, total, status) VALUES (%d, %d, 'pending');", req.GetCustomerId(), req.GetTotal())
	if result, err = tx.ExecContext(ctx, sqlStatement); err != nil {
		tx.Rollback()
		s.logger.WithFields(logrus.Fields{
			"customer_id": req.GetCustomerId(),
		}).Errorln("could not place order for customer")
		return &pb.CreateOrderResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	if err = tx.Commit(); err != nil {
		return &pb.CreateOrderResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	if orderId, err = result.LastInsertId(); err != nil {
		return &pb.CreateOrderResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	makePaymentRequest = &customerPb.MakePaymentRequest{
		OrderId:    int32(orderId),
		CustomerId: int32(req.GetCustomerId()),
		Amount:     req.GetTotal(),
	}

	if b, err = proto.Marshal(makePaymentRequest); err != nil {
		return &pb.CreateOrderResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	s.logger.WithFields(logrus.Fields{
		"customer_id": req.GetCustomerId(),
		"order_id":    orderId,
	}).Infoln("successfully placed order")

	if err = s.conn.Publish(ctx, "ORDER_CREATED_EVENT", string(b)).Err(); err != nil {
		s.logger.Errorln("could not publish order event")
		return &pb.CreateOrderResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.CreateOrderResponse{
		Response: &pb.CreateOrderResponse_Order{
			Order: &pb.Order{
				Id:         int32(orderId),
				CustomerId: req.GetCustomerId(),
				Total:      req.GetTotal(),
			},
		},
	}, nil
}

func (s *Service) GetOrderStatus(ctx context.Context, req *pb.GetOrderStatusRequest) (*pb.GetOrderStatusResponse, error) {
	var (
		rows         *sql.Rows
		sqlStatement string
		err          error
		r            struct {
			status string
		}
	)

	sqlStatement = fmt.Sprintf("SELECT status FROM orders_table WHERE id = %d", req.GetOrderId())
	if rows, err = s.db.Query(sqlStatement); err != nil {
		return &pb.GetOrderStatusResponse{}, err
	}

	for rows.Next() {
		if err = rows.Scan(&r.status); err != nil {
			return &pb.GetOrderStatusResponse{}, err
		}
	}

	rows.Close()
	if r.status == "" {
		return &pb.GetOrderStatusResponse{
			Status: "declined",
		}, nil
	}
	return &pb.GetOrderStatusResponse{
		Status: r.status,
	}, nil
}

func (s *Service) ListenForPayments(ctx context.Context) {
	paymentsChan := s.conn.Subscribe(ctx, "PAYMENT_EVENT").Channel()
	s.logger.Infoln("listening for payments...")
	for {
		var (
			tx           *sql.Tx
			payment      *redis.Message
			resp         customerPb.MakePaymentResponse
			sqlStatement string
			err          error
		)

		payment = <-paymentsChan
		if err = proto.Unmarshal([]byte(payment.Payload), &resp); err != nil {
			s.logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("unmarshaling error")
			continue
		}

		if tx, err = s.db.BeginTx(ctx, nil); err != nil {
			s.logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("could not start transaction")
			continue
		}

		switch resp.Response.(type) {
		case *customerPb.MakePaymentResponse_Error:
			s.logger.WithFields(logrus.Fields{
				"order_id": resp.GetOrderId(),
				"err":      resp.GetError().GetErrorMessage(),
			}).Errorln("an error occurred while placing an order")
			sqlStatement = fmt.Sprintf("DELETE FROM orders_table WHERE id='%d';", resp.GetOrderId())
		case *customerPb.MakePaymentResponse_Remaining:
			s.logger.WithFields(logrus.Fields{
				"order_id":            resp.GetOrderId(),
				"remaining_in_wallet": resp.GetRemaining(),
			}).Infoln("order placed")
			sqlStatement = fmt.Sprintf("UPDATE orders_table SET status='in progress' WHERE id='%d';", resp.GetOrderId())
		}

		if _, err = tx.ExecContext(ctx, sqlStatement); err != nil {
			tx.Rollback()
			s.logger.WithFields(logrus.Fields{
				"err":   err,
				"query": sqlStatement,
			}).Errorln("could not execute query")
			continue
		}
		if err = tx.Commit(); err != nil {
			s.logger.WithFields(logrus.Fields{
				"err":   err,
				"query": sqlStatement,
			}).Errorln("could not commit transaction")
		}
	}
}

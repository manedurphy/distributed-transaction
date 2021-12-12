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

	pb "dist-tranx/api/customers/v1"
	paymentv1 "dist-tranx/api/payments/v1"
)

const (
	ORDER_CREATED_EVENT = "ORDER_CREATED_EVENT"
	ADD_FUNDS_EVENT     = "ADD_FUNDS_EVENT"
)

var (
	connString = fmt.Sprintf("%s:%s@tcp(%s)/customers", os.Getenv("CUSTOMERS_MYSQL_USERNAME"), os.Getenv("CUSTOMERS_MYSQL_PASSWORD"), os.Getenv("CUSTOMERS_DB_ADDR"))
)

type (
	Config struct {
		Logger *logrus.Logger
	}

	Service struct {
		pb.UnimplementedCustomerServiceServer
		db     *sql.DB
		conn   *redis.Client
		logger *logrus.Logger
	}
)

func NewService(cfg Config) (*Service, error) {
	var (
		db           *sql.DB
		conn         *redis.Client
		sqlStatement string
		err          error
	)

	sqlStatement = "TRUNCATE TABLE customers_table"
	db, err = sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	conn = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	if _, err = db.Exec(sqlStatement); err != nil {
		return nil, err
	}

	return &Service{
		db:     db,
		conn:   conn,
		logger: cfg.Logger,
	}, err
}

func (s *Service) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	var (
		row          *sql.Row
		customer     pb.Customer
		password     string
		sqlStatement string
		err          error
	)

	sqlStatement = fmt.Sprintf("SELECT id, first_name, last_name, email, password, wallet FROM customers_table WHERE email = '%s';", req.GetEmail())
	row = s.db.QueryRowContext(ctx, sqlStatement)

	if err = row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Email, &password, &customer.Wallet); err != nil {
		s.logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("could not get data from mysql")
		return &pb.GetCustomerResponse{}, status.Error(codes.NotFound, "Customer with that email not found")
	}

	if password != req.GetPassword() {
		s.logger.Errorln("invalid password")
		return &pb.GetCustomerResponse{}, status.Error(codes.FailedPrecondition, "Invalid credentials")
	}

	s.logger.Infoln("successful login")
	return &pb.GetCustomerResponse{
		Customer: &customer,
	}, nil
}

func (s *Service) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	var (
		tx           *sql.Tx
		result       sql.Result
		customer     pb.Customer
		queryResult  *sql.Row
		query        string
		customerId   int64
		sqlStatement string
		err          error
	)

	query = fmt.Sprintf("SELECT id FROM customers.customers_table WHERE email = '%s'", req.GetEmail())
	queryResult = s.db.QueryRowContext(ctx, query)
	queryResult.Scan(&customer.Id)

	if customer.Id > 0 {
		s.logger.WithFields(logrus.Fields{
			"customer_id": customer.Id,
		}).Errorln("customer already exists")
		return &pb.CreateCustomerResponse{}, status.Error(codes.InvalidArgument, "A customer with that email already exists")
	}

	sqlStatement = fmt.Sprintf("INSERT INTO customers_table (first_name, last_name, email, password, wallet) VALUES ('%s', '%s', '%s', '%s', '%d');", req.GetFirstName(), req.GetLastName(), req.GetEmail(), req.GetPassword(), 100)
	if tx, err = s.db.BeginTx(ctx, nil); err != nil {
		return &pb.CreateCustomerResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	if result, err = tx.ExecContext(ctx, sqlStatement); err != nil {
		tx.Rollback()
		s.logger.WithFields(logrus.Fields{
			"err":   err,
			"query": sqlStatement,
		}).Errorln("could not execute query")
		return &pb.CreateCustomerResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	if err = tx.Commit(); err != nil {
		s.logger.WithFields(logrus.Fields{
			"err":   err,
			"query": sqlStatement,
		}).Errorln("could not commit transaction")
		return &pb.CreateCustomerResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	if customerId, err = result.LastInsertId(); err != nil {
		return &pb.CreateCustomerResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	s.logger.WithFields(logrus.Fields{
		"customer_id": customerId,
	}).Infoln("successfully registered a new customer")

	return &pb.CreateCustomerResponse{
		Customer: &pb.Customer{
			Id:        int32(customerId),
			FirstName: req.GetFirstName(),
			LastName:  req.GetLastName(),
			Email:     req.GetEmail(),
			Wallet:    100,
		},
	}, nil
}

func (s *Service) ListenForOrders(ctx context.Context) {
	var (
		ordersChan <-chan *redis.Message
		fundsChan  <-chan *redis.Message
	)

	ordersChan = s.conn.Subscribe(ctx, ORDER_CREATED_EVENT).Channel()
	fundsChan = s.conn.Subscribe(ctx, ADD_FUNDS_EVENT).Channel()

	s.logger.WithFields(logrus.Fields{
		"channels": []string{ORDER_CREATED_EVENT, ADD_FUNDS_EVENT},
	}).Infoln("subscribed to event channels")
	for {
		select {
		case fund := <-fundsChan:
			var (
				addFundsEvent paymentv1.AddFundsEvent
				tx            *sql.Tx
				sqlStatement  string
				err           error
			)

			s.logger.WithFields(logrus.Fields{
				"channel": ADD_FUNDS_EVENT,
			}).Infoln("event occurred")

			if err = proto.Unmarshal([]byte(fund.Payload), &addFundsEvent); err != nil {
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

			sqlStatement = fmt.Sprintf("UPDATE customers_table SET wallet=wallet+%d WHERE id='%d';", addFundsEvent.GetAmount(), addFundsEvent.GetCustomerId())
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
					"err": err,
				}).Errorln("could not commit transaction")
			}
		case order := <-ordersChan:
			var (
				req          pb.MakePaymentRequest
				resp         *pb.MakePaymentResponse
				tx           *sql.Tx
				b            []byte
				remaining    int32
				err          error
				sqlStatement string
			)

			s.logger.WithFields(logrus.Fields{
				"channel": ORDER_CREATED_EVENT,
			}).Infoln("event occurred")

			if err = proto.Unmarshal([]byte(order.Payload), &req); err != nil {
				s.logger.WithFields(logrus.Fields{
					"err": err,
				}).Errorln("unmarshaling error")
				resp = &pb.MakePaymentResponse{
					OrderId: req.GetOrderId(),
					Response: &pb.MakePaymentResponse_Error{
						Error: &pb.MakePaymentResponse_PaymentError{
							ErrorMessage: err.Error(),
						},
					},
				}
				if b, err = proto.Marshal(resp); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("marshaling error")
					continue
				}
				if err = s.conn.Publish(ctx, "PAYMENT_EVENT", string(b)).Err(); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("could not publish event")
				}
				continue
			}

			if tx, err = s.db.BeginTx(ctx, nil); err != nil {
				s.logger.WithFields(logrus.Fields{
					"err": err,
				}).Errorln("could not start transaction")
				resp = &pb.MakePaymentResponse{
					OrderId: req.GetOrderId(),
					Response: &pb.MakePaymentResponse_Error{
						Error: &pb.MakePaymentResponse_PaymentError{
							ErrorMessage: err.Error(),
						},
					},
				}
				if b, err = proto.Marshal(resp); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("marshaling error")
					continue
				}
				if err = s.conn.Publish(ctx, "PAYMENT_EVENT", string(b)).Err(); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("could not publish event")
				}
				continue
			}

			sqlStatement = fmt.Sprintf("UPDATE customers_table SET wallet=wallet-%d WHERE id='%d';", req.GetAmount(), req.GetCustomerId())
			if _, err = tx.ExecContext(ctx, sqlStatement); err != nil {
				tx.Rollback()
				resp = &pb.MakePaymentResponse{
					OrderId: req.GetOrderId(),
					Response: &pb.MakePaymentResponse_Error{
						Error: &pb.MakePaymentResponse_PaymentError{
							ErrorMessage: err.Error(),
						},
					},
				}
				if b, err = proto.Marshal(resp); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("marshaling error")
					continue
				}
				if err = s.conn.Publish(ctx, "PAYMENT_EVENT", string(b)).Err(); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("could not publish event")
				}
				continue
			}

			sqlStatement = fmt.Sprintf("SELECT wallet FROM customers_table WHERE id='%d'", req.GetCustomerId())
			if err = tx.QueryRowContext(ctx, sqlStatement).Scan(&remaining); err != nil {
				tx.Rollback()
				resp = &pb.MakePaymentResponse{
					OrderId: req.GetOrderId(),
					Response: &pb.MakePaymentResponse_Error{
						Error: &pb.MakePaymentResponse_PaymentError{
							ErrorMessage: err.Error(),
						},
					},
				}
				if b, err = proto.Marshal(resp); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("marshaling error")
					continue
				}
				if err = s.conn.Publish(ctx, "PAYMENT_EVENT", string(b)).Err(); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("could not publish event")
				}
				continue
			}

			if err = tx.Commit(); err != nil {
				resp = &pb.MakePaymentResponse{
					OrderId: req.GetOrderId(),
					Response: &pb.MakePaymentResponse_Error{
						Error: &pb.MakePaymentResponse_PaymentError{
							ErrorMessage: err.Error(),
						},
					},
				}
				if b, err = proto.Marshal(resp); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("marshaling error")
					continue
				}
				if err = s.conn.Publish(ctx, "PAYMENT_EVENT", string(b)).Err(); err != nil {
					s.logger.WithFields(logrus.Fields{
						"err": err,
					}).Errorln("could not publish event")
				}
				continue
			}

			resp = &pb.MakePaymentResponse{
				OrderId: req.GetOrderId(),
				Response: &pb.MakePaymentResponse_Remaining{
					Remaining: remaining,
				},
			}

			if b, err = proto.Marshal(resp); err != nil {
				s.logger.WithFields(logrus.Fields{
					"err": err,
				}).Errorln("marshaling error")
				continue
			}

			if err = s.conn.Publish(ctx, "PAYMENT_EVENT", string(b)).Err(); err != nil {
				s.logger.WithFields(logrus.Fields{
					"err": err,
				}).Errorln("could not publish event")
			}

		}
	}
}

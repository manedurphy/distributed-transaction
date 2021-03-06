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

	pb "dist-tranx/api/payments/v1"
)

const (
	ADD_FUNDS_EVENT = "ADD_FUNDS_EVENT"
)

var (
	connString = fmt.Sprintf("%s:%s@tcp(%s)/payments", os.Getenv("PAYMENTS_MYSQL_USERNAME"), os.Getenv("PAYMENTS_MYSQL_PASSWORD"), os.Getenv("PAYMENTS_DB_ADDR"))
)

type (
	Config struct {
		Logger *logrus.Logger
	}

	Service struct {
		pb.UnimplementedPaymentServiceServer
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

	sqlStatement = "TRUNCATE TABLE payments_table"
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

func (s *Service) AddCreditCard(ctx context.Context, req *pb.AddCreditCardRequest) (*pb.AddCreditCardResponse, error) {
	var (
		tx           *sql.Tx
		result       sql.Result
		creditCardId int64
		sqlStatement string
		err          error
	)

	if tx, err = s.db.BeginTx(ctx, nil); err != nil {
		s.logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("could not start transaction")
		return &pb.AddCreditCardResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	sqlStatement = fmt.Sprintf("INSERT INTO payments_table (credit_card_number, expiration, cvv, customer_id) VALUES ('%s', '%s', %d, %d);", req.GetCreditCardNumber(), fmt.Sprintf("%s-01", req.GetExpiration()), req.GetCvv(), req.GetCustomerId())
	if result, err = tx.ExecContext(ctx, sqlStatement); err != nil {
		tx.Rollback()
		s.logger.WithFields(logrus.Fields{
			"err":   err,
			"query": sqlStatement,
		}).Errorln("could not execute query")
		return &pb.AddCreditCardResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	if err = tx.Commit(); err != nil {
		return &pb.AddCreditCardResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	if creditCardId, err = result.LastInsertId(); err != nil {
		return &pb.AddCreditCardResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.AddCreditCardResponse{
		Id: int32(creditCardId),
	}, nil
}

func (s *Service) AddFunds(ctx context.Context, req *pb.AddFundsRequest) (*pb.AddFundsResponse, error) {
	var (
		rows               *sql.Rows
		addFundsEvent      *pb.AddFundsEvent
		addFundsEventProto []byte
		creditCardId       int32
		query              string
		err                error
	)

	s.logger.WithFields(logrus.Fields{
		"request": req,
	}).Infoln("request received")

	query = fmt.Sprintf("SELECT id FROM payments_table WHERE id=%d", req.GetCreditCardId())
	if rows, err = s.db.QueryContext(ctx, query); err != nil {
		s.logger.WithFields(logrus.Fields{
			"err":   err,
			"query": query,
		}).Errorln("could not query database")
		return &pb.AddFundsResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	for rows.Next() {
		if err = rows.Scan(&creditCardId); err != nil {
			s.logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("could not map credit card id value")
			return &pb.AddFundsResponse{}, status.Error(codes.Internal, "Internal server error")
		}
	}

	if creditCardId == 0 {
		return &pb.AddFundsResponse{}, status.Error(codes.NotFound, "Could not find credit card")
	}

	addFundsEvent = &pb.AddFundsEvent{
		CustomerId: req.GetCustomerId(),
		Amount:     req.GetAmount(),
	}

	if addFundsEventProto, err = proto.Marshal(addFundsEvent); err != nil {
		return &pb.AddFundsResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	if err = s.conn.Publish(ctx, ADD_FUNDS_EVENT, addFundsEventProto).Err(); err != nil {
		return &pb.AddFundsResponse{}, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.AddFundsResponse{
		Message: "Successfully added funds",
	}, nil
}

func (s *Service) GetCreditCards(ctx context.Context, req *pb.GetCreditCardsRequest) (*pb.GetCreditCardsResponse, error) {
	var (
		rows        *sql.Rows
		query       string
		creditCards []*pb.GetCreditCardsResponse_CreditCard
		err         error
	)

	query = fmt.Sprintf("SELECT id, credit_card_number FROM payments_table WHERE customer_id = %d", req.GetCustomerId())
	if rows, err = s.db.QueryContext(ctx, query); err != nil {
		return &pb.GetCreditCardsResponse{}, status.Error(codes.NotFound, "No credit cards found for this account")
	}

	creditCards = make([]*pb.GetCreditCardsResponse_CreditCard, 0)
	for rows.Next() {
		var (
			id  int32
			val string
		)
		if err = rows.Scan(&id, &val); err != nil {
			return &pb.GetCreditCardsResponse{}, status.Error(codes.Internal, "Internal server error")
		}

		creditCards = append(creditCards, &pb.GetCreditCardsResponse_CreditCard{
			Id:               id,
			CreditCardNumber: val[len(val)-4:],
		})
	}

	return &pb.GetCreditCardsResponse{
		CreditCards: creditCards,
	}, nil
}

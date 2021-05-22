package app

import (
	"fmt"
	"net"
	"time"

	pb "github.com/RTUITLab/Civfrovoi-Proriv-Back/protos/coordservice"

	"github.com/RTUITLab/Civfrovoi-Proriv-Back/pkg/config"
	"github.com/RTUITLab/Civfrovoi-Proriv-Back/service/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	rabbit "github.com/streadway/amqp"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Port 		string
	DB_URI		string
	AMPQ_URL	string
}

func New(
	cfg *config.Config,
) *App {
	return &App{
		Port: cfg.App.Port,
		DB_URI: cfg.DB.URI,
		AMPQ_URL: cfg.RabbitMQ.URL,
	}
}

func (a *App) Start() error {
	sqlDB, err := connectOrWait("mysql", a.DB_URI, 45 *time.Second)
	if err != nil {
		return err
	}

	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				Conn: sqlDB,
			},
		),
		&gorm.Config{
		
		},
	)

	if err != nil {
		panic(err)
	}

	rabbitConn, err := connToAMPQ(a.AMPQ_URL)
	if err != nil {
		return err
	}

	Server := &server.Server{
		DB: db,
		GetCoordTime: time.Second,
		RabbitConn: rabbitConn,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", a.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCoordsServiceServer(
		grpcServer,
		Server,
	)

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func connectOrWait(driverName, URI string, WaitTime time.Duration) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(
		driverName,
		URI,
	)
	if err != nil {
		for i := 0; i < int(WaitTime.Seconds()); i++ {
			t := time.NewTimer(time.Second)
			<-t.C
			conn, err := sqlx.Connect(driverName, URI)
			if err != nil {
				continue
			} else {
				return conn, nil
			}
		}
		return nil, err
	}
	return conn, nil

}

func connToAMPQ(AMPQ_URI string) (*rabbit.Connection, error) {
	conn, err := rabbit.Dial(
		AMPQ_URI,
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
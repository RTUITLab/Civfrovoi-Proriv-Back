package app

import (
	"time"
	"fmt"
	"net"

	pb "github.com/RTUITLab/Civfrovoi-Proriv-Back/protos/coordservice"

	"github.com/RTUITLab/Civfrovoi-Proriv-Back/pkg/config"
	"github.com/RTUITLab/Civfrovoi-Proriv-Back/service/server"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Port 		string
	DB_URI		string
}

func New(
	cfg *config.Config,
) *App {
	return &App{
		Port: cfg.App.Port,
		DB_URI: cfg.DB.URI,
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

	Server := &server.Server{
		DB: db,
		GetCoordTime: time.Second,
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
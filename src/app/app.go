package app

import (
	pb "github.com/RTUITLab/Civfrovoi-Proriv-Back/protos/coordservice"
	"fmt"
	"net"

	"github.com/RTUITLab/Civfrovoi-Proriv-Back/pkg/config"
	"github.com/RTUITLab/Civfrovoi-Proriv-Back/service/server"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	db, err := gorm.Open(mysql.Open(a.DB_URI), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Server := &server.Server{
		DB: db,
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
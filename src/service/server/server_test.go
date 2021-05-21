package server_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/RTUITLab/Civfrovoi-Proriv-Back/pkg/models/coordinates"
	pb "github.com/RTUITLab/Civfrovoi-Proriv-Back/protos/coordservice"

	"github.com/RTUITLab/Civfrovoi-Proriv-Back/service/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var _server *server.Server
var client *pb.CoordsServiceClient

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}
	dsn, ok := os.LookupEnv("CP_DATABASE_URL")
	if !ok {
		panic("Not ok")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db

	_server = &server.Server{
		DB: db,
	}

	if err := coordinates.InitTable(db); err != nil {
		panic(err)
	}

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "8080"))
		if err != nil {
			log.Panicln(err)
		}

		grpcServer := grpc.NewServer()

		pb.RegisterCoordsServiceServer(
			grpcServer,
			_server,
		)

		if err := grpcServer.Serve(lis); err != nil {
			log.Panic(err)
		}
	}()
}

func TestFunc_Init(t *testing.T) {
	t.Log("Init")
}

func TestFunc_WriteCoords(t *testing.T) {
	_, err := _server.WriteCoords(
		context.Background(),
		&pb.WriteCoordsReq{
			Id:   "1srJx0G8gmTEO9GrldnSbkCPT0U",
			Lat:  12,
			Long: 15,
		},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

}

func TestFunc_InitApp(t *testing.T) {
	id, err := _server.InitApp(
		context.Background(),
		&pb.Empty{},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(id)

}

func TestFunc_GetCoords(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	defer conn.Close()

	client := pb.NewCoordsServiceClient(conn)
	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	cords, err := client.GetCoords(
		ctx,
		&pb.Units{
			Ids: []string{"1srIn1j3TryJBp2jdSmymjpkNjK", "1srJx0G8gmTEO9GrldnSbkCPT0U"},
		},
	)
	var count int
	for {
		count++
		c, err := cords.Recv()
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		t.Log(c.String())
		if count == 100 {
			break
		}
	}

	cancel()
	time.Sleep(2 * time.Second)
}

func TestFunc_UpdateUnits(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	defer conn.Close()

	client := pb.NewCoordsServiceClient(conn)

	units, err := client.UpdateUnits(
		context.Background(),
		&pb.UpdateUnitsReq{},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	for _, u := range units.Coords {
		t.Log(u.String())
	}
}

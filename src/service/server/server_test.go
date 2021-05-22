package server_test

import (
	"context"

	"testing"
	"time"

	"google.golang.org/grpc"

	pb "github.com/RTUITLab/Civfrovoi-Proriv-Back/protos/coordservice"

	"github.com/RTUITLab/Civfrovoi-Proriv-Back/service/server"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var _server *server.Server
var client *pb.CoordsServiceClient

// func init() {
// 	if err := godotenv.Load("../../.env"); err != nil {
// 		panic(err)
// 	}
// 	dsn, ok := os.LookupEnv("CP_DATABASE_URL")
// 	if !ok {
// 		panic("Not ok")
// 	}

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	DB = db

// 	_server = &server.Server{
// 		DB: db,
// 	}

// 	if err := coordinates.InitTable(db); err != nil {
// 		panic(err)
// 	}

// 	// go func() {
// 	// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "8080"))
// 	// 	if err != nil {
// 	// 		log.Panicln(err)
// 	// 	}

// 	// 	grpcServer := grpc.NewServer()

// 	// 	pb.RegisterCoordsServiceServer(
// 	// 		grpcServer,
// 	// 		_server,
// 	// 	)

// 	// 	if err := grpcServer.Serve(lis); err != nil {
// 	// 		log.Panic(err)
// 	// 	}
// 	// }()
// }

func TestFunc_Init(t *testing.T) {
	t.Log("Init")
	conn, err := grpc.Dial("82.146.61.131:8081", grpc.WithInsecure())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	defer conn.Close()

	client := pb.NewCoordsServiceClient(conn)

	id, err := client.InitApp(
		context.Background(),
		&pb.InitReq{
			Type: pb.Resuource_HUMAN,
		},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(id)
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
		&pb.InitReq{},
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
		t.Log("Recived")

		t.Log(c.String())
		if count == 100 {
			break
		}
	}

	cancel()
	time.Sleep(2 * time.Second)
}

func TestFunc_UpdateUnits(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
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

func TestFunc_SweepingUp(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	defer conn.Close()

	client := pb.NewCoordsServiceClient(conn)

	_, err = client.SweepingUp(
		context.Background(),
		&pb.OpeataionOn{
			Units: &pb.Units{
				Ids: []string{"1srIn1j3TryJBp2jdSmymjpkNjK", "1srJx0G8gmTEO9GrldnSbkCPT0U"},
			},
			Coords: &pb.Coords{
				Lat:  20,
				Long: 20,
			},
		},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

}

func TestFunc_ListenCommand(t *testing.T) {
	conn, err := grpc.Dial("82.146.61.131:8081", grpc.WithInsecure())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	defer conn.Close()

	client := pb.NewCoordsServiceClient(conn)

	go func() {
		for {
			client.ClearTemp(
				context.Background(),
				&pb.OpeataionOn{
					Units: &pb.Units{
						Ids: []string{"1srJx0G8gmTEO9GrldnSbkCPT0U"},
					},
					Coords: &pb.Coords{
						Lat:  10,
						Long: 20,
					},
				},
			)

			client.ExportSnowFromTemp(
				context.Background(),
				&pb.OperationFromTo{
					Units: &pb.Units{
						Ids: []string{"1srJx0G8gmTEO9GrldnSbkCPT0U"},
					},
					From: &pb.Coords{
						Lat:  9,
						Long: 17,
					},
					To: &pb.Coords{
						Lat:  7,
						Long: 12,
					},
				},
			)

			client.LoadingSnowFromTemp(
				context.Background(),
				&pb.OpeataionOn{
					Units: &pb.Units{
						Ids: []string{"1srJx0G8gmTEO9GrldnSbkCPT0U"},
					},
					Coords: &pb.Coords{
						Lat:  13,
						Long: 207,
					},
				},
			)

			client.MoveSnowToTemp(
				context.Background(),
				&pb.OperationFromTo{
					Units: &pb.Units{
						Ids: []string{"1srJx0G8gmTEO9GrldnSbkCPT0U"},
					},
					From: &pb.Coords{
						Lat:  9,
						Long: 17,
					},
					To: &pb.Coords{
						Lat:  7,
						Long: 12,
					},
				},
			)

			client.ReagentTreatment(
				context.Background(),
				&pb.OpeataionOn{
					Units: &pb.Units{
						Ids: []string{"1srJx0G8gmTEO9GrldnSbkCPT0U"},
					},
					Coords: &pb.Coords{
						Lat:  13,
						Long: 207,
					},
				},
			)

			client.ShaftFormation(
				context.Background(),
				&pb.OpeataionOn{
					Units: &pb.Units{
						Ids: []string{"1srJx0G8gmTEO9GrldnSbkCPT0U"},
					},
					Coords: &pb.Coords{
						Lat:  13,
						Long: 207,
					},
				},
			)

			client.SweepingUp(
				context.Background(),
				&pb.OpeataionOn{
					Units: &pb.Units{
						Ids: []string{"1srJx0G8gmTEO9GrldnSbkCPT0U"},
					},
					Coords: &pb.Coords{
						Lat:  13,
						Long: 207,
					},
				},
			)
		}
	}()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		25*time.Second,
	)
	defer cancel()

	cmds, err := client.ListenCommands(
		ctx,
		&pb.Unit{
			Id: "1srJx0G8gmTEO9GrldnSbkCPT0U",
		},
	)
	for {
		op, err := cmds.Recv()
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		t.Log(op.String())
	}

}

func TestFunc_OpenTask(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	defer conn.Close()

	client := pb.NewCoordsServiceClient(conn)

	_, err = client.OpenTask(
		context.Background(),
		&pb.Task{
			On: &pb.Coords{
				Lat:  20,
				Long: 15,
			},
			OperationName: "Почисти снег 41792",
		},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

}

func TestFunc_GetTasks(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	defer conn.Close()

	client := pb.NewCoordsServiceClient(conn)

	tasks, err := client.GetTask(
		context.Background(),
		&pb.Empty{},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(len(tasks.Tasks))

	for _, task := range tasks.Tasks {
		t.Log(task.String())
	}

}

package server

import (
	"context"
	"time"

	model "github.com/RTUITLab/Civfrovoi-Proriv-Back/pkg/models/coordinates"
	pb "github.com/RTUITLab/Civfrovoi-Proriv-Back/protos/coordservice"
	"github.com/segmentio/ksuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"github.com/streadway/amqp"
)

type Config struct {
	GetCoordsTime	string
}

type Server struct {
	pb.UnimplementedCoordsServiceServer
	DB *gorm.DB
	RabbitConn	*amqp.Connection
	GetCoordTime	time.Duration
}

func New() *Server {
	return &Server{}
}

func(s *Server)	WriteCoords(ctx context.Context, req *pb.WriteCoordsReq) (*pb.WriteCoordsResp, error) {
	logger := prepareEntrty("WriteCoords")
	Get := &model.Coordinates{
		ID: req.Id,
	}

	if err := s.DB.First(Get).Error; err == gorm.ErrRecordNotFound {
		return nil, status.Error(codes.Unauthenticated, "Your id is not exist")
	} else if err != nil {
		logger.Error("Failed to wrire coords: ", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	Get.DateTime = time.Now()
	Get.Lat = req.Lat
	Get.Long = req.Long

	if err := s.DB.Updates(Get).Error; err != nil {
		logger.Error("Failed to wrire coords: ", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.WriteCoordsResp{}, nil
}

func(s *Server)	GetCoords(req *pb.Units, stream pb.CoordsService_GetCoordsServer) error {
	logger := prepareEntrty("GetCoords")
	timer := time.NewTimer(s.GetCoordTime)
	for {
		select {
		case <- timer.C:
			timer.Stop()
			var Coords []*model.Coordinates
			if err := s.DB.Find(&Coords, req.Ids).Error; err == gorm.ErrRecordNotFound {
				return status.Error(codes.InvalidArgument, "This ids is not exsist")
			}
			resp := &pb.GetCoordsResp{}
			for _, c := range Coords {
				resp.Coords = append(
					resp.Coords, 
					&pb.UnitWithCords{
						Id: c.ID,
						Coords: &pb.Coords{
							Lat: c.Lat,
							Long: c.Long,
						},
					},
				)
			}
			err := stream.Send(
				resp,
			)

			if stats, ok := status.FromError(err); ok && stats.Code() == codes.Canceled {
				logger.Info("Canceled")
				return nil
			} else if err != nil {
				logger.Error("Error on streaming: ", err)
				return status.Error(
					codes.Internal,
					err.Error(),
				)
			}
			timer.Reset(s.GetCoordTime)
		}
		
	}
}

func(s *Server) UpdateUnits(ctx context.Context, req *pb.UpdateUnitsReq) (*pb.UpdateUnitsResp, error) {
	logger := prepareEntrty("UpdateUnits")
	var Coor []*model.Coordinates

	if err := s.DB.Find(&Coor).Error; err == gorm.ErrRecordNotFound {
		// Pass
	} else if err != nil {
		logger.Error("Faield to updated units", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.UpdateUnitsResp{}

	for _, c := range Coor {
		resp.Coords = append(resp.Coords, c.ToPB())
	}

	return resp, nil
}
func(s *Server)	SweepingUp(ctx context.Context, req *pb.OpeataionOn) (*pb.Empty, error){
	panic("Not imolemented")
}
func(s *Server)	ReagentTreatment(ctx context.Context, req *pb.OpeataionOn) (*pb.Empty, error){
	panic("Not imolemented")
}
func(s *Server)	ShaftFormation(ctx context.Context, req *pb.OpeataionOn) (*pb.Empty, error){
	panic("Not imolemented")
}
func(s *Server)	MoveSnowToTemp(ctx context.Context, req *pb.OperationFromTo) (*pb.Empty, error) {
	panic("Not imolemented")
}
func(s *Server)	LoadingSnowFromTemp(ctx context.Context, req *pb.OpeataionOn) (*pb.Empty, error) {
	panic("Not imolemented")
}
func(s *Server)	ExportSnowFromTemp(ctx context.Context, req  *pb.OperationFromTo) (*pb.Empty, error) {
	panic("Not imolemented")
}
func(s *Server)	ClearTemp(ctx context.Context, req *pb.OpeataionOn) (*pb.Empty, error) {
	panic("Not imolemented")
}

func(s *Server)	ListenCommands(req *pb.Unit, resp pb.CoordsService_ListenCommandsServer) error{
	panic("Not imolemented")
}

func(s *Server) InitApp(ctx context.Context, req *pb.InitReq) (*pb.ID, error) {
	logger := prepareEntrty("InitApp")
	uuid := ksuid.New()

	c := &model.Coordinates{
		ID: uuid.String(),
		DateTime: time.Now(),
		Lat: 0,
		Long: 0,
		Resource: req.Type,
	}

	if err := s.DB.Create(c).Error; err != nil {
		logger.Error("Faield to init App: ", err)
		return nil, status.Error(codes.Internal, "faield to get")
	}

	return &pb.ID{Id: uuid.String()}, nil
}

func prepareEntrty(
	Method 	string,
) *log.Entry {
	return log.WithFields(
		log.Fields{
			"method": Method,
		},
	)
}
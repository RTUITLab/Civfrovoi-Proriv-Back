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
)

type Server struct {
	pb.UnimplementedCoordsServiceServer
	DB *gorm.DB
}

func(s *Server)	WriteCoords(ctx context.Context, req *pb.WriteCoordsReq) (*pb.WriteCoordsResp, error) {
	Get := &model.Coordinates{
		ID: req.Id,
	}

	if err := s.DB.First(Get).Error; err == gorm.ErrRecordNotFound {
		return nil, status.Error(codes.Unauthenticated, "Your id is not exist")
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	Get.DateTime = time.Now()
	Get.Lat = req.Lat
	Get.Long = req.Long

	if err := s.DB.Updates(Get).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.WriteCoordsResp{}, nil
}

func(s *Server)	GetCoords(req *pb.Units, stream pb.CoordsService_GetCoordsServer) error {
	for {
		var Coords []*model.Coordinates
		if err := s.DB.Find(&Coords, req.Ids).Error; err == gorm.ErrRecordNotFound {
			return status.Error(codes.InvalidArgument, "This ids is not exsist")
		}
		
		for _, c := range Coords {
			err := stream.Send(
				&pb.GetCoordsResp{
					Id: c.ID,
					Lat: c.Lat,
					Long: c.Long,
				},
			)

			if stats, ok := status.FromError(err); ok && stats.Code() == codes.Canceled {
				log.Info("Canceled")
				return nil
			} else if err != nil {
				log.Error("Err on streaming")
				return status.Error(
					codes.Internal,
					err.Error(),
				)
			}
		}
	}
}

func(s *Server) UpdateUnits(ctx context.Context, req *pb.UpdateUnitsReq) (*pb.UpdateUnitsResp, error) {
	var Coor []*model.Coordinates

	if err := s.DB.Find(&Coor).Error; err == gorm.ErrRecordNotFound {
		// Pass
	} else if err != nil {
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

func(s *Server)	ListenCommands(*pb.Unit, pb.CoordsService_ListenCommandsServer) error{
	panic("Not imolemented")
}

func(s *Server) InitApp(ctx context.Context, req *pb.Empty) (*pb.ID, error) {
	uuid := ksuid.New()

	c := &model.Coordinates{
		ID: uuid.String(),
		DateTime: time.Now(),
		Lat: 0,
		Long: 0,
	}

	if err := s.DB.Create(c).Error; err != nil {
		log.Error("Faield to init App")
		return nil, status.Error(codes.Internal, "faield to get")
	}

	return &pb.ID{Id: uuid.String()}, nil
}
package server

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	model "github.com/RTUITLab/Civfrovoi-Proriv-Back/pkg/models/coordinates"
	pb "github.com/RTUITLab/Civfrovoi-Proriv-Back/protos/coordservice"
	"github.com/segmentio/ksuid"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm" 
)

const (
	DISPETCHER_QUEUE = "dispetcher_queue"
)

type Config struct {
	GetCoordsTime	string
}

type Server struct {
	pb.UnimplementedCoordsServiceServer
	DB 				*gorm.DB
	RabbitConn		*amqp.Connection
	GetCoordTime	time.Duration
}

func New(
	DB 				*gorm.DB,
	RabbitConn		*amqp.Connection,
	GetCoordTime	time.Duration,
) *Server {
	
	s := &Server{
		DB: DB,
		RabbitConn: RabbitConn,
		GetCoordTime: GetCoordTime,
	}

	if err := s.createQueueForID(DISPETCHER_QUEUE); err != nil {
		log.Error("Failed to start Queue for dispethcer")
	}

	return s
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
	logger := prepareEntrty("SweepingUp")
	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error("Faield to open channle on rabbit mq", err)
		return nil, status.Error(codes.Internal, "Faield to send operation")
	}
	defer channel.Close()

	wg := sync.WaitGroup{}
	for _, u := range req.Units.Ids {
		wg.Add(1)
		go func(id string, wg *sync.WaitGroup) {
			defer wg.Done()
			q, err := channel.QueueDeclare(
				id,
				true,
				false,
				false,
				true,
				nil,	
			)
			if err != nil {
				logger.Errorf("Failed to open queue for %s: %v", id, err)
				return
			}
			ops := &pb.Operaions{
				OperationName: "Sweeping Up",
				From: req.Coords,
			}
			data, err := json.Marshal(ops)
			if err != nil {
				logger.Error("Failed to marshall: ", err)
				return
			}

			err = channel.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body: data,
				},
			)
			if err != nil {
				logger.Error("Failed to send to AMQP: ", err)
			}
		}(u, &wg)
	}
	wg.Wait()

	return &pb.Empty{}, status.Error(codes.OK, "")
}
func(s *Server)	ReagentTreatment(ctx context.Context, req *pb.OpeataionOn) (*pb.Empty, error){
	logger := prepareEntrty("ReagentTreatment")
	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error("Faield to open channle on rabbit mq", err)
		return nil, status.Error(codes.Internal, "Faield to send operation")
	}
	defer channel.Close()

	wg := sync.WaitGroup{}
	for _, u := range req.Units.Ids {
		wg.Add(1)
		go func(id string, wg *sync.WaitGroup) {
			defer wg.Done()
			q, err := channel.QueueDeclare(
				id,
				true,
				false,
				false,
				true,
				nil,	
			)
			if err != nil {
				logger.Errorf("Failed to open queue for %s: %v", id, err)
				return
			}
			ops := &pb.Operaions{
				OperationName: "Reagent Treatment",
				From: req.Coords,
			}
			data, err := json.Marshal(ops)
			if err != nil {
				logger.Error("Failed to marshall: ", err)
				return
			}

			err = channel.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body: data,
				},
			)
			if err != nil {
				logger.Error("Failed to send to AMQP: ", err)
			}
		}(u, &wg)
	}
	wg.Wait()

	return &pb.Empty{}, status.Error(codes.OK, "")
}
func(s *Server)	ShaftFormation(ctx context.Context, req *pb.OpeataionOn) (*pb.Empty, error){
	logger := prepareEntrty("ShaftFormation")
	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error("Faield to open channle on rabbit mq", err)
		return nil, status.Error(codes.Internal, "Faield to send operation")
	}
	defer channel.Close()

	wg := sync.WaitGroup{}
	for _, u := range req.Units.Ids {
		wg.Add(1)
		go func(id string, wg *sync.WaitGroup) {
			defer wg.Done()
			q, err := channel.QueueDeclare(
				id,
				true,
				false,
				false,
				true,
				nil,	
			)
			if err != nil {
				logger.Errorf("Failed to open queue for %s: %v", id, err)
				return
			}
			ops := &pb.Operaions{
				OperationName: "Shaft Formation",
				From: req.Coords,
			}
			data, err := json.Marshal(ops)
			if err != nil {
				logger.Error("Failed to marshall: ", err)
				return
			}

			err = channel.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body: data,
				},
			)
			if err != nil {
				logger.Error("Failed to send to AMQP: ", err)
			}
		}(u, &wg)
	}
	wg.Wait()

	return &pb.Empty{}, status.Error(codes.OK, "")
}
func(s *Server)	MoveSnowToTemp(ctx context.Context, req *pb.OperationFromTo) (*pb.Empty, error) {
	logger := prepareEntrty("MoveSnowToTemp")
	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error("Faield to open channle on rabbit mq", err)
		return nil, status.Error(codes.Internal, "Faield to send operation")
	}
	defer channel.Close()

	wg := sync.WaitGroup{}
	for _, u := range req.Units.Ids {
		wg.Add(1)
		go func(id string, wg *sync.WaitGroup) {
			defer wg.Done()
			q, err := channel.QueueDeclare(
				id,
				true,
				false,
				false,
				true,
				nil,	
			)
			if err != nil {
				logger.Errorf("Failed to open queue for %s: %v", id, err)
				return
			}
			ops := &pb.Operaions{
				OperationName: "Move Snow To Temp",
				From: req.From,
				To: req.To,
			}
			data, err := json.Marshal(ops)
			if err != nil {
				logger.Error("Failed to marshall: ", err)
				return
			}

			err = channel.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body: data,
				},
			)
			if err != nil {
				logger.Error("Failed to send to AMQP: ", err)
			}
		}(u, &wg)
	}
	wg.Wait()

	return &pb.Empty{}, status.Error(codes.OK, "")
}
func(s *Server)	LoadingSnowFromTemp(ctx context.Context, req *pb.OpeataionOn) (*pb.Empty, error) {
	logger := prepareEntrty("LoadingSnowFromTemp")
	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error("Faield to open channle on rabbit mq", err)
		return nil, status.Error(codes.Internal, "Faield to send operation")
	}
	defer channel.Close()

	wg := sync.WaitGroup{}
	for _, u := range req.Units.Ids {
		wg.Add(1)
		go func(id string, wg *sync.WaitGroup) {
			defer wg.Done()
			q, err := channel.QueueDeclare(
				id,
				true,
				false,
				false,
				true,
				nil,	
			)
			if err != nil {
				logger.Errorf("Failed to open queue for %s: %v", id, err)
				return
			}
			ops := &pb.Operaions{
				OperationName: "Loading Snow From Temp",
				From: req.Coords,
			}
			data, err := json.Marshal(ops)
			if err != nil {
				logger.Error("Failed to marshall: ", err)
				return
			}

			err = channel.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body: data,
				},
			)
			if err != nil {
				logger.Error("Failed to send to AMQP: ", err)
			}
		}(u, &wg)
	}
	wg.Wait()

	return &pb.Empty{}, status.Error(codes.OK, "")
}
func(s *Server)	ExportSnowFromTemp(ctx context.Context, req  *pb.OperationFromTo) (*pb.Empty, error) {
	logger := prepareEntrty("ExportSnowFromTemp")
	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error("Faield to open channle on rabbit mq", err)
		return nil, status.Error(codes.Internal, "Faield to send operation")
	}
	defer channel.Close()

	wg := sync.WaitGroup{}
	for _, u := range req.Units.Ids {
		wg.Add(1)
		go func(id string, wg *sync.WaitGroup) {
			defer wg.Done()
			q, err := channel.QueueDeclare(
				id,
				true,
				false,
				false,
				true,
				nil,	
			)
			if err != nil {
				logger.Errorf("Failed to open queue for %s: %v", id, err)
				return
			}
			ops := &pb.Operaions{
				OperationName: "Export Snow From Temp",
				From: req.From,
				To: req.To,
			}
			data, err := json.Marshal(ops)
			if err != nil {
				logger.Error("Failed to marshall: ", err)
				return
			}

			err = channel.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body: data,
				},
			)
			if err != nil {
				logger.Error("Failed to send to AMQP: ", err)
			}
		}(u, &wg)
	}
	wg.Wait()

	return &pb.Empty{}, status.Error(codes.OK, "")
}
func(s *Server)	ClearTemp(ctx context.Context, req *pb.OpeataionOn) (*pb.Empty, error) {
	logger := prepareEntrty("ClearTemp")
	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error("Faield to open channle on rabbit mq", err)
		return nil, status.Error(codes.Internal, "Faield to send operation")
	}
	defer channel.Close()

	wg := sync.WaitGroup{}
	for _, u := range req.Units.Ids {
		wg.Add(1)
		go func(id string, wg *sync.WaitGroup) {
			defer wg.Done()
			q, err := channel.QueueDeclare(
				id,
				true,
				false,
				false,
				true,
				nil,	
			)
			if err != nil {
				logger.Errorf("Failed to open queue for %s: %v", id, err)
				return
			}
			ops := &pb.Operaions{
				OperationName: "Clear Temp",
				From: req.Coords,
			}
			data, err := json.Marshal(ops)
			if err != nil {
				logger.Error("Failed to marshall: ", err)
				return
			}

			err = channel.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body: data,
				},
			)
			if err != nil {
				logger.Error("Failed to send to AMQP: ", err)
			}
		}(u, &wg)
	}
	wg.Wait()

	return &pb.Empty{}, status.Error(codes.OK, "")
}

func(s *Server)	ListenCommands(req *pb.Unit, resp pb.CoordsService_ListenCommandsServer) error{
	logger := prepareEntrty("ListenCommands")
	ch, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error("Faield to get chan: ", err)
		return status.Error(codes.Internal, "")
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		req.Id,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logger.Error("Faield to consume: ", err)
		return status.Error(codes.Internal, err.Error())
	}

	for d := range msgs {
		ops := &pb.Operaions{}
		err := json.Unmarshal(d.Body, ops)
		if err != nil {
			logger.Error("Faield to unmarhsall: ", err)
			// TODO think about error handling
			continue
		}

		err = resp.Send(ops)
		if stats, ok := status.FromError(err); ok && stats.Code() == codes.Canceled {
			logger.Info("Cancelen")
			return nil
		} else if ok && stats.Code() == codes.DeadlineExceeded {
			logger.Info("Deadline")
			return nil
		}else if err != nil {
			logger.Error("Error on streaming", err)
			return status.Error(codes.Aborted, "")
		}
	}

	return nil
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

func (s *Server) OpenTask(ctx context.Context, req *pb.Task) (*pb.Empty, error) {
	logger := prepareEntrty("OpenTask")

	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error(err)
		return nil, status.Error(codes.Internal, "Faield to open task") 
	}
	defer channel.Close()
	
	data, err := json.Marshal(req)
	if err != nil {
		logger.Error("Faield to marshall messgae ", err)
		return nil, status.Error(codes.Internal, "Faield to marshall messgae")
	}

	err = channel.Publish(
		"",
		DISPETCHER_QUEUE,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: data,
		},
	)
	if err != nil {
		logger.Error("Failed to publish ", err)
		return nil, status.Error(codes.Internal, "Faield to publish task")
	}

	return &pb.Empty{}, nil
}

func (s *Server) GetTask(ctx context.Context, req *pb.Empty) (*pb.Tasks, error) {
	logger := prepareEntrty("OpenTask")
	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error(err)
		return nil, status.Error(codes.Internal, "Faield to open task") 
	}
	defer channel.Close()

	msgs, err := channel.Consume(
		DISPETCHER_QUEUE,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	tasks := []*pb.Task{}
	for i := 0; i < 10; i++ {
		select {
		case msg := <- msgs:
			task := &pb.Task{}
			if err := json.Unmarshal(msg.Body, task); err != nil {
				logger.Error("Faield to unmarhsal: ", err)
				continue
			}
			tasks = append(tasks, task)
		case <- time.After(30*time.Millisecond):
			logger.Info("Time after")
			continue
		}
	}

	return &pb.Tasks{Tasks: tasks}, nil
}

func (s *Server) createQueueForID(id string) error {
	logger := prepareEntrty("createQueueForID")
	channel, err := s.RabbitConn.Channel()
	if err != nil {
		logger.Error("Faield to create Queue: ", err)
		return err
	}
	defer channel.Close()

	_, err = channel.QueueDeclare(
		id,
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		logger.Error("Faield to create Queue: ", err)
		return err
	}
	
	return nil
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
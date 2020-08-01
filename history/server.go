package history

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/history/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	proto.RegisterHistoryServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) AddHistory(ctx context.Context, r *proto.AddHistoryRequest) (*proto.AddHistoryResponse, error) {
	err := s.service.AddHistory(ctx, r.UserId, r.DiseaseId, r.Note)
	if err != nil {
		return nil, err
	}
	return &proto.AddHistoryResponse{History: &proto.History{
		UserId:    r.UserId,
		DiseaseId: r.DiseaseId,
		Note:      r.Note,
	}}, nil
}

func (s *grpcServer) GetHistoryByID(ctx context.Context, r *proto.GetHistoryByIDRequest) (*proto.GetHistoryByIDResponse, error) {
	a, err := s.service.GetHistoryByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &proto.GetHistoryByIDResponse{
		History: &proto.History{
			Id:        a.ID,
			UserId:    a.UserID,
			DiseaseId: a.DiseaseID,
			Note:      a.Note,
			CreatedAt: a.CreatedAt,
		},
	}, nil
}

func (s *grpcServer) GetHistories(ctx context.Context, _ *proto.GetHistoryRequest) (*proto.GetHistoryResponse, error) {
	res, err := s.service.GetHistories(ctx)
	if err != nil {
		return nil, err
	}
	var genders []*proto.History
	for _, p := range res {
		genders = append(
			genders,
			&proto.History{
				Id:        p.ID,
				UserId:    p.UserID,
				DiseaseId: p.DiseaseID,
				Note:      p.Note,
				CreatedAt: p.CreatedAt,
			},
		)
	}
	return &proto.GetHistoryResponse{Histories: genders}, nil
}

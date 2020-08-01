package blood

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/blood/proto"
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
	proto.RegisterBloodServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) AddBlood(ctx context.Context, r *proto.AddBloodRequest) (*proto.AddBloodResponse, error) {
	err := s.service.AddBlood(ctx, r.Name, r.Description)
	if err != nil {
		return nil, err
	}
	return &proto.AddBloodResponse{Blood: &proto.Blood{
		Name:        r.Name,
		Description: r.Description,
	}}, nil
}

func (s *grpcServer) GetBloodByID(ctx context.Context, r *proto.GetBloodByIDRequest) (*proto.GetBloodByIDResponse, error) {
	a, err := s.service.GetBloodByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &proto.GetBloodByIDResponse{
		Blood: &proto.Blood{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetBloods(ctx context.Context, _ *proto.GetBloodRequest) (*proto.GetBloodResponse, error) {
	res, err := s.service.GetBloods(ctx)
	if err != nil {
		return nil, err
	}
	var bloods []*proto.Blood
	for _, p := range res {
		bloods = append(
			bloods,
			&proto.Blood{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}
	return &proto.GetBloodResponse{Bloods: bloods}, nil
}

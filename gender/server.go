package gender

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/gender/proto"
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
	proto.RegisterGenderServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) AddGender(ctx context.Context, r *proto.AddGenderRequest) (*proto.AddGenderResponse, error) {
	err := s.service.AddGender(ctx, r.Name, r.Description)
	if err != nil {
		return nil, err
	}
	return &proto.AddGenderResponse{Gender: &proto.Gender{
		Name:        r.Name,
		Description: r.Description,
	}}, nil
}

func (s *grpcServer) GetGenderByID(ctx context.Context, r *proto.GetGenderByIDRequest) (*proto.GetGenderByIDResponse, error) {
	a, err := s.service.GetGenderByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &proto.GetGenderByIDResponse{
		Gender: &proto.Gender{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetGenders(ctx context.Context, _ *proto.GetGenderRequest) (*proto.GetGenderResponse, error) {
	res, err := s.service.GetGenders(ctx)
	if err != nil {
		return nil, err
	}
	var genders []*proto.Gender
	for _, p := range res {
		genders = append(
			genders,
			&proto.Gender{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}
	return &proto.GetGenderResponse{Genders: genders}, nil
}

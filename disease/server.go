package disease

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/disease/proto"
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
	proto.RegisterDiseaseServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) AddDisease(ctx context.Context, r *proto.AddDiseaseRequest) (*proto.AddDiseaseResponse, error) {
	err := s.service.AddDisease(ctx, r.Name, r.Description)
	if err != nil {
		return nil, err
	}
	return &proto.AddDiseaseResponse{Disease: &proto.Disease{
		Name:        r.Name,
		Description: r.Description,
	}}, nil
}

func (s *grpcServer) GetDiseaseByID(ctx context.Context, r *proto.GetDiseaseByIDRequest) (*proto.GetDiseaseByIDResponse, error) {
	a, err := s.service.GetDiseaseByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &proto.GetDiseaseByIDResponse{
		Disease: &proto.Disease{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetDiseases(ctx context.Context, r *proto.GetDiseaseRequest) (*proto.GetDiseaseResponse, error) {

	var res []Disease
	var err error
	if r.Query != "" {
		res, err = s.service.SearchDiseases(ctx, r.Query)
	} else {
		res, err = s.service.GetDiseases(ctx)
	}
	if err != nil {
		return nil, err
	}
	var genders []*proto.Disease
	for _, p := range res {
		genders = append(
			genders,
			&proto.Disease{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}
	return &proto.GetDiseaseResponse{Diseases: genders}, nil
}

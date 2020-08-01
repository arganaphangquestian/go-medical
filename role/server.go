package role

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/role/proto"
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
	proto.RegisterRoleServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) AddRole(ctx context.Context, r *proto.AddRoleRequest) (*proto.AddRoleResponse, error) {
	err := s.service.AddRole(ctx, r.Name, r.Description)
	if err != nil {
		return nil, err
	}
	return &proto.AddRoleResponse{Role: &proto.Role{
		Name:        r.Name,
		Description: r.Description,
	}}, nil
}

func (s *grpcServer) GetRoleByID(ctx context.Context, r *proto.GetRoleByIDRequest) (*proto.GetRoleByIDResponse, error) {
	a, err := s.service.GetRoleByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &proto.GetRoleByIDResponse{
		Role: &proto.Role{
			Id:   a.ID,
			Name: a.Name,
		},
	}, nil
}

func (s *grpcServer) GetRoles(ctx context.Context, _ *proto.GetRoleRequest) (*proto.GetRoleResponse, error) {
	res, err := s.service.GetRoles(ctx)
	if err != nil {
		return nil, err
	}
	var roles []*proto.Role
	for _, p := range res {
		roles = append(
			roles,
			&proto.Role{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}
	return &proto.GetRoleResponse{Roles: roles}, nil
}

package user

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/user/proto"
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
	proto.RegisterUserServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) AddUser(ctx context.Context, r *proto.AddUserRequest) (*proto.AddUserResponse, error) {
	err := s.service.AddUser(
		ctx,
		r.Name,
		r.Email,
		r.Address,
		r.RoleId,
		r.GenderId,
		r.BloodId,
		r.BirthOfDate,
		r.Contact,
	)
	if err != nil {
		return nil, err
	}
	return &proto.AddUserResponse{User: &proto.User{
		Name:        r.Name,
		Email:       r.Email,
		Address:     r.Address,
		RoleId:      r.RoleId,
		GenderId:    r.GenderId,
		BloodId:     r.BloodId,
		BirthOfDate: r.BirthOfDate,
		Contact:     r.Contact,
	}}, nil
}

func (s *grpcServer) GetUserByID(ctx context.Context, r *proto.GetUserByIDRequest) (*proto.GetUserByIDResponse, error) {
	a, err := s.service.GetUserByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &proto.GetUserByIDResponse{
		User: &proto.User{
			Id:          a.ID.String(),
			Name:        a.Name,
			Email:       a.Email,
			Address:     a.Address,
			RoleId:      a.RoleID,
			GenderId:    a.GenderID,
			BloodId:     a.BloodID,
			BirthOfDate: a.BirthOfDate,
			Contact:     a.Contact,
		},
	}, nil
}

func (s *grpcServer) GetUsers(ctx context.Context, _ *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	res, err := s.service.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	var roles []*proto.User
	for _, p := range res {
		roles = append(
			roles,
			&proto.User{
				Id:          p.ID.String(),
				Name:        p.Name,
				Email:       p.Email,
				Address:     p.Address,
				RoleId:      p.RoleID,
				GenderId:    p.GenderID,
				BloodId:     p.BloodID,
				BirthOfDate: p.BirthOfDate,
				Contact:     p.Contact,
			},
		)
	}
	return &proto.GetUserResponse{Users: roles}, nil
}

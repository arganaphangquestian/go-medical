package main

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/gateway/graph/model"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateBlood(ctx context.Context, blood *model.InputBlood) (*model.Blood, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateGender(ctx context.Context, gender *model.InputGender) (*model.Gender, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateHistory(ctx context.Context, history *model.InputHistory) (*model.History, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRole(ctx context.Context, role *model.InputRole) (*model.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, user *model.InputUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDisease(ctx context.Context, disease *model.InputDisease) (*model.Disease, error) {
	panic(fmt.Errorf("not implemented"))
}

package main

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/gateway/graph/model"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Blood(ctx context.Context, id *string) ([]*model.Blood, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Gender(ctx context.Context, id *string) ([]*model.Gender, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) History(ctx context.Context, id *string) ([]*model.History, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Role(ctx context.Context, id *string) ([]*model.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id *string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Disease(ctx context.Context, query *string, id *string) ([]*model.Disease, error) {
	panic(fmt.Errorf("not implemented"))
}

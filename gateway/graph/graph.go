package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/arganaphangquestian/go-medical/blood"
	"github.com/arganaphangquestian/go-medical/disease"
	"github.com/arganaphangquestian/go-medical/gateway/graph/generated"
	"github.com/arganaphangquestian/go-medical/gender"
	"github.com/arganaphangquestian/go-medical/history"
	"github.com/arganaphangquestian/go-medical/role"
	"github.com/arganaphangquestian/go-medical/user"
)

type Server struct {
	bloodClient   *blood.Client
	diseaseClient *disease.Client
	genderClient  *gender.Client
	historyClient *history.Client
	roleClient    *role.Client
	userClient    *user.Client
}

func NewGraphQLServer(bloodUrl, diseaseUrl, genderUrl, historyUrl, roleUrl, userUrl string) (*Server, error) {
	bloodClient, err := blood.NewClient(bloodUrl)
	if err != nil {
		return nil, err
	}
	diseaseClient, err := disease.NewClient(diseaseUrl)
	if err != nil {
		return nil, err
	}
	genderClient, err := gender.NewClient(genderUrl)
	if err != nil {
		return nil, err
	}
	roleClient, err := role.NewClient(roleUrl)
	if err != nil {
		return nil, err
	}
	userClient, err := user.NewClient(userUrl)
	if err != nil {
		return nil, err
	}
	historyClient, err := history.NewClient(historyUrl)
	if err != nil {
		return nil, err
	}
	return &Server{
		bloodClient:   bloodClient,
		diseaseClient: diseaseClient,
		genderClient:  genderClient,
		historyClient: historyClient,
		roleClient:    roleClient,
		userClient:    userClient,
	}, nil
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: s,
	})
}

func (s *Server) Mutation() generated.MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() generated.QueryResolver {
	return &queryResolver{
		server: s,
	}
}

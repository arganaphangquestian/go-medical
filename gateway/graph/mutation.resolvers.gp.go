package main

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/gateway/graph/model"
	"log"
	"strconv"
	"time"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateBlood(ctx context.Context, blood *model.InputBlood) (*model.Blood, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	a, err := r.server.bloodClient.AddBlood(ctx, blood.Name, blood.Rhesus, *blood.Description)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.Blood{
		ID:          a.ID,
		Name:        a.Name,
		Rhesus:      a.Rhesus,
		Description: &a.Description,
		CreatedAt:   &a.CreatedAt,
	}, nil
}

func (r *mutationResolver) CreateGender(ctx context.Context, gender *model.InputGender) (*model.Gender, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	a, err := r.server.genderClient.AddGender(ctx, gender.Name, *gender.Description)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.Gender{
		ID:          a.ID,
		Name:        a.Name,
		Description: &a.Description,
		CreatedAt:   &a.CreatedAt,
	}, nil
}

func (r *mutationResolver) CreateHistory(ctx context.Context, history *model.InputHistory) (*model.History, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	a, err := r.server.historyClient.AddHistory(ctx, history.UserID, history.DiseaseID, *history.Note)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.History{
		ID:        a.ID,
		UserID:    a.UserID,
		DiseaseID: a.DiseaseID,
		Note:      &a.Note,
		CreatedAt: &a.CreatedAt,
	}, nil
}

func (r *mutationResolver) CreateRole(ctx context.Context, role *model.InputRole) (*model.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	a, err := r.server.roleClient.AddRole(ctx, role.Name, *role.Description)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.Role{
		ID:          a.ID,
		Name:        a.Name,
		Description: &a.Description,
		CreatedAt:   &a.CreatedAt,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, user *model.InputUser) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	roleID, err := strconv.ParseUint(user.RoleID, 10, 32)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	genderID, err := strconv.ParseUint(user.GenderID, 10, 32)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bloodID, err := strconv.ParseUint(user.GenderID, 10, 32)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	a, err := r.server.userClient.AddUser(ctx, user.Name, user.Email, user.Address, uint32(roleID), uint32(genderID), uint32(bloodID), *user.BirthOfDate, *user.Contact)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.User{
		ID:          a.ID.String(),
		Name:        a.Name,
		Email:       a.Email,
		Address:     a.Address,
		RoleID:      fmt.Sprint(a.RoleID),
		GenderID:    fmt.Sprint(a.GenderID),
		BloodID:     fmt.Sprint(a.BloodID),
		BirthOfDate: &a.BirthOfDate,
		Contact:     &a.Contact,
	}, nil
}

func (r *mutationResolver) CreateDisease(ctx context.Context, disease *model.InputDisease) (*model.Disease, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	a, err := r.server.diseaseClient.AddDisease(ctx, disease.Name, *disease.Description)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.Disease{
		ID:          a.ID,
		Name:        a.Name,
		Description: &a.Description,
		CreatedAt:   &a.CreatedAt,
	}, nil
}

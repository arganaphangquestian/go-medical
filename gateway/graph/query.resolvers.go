package main

import (
	"context"
	"fmt"
	"github.com/arganaphangquestian/go-medical/gateway/graph/model"
	"log"
	"time"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Blood(ctx context.Context, id *string) ([]*model.Blood, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if id != nil {
		r, err := r.server.bloodClient.GetBloodByID(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*model.Blood{{
			ID:          r.ID,
			Name:        r.Name,
			Rhesus:      r.Rhesus,
			Description: &r.Description,
			CreatedAt:   &r.CreatedAt,
		}}, nil
	}
	bloodList, err := r.server.bloodClient.GetBloods(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var bloods []*model.Blood
	for _, a := range bloodList {
		blood := &model.Blood{
			ID:          a.ID,
			Name:        a.Name,
			Rhesus:      a.Rhesus,
			Description: &a.Description,
			CreatedAt:   &a.CreatedAt,
		}
		bloods = append(bloods, blood)
	}

	return bloods, nil
}

func (r *queryResolver) Gender(ctx context.Context, id *string) ([]*model.Gender, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if id != nil {
		r, err := r.server.genderClient.GetGenderByID(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*model.Gender{{
			ID:          r.ID,
			Name:        r.Name,
			Description: &r.Description,
			CreatedAt:   &r.CreatedAt,
		}}, nil
	}
	genderList, err := r.server.genderClient.GetGenders(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var genders []*model.Gender
	for _, a := range genderList {
		gender := &model.Gender{
			ID:          a.ID,
			Name:        a.Name,
			Description: &a.Description,
			CreatedAt:   &a.CreatedAt,
		}
		genders = append(genders, gender)
	}

	return genders, nil
}

func (r *queryResolver) History(ctx context.Context, id *string) ([]*model.History, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if id != nil {
		r, err := r.server.historyClient.GetHistoryByID(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*model.History{{
			ID:        r.ID,
			UserID:    r.UserID,
			DiseaseID: r.DiseaseID,
			Note:      &r.Note,
			CreatedAt: &r.CreatedAt,
		}}, nil
	}
	historyList, err := r.server.historyClient.GetHistories(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var histories []*model.History
	for _, a := range historyList {
		history := &model.History{
			ID:        a.ID,
			UserID:    a.UserID,
			DiseaseID: a.DiseaseID,
			Note:      &a.Note,
			CreatedAt: &a.CreatedAt,
		}
		histories = append(histories, history)
	}

	return histories, nil
}

func (r *queryResolver) Role(ctx context.Context, id *string) ([]*model.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if id != nil {
		r, err := r.server.roleClient.GetRoleByID(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*model.Role{{
			ID:          r.ID,
			Name:        r.Name,
			Description: &r.Description,
			CreatedAt:   &r.CreatedAt,
		}}, nil
	}
	roleList, err := r.server.roleClient.GetRoles(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var roles []*model.Role
	for _, a := range roleList {
		role := &model.Role{
			ID:          a.ID,
			Name:        a.Name,
			Description: &a.Description,
			CreatedAt:   &a.CreatedAt,
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *queryResolver) User(ctx context.Context, id *string) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if id != nil {
		r, err := r.server.userClient.GetUserByID(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*model.User{{
			ID:          r.ID.String(),
			Name:        r.Name,
			Email:       r.Email,
			Address:     r.Address,
			RoleID:      fmt.Sprint(r.RoleID),
			GenderID:    fmt.Sprint(r.GenderID),
			BloodID:     fmt.Sprint(r.BloodID),
			BirthOfDate: &r.BirthOfDate,
			Contact:     &r.Contact,
		}}, nil
	}
	userList, err := r.server.userClient.GetUsers(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var users []*model.User
	for _, a := range userList {
		user := &model.User{
			ID:          a.ID.String(),
			Name:        a.Name,
			Email:       a.Email,
			Address:     a.Address,
			RoleID:      fmt.Sprint(a.RoleID),
			GenderID:    fmt.Sprint(a.GenderID),
			BloodID:     fmt.Sprint(a.BloodID),
			BirthOfDate: &a.BirthOfDate,
			Contact:     &a.Contact,
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *queryResolver) Disease(ctx context.Context, query *string, id *string) ([]*model.Disease, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if id != nil {
		r, err := r.server.diseaseClient.GetDiseaseByID(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*model.Disease{{
			ID:          r.ID,
			Name:        r.Name,
			Description: &r.Description,
			CreatedAt:   &r.CreatedAt,
		}}, nil
	}
	q := ""
	if query != nil {
		q = *query
	}
	diseaseList, err := r.server.diseaseClient.GetDiseases(ctx, q)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var diseases []*model.Disease
	for _, a := range diseaseList {
		disease := &model.Disease{
			ID:          a.ID,
			Name:        a.Name,
			Description: &a.Description,
			CreatedAt:   &a.CreatedAt,
		}
		diseases = append(diseases, disease)
	}

	return diseases, nil
}

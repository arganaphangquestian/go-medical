package user

import (
	"context"
	"github.com/arganaphangquestian/go-medical/user/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service proto.UserServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := proto.NewUserServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	_ = c.conn.Close()
}

func (c *Client) AddUser(ctx context.Context, name string, email string, address string, roleID uint32, genderID uint32, bloodID uint32, birthOfDate string, contact string) (*User, error) {
	r, err := c.service.AddUser(
		ctx,
		&proto.AddUserRequest{
			Name:        name,
			Email:       email,
			Address:     address,
			RoleId:      roleID,
			GenderId:    genderID,
			BloodId:     bloodID,
			BirthOfDate: birthOfDate,
			Contact:     contact,
		},
	)
	if err != nil {
		return nil, err
	}
	userID, _ := primitive.ObjectIDFromHex(r.User.Id)
	return &User{
		ID:          userID,
		Name:        r.User.Name,
		Email:       r.User.Email,
		Address:     r.User.Address,
		RoleID:      r.User.RoleId,
		GenderID:    r.User.GenderId,
		BloodID:     r.User.BloodId,
		BirthOfDate: r.User.BirthOfDate,
		Contact:     r.User.Contact,
	}, nil
}

func (c *Client) GetUsers(ctx context.Context) ([]User, error) {
	r, err := c.service.GetUsers(
		ctx,
		&proto.GetUserRequest{},
	)
	if err != nil {
		return nil, err
	}
	var bloods []User
	for _, a := range r.Users {
		userID, _ := primitive.ObjectIDFromHex(a.Id)
		bloods = append(bloods, User{
			ID:   userID,
			Name: a.Name,
		})
	}
	return bloods, nil
}

func (c *Client) GetUserByID(ctx context.Context, id string) (*User, error) {
	r, err := c.service.GetUserByID(
		ctx,
		&proto.GetUserByIDRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	userID, _ := primitive.ObjectIDFromHex(r.User.Id)
	return &User{
		ID:        userID,
		Name:      r.User.Name,
		CreatedAt: r.User.CreatedAt,
	}, nil
}

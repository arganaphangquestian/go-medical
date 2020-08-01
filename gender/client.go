package gender

import (
	"context"
	"github.com/arganaphangquestian/go-medical/gender/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service proto.GenderServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := proto.NewGenderServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	_ = c.conn.Close()
}

func (c *Client) AddGender(ctx context.Context, name string, description string) (*Gender, error) {
	r, err := c.service.AddGender(
		ctx,
		&proto.AddGenderRequest{Name: name, Description: description},
	)
	if err != nil {
		return nil, err
	}
	return &Gender{
		ID:          r.Gender.Id,
		Name:        r.Gender.Name,
		Description: r.Gender.Description,
		CreatedAt:   r.Gender.CreatedAt,
	}, nil
}

func (c *Client) GetGenders(ctx context.Context) ([]Gender, error) {
	r, err := c.service.GetGenders(
		ctx,
		&proto.GetGenderRequest{},
	)
	if err != nil {
		return nil, err
	}
	var genders []Gender
	for _, a := range r.Genders {
		genders = append(genders, Gender{
			ID:   a.Id,
			Name: a.Name,
		})
	}
	return genders, nil
}

func (c *Client) GetGenderByID(ctx context.Context, id string) (*Gender, error) {
	r, err := c.service.GetGenderByID(
		ctx,
		&proto.GetGenderByIDRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &Gender{
		ID:          r.Gender.Id,
		Name:        r.Gender.Name,
		Description: r.Gender.Description,
		CreatedAt:   r.Gender.CreatedAt,
	}, nil
}

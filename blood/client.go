package blood

import (
	"context"
	"github.com/arganaphangquestian/go-medical/blood/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service proto.BloodServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := proto.NewBloodServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	_ = c.conn.Close()
}

func (c *Client) AddBlood(ctx context.Context, name string, description string) (*Blood, error) {
	r, err := c.service.AddBlood(
		ctx,
		&proto.AddBloodRequest{Name: name, Description: description},
	)
	if err != nil {
		return nil, err
	}
	return &Blood{
		ID:          r.Blood.Id,
		Name:        r.Blood.Name,
		Description: r.Blood.Description,
		CreatedAt:   r.Blood.CreatedAt,
	}, nil
}

func (c *Client) GetBloods(ctx context.Context) ([]Blood, error) {
	r, err := c.service.GetBloods(
		ctx,
		&proto.GetBloodRequest{},
	)
	if err != nil {
		return nil, err
	}
	var bloods []Blood
	for _, a := range r.Bloods {
		bloods = append(bloods, Blood{
			ID:   a.Id,
			Name: a.Name,
		})
	}
	return bloods, nil
}

func (c *Client) GetBloodByID(ctx context.Context, id string) (*Blood, error) {
	r, err := c.service.GetBloodByID(
		ctx,
		&proto.GetBloodByIDRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &Blood{
		ID:          r.Blood.Id,
		Name:        r.Blood.Name,
		Description: r.Blood.Description,
		CreatedAt:   r.Blood.CreatedAt,
	}, nil
}

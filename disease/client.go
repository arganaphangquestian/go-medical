package disease

import (
	"context"
	"github.com/arganaphangquestian/go-medical/disease/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service proto.DiseaseServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := proto.NewDiseaseServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	_ = c.conn.Close()
}

func (c *Client) AddDisease(ctx context.Context, name string, description string) (*Disease, error) {
	r, err := c.service.AddDisease(
		ctx,
		&proto.AddDiseaseRequest{Name: name, Description: description},
	)
	if err != nil {
		return nil, err
	}
	return &Disease{
		ID:          r.Disease.Id,
		Name:        r.Disease.Name,
		Description: r.Disease.Description,
		CreatedAt:   r.Disease.CreatedAt,
	}, nil
}

func (c *Client) GetDiseases(ctx context.Context) ([]Disease, error) {
	r, err := c.service.GetDiseases(
		ctx,
		&proto.GetDiseaseRequest{},
	)
	if err != nil {
		return nil, err
	}
	var diseases []Disease
	for _, a := range r.Diseases {
		diseases = append(diseases, Disease{
			ID:   a.Id,
			Name: a.Name,
		})
	}
	return diseases, nil
}

func (c *Client) GetDiseaseByID(ctx context.Context, id string) (*Disease, error) {
	r, err := c.service.GetDiseaseByID(
		ctx,
		&proto.GetDiseaseByIDRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &Disease{
		ID:          r.Disease.Id,
		Name:        r.Disease.Name,
		Description: r.Disease.Description,
		CreatedAt:   r.Disease.CreatedAt,
	}, nil
}

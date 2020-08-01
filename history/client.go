package history

import (
	"context"
	"github.com/arganaphangquestian/go-medical/history/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service proto.HistoryServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := proto.NewHistoryServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	_ = c.conn.Close()
}

func (c *Client) AddHistory(ctx context.Context, userID string, diseaseID string, note string) (*History, error) {
	r, err := c.service.AddHistory(
		ctx,
		&proto.AddHistoryRequest{UserId: userID, DiseaseId: diseaseID, Note: note},
	)
	if err != nil {
		return nil, err
	}
	return &History{
		ID:        r.History.Id,
		UserID:    r.History.UserId,
		DiseaseID: r.History.DiseaseId,
		Note:      r.History.Note,
		CreatedAt: r.History.CreatedAt,
	}, nil
}

func (c *Client) GetHistories(ctx context.Context) ([]History, error) {
	r, err := c.service.GetHistories(
		ctx,
		&proto.GetHistoryRequest{},
	)
	if err != nil {
		return nil, err
	}
	var histories []History
	for _, a := range r.Histories {
		histories = append(histories, History{
			ID:        a.Id,
			UserID:    a.UserId,
			DiseaseID: a.DiseaseId,
			Note:      a.Note,
			CreatedAt: a.CreatedAt,
		})
	}
	return histories, nil
}

func (c *Client) GetHistoryByID(ctx context.Context, id string) (*History, error) {
	r, err := c.service.GetHistoryByID(
		ctx,
		&proto.GetHistoryByIDRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &History{
		ID:        r.History.Id,
		UserID:    r.History.UserId,
		DiseaseID: r.History.DiseaseId,
		Note:      r.History.Note,
		CreatedAt: r.History.CreatedAt,
	}, nil
}

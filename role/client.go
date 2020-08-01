package role

import (
	"context"
	"github.com/arganaphangquestian/go-medical/role/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service proto.RoleServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := proto.NewRoleServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	_ = c.conn.Close()
}

func (c *Client) AddRole(ctx context.Context, name string, description string) (*Role, error) {
	r, err := c.service.AddRole(
		ctx,
		&proto.AddRoleRequest{Name: name, Description: description},
	)
	if err != nil {
		return nil, err
	}
	return &Role{
		ID:          r.Role.Id,
		Name:        r.Role.Name,
		Description: r.Role.Description,
		CreatedAt:   r.Role.CreatedAt,
	}, nil
}

func (c *Client) GetRoles(ctx context.Context) ([]Role, error) {
	r, err := c.service.GetRoles(
		ctx,
		&proto.GetRoleRequest{},
	)
	if err != nil {
		return nil, err
	}
	var roles []Role
	for _, a := range r.Roles {
		roles = append(roles, Role{
			ID:   a.Id,
			Name: a.Name,
		})
	}
	return roles, nil
}

func (c *Client) GetRoleByID(ctx context.Context, id string) (*Role, error) {
	r, err := c.service.GetRoleByID(
		ctx,
		&proto.GetRoleByIDRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &Role{
		ID:          r.Role.Id,
		Name:        r.Role.Name,
		Description: r.Role.Description,
		CreatedAt:   r.Role.CreatedAt,
	}, nil
}

package role

import "context"

// Service for Role Module
type Service interface {
	AddRole(ctx context.Context, name string, description string) error
	GetRoles(ctx context.Context) ([]Role, error)
	GetRoleByID(ctx context.Context, id string) (*Role, error)
}

// Role struct Model
type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

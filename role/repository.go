package role

import "context"

// Repository Role
type Repository interface {
	Close()

	AddRole(ctx context.Context, name string, description string) error
	GetRoles(ctx context.Context) ([]Role, error)
	GetRoleByID(ctx context.Context, id string) (*Role, error)
}

package role

import "context"

type roleService struct {
	repository Repository
}

// NewService methods
func NewService(r Repository) Service {
	return &roleService{r}
}

// AddRole service implementation
func (s *roleService) AddRole(ctx context.Context, name string, description string) error {
	return s.repository.AddRole(ctx, name, description)
}

// GetRoles service implementation
func (s *roleService) GetRoles(ctx context.Context) ([]Role, error) {
	return s.repository.GetRoles(ctx)
}

// GetRoleByID service implementation
func (s *roleService) GetRoleByID(ctx context.Context, id string) (*Role, error) {
	return s.repository.GetRoleByID(ctx, id)
}

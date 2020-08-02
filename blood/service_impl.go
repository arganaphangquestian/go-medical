package blood

import "context"

type bloodService struct {
	repository Repository
}

// NewService methods
func NewService(r Repository) Service {
	return &bloodService{r}
}

// AddBlood service implementation
func (s *bloodService) AddBlood(ctx context.Context, name string, rhesus bool, description string) error {
	return s.AddBlood(ctx, name, rhesus, description)
}

// GetBloods service implementation
func (s *bloodService) GetBloods(ctx context.Context) ([]Blood, error) {
	return s.GetBloods(ctx)
}

// GetBloodByID service implementation
func (s *bloodService) GetBloodByID(ctx context.Context, id string) (*Blood, error) {
	return s.GetBloodByID(ctx, id)
}

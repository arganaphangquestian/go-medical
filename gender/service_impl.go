package gender

import "context"

type genderService struct {
	repository Repository
}

// NewService methods
func NewService(r Repository) Service {
	return &genderService{r}
}

// AddGender service implementation
func (s *genderService) AddGender(ctx context.Context, name string, description string) error {
	return s.AddGender(ctx, name, description)
}

// GetGenders service implementation
func (s *genderService) GetGenders(ctx context.Context) ([]Gender, error) {
	return s.GetGenders(ctx)
}

// GetGenderByID service implementation
func (s *genderService) GetGenderByID(ctx context.Context, id string) (*Gender, error) {
	return s.GetGenderByID(ctx, id)
}

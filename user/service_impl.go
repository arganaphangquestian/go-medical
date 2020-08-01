package user

import "context"

type userService struct {
	repository Repository
}

// NewService methods
func NewService(r Repository) Service {
	return &userService{r}
}

// AddUser service implementation
func (s *userService) AddUser(ctx context.Context, name string, email string, address string, roleID uint32, genderID uint32, bloodID uint32, BirthOfDate string, contact string) error {
	return s.repository.AddUser(ctx, name, email, address, roleID, genderID, bloodID, BirthOfDate, contact)
}

// GetUsers service implementation
func (s *userService) GetUsers(ctx context.Context) ([]User, error) {
	return s.repository.GetUsers(ctx)
}

// GetUserByID service implementation
func (s *userService) GetUserByID(ctx context.Context, id string) (*User, error) {
	return s.repository.GetUserByID(ctx, id)
}

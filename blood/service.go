package blood

import "context"

// Service for blood Module
type Service interface {
	AddBlood(ctx context.Context, name string, rhesus bool, description string) error
	GetBloods(ctx context.Context) ([]Blood, error)
	GetBloodByID(ctx context.Context, id string) (*Blood, error)
}

// Blood struct Model
type Blood struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Rhesus      bool   `json:"rhesus"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

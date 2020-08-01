package disease

import "context"

// Service for Disease Module
type Service interface {
	AddDisease(ctx context.Context, name string, description string) error
	GetDiseases(ctx context.Context) ([]Disease, error)
	GetDiseaseByID(ctx context.Context, id string) (*Disease, error)
	SearchProducts(ctx context.Context, query string) ([]Disease, error)
}

// Disease struct Model
type Disease struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

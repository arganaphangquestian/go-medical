package disease

import "context"

type Repository interface {
	Close()

	AddDisease(ctx context.Context, name string, description string) error
	GetDiseases(ctx context.Context) ([]Disease, error)
	GetDiseaseByID(ctx context.Context, id string) (*Disease, error)
	SearchProducts(ctx context.Context, query string) ([]Disease, error)
}

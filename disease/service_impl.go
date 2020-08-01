package disease

import "context"

type diseaseService struct {
	repository Repository
}

// NewService methods
func NewService(r Repository) Service {
	return &diseaseService{r}
}

func (d *diseaseService) AddDisease(ctx context.Context, name string, description string) error {
	return d.repository.AddDisease(ctx, name, description)
}

func (d *diseaseService) GetDiseases(ctx context.Context) ([]Disease, error) {
	return d.repository.GetDiseases(ctx)
}

func (d *diseaseService) GetDiseaseByID(ctx context.Context, id string) (*Disease, error) {
	return d.repository.GetDiseaseByID(ctx, id)
}

func (d *diseaseService) SearchDiseases(ctx context.Context, query string) ([]Disease, error) {
	return d.repository.SearchProducts(ctx, query)
}

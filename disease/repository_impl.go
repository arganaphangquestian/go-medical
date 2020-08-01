package disease

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/olivere/elastic"
	"github.com/segmentio/ksuid"
)

var (
	ErrNotFound = errors.New("entity not found")
)

type elasticRepository struct {
	client *elastic.Client
}

type diseaseDocument struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NewElasticRepository methods
func NewElasticRepository(url string) (Repository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &elasticRepository{client}, nil
}

func (r *elasticRepository) Close() {

}

func (r *elasticRepository) AddDisease(ctx context.Context, name string, description string) error {
	_, err := r.client.Index().
		Index("disease").
		Type("disease").
		Id(ksuid.New().String()).
		BodyJson(diseaseDocument{
			Name:        name,
			Description: description,
		}).Do(ctx)

	return err

}

func (r *elasticRepository) GetDiseases(ctx context.Context) ([]Disease, error) {
	res, err := r.client.Search().
		Index("disease").
		Type("disease").
		Query(elastic.NewMatchAllQuery()).
		Do(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var diseases []Disease
	for _, hit := range res.Hits.Hits {
		d := diseaseDocument{}
		if err = json.Unmarshal(*hit.Source, &d); err == nil {
			diseases = append(diseases, Disease{
				ID:          hit.Id,
				Name:        d.Name,
				Description: d.Description,
			})
		}
	}
	return diseases, err
}

func (r *elasticRepository) GetDiseaseByID(ctx context.Context, id string) (*Disease, error) {
	res, err := r.client.Get().
		Index("disease").
		Type("disease").
		Id(id).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	if !res.Found {
		return nil, ErrNotFound
	}
	p := diseaseDocument{}
	if err = json.Unmarshal(*res.Source, &p); err != nil {
		return nil, err
	}
	return &Disease{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
	}, err
}

func (r *elasticRepository) SearchProducts(ctx context.Context, query string) ([]Disease, error) {
	res, err := r.client.Search().
		Index("disease").
		Type("disease").
		Query(elastic.NewMultiMatchQuery(query, "name", "description")).
		Do(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var diseases []Disease
	for _, hit := range res.Hits.Hits {
		p := diseaseDocument{}
		if err = json.Unmarshal(*hit.Source, &p); err == nil {
			diseases = append(diseases, Disease{
				ID:          hit.Id,
				Name:        p.Name,
				Description: p.Description,
			})
		}
	}
	return diseases, err
}

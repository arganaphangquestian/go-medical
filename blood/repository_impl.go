package blood

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/segmentio/ksuid"
	"time"
)

type postgresRepository struct {
	db *pgx.Conn
}

// NewPostgres methods
func NewPostgres(url string) (Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = r.db.Close(ctx)
}

// AddBlood repository implementation
func (r *postgresRepository) AddBlood(ctx context.Context, name string, rhesus bool, description string) error {
	a := &Blood{
		Name:        name,
		Rhesus:      rhesus,
		Description: description,
	}
	res, err := r.db.Query(ctx, "INSERT INTO bloods(id, name, rhesus, description) VALUES($1, $2, $3)", ksuid.New().String(), a.Name, a.Rhesus, a.Description)
	if res != nil {
		res.Close()
	}
	return err
}

// GetBloods repository implementation
func (r *postgresRepository) GetBloods(ctx context.Context) ([]Blood, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT id, name, description FROM bloods",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bloods []Blood
	for rows.Next() {
		a := &Blood{}
		if err = rows.Scan(&a.ID, &a.Name, &a.Description); err == nil {
			bloods = append(bloods, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bloods, nil
}

// GetBloodByID repository implementation
func (r *postgresRepository) GetBloodByID(ctx context.Context, id string) (*Blood, error) {
	a := &Blood{}
	row := r.db.QueryRow(ctx, `SELECT id, name, description FROM bloods WHERE id=$1;`, id)

	err := row.Scan(&a.ID, &a.Name, &a.Description)

	if err != nil {
		return nil, err
	}

	return a, nil

}

package blood

import (
	"context"
	"database/sql"
)

type postgresRepository struct {
	db *sql.DB
}

// NewPostgres methods
func NewPostgres(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

// AddBlood repository implementation
func (r *postgresRepository) AddBlood(ctx context.Context, name string, description string) error {
	a := &Blood{
		Name:        name,
		Description: description,
	}
	_, err := r.db.ExecContext(ctx, "INSERT INTO bloods(name, description) VALUES($1, $2)", a.Name, a.Description)
	return err
}

// GetBloods repository implementation
func (r *postgresRepository) GetBloods(ctx context.Context) ([]Blood, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, description FROM bloods",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bloods := []Blood{}
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
	row := r.db.QueryRow(`SELECT id, name, description FROM bloods WHERE id=$1;`, id)

	err := row.Scan(&a.ID, &a.Name, &a.Description)

	if err != nil {
		return nil, err
	}

	return a, nil

}

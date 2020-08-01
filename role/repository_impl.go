package role

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
	_ = r.db.Close()
}

// AddRole repository implementation
func (r *postgresRepository) AddRole(ctx context.Context, name string, description string) error {
	a := &Role{
		Name:        name,
		Description: description,
	}
	_, err := r.db.ExecContext(ctx, "INSERT INTO roles(name, description) VALUES($1, $2)", a.Name, a.Description)
	return err
}

// GetRoles repository implementation
func (r *postgresRepository) GetRoles(ctx context.Context) ([]Role, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, description FROM roles",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		a := &Role{}
		if err = rows.Scan(&a.ID, &a.Name, &a.Description); err == nil {
			roles = append(roles, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

// GetRoleByID repository implementation
func (r *postgresRepository) GetRoleByID(ctx context.Context, id string) (*Role, error) {
	a := &Role{}
	row := r.db.QueryRow(`SELECT id, name, description FROM roles WHERE id=$1;`, id)

	err := row.Scan(&a.ID, &a.Name, &a.Description)

	if err != nil {
		return nil, err
	}

	return a, nil

}

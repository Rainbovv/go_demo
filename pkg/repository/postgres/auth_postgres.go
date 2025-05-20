package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-demo/types"
)

type AuthPostgresRepository struct {
	db *sqlx.DB
}

func (r *AuthPostgresRepository) GetUserByUsername(username string) (types.User, error) {
	var user types.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", userTable)
	err := r.db.Get(&user, query, username)

	return user, err
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgresRepository {
	return &AuthPostgresRepository{db: db}
}

func (r *AuthPostgresRepository) CreateUser(user types.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password) VALUES ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userTable     = "demo_app.app_user"
	todoListTable = "demo_app.todo_list"
	todoUserList  = "demo_app.user_list"
	todoTodoItem  = "demo_app.todo_item"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
	SSLMode  string
}

func NewPostgresRepository(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.UserName, config.Database, config.Password, config.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

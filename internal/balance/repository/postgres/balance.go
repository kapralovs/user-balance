package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type PostgresRepo struct {
	Url  string
	Conn *pgx.Conn
}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{
		Url: os.Getenv("DATABASE_URL"),
	}
}

func (pr *PostgresRepo) GetBalanceInfo(userId int) ([]byte, error) {
	row := pr.Conn.QueryRow(context.Background(), fmt.Sprintf("select * from users where id=%d", userId))
	return info, nil
}

func (pr *PostgresRepo) Crediting(userId int, sum int) error {

}

func (pr *PostgresRepo) Debiting(userId, sum int) error {

}

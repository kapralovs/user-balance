package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/kapralovs/user-balance/internal/models"
)

type PostgresRepo struct {
	conn *pgx.Conn
}

func NewConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_POSTGRES_USER"),
		os.Getenv("DB_POSTGRES_PASSWORD"),
		os.Getenv("POOLER_PGBOUNCER_ADDRESS"),
		os.Getenv("POOLER_PGBOUNCER_GO_PORT"),
		os.Getenv("DB_POSTGRES_DB")))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func NewRepo(conn *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{
		conn: conn,
	}
}

// Реализация метода получения информации о балансе пользователя
func (pr *PostgresRepo) GetBalanceInfo(userId int) (*models.User, error) {
	info := &models.User{}
	err := pr.conn.QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", userId).Scan(info)
	if err != nil {
		return nil, err
	}
	// err := pgxscan.Get(context.Background(), pr.conn, info, "SELECT * FROM users WHERE user_id=$1", userId)
	// if err != nil {
	// 	return nil, err
	// }
	return info, nil
}

// Реализация метода зачисления средств
func (pr *PostgresRepo) Crediting(userId, sum int) error {
	return nil
}

// Реализация метода списания средств
func (pr *PostgresRepo) Debiting(userId, sum int) error {
	return nil
}

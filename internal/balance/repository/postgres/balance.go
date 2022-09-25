package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/kapralovs/user-balance/internal/models"
)

type PostgresRepo struct {
	cfg *models.PostgresConfig
	// Url  string
	Conn *pgx.Conn
}

func NewPostgresRepo() *pgx.Conn {
	r:=&PostgresRepo{
		cfg:&models.PostgresConfig{
			Host:     os.Getenv("POOLER_PGBOUNCER_ADDRESS"),
			Port:     os.Getenv("POOLER_PGBOUNCER_GO_PORT"),
			Username: os.Getenv("DB_POSTGRES_USER"),
			DBName:   os.Getenv("DB_POSTGRES_DB"),
			SSLMode:  "disable",
			TimeZone: "Europe/Moscow",
			Password: os.Getenv("DB_POSTGRES_PASSWORD"),
		},
	}

	conn,err:=pgx.Connect(context.Background(),fmt.Sprintf("postgres://%s:%s@%s:%s/%s",r.cfg.Username,r.cfg.Password,r.cfg.Host,r.cfg.Port,r.cfg.DBName))
	if err!=nil{
		log.Fatal(err)
	}
	return conn
}

// Реализация метода получения информации о балансе пользователя
func (pr *PostgresRepo) GetBalanceInfo(userId int) ([]byte, error) {
	info:=models.Users
	// err := pr.Conn.QueryRow(context.Background(), fmt.Sprintf("SELECT * FROM users WHERE id=%d", userId)).Scan(&)
	if err != nil {
		return nil, err
	}
	pgxscan.Get(context.Background(),pr.Conn,,"SELECT * FROM users WHERE user_id=$1",userId)

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

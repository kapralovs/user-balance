package postgres

import "os"

type PostgresRepo struct {
	Url string
}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{
		Url: os.Getenv("DATABASE_URL"),
	}
}

func (pr *PostgresRepo) GetBalanceInfo() error {
	return info, nil
}

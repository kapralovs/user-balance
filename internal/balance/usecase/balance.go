package usecase

import "github.com/kapralovs/user-balance/internal/balance"

type Usecase struct {
	repo balance.Repository
}

func New(r balance.Repository) *Usecase {
	return &Usecase{
		repo: r,
	}
}

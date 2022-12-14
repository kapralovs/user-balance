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

func (uc *Usecase) GetBalanceInfo(userId int) error {
	return nil
}

func (uc *Usecase) Transfer(fromId, toId int) error {
	return nil
}

func (uc *Usecase) Debiting(userId, sum int) error {
	return nil
}

func (uc *Usecase) Crediting(userId, sum int) error {
	return nil
}

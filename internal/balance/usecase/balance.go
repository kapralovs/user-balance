package usecase

import (
	"log"

	"github.com/kapralovs/user-balance/internal/balance"
	"github.com/kapralovs/user-balance/internal/models"
)

type Usecase struct {
	repo balance.Repository
}

func New(r balance.Repository) *Usecase {
	return &Usecase{
		repo: r,
	}
}

func (uc *Usecase) GetBalanceInfo(userId int) (*models.User, error) {
	user, err := uc.repo.GetBalanceInfo(userId)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
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

package balance

import "github.com/kapralovs/user-balance/internal/models"

type Repository interface {
	GetBalanceInfo(userId int) (*models.User, error)
	Debiting(userId, sum int) error
	Crediting(userId, sum int) error
}
type Usecase interface {
	GetBalanceInfo(userId int) error
	Transfer(fromId, toId int) error
	Debiting(userId, sum int) error
	Crediting(userId, sum int) error
}

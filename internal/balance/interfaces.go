package balance

type Repository interface {
	GetBalanceInfo(userId int) ([]byte, error)
	Debiting(userId, sum int) error
	Crediting(userId, sum int) error
}
type Usecase interface {
	GetBalanceInfo(userId int) error
	Transfer(fromId, toId int) error
	Debiting(userId, sum int) error
	Crediting(userId, sum int) error
}

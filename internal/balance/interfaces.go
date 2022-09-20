package balance

type Repository interface {
	GetBalanceInfo()
}
type Usecase interface {
	GetBalanceInfo() error
	Transfer(userId int) error
}

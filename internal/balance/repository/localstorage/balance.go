package localstorage

import (
	"sync"

	"github.com/kapralovs/user-balance/internal/models"
)

type LocalRepo struct {
	users map[int]*models.User
	mu    *sync.Mutex
}

func (ls *LocalRepo) Collect() {
	ls.users[1] = &models.User{
		ID:         1,
		FirstName:  "Сергей",
		MiddleName: "Андреевич",
		LastName:   "Капралов",
		Balance:    3000,
	}
	ls.users[2] = &models.User{
		ID:         2,
		FirstName:  "Джефф",
		MiddleName: "Батькович",
		LastName:   "Безос",
		Balance:    15000000000,
	}
}

func NewLocalRepo() *LocalRepo {
	return &LocalRepo{
		users: make(map[int]*models.User),
		mu:    new(sync.Mutex),
	}
}

func (lr *LocalRepo) GetBalanceInfo(userId int) (*models.User, error) {
	user := lr.users[userId]
	return user, nil
}

// Имплементация метода для списания средств
func (lr *LocalRepo) Debiting(userId, sum int) error {
	
	return nil
}

// Имплементация метода для зачисления средств
func (lr *LocalRepo) Crediting(userId, sum int) error {
	return nil
}

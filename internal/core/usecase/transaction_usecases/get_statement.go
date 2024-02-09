package transaction_usecases

import (
	"github.com/reonardoleis/banky/internal/core/domain"
)

func (u usecase) GetStatement(userId uint) (*domain.Statement, error) {
	statement, err := u.r.GetStatement(userId)
	if err != nil {
		return nil, err
	}

	return statement, nil
}

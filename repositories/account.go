package repositories

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
)

type (
	AccountRepository interface {
		GetAccountPaginate(ctx context.Context, pagination database.Pagination, search string) (*database.Pagination, error)
		GetAccountByID(ctx context.Context, id int) (models.Account, error)
		CreateAccount(ctx context.Context, account *models.Account) error
		UpdateAccount(ctx context.Context, id int, account *models.Account) error
		DeleteAccount(ctx context.Context, id int) error
	}
)

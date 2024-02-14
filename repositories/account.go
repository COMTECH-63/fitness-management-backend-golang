package repositories

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/getsentry/sentry-go"
)

type (
	AccountRepository interface {
		GetAccountPaginate(ctx context.Context, span *sentry.Span, pagination database.Pagination, search string) (*database.Pagination, error)
		GetAccountByID(ctx context.Context, span *sentry.Span, id int) (models.Account, error)
		CreateAccount(ctx context.Context, span *sentry.Span, account *models.Account) error
		UpdateAccount(ctx context.Context, span *sentry.Span, id int, account *models.Account) error
		DeleteAccount(ctx context.Context, span *sentry.Span, id int) error
	}
)

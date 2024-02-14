package services

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/getsentry/sentry-go"
)

type (
	AccountService interface {
		GetAccounts(ctx context.Context, span *sentry.Span, paginate database.Pagination, search string) (*database.Pagination, error)
		GetAccount(ctx context.Context, span *sentry.Span, id int) (map[string]interface{}, error)
		CreateAccount(ctx context.Context, span *sentry.Span, accountDto *AccountDto) error
		UpdateAccount(ctx context.Context, span *sentry.Span, id int, accountDto *AccountDto) error
		DeleteAccount(ctx context.Context, span *sentry.Span, id int) error
	}

	AccountDto struct {
		Username string `json:"username" form:"username" query:"username" validate:"required,max=55"`
		Password string `json:"password" form:"password" query:"password" validate:"required,max=100"`
	}
)

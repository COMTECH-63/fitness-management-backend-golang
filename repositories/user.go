package repositories

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/getsentry/sentry-go"
)

type (
	UserRepository interface {
		GetUserPaginate(ctx context.Context, span *sentry.Span, pagination database.Pagination, search string) (*database.Pagination, error)
		GetUserByID(ctx context.Context, span *sentry.Span, id int) (models.User, error)
		CreateUser(ctx context.Context, span *sentry.Span, user *models.User) error
		UpdateUser(ctx context.Context, span *sentry.Span, id int, user *models.User) error
		UpdatePasswordUser(ctx context.Context, span *sentry.Span, id int, user *models.User) error
		DeleteUser(ctx context.Context, span *sentry.Span, id int) error
	}
)

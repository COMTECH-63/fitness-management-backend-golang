package repositories

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/getsentry/sentry-go"
)

type (
	PermissionRepository interface {
		GetPermissionPaginate(ctx context.Context, span *sentry.Span, pagination database.Pagination, search string) (*database.Pagination, error)
		GetPermissionByID(ctx context.Context, span *sentry.Span, id int) (models.Permission, error)
		CreatePermission(ctx context.Context, span *sentry.Span, permission *models.Permission) error
		UpdatePermission(ctx context.Context, span *sentry.Span, id int, permission *models.Permission) error
		DeletePermission(ctx context.Context, span *sentry.Span, id int) error
	}
)

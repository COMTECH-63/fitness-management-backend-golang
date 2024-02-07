package repositories

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/getsentry/sentry-go"
)

type (
	RoleRepository interface {
		GetRolePaginate(ctx context.Context, span *sentry.Span, pagination database.Pagination, search string) (*database.Pagination, error)
		GetRoleByID(ctx context.Context, span *sentry.Span, id int) (models.Role, error)
		CreateRole(ctx context.Context, span *sentry.Span, role *models.Role) error
		UpdateRole(ctx context.Context, span *sentry.Span, id int, role *models.Role) error
		DeleteRole(ctx context.Context, span *sentry.Span, id int) error
	}
)

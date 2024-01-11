package repositories

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
)

type (
	RoleRepository interface {
		GetRolePaginate(ctx context.Context, pagination database.Pagination, search string) (*database.Pagination, error)
		GetRoleByID(ctx context.Context, id int) (models.Role, error)
		CreateRole(ctx context.Context, role *models.Role) error
		UpdateRole(ctx context.Context, id int, role *models.Role) error
		DeleteRole(ctx context.Context, id int) error
	}
)

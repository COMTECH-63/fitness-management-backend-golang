package repositories

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
)

type (
	UserRepository interface {
		GetUserPaginate(ctx context.Context, pagination database.Pagination, search string) (*database.Pagination, error)
		GetUserByID(ctx context.Context, id int) (models.User, error)
		CreateUser(ctx context.Context, user *models.User) error
		UpdateUser(ctx context.Context, id int, user *models.User) error
		DeleteUser(ctx context.Context, id int) error
	}
)

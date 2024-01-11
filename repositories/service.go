package repositories

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
)

type (
	ServiceRepository interface {
		GetServicePaginate(ctx context.Context, pagination database.Pagination, search string) (*database.Pagination, error)
		GetServiceByID(ctx context.Context, id int) (models.Service, error)
		CreateService(ctx context.Context, service *models.Service) error
		UpdateService(ctx context.Context, id int, service *models.Service) error
		DeleteService(ctx context.Context, id int) error
	}
)

package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/COMTECH-63/fitness-management/pkg/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return serviceRepository{db: db}
}

func (r serviceRepository) GetServicePaginate(ctx context.Context, pagination database.Pagination, search string) (*database.Pagination, error) {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "GetServicePaginateRepository", trace.WithAttributes(attribute.String("repository", "GetServicePaginate"), attribute.String("search", search)))
		services     []models.Service
		err          error
	)

	// Pagination query
	if search != "" {
		if err = r.db.Scopes(database.Paginate(services, &pagination, r.db)).
			Where(`name LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Find(&services).Error; err != nil {
			log.Println(err)
			return nil, errors.New("GetServicePaginateError")
		}
	} else {
		if err = r.db.Scopes(database.Paginate(services, &pagination, r.db)).
			Preload("Users").Find(&services).Error; err != nil {
			return nil, err
		}
	}

	// Set data
	pagination.Data = services

	childSpan.End()

	return &pagination, nil
}

func (r serviceRepository) GetServiceByID(ctx context.Context, id int) (models.Service, error) {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "GetServiceByIDRepository", trace.WithAttributes(attribute.String("repository", "GetServiceByID")))
		service      models.Service
		err          error
	)

	// Query
	if err = r.db.Preload("Users").Find(&service, id).Error; err != nil {
		return service, err
	}

	childSpan.End()

	return service, nil
}

func (r serviceRepository) CreateService(ctx context.Context, service *models.Service) error {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "CreateServiceRepository", trace.WithAttributes(attribute.String("repository", "CreateService")))
		err          error
	)

	// Execute
	if err = r.db.Create(&service).Error; err != nil {
		return err
	}

	childSpan.End()

	return nil
}

func (r serviceRepository) UpdateService(ctx context.Context, id int, service *models.Service) error {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "UpdateServiceRepository", trace.WithAttributes(attribute.String("repository", "UpdateService")))
		existService *models.Service
		err          error
	)

	// Get model
	r.db.First(&existService)

	// Clear existing associations
	r.db.Model(&existService).Association("Users").Clear()

	// Set attributes
	existService.Name = service.Name
	existService.Description = service.Description
	existService.Price = service.Price

	existService.Users = service.Users

	// Execute
	if err = r.db.Save(&existService).Error; err != nil {
		return err
	}

	childSpan.End()

	return nil
}

func (r serviceRepository) DeleteService(ctx context.Context, id int) error {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "DeleteServiceRepository", trace.WithAttributes(attribute.String("repository", "DeleteService")))
		err          error
	)

	// Execute
	if err = r.db.Delete(&models.Service{}, id).Error; err != nil {
		return err
	}

	childSpan.End()

	return nil
}

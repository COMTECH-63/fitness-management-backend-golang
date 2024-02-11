package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/COMTECH-63/fitness-management/pkg/tracing"
	"github.com/getsentry/sentry-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepository{db: db}
}

func (r roleRepository) GetRolePaginate(ctx context.Context, span *sentry.Span, pagination database.Pagination, search string) (*database.Pagination, error) {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "GetRolePaginateRepository", trace.WithAttributes(attribute.String("repository", "GetRolePaginate"), attribute.String("search", search)))
		roles        []models.Role
		err          error
	)

	// Pagination query
	if search != "" {
		if err = r.db.Scopes(database.Paginate(roles, &pagination, r.db)).
			Where(`name LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Find(&roles).Error; err != nil {
			log.Println(err)
			return nil, errors.New("GetRolePaginateError")
		}
	} else {
		if err = r.db.Scopes(database.Paginate(roles, &pagination, r.db)).
			// Preload("Roles").Preload("Permissions").Preload("Services").Preload("Classes").Preload("Orders").Preload("Bookings").Preload("BookingClasses").Preload("BookingPersonalTrainers").Find(&roles).Error; err != nil {
			Preload("Users").Preload("Permissions").Find(&roles).Error; err != nil {
			return nil, err
		}
	}

	// Set data
	pagination.Data = roles

	childSpan.End()

	return &pagination, nil
}

func (r roleRepository) GetRoleByID(ctx context.Context, span *sentry.Span, id int) (models.Role, error) {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "GetRoleByIDRepository", trace.WithAttributes(attribute.String("repository", "GetRoleByID")))
		role         models.Role
		err          error
	)

	// Query
	// if err = r.db.Preload("Roles").Preload("Permissions").Preload("Services").Preload("Classes").Preload("Orders").Preload("Bookings").Preload("BookingClasses").Preload("BookingPersonalTrainers").Find(&role, id).Error; err != nil {
	// 	return role, err
	// }

	if err = r.db.Preload("Users").Preload("Permissions").Find(&role, id).Error; err != nil {
		return role, err
	}

	childSpan.End()

	return role, nil
}

func (r roleRepository) CreateRole(ctx context.Context, span *sentry.Span, role *models.Role) error {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "CreateRoleRepository", trace.WithAttributes(attribute.String("repository", "CreateRole")))
		err          error
	)

	// Execute
	if err = r.db.Create(&role).Error; err != nil {
		return err
	}

	childSpan.End()

	return nil
}

func (r roleRepository) UpdateRole(ctx context.Context, span *sentry.Span, id int, role *models.Role) error {
	var (
		childSpan = span.StartChild("UpdateRoleRepository")
		existRole *models.Role
	)

	// Get model
	r.db.Find(&existRole, id)

	// Clear existing associations
	r.db.Model(&existRole).Association("Users").Clear()
	r.db.Model(&existRole).Association("Permissions").Clear()

	// Set attributes
	existRole.Name = role.Name

	existRole.Users = role.Users
	existRole.Permissions = role.Permissions

	// Execute
	if err := r.db.Save(&existRole).Error; err != nil {
		return err
	}

	childSpan.Finish()

	return nil
}

func (r roleRepository) DeleteRole(ctx context.Context, span *sentry.Span, id int) error {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "DeleteRoleRepository", trace.WithAttributes(attribute.String("repository", "DeleteRole")))
		err          error
	)

	// Execute
	if err = r.db.Delete(&models.Role{}, id).Error; err != nil {
		return err
	}

	childSpan.End()

	return nil
}

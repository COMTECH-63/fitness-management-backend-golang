package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/getsentry/sentry-go"
	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return permissionRepository{db: db}
}

func (r permissionRepository) GetPermissionPaginate(ctx context.Context, span *sentry.Span, pagination database.Pagination, search string) (*database.Pagination, error) {
	var (
		childSpan   = span.StartChild("GetPermissionPaginateRepository")
		permissions []models.Permission
		err         error
	)

	// Pagination query
	if search != "" {
		if err = r.db.Scopes(database.Paginate(permissions, &pagination, r.db)).
			Where(`email LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Or(`first_name LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Or(`last_name LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Or(`phone LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Find(&permissions).Error; err != nil {
			log.Println(err)
			return nil, errors.New("GetPermissionPaginateError")
		}
	} else {
		if err = r.db.Scopes(database.Paginate(permissions, &pagination, r.db)).
			Preload("Teams").Preload("Roles").Preload("Permissions").Preload("Projects").Find(&permissions).Error; err != nil {
			return nil, err
		}
	}

	// Set data
	pagination.Data = permissions

	childSpan.Finish()

	return &pagination, nil
}

func (r permissionRepository) GetPermissionByID(ctx context.Context, span *sentry.Span, id int) (models.Permission, error) {
	var (
		childSpan  = span.StartChild("GetPermissionByIDRepository")
		permission models.Permission
		err        error
	)

	// Query
	if err = r.db.First(&permission, id).Error; err != nil {
		return permission, err
	}

	if err = r.db.Preload("Teams").Preload("Roles").Preload("Permissions").Preload("Projects").Find(&permission).Error; err != nil {
		return permission, err
	}

	childSpan.Finish()

	return permission, nil
}

func (r permissionRepository) CreatePermission(ctx context.Context, span *sentry.Span, permission *models.Permission) error {
	var (
		childSpan = span.StartChild("CreatePermissionRepository")
		err       error
	)

	// Execute
	if err = r.db.Create(&permission).Error; err != nil {
		return err
	}

	fmt.Println(&permission)

	childSpan.Finish()

	return nil
}

func (r permissionRepository) UpdatePermission(ctx context.Context, span *sentry.Span, id int, permission *models.Permission) error {
	var (
		childSpan       = span.StartChild("UpdatePermissionRepository")
		existPermission *models.Permission
		err             error
	)

	// Get model
	r.db.Find(&existPermission, id)

	r.db.Model(&existPermission).Association("Users").Clear()
	r.db.Model(&existPermission).Association("Roles").Clear()

	// Set attributes
	existPermission.Name = permission.Name
	existPermission.Description = permission.Description

	existPermission.Users = permission.Users
	existPermission.Roles = permission.Roles

	// Execute
	if err = r.db.Save(&existPermission).Error; err != nil {
		return err
	}

	childSpan.Finish()

	return nil
}

func (r permissionRepository) DeletePermission(ctx context.Context, span *sentry.Span, id int) error {
	var (
		childSpan = span.StartChild("DeletePermissionRepository")
		err       error
	)

	// Execute
	if err = r.db.Delete(&models.Permission{}, id).Error; err != nil {
		return err
	}

	childSpan.Finish()

	return nil
}

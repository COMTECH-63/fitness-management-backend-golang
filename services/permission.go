package services

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/getsentry/sentry-go"
)

type (
	PermissionService interface {
		GetPermissions(ctx context.Context, span *sentry.Span, paginate database.Pagination, search string) (*database.Pagination, error)
		GetPermission(ctx context.Context, span *sentry.Span, id int) (map[string]interface{}, error)
		CreatePermission(ctx context.Context, span *sentry.Span, permissionDto *CreatePermissionDto) error
		UpdatePermission(ctx context.Context, span *sentry.Span, id int, permissionDto *UpdatePermissionDto) error
		DeletePermission(ctx context.Context, span *sentry.Span, id int) error
	}

	// Create Permission
	CreatePermissionUserDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	CreatePermissionRoleDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	CreatePermissionDto struct {
		Name        string `json:"name" form:"name" query:"name" validate:"required,max=25"`
		Description string `json:"description" form:"description" query:"description" validate:"required,max=50"`

		UserDtos []CreatePermissionUserDto `json:"users" validate:"dive"`
		RoleDtos []CreatePermissionRoleDto `json:"roles" validate:"dive"`
	}

	// Update Permission
	UpdatePermissionUserDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	UpdatePermissionRoleDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}
	UpdatePermissionDto struct {
		Name        string `json:"name" form:"name" query:"name" validate:"required,max=25"`
		Description string `json:"description" form:"description" query:"description" validate:"required,max=50"`

		UserDtos []UpdatePermissionUserDto `json:"users" validate:"dive"`
		RoleDtos []UpdatePermissionRoleDto `json:"roles" validate:"dive"`
	}
)

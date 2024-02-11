package services

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/getsentry/sentry-go"
)

type (
	RoleService interface {
		GetRoles(ctx context.Context, span *sentry.Span, paginate database.Pagination, search string) (*database.Pagination, error)
		GetRole(ctx context.Context, span *sentry.Span, id int) (map[string]interface{}, error)
		CreateRole(ctx context.Context, span *sentry.Span, roleDto *CreateRoleDto) error
		UpdateRole(ctx context.Context, span *sentry.Span, id int, roleDto *UpdateRoleDto) error
		DeleteRole(ctx context.Context, span *sentry.Span, id int) error
	}

	// Create Role
	CreateRoleUserDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	CreateRolePermissionDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	CreateRoleDto struct {
		Name string `json:"name" form:"name" query:"name" validate:"required,max=10"`

		UserDtos       []CreateRoleUserDto       `json:"users" validate:"dive"`
		PermissionDtos []CreateRolePermissionDto `json:"permissions" validate:"dive"`
	}

	// Update Role
	UpdateRoleUserDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	UpdateRolePermissionDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	UpdateRoleDto struct {
		Name string `json:"name" form:"name" query:"name" validate:"required,max=10"`

		UserDtos       []UpdateRoleUserDto       `json:"users" validate:"dive"`
		PermissionDtos []UpdateRolePermissionDto `json:"permissions" validate:"dive"`
	}
)

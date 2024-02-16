package services

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/getsentry/sentry-go"
)

type (
	UserService interface {
		GetUsers(ctx context.Context, span *sentry.Span, paginate database.Pagination, search string) (*database.Pagination, error)
		GetUser(ctx context.Context, span *sentry.Span, id int) (map[string]interface{}, error)
		CreateUser(ctx context.Context, span *sentry.Span, userDto *CreateUserDto) error
		UpdateUser(ctx context.Context, span *sentry.Span, id int, userDto *UpdateUserDto) error
		UpdatePasswordUser(ctx context.Context, span *sentry.Span, id int, userDto *UpdatePasswordUserdDto) error
		DeleteUser(ctx context.Context, span *sentry.Span, id int) error
	}

	// Create User
	CreateUserRoleDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	CreateUserPermissionDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	CreateUserServiceDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	CreateUserDto struct {
		Username    string `json:"username" form:"username" query:"username" validate:"required,max=55"`
		Password    string `json:"password" form:"password" query:"password" validate:"required,max=100"`
		FirstName   string `json:"first_name" form:"first_name" query:"first_name" validate:"required,max=30"`
		LastName    string `json:"last_name" form:"last_name" query:"last_name" validate:"required,max=50"`
		IDCard      string `json:"id_card" form:"id_card" query:"id_card" validate:"required,max=13"`
		Email       string `json:"email" form:"email" query:"email" validate:"required,max=100"`
		PhoneNumber string `json:"phone_number" form:"phone_number" query:"phone_number" validate:"required,max=10"`
		Address     string `json:"address" form:"address" query:"address" validate:"required"`
		Sex         string `json:"sex" form:"sex" query:"sex" validate:"required"`
		ImageURL    string `json:"image_url" form:"image_url" query:"image_url" validate:"required"`
		MemberID    string `json:"member_id" form:"member_id" query:"member_id" validate:"required,max=5"`

		RoleDtos       []CreateUserRoleDto       `json:"roles" validate:"dive"`
		PermissionDtos []CreateUserPermissionDto `json:"permissions" validate:"dive"`
		ServiceDtos    []CreateUserServiceDto    `json:"services" validate:"dive"`
	}

	// Update User
	UpdateUserRoleDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}
	UpdateUserPermissionDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	UpdateUserServiceDto struct {
		ID uint `json:"id" form:"id" query:"id" validate:"required"`
	}

	UpdateUserDto struct {
		// Username    string `json:"username" form:"username" query:"username" validate:"required,max=55"`
		// Password    string `json:"password" form:"password" query:"password" validate:"required,max=100"`
		FirstName   string `json:"first_name" form:"first_name" query:"first_name" validate:"required,max=30"`
		LastName    string `json:"last_name" form:"last_name" query:"last_name" validate:"required,max=50"`
		IDCard      string `json:"id_card" form:"id_card" query:"id_card" validate:"required,max=13"`
		Email       string `json:"email" form:"email" query:"email" validate:"required,max=100"`
		PhoneNumber string `json:"phone_number" form:"phone_number" query:"phone_number" validate:"required,max=10"`
		Address     string `json:"address" form:"address" query:"address" validate:"required"`
		Sex         string `json:"sex" form:"sex" query:"sex" validate:"required"`
		ImageURL    string `json:"image_url" form:"image_url" query:"image_url" validate:"required"`
		MemberID    string `json:"member_id" form:"member_id" query:"member_id" validate:"required,max=5"`

		RoleDtos       []UpdateUserRoleDto       `json:"roles" validate:"dive"`
		PermissionDtos []UpdateUserPermissionDto `json:"permissions" validate:"dive"`
		ServiceDtos    []UpdateUserServiceDto    `json:"services" validate:"dive"`
	}

	UpdatePasswordUserdDto struct {
		Username string `json:"username" form:"username" query:"username" validate:"required,max=55"`
		Password string `json:"password" form:"password" query:"password" validate:"required,max=100"`
	}
)

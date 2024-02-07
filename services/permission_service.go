package services

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/COMTECH-63/fitness-management/repositories"
	"github.com/getsentry/sentry-go"
)

type (
	permissionService struct {
		permissionRepository repositories.PermissionRepository
	}
)

func NewPermissionService(
	permissionRepo repositories.PermissionRepository,
) PermissionService {
	return &permissionService{
		permissionRepository: permissionRepo,
	}
}

func (s permissionService) GetPermissions(ctx context.Context, span *sentry.Span, paginate database.Pagination, search string) (*database.Pagination, error) {
	childSpan := span.StartChild("GetPermissionsService")
	result, err := s.permissionRepository.GetPermissionPaginate(ctx, childSpan, paginate, search)
	childSpan.Finish()
	return result, err
}

func (s permissionService) GetPermission(ctx context.Context, span *sentry.Span, id int) (map[string]interface{}, error) {
	childSpan := span.StartChild("GetPermissionService")
	permission, err := s.permissionRepository.GetPermissionByID(ctx, childSpan, id)
	childSpan.Finish()

	return map[string]interface{}{"data": permission}, err
}

func (s permissionService) CreatePermission(ctx context.Context, span *sentry.Span, permissionDto *CreatePermissionDto) error {
	childSpan := span.StartChild("CreatePermissionService")
	permission := new(models.Permission)

	permission.Name = permissionDto.Name
	permission.Description = permissionDto.Description

	for _, userDto := range permissionDto.UserDtos {
		var user models.User

		user.ID = userDto.ID
		permission.Users = append(permission.Users, user)
	}

	for _, roleDto := range permissionDto.RoleDtos {
		var role models.Role

		role.ID = roleDto.ID
		permission.Roles = append(permission.Roles, role)
	}

	err := s.permissionRepository.CreatePermission(ctx, childSpan, permission)
	childSpan.Finish()

	return err
}

func (s permissionService) UpdatePermission(ctx context.Context, span *sentry.Span, id int, permissionDto *UpdatePermissionDto) error {
	childSpan := span.StartChild("UpdatePermissionService")
	permission := new(models.Permission)

	permission.Name = permissionDto.Name
	permission.Description = permissionDto.Description

	for _, userDto := range permissionDto.UserDtos {
		var user models.User

		user.ID = userDto.ID
		permission.Users = append(permission.Users, user)
	}

	for _, roleDto := range permissionDto.RoleDtos {
		var role models.Role

		role.ID = roleDto.ID
		permission.Roles = append(permission.Roles, role)
	}

	err := s.permissionRepository.UpdatePermission(ctx, childSpan, id, permission)
	childSpan.Finish()

	return err
}

func (s permissionService) DeletePermission(ctx context.Context, span *sentry.Span, id int) error {
	childSpan := span.StartChild("DeletePermissionService")
	err := s.permissionRepository.DeletePermission(ctx, childSpan, id)
	childSpan.Finish()

	return err
}

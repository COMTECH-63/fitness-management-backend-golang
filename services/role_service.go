package services

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/COMTECH-63/fitness-management/repositories"
	"github.com/getsentry/sentry-go"
)

type (
	roleService struct {
		roleRepository repositories.RoleRepository
	}
)

func NewRoleService(
	roleRepo repositories.RoleRepository,
) RoleService {
	return &roleService{
		roleRepository: roleRepo,
	}
}

func (s roleService) GetRoles(ctx context.Context, span *sentry.Span, paginate database.Pagination, search string) (*database.Pagination, error) {
	childSpan := span.StartChild("GetRolesService")
	result, err := s.roleRepository.GetRolePaginate(ctx, childSpan, paginate, search)
	childSpan.Finish()
	return result, err
}

func (s roleService) GetRole(ctx context.Context, span *sentry.Span, id int) (map[string]interface{}, error) {
	childSpan := span.StartChild("GetRoleService")
	role, err := s.roleRepository.GetRoleByID(ctx, childSpan, id)
	childSpan.Finish()

	return map[string]interface{}{"data": role}, err
}

func (s roleService) CreateRole(ctx context.Context, span *sentry.Span, roleDto *CreateRoleDto) error {

	childSpan := span.StartChild("CreateRoleService")

	role := new(models.Role)

	role.Name = roleDto.Name

	for _, userDto := range roleDto.UserDtos {
		var user models.User

		user.ID = userDto.ID
		role.Users = append(role.Users, user)
	}

	for _, permissionDto := range roleDto.PermissionDtos {
		var permission models.Permission

		permission.ID = permissionDto.ID
		role.Permissions = append(role.Permissions, permission)
	}

	err := s.roleRepository.CreateRole(ctx, childSpan, role)
	childSpan.Finish()

	return err
}

func (s roleService) UpdateRole(ctx context.Context, span *sentry.Span, id int, roleDto *UpdateRoleDto) error {
	childSpan := span.StartChild("UpdateRoleService")
	role := new(models.Role)

	role.Name = roleDto.Name

	for _, userDto := range roleDto.UserDtos {
		var user models.User

		user.ID = userDto.ID
		role.Users = append(role.Users, user)
	}

	for _, permissionDto := range roleDto.PermissionDtos {
		var permission models.Permission

		permission.ID = permissionDto.ID
		role.Permissions = append(role.Permissions, permission)
	} 
	
	err := s.roleRepository.UpdateRole(ctx, childSpan, id, role)
	childSpan.Finish()

	return err
}

func (s roleService) DeleteRole(ctx context.Context, span *sentry.Span, id int) error {
	childSpan := span.StartChild("DeleteRoleService")
	err := s.roleRepository.DeleteRole(ctx, childSpan, id)
	childSpan.Finish()

	return err
}

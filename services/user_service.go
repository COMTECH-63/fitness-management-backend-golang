package services

import (
	"context"

	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/models"
	"github.com/COMTECH-63/fitness-management/repositories"
	"github.com/getsentry/sentry-go"
)

type (
	userService struct {
		userRepository repositories.UserRepository
	}
)

func NewUserService(
	userRepo repositories.UserRepository,
) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s userService) GetUsers(ctx context.Context, span *sentry.Span, paginate database.Pagination, search string) (*database.Pagination, error) {
	childSpan := span.StartChild("GetUsersService")
	result, err := s.userRepository.GetUserPaginate(ctx, childSpan, paginate, search)
	childSpan.Finish()
	return result, err
}

func (s userService) GetUser(ctx context.Context, span *sentry.Span, id int) (map[string]interface{}, error) {
	childSpan := span.StartChild("GetUserService")
	user, err := s.userRepository.GetUserByID(ctx, childSpan, id)
	childSpan.Finish()

	return map[string]interface{}{"data": user}, err
}

func (s userService) CreateUser(ctx context.Context, span *sentry.Span, userDto *CreateUserDto) error {

	childSpan := span.StartChild("CreateUserService")

	user := new(models.User)

	// receiver method
	user.SetSex(userDto.Sex)

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.IDCard = userDto.IDCard
	user.Email = userDto.Email
	user.PhoneNumber = userDto.PhoneNumber
	user.Address = userDto.Address
	user.ImageURL = userDto.ImageURL
	user.MemberID = userDto.MemberID

	for _, roleDto := range userDto.RoleDtos {
		var role models.Role

		role.ID = roleDto.ID
		user.Roles = append(user.Roles, role)
	}

	for _, permissionDto := range userDto.PermissionDtos {
		var permission models.Permission

		permission.ID = permissionDto.ID
		user.Permissions = append(user.Permissions, permission)
	}

	for _, serviceDto := range userDto.ServiceDtos {
		var service models.Service

		service.ID = serviceDto.ID
		user.Services = append(user.Services, service)
	}

	err := s.userRepository.CreateUser(ctx, childSpan, user)
	childSpan.Finish()

	return err
}

func (s userService) UpdateUser(ctx context.Context, span *sentry.Span, id int, userDto *UpdateUserDto) error {
	childSpan := span.StartChild("UpdateUserService")
	user := new(models.User)

	// receiver method
	user.SetSex(userDto.Sex)

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.IDCard = userDto.IDCard
	user.Email = userDto.Email
	user.PhoneNumber = userDto.PhoneNumber
	user.Address = userDto.Address
	user.ImageURL = userDto.ImageURL

	for _, roleDto := range userDto.RoleDtos {
		var role models.Role

		role.ID = roleDto.ID
		user.Roles = append(user.Roles, role)
	}

	for _, permissionDto := range userDto.PermissionDtos {
		var permission models.Permission

		permission.ID = permissionDto.ID
		user.Permissions = append(user.Permissions, permission)
	}

	for _, serviceDto := range userDto.ServiceDtos {
		var service models.Service

		service.ID = serviceDto.ID
		user.Services = append(user.Services, service)
	}

	err := s.userRepository.UpdateUser(ctx, childSpan, id, user)
	childSpan.Finish()

	return err
}

func (s userService) DeleteUser(ctx context.Context, span *sentry.Span, id int) error {
	childSpan := span.StartChild("DeleteUserService")
	err := s.userRepository.DeleteUser(ctx, childSpan, id)
	childSpan.Finish()

	return err
}

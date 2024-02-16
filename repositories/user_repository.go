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

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}

func (r userRepository) GetUserPaginate(ctx context.Context, span *sentry.Span, pagination database.Pagination, search string) (*database.Pagination, error) {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "GetUserPaginateRepository", trace.WithAttributes(attribute.String("repository", "GetUserPaginate"), attribute.String("search", search)))
		users        []models.User
		err          error
	)

	// Pagination query
	if search != "" {
		if err = r.db.Scopes(database.Paginate(users, &pagination, r.db)).
			Where(`email LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Or(`first_name LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Or(`last_name LIKE ?`, fmt.Sprintf(`%%%s%%`, search)).
			Find(&users).Error; err != nil {
			log.Println(err)
			return nil, errors.New("GetUserPaginateError")
		}
	} else {
		if err = r.db.Scopes(database.Paginate(users, &pagination, r.db)).
			// Preload("Roles").Preload("Permissions").Preload("Services").Preload("Classes").Preload("Orders").Preload("Bookings").Preload("BookingClasses").Preload("BookingPersonalTrainers").Find(&users).Error; err != nil {
			Preload("Roles").Preload("Permissions").Preload("Services").Find(&users).Error; err != nil {
			return nil, err
		}
	}

	// Set data
	pagination.Data = users

	childSpan.End()

	return &pagination, nil
}

func (r userRepository) GetUserByID(ctx context.Context, span *sentry.Span, id int) (models.User, error) {
	var (
		_, childSpan = tracing.Tracer.Start(ctx, "GetUserByIDRepository", trace.WithAttributes(attribute.String("repository", "GetUserByID")))
		user         models.User
		err          error
	)

	// Query
	// if err = r.db.Preload("Roles").Preload("Permissions").Preload("Services").Preload("Classes").Preload("Orders").Preload("Bookings").Preload("BookingClasses").Preload("BookingPersonalTrainers").Find(&user, id).Error; err != nil {
	// 	return user, err
	// }
	if err = r.db.Preload("Roles").Preload("Permissions").Preload("Services").Find(&user, id).Error; err != nil {
		return user, err
	}

	childSpan.End()

	return user, nil
}

func (r userRepository) CreateUser(ctx context.Context, span *sentry.Span, user *models.User) error {
	var (
		childSpan = span.StartChild("CreateUserRepository")
		err       error
	)

	// Execute
	if err = r.db.Create(&user).Error; err != nil {
		return err
	}

	childSpan.Finish()

	return nil
}

func (r userRepository) UpdateUser(ctx context.Context, span *sentry.Span, id int, user *models.User) error {
	var (
		childSpan = span.StartChild("UpdateUserRepository")
		existUser *models.User
	)

	// Get model
	r.db.Find(&existUser, id)

	// Clear existing associations
	r.db.Model(&existUser).Association("Roles").Clear()
	r.db.Model(&existUser).Association("Permissions").Clear()
	r.db.Model(&existUser).Association("Services").Clear()
	// r.db.Model(&existUser).Association("Classes").Clear()
	// r.db.Model(&existUser).Association("Orders").Clear()
	// r.db.Model(&existUser).Association("Bookings").Clear()
	// r.db.Model(&existUser).Association("BookingClasses").Clear()
	// r.db.Model(&existUser).Association("BookingPersonalTrainers").Clear()

	// Set attributes
	// existUser.Username = user.Username
	// existUser.Password = user.Password
	existUser.FirstName = user.FirstName
	existUser.LastName = user.LastName
	existUser.IDCard = user.IDCard
	existUser.Email = user.Email
	existUser.PhoneNumber = user.PhoneNumber
	existUser.Address = user.Address
	existUser.Sex = user.Sex
	existUser.ImageURL = user.ImageURL
	existUser.MemberID = user.MemberID

	existUser.Roles = user.Roles
	existUser.Permissions = user.Permissions
	existUser.Services = user.Services
	// existUser.Classes = user.Classes
	// existUser.Orders = user.Orders
	// existUser.Bookings = user.Bookings
	// existUser.BookingClasses = user.BookingClasses
	// existUser.BookingPersonalTrainers = user.BookingPersonalTrainers

	// Execute
	if err := r.db.Save(&existUser).Error; err != nil {
		return err
	}

	childSpan.Finish()

	return nil
}

func (r userRepository) UpdatePasswordUser(ctx context.Context, span *sentry.Span, id int, user *models.User) error {
	var (
		childSpan = span.StartChild("UpdateUserRepository")
		existUser *models.User
	)

	// Get model
	r.db.Find(&existUser, id)

	// Set attributes
	existUser.Username = user.Username
	existUser.Password = user.Password

	// Execute
	if err := r.db.Save(&existUser).Error; err != nil {
		return err
	}

	childSpan.Finish()

	return nil
}

func (r userRepository) DeleteUser(ctx context.Context, span *sentry.Span, id int) error {
	var (
		childSpan = span.StartChild("DeleteUserRepository")
		err       error
	)

	// Execute
	if err = r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	childSpan.Finish()

	return nil
}

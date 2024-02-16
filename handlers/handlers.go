package handlers

import (
	"context"
	"log"

	"github.com/COMTECH-63/fitness-management/cache"
	"github.com/COMTECH-63/fitness-management/config"
	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/services"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

type (
	// Register handler services
	handler struct {
		cacher            *cache.Cache
		userService       services.UserService
		roleService       services.RoleService
		permissionService services.PermissionService
	}
	// Register handler interfaces
	Handler interface {
		UserHandler
		RoleHandler
		PermissionHandler
	}
)

func NewHandler(
	cacher *cache.Cache,
	userService services.UserService,
	roleService services.RoleService,
	permissionService services.PermissionService,
) handler {
	return handler{
		cacher:            cacher,
		userService:       userService,
		roleService:       roleService,
		permissionService: permissionService,
	}
}

type ServicePaginationFunc func(ctx context.Context, span *sentry.Span, paginate database.Pagination, search string) (*database.Pagination, error)
type ServiceQueryFunc func(ctx context.Context, span *sentry.Span, id int) (map[string]interface{}, error)
type ServiceQueryByStringParamFunc func(ctx context.Context, span *sentry.Span, id string) (map[string]interface{}, error)

func (h handler) PaginationCache(c *fiber.Ctx, span *sentry.Span, key string, tags []string, paginate database.Pagination, search string, f ServicePaginationFunc) (*database.Pagination, error) {
	var (
		responseData *database.Pagination
		err          error
	)

	// Get the cached attributes object
	err = cache.Cacher.Get(c.Context(), key, &responseData)
	if err != nil {
		return nil, err
	}

	if responseData == nil {
		// Call service function
		responseData, err = f(c.Context(), span, paginate, search)
		if err != nil {
			return nil, err
		}

		// Set cache
		err = cache.Cacher.Tag(tags...).Set(c.Context(), key, &responseData)
		if err != nil {
			return nil, err
		}
	}

	return responseData, nil
}

func (h handler) QueryCache(c *fiber.Ctx, span *sentry.Span, key string, tags []string, id int, f ServiceQueryFunc) (map[string]interface{}, error) {
	var (
		responseData map[string]interface{}
		err          error
	)

	// Get the cached attributes object
	err = cache.Cacher.Get(c.Context(), key, &responseData)
	if err != nil {
		return nil, err
	}

	if responseData == nil {
		// Call service function
		responseData, err = f(c.Context(), span, id)
		if err != nil {
			return nil, err
		}

		// Set cache
		err = cache.Cacher.Tag(tags...).Set(c.Context(), key, &responseData)
		if err != nil {
			return nil, err
		}
	}

	return responseData, nil
}

// Root handlers  ------------------------------------------------------------------

func GetRootPath(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString(config.AppConfig.AppName)
}

func GetHealthCheck(c *fiber.Ctx) error {
	sqlDB, _ := database.DBConn.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("DatabaseError:", err)
	}

	return c.SendStatus(fiber.StatusOK)
}

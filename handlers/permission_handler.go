package handlers

import (
	"fmt"

	"github.com/COMTECH-63/fitness-management/cache"
	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/services"
	"github.com/COMTECH-63/fitness-management/utils"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

type (
	PermissionHandler interface {
		// Permission handlers
		GetPermissions(c *fiber.Ctx) error
		GetPermission(c *fiber.Ctx) error
		CreatePermission(c *fiber.Ctx) error
		UpdatePermission(c *fiber.Ctx) error
		DeletePermission(c *fiber.Ctx) error
	}
)

func (h handler) GetPermissions(c *fiber.Ctx) error {
	var (
		span         = sentry.StartSpan(c.Context(), "GetPermissionsHandler", sentry.WithTransactionName("GetPermissions"))
		responseData *database.Pagination
	)

	// Get paginate values
	paginate := database.Pagination{
		Page:  c.QueryInt("page", 1),
		Limit: c.QueryInt("limit", 20),
	}
	search := c.Query("search")

	// Make cache key
	cacheTags := []string{"Permissions"}
	cacheKey := fmt.Sprintf("GetPermissions_%d_%d", paginate.Page, paginate.Limit)
	if search != "" {
		cacheKey = fmt.Sprintf(`%s_%s`, cacheKey, search)
	}

	responseData, err := h.PaginationCache(c, span, cacheKey, cacheTags, paginate, search, h.permissionService.GetPermissions)
	if err != nil {
		utils.HandleErrors(err)
		return fiber.ErrInternalServerError
	}

	span.Finish()
	return c.JSON(responseData)
}

func (h handler) GetPermission(c *fiber.Ctx) error {
	var (
		span         = sentry.StartSpan(c.Context(), "GetPermissionHandler", sentry.WithTransactionName("GetPermission"))
		responseData map[string]interface{}
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Make cache key
	cacheTags := []string{"Permissions"}
	cacheKey := fmt.Sprintf("GetPermission_%d", id)

	responseData, err := h.QueryCache(c, span, cacheKey, cacheTags, id, h.permissionService.GetPermission)
	if err != nil {
		utils.HandleErrors(err)
		return fiber.ErrInternalServerError
	}

	span.Finish()
	return c.JSON(responseData)
}

func (h handler) CreatePermission(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "CreatePermissionHandler", sentry.WithTransactionName("CreatePermission"))
	)

	// Create data transfer object
	PermissionDto := new(services.CreatePermissionDto)

	// Parse HTTP request body to struct variable
	if err := c.BodyParser(PermissionDto); err != nil {
		return err
	}

	// Form request validation
	errors := utils.Validate(*PermissionDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Call service function
	err := h.permissionService.CreatePermission(c.Context(), span, PermissionDto)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear Permission cache
	cache.Cacher.Tag("Permissions").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

func (h handler) UpdatePermission(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "UpdatePermissionHandler", sentry.WithTransactionName("UpdatePermission"))
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Create data transfer object
	PermissionDto := new(services.UpdatePermissionDto)

	// Parse HTTP request body to struct variable
	if err := c.BodyParser(PermissionDto); err != nil {
		return err
	}

	// Form request validation
	errors := utils.Validate(*PermissionDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Call service function
	err := h.permissionService.UpdatePermission(c.Context(), span, id, PermissionDto)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear Permission cache
	cache.Cacher.Tag("Permissions").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

func (h handler) DeletePermission(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "DeletePermissionHandler", sentry.WithTransactionName("DeletePermission"))
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Call service function
	err := h.permissionService.DeletePermission(c.Context(), span, id)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear Permission cache
	cache.Cacher.Tag("Permissions").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

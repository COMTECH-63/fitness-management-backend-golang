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
	RoleHandler interface {
		// Role handlers
		GetRoles(c *fiber.Ctx) error
		GetRole(c *fiber.Ctx) error
		CreateRole(c *fiber.Ctx) error
		UpdateRole(c *fiber.Ctx) error
		DeleteRole(c *fiber.Ctx) error
	}
)

func (h handler) GetRoles(c *fiber.Ctx) error {
	var (
		span         = sentry.StartSpan(c.Context(), "GetRolesHandler", sentry.WithTransactionName("GetRoles"))
		responseData *database.Pagination
	)

	// Get paginate values
	paginate := database.Pagination{
		Page:  c.QueryInt("page", 1),
		Limit: c.QueryInt("limit", 20),
	}
	search := c.Query("search")

	// Make cache key
	cacheTags := []string{"roles"}
	cacheKey := fmt.Sprintf("GetRoles_%d_%d", paginate.Page, paginate.Limit)
	if search != "" {
		cacheKey = fmt.Sprintf(`%s_%s`, cacheKey, search)
	}

	responseData, err := h.PaginationCache(c, span, cacheKey, cacheTags, paginate, search, h.roleService.GetRoles)
	if err != nil {
		utils.HandleErrors(err)
		return fiber.ErrInternalServerError
	}

	span.Finish()
	return c.JSON(responseData)
}

func (h handler) GetRole(c *fiber.Ctx) error {
	var (
		span         = sentry.StartSpan(c.Context(), "GetRoleHandler", sentry.WithTransactionName("GetRole"))
		responseData map[string]interface{}
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Make cache key
	cacheTags := []string{"roles"}
	cacheKey := fmt.Sprintf("GetRole_%d", id)

	responseData, err := h.QueryCache(c, span, cacheKey, cacheTags, id, h.roleService.GetRole)
	if err != nil {
		utils.HandleErrors(err)
		return fiber.ErrInternalServerError
	}

	span.Finish()
	return c.JSON(responseData)
}

func (h handler) CreateRole(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "CreateRoleHandler", sentry.WithTransactionName("CreateRole"))
	)

	// Create data transfer object
	roleDto := new(services.CreateRoleDto)

	// Parse HTTP request body to struct variable
	if err := c.BodyParser(roleDto); err != nil {
		return err
	}

	// Form request validation
	errors := utils.Validate(*roleDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Call service function
	err := h.roleService.CreateRole(c.Context(), span, roleDto)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear role cache
	cache.Cacher.Tag("roles").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

func (h handler) UpdateRole(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "UpdateRoleHandler", sentry.WithTransactionName("UpdateRole"))
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Create data transfer object
	roleDto := new(services.UpdateRoleDto)

	// Parse HTTP request body to struct variable
	if err := c.BodyParser(roleDto); err != nil {
		return err
	}

	// Form request validation
	errors := utils.Validate(*roleDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Call service function
	err := h.roleService.UpdateRole(c.Context(), span, id, roleDto)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear role cache
	cache.Cacher.Tag("roles").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

func (h handler) DeleteRole(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "DeleteRoleHandler", sentry.WithTransactionName("DeleteRole"))
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Call service function
	err := h.roleService.DeleteRole(c.Context(), span, id)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear role cache
	cache.Cacher.Tag("roles").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

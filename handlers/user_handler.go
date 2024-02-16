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
	UserHandler interface {
		// User handlers
		GetUsers(c *fiber.Ctx) error
		GetUser(c *fiber.Ctx) error
		CreateUser(c *fiber.Ctx) error
		UpdateUser(c *fiber.Ctx) error
		UpdatePasswordUser(c *fiber.Ctx) error
		DeleteUser(c *fiber.Ctx) error
	}
)

func (h handler) GetUsers(c *fiber.Ctx) error {
	var (
		span         = sentry.StartSpan(c.Context(), "GetUsersHandler", sentry.WithTransactionName("GetUsers"))
		responseData *database.Pagination
	)

	// Get paginate values
	paginate := database.Pagination{
		Page:  c.QueryInt("page", 1),
		Limit: c.QueryInt("limit", 20),
	}
	search := c.Query("search")

	// Make cache key
	cacheTags := []string{"users"}
	cacheKey := fmt.Sprintf("GetUsers_%d_%d", paginate.Page, paginate.Limit)
	if search != "" {
		cacheKey = fmt.Sprintf(`%s_%s`, cacheKey, search)
	}

	responseData, err := h.PaginationCache(c, span, cacheKey, cacheTags, paginate, search, h.userService.GetUsers)
	if err != nil {
		utils.HandleErrors(err)
		return fiber.ErrInternalServerError
	}

	span.Finish()
	return c.JSON(responseData)
}

func (h handler) GetUser(c *fiber.Ctx) error {
	var (
		span         = sentry.StartSpan(c.Context(), "GetUserHandler", sentry.WithTransactionName("GetUser"))
		responseData map[string]interface{}
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Make cache key
	cacheTags := []string{"users"}
	cacheKey := fmt.Sprintf("GetUser_%d", id)

	responseData, err := h.QueryCache(c, span, cacheKey, cacheTags, id, h.userService.GetUser)
	if err != nil {
		utils.HandleErrors(err)
		return fiber.ErrInternalServerError
	}

	span.Finish()
	return c.JSON(responseData)
}

func (h handler) CreateUser(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "CreateUserHandler", sentry.WithTransactionName("CreateUser"))
	)

	// Create data transfer object
	userDto := new(services.CreateUserDto)

	// Parse HTTP request body to struct variable
	if err := c.BodyParser(userDto); err != nil {
		return err
	}

	// Form request validation
	errors := utils.Validate(*userDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Call service function
	err := h.userService.CreateUser(c.Context(), span, userDto)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear user cache
	cache.Cacher.Tag("users").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

func (h handler) UpdateUser(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "UpdateUserHandler", sentry.WithTransactionName("UpdateUser"))
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Create data transfer object
	userDto := new(services.UpdateUserDto)

	// Parse HTTP request body to struct variable
	if err := c.BodyParser(userDto); err != nil {
		return err
	}

	// Form request validation
	errors := utils.Validate(*userDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Call service function
	err := h.userService.UpdateUser(c.Context(), span, id, userDto)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear user cache
	cache.Cacher.Tag("users").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

func (h handler) UpdatePasswordUser(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "UpdateUserHandler", sentry.WithTransactionName("UpdateUser"))
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Create data transfer object
	userDto := new(services.UpdatePasswordUserdDto)

	// Parse HTTP request body to struct variable
	if err := c.BodyParser(userDto); err != nil {
		return err
	}

	// Form request validation
	errors := utils.Validate(*userDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Call service function
	err := h.userService.UpdatePasswordUser(c.Context(), span, id, userDto)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear user cache
	cache.Cacher.Tag("users").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

func (h handler) DeleteUser(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "DeleteUserHandler", sentry.WithTransactionName("DeleteUser"))
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Call service function
	err := h.userService.DeleteUser(c.Context(), span, id)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear user cache
	cache.Cacher.Tag("users").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

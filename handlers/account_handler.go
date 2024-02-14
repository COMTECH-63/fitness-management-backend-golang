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
	AccountHandler interface {
		// Account handlers
		GetAccounts(c *fiber.Ctx) error
		GetAccount(c *fiber.Ctx) error
		CreateAccount(c *fiber.Ctx) error
		UpdateAccount(c *fiber.Ctx) error
		DeleteAccount(c *fiber.Ctx) error
	}
)

func (h handler) GetAccounts(c *fiber.Ctx) error {
	var (
		span         = sentry.StartSpan(c.Context(), "GetAccountsHandler", sentry.WithTransactionName("GetAccounts"))
		responseData *database.Pagination
	)

	// Get paginate values
	paginate := database.Pagination{
		Page:  c.QueryInt("page", 1),
		Limit: c.QueryInt("limit", 20),
	}
	search := c.Query("search")

	// Make cache key
	cacheTags := []string{"Accounts"}
	cacheKey := fmt.Sprintf("GetAccounts_%d_%d", paginate.Page, paginate.Limit)
	if search != "" {
		cacheKey = fmt.Sprintf(`%s_%s`, cacheKey, search)
	}

	responseData, err := h.PaginationCache(c, span, cacheKey, cacheTags, paginate, search, h.accountService.GetAccounts)
	if err != nil {
		utils.HandleErrors(err)
		return fiber.ErrInternalServerError
	}

	span.Finish()
	return c.JSON(responseData)
}

func (h handler) GetAccount(c *fiber.Ctx) error {
	var (
		span         = sentry.StartSpan(c.Context(), "GetAccountHandler", sentry.WithTransactionName("GetAccount"))
		responseData map[string]interface{}
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Make cache key
	cacheTags := []string{"Accounts"}
	cacheKey := fmt.Sprintf("GetAccount_%d", id)

	responseData, err := h.QueryCache(c, span, cacheKey, cacheTags, id, h.accountService.GetAccount)
	if err != nil {
		utils.HandleErrors(err)
		return fiber.ErrInternalServerError
	}

	span.Finish()
	return c.JSON(responseData)
}

func (h handler) CreateAccount(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "CreateAccountHandler", sentry.WithTransactionName("CreateAccount"))
	)

	// Create data transfer object
	AccountDto := new(services.AccountDto)

	// Parse HTTP request body to struct variable
	if err := c.BodyParser(AccountDto); err != nil {
		return err
	}

	// Form request validation
	errors := utils.Validate(*AccountDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Call service function
	err := h.accountService.CreateAccount(c.Context(), span, AccountDto)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear Account cache
	cache.Cacher.Tag("Accounts").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

func (h handler) UpdateAccount(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "UpdateAccountHandler", sentry.WithTransactionName("UpdateAccount"))
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Create data transfer object
	AccountDto := new(services.AccountDto)

	// Parse HTTP request body to struct variable
	if err := c.BodyParser(AccountDto); err != nil {
		return err
	}

	// Form request validation
	errors := utils.Validate(*AccountDto)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Call service function
	err := h.accountService.UpdateAccount(c.Context(), span, id, AccountDto)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear Account cache
	cache.Cacher.Tag("Accounts").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

func (h handler) DeleteAccount(c *fiber.Ctx) error {
	var (
		span = sentry.StartSpan(c.Context(), "DeleteAccountHandler", sentry.WithTransactionName("DeleteAccount"))
	)

	// Get route parameter
	id, _ := c.ParamsInt("id")

	// Call service function
	err := h.accountService.DeleteAccount(c.Context(), span, id)
	if err != nil {
		utils.HandleErrors(err)
		return err
	}

	// Clear Account cache
	cache.Cacher.Tag("Accounts").Flush(c.Context())

	span.Finish()
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"code":    "0",
		"message": "OK",
	})
}

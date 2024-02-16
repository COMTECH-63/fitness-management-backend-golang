package routes

import (
	"github.com/COMTECH-63/fitness-management/cache"
	"github.com/COMTECH-63/fitness-management/config"
	"github.com/COMTECH-63/fitness-management/database"
	"github.com/COMTECH-63/fitness-management/handlers"
	"github.com/COMTECH-63/fitness-management/microservices"
	"github.com/COMTECH-63/fitness-management/repositories"
	"github.com/COMTECH-63/fitness-management/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func HTTPRootMiddleware(ms *microservices.Microservice) {
	// Default middleware config
	fiberRequestID := requestid.New()
	fiberLogger := logger.New(config.LoggerConfig)
	fiberRecover := recover.New()
	fiberLimiter := limiter.New(config.LimiterConfig)
	fiberETag := etag.New(config.ETagConfig)
	fiberCors := cors.New(config.CorsConfig)
	fiberFavicon := favicon.New(config.FaviconConfig)

	ms.Use(fiberRequestID)
	ms.Use(fiberLogger)
	ms.Use(fiberRecover)
	ms.Use(fiberLimiter)
	ms.Use(fiberETag)
	ms.Use(fiberCors)
	ms.Use(fiberFavicon)
}

func HTTPRootRoute(ms *microservices.Microservice) {
	HTTPRootMiddleware(ms)

	ms.GET("/", func(c *fiber.Ctx) error {
		return handlers.GetRootPath(c)
	})
	ms.GET("/health", func(c *fiber.Ctx) error {
		return handlers.GetHealthCheck(c)
	})
	ms.GET("/monitor", monitor.New(monitor.Config{Title: "Fiber"}))
}

func HTTPRoutes(ms *microservices.Microservice) {
	// Initialize repositories, services, and handlers
	userRepo := repositories.NewUserRepository(database.DBConn)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize repositories, services, and handlers
	roleRepo := repositories.NewRoleRepository(database.DBConn)

	// Initialize services
	roleService := services.NewRoleService(roleRepo)

	// Initialize repositories, services, and handlers
	permissionRepo := repositories.NewPermissionRepository(database.DBConn)

	// Initialize services
	permissionService := services.NewPermissionService(permissionRepo)

	// Initialize handlers
	handler := handlers.NewHandler(
		cache.Cacher,
		userService,
		roleService,
		permissionService,
	)

	// REST API endpoint ------------------------------------------------------------------

	// Example route grouping
	api := ms.Group("api")
	apiV1 := api.Group("v1")

	// User service routes
	apiV1.Get("/users", func(c *fiber.Ctx) error { return handler.GetUsers(c) })
	apiV1.Get("/users/:id", func(c *fiber.Ctx) error { return handler.GetUser(c) })
	apiV1.Post("/users", func(c *fiber.Ctx) error { return handler.CreateUser(c) })
	apiV1.Put("/users/:id", func(c *fiber.Ctx) error { return handler.UpdateUser(c) })
	apiV1.Put("/users/passwords/:id", func(c *fiber.Ctx) error { return handler.UpdatePasswordUser(c) })
	apiV1.Delete("/users/:id", func(c *fiber.Ctx) error { return handler.DeleteUser(c) })

	// Role service routes
	apiV1.Get("/roles", func(c *fiber.Ctx) error { return handler.GetRoles(c) })
	apiV1.Get("/roles/:id", func(c *fiber.Ctx) error { return handler.GetRole(c) })
	apiV1.Post("/roles", func(c *fiber.Ctx) error { return handler.CreateRole(c) })
	apiV1.Put("/roles/:id", func(c *fiber.Ctx) error { return handler.UpdateRole(c) })
	apiV1.Delete("/roles/:id", func(c *fiber.Ctx) error { return handler.DeleteRole(c) })

	// Permission service routes
	apiV1.Get("/permissions", func(c *fiber.Ctx) error { return handler.GetPermissions(c) })
	apiV1.Get("/permissions/:id", func(c *fiber.Ctx) error { return handler.GetPermission(c) })
	apiV1.Post("/permissions", func(c *fiber.Ctx) error { return handler.CreatePermission(c) })
	apiV1.Put("/permissions/:id", func(c *fiber.Ctx) error { return handler.UpdatePermission(c) })
	apiV1.Delete("/permissions/:id", func(c *fiber.Ctx) error { return handler.DeletePermission(c) })

}

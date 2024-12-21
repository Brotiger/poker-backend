package router

import (
	swagger "github.com/Brotiger/per-painted_poker-backend/docs/swagger"
	"github.com/Brotiger/per-painted_poker-backend/internal/config"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRouter(app *fiber.App) *fiber.Router {
	swagger.SwaggerInfo.Host = config.Cfg.Swagger.Host
	swagger.SwaggerInfo.Version = config.TagVersion

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	router := app.Group("/api")
	router = setupAuthRouter(router)

	return &router
}

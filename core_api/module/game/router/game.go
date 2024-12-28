package router

import (
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/game/handler"
	"github.com/Brotiger/per-painted_poker-backend/core_api/module/game/middleware"
	sharedMiddleware "github.com/Brotiger/per-painted_poker-backend/core_api/shared/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(api fiber.Router) {
	gameHandler := handler.NewGame()

	authMiddleware := sharedMiddleware.NewShared()
	gameMiddleware := middleware.NewGameMiddleware()

	game := api.Group("/game")
	game.Use(authMiddleware.Token)
	game.Get("", gameHandler.List)
	game.Post("", gameMiddleware.AlreadyHasGame, gameHandler.Create)
	game.Post("/start", gameHandler.Start)
}
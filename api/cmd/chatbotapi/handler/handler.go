package handlers

import (
	"easyaichat/cmd/chatbotapi/handler/aigenerate_hdl"
	"easyaichat/internal/web"
	"easyaichat/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// Add this import
)

func CreateHandlers() *fiber.App {
	app := fiber.New()

	// Enable CORS for all routes
	app.Use(cors.New())
	// Public routes
	public := app.Group("/api")
	registerPublicHandlers(public)

	private := app.Group("/api/private", middleware.AuthRequired)
	registerPrivateHandlers(private)

	return app
}

// registerPublicHandlers register public handlers and Init them
func registerPublicHandlers(root fiber.Router) {

	handlers := web.HandlerRegistrator{}

	// TODO: register your handlers here
	handlers.Register(
		new(PublicHandler),
		new(aigenerate_hdl.AiGenerateHandler),
	)

	handlers.Init(root)
}

// registerPrivateHandlers register private handlers and Init them
func registerPrivateHandlers(root fiber.Router) {

	handlers := web.HandlerRegistrator{}
	handlers.Register(
		new(PrivateHandler),
	)

	handlers.Init(root)

}

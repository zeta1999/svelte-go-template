package routes

import (
	config "homefill/backend/config"
	"homefill/backend/logs"

	"github.com/gofiber/fiber/v2"
)

// StartServer : start server on port and listen to routes
func StartServer() {
	app := fiber.New()

	// Universally Accessible Routes
	app.Get("/", HomeRoute)
	app.Get("/login", LoginRoute)
	app.Get("/callback", AuthCallBack)

	// -------------- PROTECTED --------------

	// User Routes
	app.Get("/user", GetUserInfo)

	// ---------------------------------------

	// Starting Server
	err := app.Listen(config.PORT)
	if err != nil {
		logs.LogIt(logs.LogFatal, "StartServer", "unable to start server", err)
	}
}

package routes

import (
	"api/middleware"
	"api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	//users
	app.Get("/api/users", controllers.AllUsers )
	app.Get("/api/users/:id", controllers.GetUser )
	app.Post("/api/users", controllers.CreateUser )
	app.Put("/api/users/:id", controllers.UpdateUser )
	app.Delete("/api/users/:id", controllers.DeleteUser )

	//roles
	app.Get("/api/roles", controllers.AllRoles )
	app.Get("/api/roles/:id", controllers.GetRole )
	app.Post("/api/roles", controllers.CreateRole )
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole )

	//permissions
	app.Get("/api/permissions", controllers.AllPermissions)
	// app.Post("/api/permissions", controllers.CreatePermission)
}
package routes

import (
	"api/controllers"
	"api/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// app.Post("/api/roles", controllers.CreateRole )
	app.Use(middleware.IsAuthenticated)

	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	//users
	app.Get("/api/users", controllers.AllUsers)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Post("/api/users", controllers.CreateUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	//roles
	app.Get("/api/roles", controllers.AllRoles)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Post("/api/roles", controllers.CreateRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	//permissions
	app.Get("/api/permissions", controllers.AllPermissions)
	app.Post("/api/permissions", controllers.CreatePermission)

	//Products
	app.Get("/api/products", controllers.AllProducts)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Post("/api/products", controllers.CreateProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

	//Orders
	app.Get("/api/orders", controllers.AllOrders)
	app.Post("/api/export", controllers.ExportCsv)
	app.Get("/api/chart", controllers.Chart)

	app.Post("/api/upload", controllers.UploadImage)
	app.Static("/api/uploads", "./uploads")

	
}

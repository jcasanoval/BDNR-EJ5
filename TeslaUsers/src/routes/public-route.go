package routes

import (
	"teslaUsers/src/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App, controller *controllers.UserController) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/users", controller.GetUsers)      // get list of all users

	route.Get("/users/:email", controller.GetUser) // get one user by email

	// Routes for POST method:
	route.Post("/users", controller.AddUser) // register a new user
	
	route.Post("/payment_methods", controller.AddPaymentMethod) // register a new payment method
}

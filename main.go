package main

import (
	"os"
	"context"

	"teslaUsers/src/configs"
	"teslaUsers/src/controllers"
	"teslaUsers/src/controllers/repointerfaces"
	"teslaUsers/src/datasources"
	"teslaUsers/src/middlewares"
	"teslaUsers/src/models"
	"teslaUsers/src/repositories"
	"teslaUsers/src/routes"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middlewares.FiberMiddleware(app) // Register Fiber's middleware for app.

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mongoURI := os.Getenv("MONGO_URI")
	mongoClient, err := datasources.NewMongoDataSource(mongoURI)

	if err != nil {
		panic(err)
	}

	defer mongoClient.Disconnect(context.TODO())

	repo := repositories.NewUserMongoRepo(mongoClient, "TeslaDB")

	// Esto es para cargar la db con algunos datos.
	// loadUsers(repo)

	// Creo una instancia de mis controladores con mi instancia de repo
	controller := controllers.NewUserController(repo)

	// Routes.
	routes.PublicRoutes(app, controller)

	// Aqui inicializamos el servidor en el puerto 8080
	app.Listen(":8080")
}

func loadUsers(repo repointerfaces.UserRepo) {
	repo.AddUser(models.User{
		EMail    : "juan@gmail.com",
		UserName : "BigJuan",
		Name     : "Juan",
		LastName : "Casanova",
		BirthDay : "26-01-1999",
		Country  : "Uruguay",
		Language : "es",
		PaymentMethods : []string{"visa", "mastercard"},
	})
}

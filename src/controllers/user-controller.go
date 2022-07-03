package controllers

import (
	"teslaUsers/src/controllers/repointerfaces"
	"teslaUsers/src/models"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	repo repointerfaces.UserRepo
}

func NewUserController(repo repointerfaces.UserRepo) *UserController {
	return &UserController{repo: repo}
}

func (controller *UserController) GetUsers(c *fiber.Ctx) error {

	// Get all users.
	users, err := controller.repo.GetUsers()
	if err != nil {
		// Return, if users not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"count": 0,
			"books": nil,
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(users),
		"users": users,
	})
}

func (controller *UserController) GetUser(c *fiber.Ctx) error {

	user, err := controller.repo.GetUser(c.Params("email"))
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"user":  nil,
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

func (controller *UserController) AddUser(c *fiber.Ctx) error {
	var user models.User
	// Tengo que pasar la direccion de memoria asi va a pone las variables
	// Que vienen en el contexto declaradas segun el json del modelo
	fmt.Println(string(c.Body()))

	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		// Return, if user has invalid fields
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"user":  nil,
		})
	}

	// Add new user
	user, err = controller.repo.AddUser(user)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"user":  nil,
		})
	}

	// Return status 201 Created.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "User created succesfully!",
		"user":  user,
	})
}

func (controller *UserController) AddPaymentMethod(c *fiber.Ctx) error {
	var paymentMethodInsert models.PaymentMethodInsert
	fmt.Println(string(c.Body()))

	err := json.Unmarshal(c.Body(), &paymentMethodInsert)
	if err != nil {
		// Return, if user has invalid fields
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Add new payment method
	err = controller.repo.AddPaymentMethod(paymentMethodInsert.EMail, paymentMethodInsert.PaymentMethod)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201 Created.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "Payment method added succesfully!",
	})
}

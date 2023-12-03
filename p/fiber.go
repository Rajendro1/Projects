package main

import "github.com/gofiber/fiber/v2"

func fiberRoute() {
	app := fiber.New()

	// Routes
	app.Get("/users", GetUsers)
	app.Post("/users", CreateUser)

	// Start the Fiber app
	err := app.Listen(":3000")
	if err != nil {
		panic("Failed to start the server")
	}
}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetUsers handles the GET request to retrieve a list of users
func GetUsers(c *fiber.Ctx) error {
	users := []User{
		{ID: 1, Name: "John Doe", Age: 25},
		{ID: 2, Name: "Raj Smith", Age: 30},
		// Add more users as needed
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": users, "message": "SUccessfully get users"})
}

// CreateUser handles the POST request to create a new user
func CreateUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// For demonstration purposes, just return the received user without storing it
	return c.Status(fiber.StatusCreated).JSON(user)
}

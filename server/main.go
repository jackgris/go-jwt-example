package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	app := fiber.New()

	// Login route
	app.Post("/login", login)

	// Unauthenticated route
	app.Get("/", accessible)

	// JWT Middleware
	// You have to know that the Config object we are passing to the middleware says this:
	// Context key to store user information from the token into context.
	// Optional. Default: "user".
	// ContextKey string
	// And in our middelware we have the following code:
	// if err == nil && token.Valid {
	// 	// Store user information from token into context.
	// 	c.Locals(cfg.ContextKey, token)
	// 	return cfg.SuccessHandler(c)
	// }
	// For that reason, we can use this code:
	// user := c.Locals("user").(*jwt.Token)
	// And with that, our restricted handler checks if the token has permission.
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)

	fmt.Println(app.Listen(":3000"))
}

func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	fmt.Printf("/login we received user %s and pass %s\n", user, pass)

	// Off course, here you should use a database or some service where you have the user data saved.
	if user != "john" || pass != "doe" {
		// Throws Unauthorized error
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		// Here you config the time that the token will be ok before being invalidated.
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	fmt.Printf("/restricted receive user token: %s\n", user.Raw)
	fmt.Printf("And we know that his name is: %s\n", name)

	return c.SendString("Welcome " + name + "\n")
}

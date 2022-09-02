package main

import (
	"StationedAtAuth/sms"

	"github.com/golang-jwt/jwt"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type SignupRequest struct {
	Name     string
	Email    string
	Password string
}

func main() {
	if err := sms.Notify("Hey..."); err != nil {
		panic(err)
	}
	/*app := fiber.New()

	user := user.NewUser{Email: "no@email.com", Password: "Noicann0t!"}
	user.CreateNewUser()

	_, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	app.Post("/signup", func(c *fiber.Ctx) error {
		req := new(SignupRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		if req.Name == "" || req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid signup credentials")
		}

		// save this info in the database

		//create a jwt token

		return nil
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/private", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path": "private"})
	})

	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "path": "public"})
	})

	app.Post("/notify_error", func(c *fiber.Ctx) error {
		if err := sms.Notify("Hello from Go!"); err != nil {
			return c.JSON(fiber.Map{"success": false})
		}
		return c.JSON(fiber.Map{"success": true})
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}*/

}

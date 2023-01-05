package controllers

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/francis560/api-go/models"
	"github.com/gofiber/fiber/v2"
)

func RandUsersHandler(c *fiber.Ctx) error {

	quantity, _ := c.ParamsInt("quantity")

	const maxQuantity int = 350

	var randUsers []models.User

	if quantity > maxQuantity {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "The limit of users that can bring is 100",
		})
	}

	for i := 0; i < quantity; i++ {

		data := models.User{
			First_Name: gofakeit.FirstName(),
			Last_Name:  gofakeit.LastName(),
			Email:      gofakeit.Email(),
			Gender:     gofakeit.Gender(),
			Phone:      gofakeit.Phone(),
			City:       gofakeit.City(),
			Country:    gofakeit.Country(),
			State:      gofakeit.State(),
			Street:     gofakeit.Street(),
			Profile:    gofakeit.ImageURL(500, 500),
		}

		randUsers = append(randUsers, data)
	}

	return c.Status(fiber.StatusOK).JSON(randUsers)
}

package controllers

import (
	"github.com/faizan-glitch/url-sh/pkg/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Resolve(c *fiber.Ctx) error {

	shortCode := c.Params("short")

	password := c.Query("p")

	var model models.UrlSchema

	if err := model.FindByShortCode(shortCode); err != nil {
		return c.Status(404).SendString(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(password)); err != nil {
		return c.Status(403).SendString("Forbidden")
	}

	return c.Redirect(model.LongUrl)
}

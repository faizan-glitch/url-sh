package controllers

import (
	"github.com/faizan-glitch/url-sh/pkg/models"
	"github.com/faizan-glitch/url-sh/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type ShortenRequest struct {
	Url       string `json:"url"`
	Protected bool   `json:"protected"`
	Password  string `json:"password"`
}

func Shorten(c *fiber.Ctx) error {
	body := ShortenRequest{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	code := utils.ShortCode(6)

	shortUrl := utils.ShortUrl(code)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), viper.GetInt("BCRYPT_ROUNDS"))

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	model := &models.UrlSchema{
		LongUrl:   body.Url,
		ShortUrl:  shortUrl,
		ShortCode: code,
		Password:  string(hashedPassword),
	}

	if err := model.Insert(); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(model)
}

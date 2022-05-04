package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func ShortCode(n uint8) string {

	rand.Seed(time.Now().UnixMicro())

	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func ShortUrl(shortCode string) string {
	return strings.Trim(viper.GetString("DOMAIN"), "/") + "/" + shortCode
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), viper.GetInt("BCRYPT_ROUNDS"))
}

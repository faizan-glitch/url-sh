package commands

import (
	"log"

	"github.com/faizan-glitch/url-sh/pkg/models"
	"github.com/faizan-glitch/url-sh/pkg/utils"
	"github.com/spf13/cobra"
)

var shorten = cobra.Command(cobra.Command{
	Use:       "shorten [flags] url",
	Args:      cobra.ExactArgs(1),
	Short:     "Shorten a url",
	Long:      `Generates a random code and returns a url with the short code.`,
	Aliases:   []string{"s"},
	ValidArgs: []string{"-p", "--password"},
	Example:   `url-sh shorten https://example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		password := cmd.Flag("password").Value.String()

		if password != "" {
			hp, _ := utils.HashPassword(password)
			password = string(hp)
		}

		code := utils.ShortCode(6)

		shortUrl := utils.ShortUrl(code)

		model := &models.UrlSchema{
			LongUrl:   args[0],
			ShortUrl:  shortUrl,
			ShortCode: code,
			Password:  password,
		}

		if err := model.Insert(); err != nil {
			log.Println("Could not insert url", err.Error())
			return
		}

		log.Println("Shortened url:", model.ShortUrl)
	},
})

func GetShortenCommand() *cobra.Command {
	shorten.PersistentFlags().StringP("password", "p", "", "Password to protect the url")
	return &shorten
}

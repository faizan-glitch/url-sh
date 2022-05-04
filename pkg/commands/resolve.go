package commands

import (
	"log"
	"sync"

	"github.com/faizan-glitch/url-sh/pkg/models"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var (
	resolve = cobra.Command(cobra.Command{
		Use:       "resolve [flags] code",
		Args:      cobra.ExactArgs(1),
		Short:     "Resolves a url",
		Long:      `Returns the full original url that corresponds to the given short code.`,
		Aliases:   []string{"r"},
		ValidArgs: []string{"-p", "--password"},
		Example:   `url-sh resolve ??????`,
		Run: func(cmd *cobra.Command, args []string) {
			password := cmd.Flag("password").Value.String()

			var model models.UrlSchema

			if err := model.FindByShortCode(args[0]); err != nil {
				log.Println("Could not find url", err.Error())
				return
			}

			if err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(password)); err != nil {
				log.Println("Invalid Password")
				return
			}

			log.Println("Original url:", model.LongUrl)
		},
	})

	resolveOnce sync.Once
)

func GetResolveCommand() *cobra.Command {
	resolveOnce.Do(func() {
		resolve.PersistentFlags().StringP("password", "p", "", "Password to protect the url")
	})
	return &resolve
}

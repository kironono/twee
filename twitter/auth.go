package twitter

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dghubble/oauth1"
	twoauth "github.com/dghubble/oauth1/twitter"
	"github.com/kironono/twee/model"
	"github.com/manifoldco/promptui"
)

func RunAuth(config *model.Config) error {
	authConfig := oauth1.Config{
		ConsumerKey:    config.TwitterSettings.ApiKey,
		ConsumerSecret: config.TwitterSettings.ApiSecret,
		CallbackURL:    "oob",
		Endpoint:       twoauth.AuthorizeEndpoint,
	}

	requestToken, _, err := authConfig.RequestToken()
	if err != nil {
		return fmt.Errorf("failed to request token: %w", err)
	}

	authURL, err := authConfig.AuthorizationURL(requestToken)
	if err != nil {
		return fmt.Errorf("failed to authorization URL: %w", err)
	}

	fmt.Println("Auth URL:")
	fmt.Println(authURL.String())
	fmt.Println()

	verifier, err := inputPIN()
	if err != nil {
		return fmt.Errorf("failed to read PIN: %w", err)
	}

	accessToken, accessSecret, err := authConfig.AccessToken(requestToken, "", verifier)
	if err != nil {
		return fmt.Errorf("failed to obtain token: %w", err)
	}

	newToken := oauth1.NewToken(accessToken, accessSecret)

	fmt.Println()
	fmt.Println("Please set the following access token in the configuration file.")
	fmt.Println("twitter_settings:")
	fmt.Printf("  access_token: %s\n", newToken.Token)
	fmt.Printf("  access_secret: %s\n", newToken.TokenSecret)

	return nil
}

func inputPIN() (string, error) {
	prompt := promptui.Prompt{
		Label: "PIN",
		Validate: func(s string) error {
			if _, err := strconv.Atoi(s); err != nil {
				return errors.New("please enter a number")
			}
			return nil
		},
	}

	return prompt.Run()
}

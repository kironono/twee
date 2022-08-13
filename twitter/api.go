package twitter

import (
	"net/http"

	"github.com/dghubble/oauth1"
	gotwitter "github.com/g8rswimmer/go-twitter/v2"
	"github.com/kironono/twee/model"
)

type Authorizer struct{}

func (a *Authorizer) Add(req *http.Request) {}

type API struct {
	client *gotwitter.Client
}

func NewAPI(config *model.Config) (*API, error) {
	client, err := newClient(config)
	if err != nil {
		return nil, err
	}

	return &API{
		client,
	}, nil
}

func newClient(config *model.Config) (*gotwitter.Client, error) {

	authConfig := oauth1.NewConfig(config.TwitterSettings.ApiKey, config.TwitterSettings.ApiSecret)
	accessToken := oauth1.NewToken(config.TwitterSettings.AccessToken, config.TwitterSettings.AccessSecret)

	httpClient := authConfig.Client(oauth1.NoContext, accessToken)

	return &gotwitter.Client{
		Authorizer: &Authorizer{},
		Client:     httpClient,
		Host:       "https://api.twitter.com",
	}, nil
}

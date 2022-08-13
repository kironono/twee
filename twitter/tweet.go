package twitter

import (
	"context"

	gotwitter "github.com/g8rswimmer/go-twitter/v2"
	"github.com/kironono/twee/model"
)

func RunTweet(config *model.Config, text string) error {
	api, err := NewAPI(config)
	if err != nil {
		return err
	}

	req := gotwitter.CreateTweetRequest{
		Text: text,
	}

	_, err = api.client.CreateTweet(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}

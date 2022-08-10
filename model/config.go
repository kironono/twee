package model

type TwitterSettings struct {
	ConsumerKey    string `yaml:"consumer_key"`
	ConsumerSecret string `yaml:"consumer_secret"`
}

type Config struct {
	TwitterSettings TwitterSettings `yaml:"twitter_settings"`
}

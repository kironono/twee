package model

type TwitterSettings struct {
	ApiKey       string `yaml:"api_key"`
	ApiSecret    string `yaml:"api_secret"`
	AccessToken  string `yaml:"access_token"`
	AccessSecret string `yaml:"access_secret"`
}

type Config struct {
	TwitterSettings TwitterSettings `yaml:"twitter_settings"`
}

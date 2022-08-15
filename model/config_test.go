package model

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  Config
	}{
		{
			name: "Normal case",
			input: []byte(heredoc.Doc(`
				twitter_settings:
				  api_key: a
				  api_secret: b
				  access_token: c
				  access_secret: d
				`)),
			want: Config{
				TwitterSettings: TwitterSettings{
					ApiKey:       "a",
					ApiSecret:    "b",
					AccessToken:  "c",
					AccessSecret: "d",
				},
			},
		},
		{
			name: "Access token missing",
			input: []byte(heredoc.Doc(`
				twitter_settings:
				  api_key: a
				  api_secret: b
			`)),
			want: Config{
				TwitterSettings{
					ApiKey:       "a",
					ApiSecret:    "b",
					AccessToken:  "",
					AccessSecret: "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parseConfig(tt.input)
			require.NoError(t, err)
			require.NotNil(t, actual)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func parseConfig(body []byte) (Config, error) {
	var c Config
	err := yaml.Unmarshal(body, &c)
	return c, err
}

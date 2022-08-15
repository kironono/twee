package config

import (
	"testing"

	"github.com/kironono/twee/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFileConfig(t *testing.T) {
	tests := []struct {
		name string
		path string
		want *FileConfig
	}{
		{
			name: "Path required",
			path: "twee.yml",
			want: &FileConfig{
				path: "twee.yml",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewFileConfig(tt.path)
			require.NotNil(t, actual)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name string
		path string
		want *model.Config
	}{
		{
			name: "Load testdata",
			path: "testdata/twee.yml",
			want: &model.Config{
				TwitterSettings: model.TwitterSettings{
					ApiKey:       "api_key",
					ApiSecret:    "api_secret",
					AccessToken:  "access_token",
					AccessSecret: "access_secret",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := NewFileConfig(tt.path)
			actual, err := fc.Load()
			require.NoError(t, err)
			require.NotNil(t, actual)
			assert.Equal(t, tt.want, actual)
		})
	}
}

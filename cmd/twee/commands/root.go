package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kironono/twee/config"
	"github.com/kironono/twee/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "twee",
	Short: "Simple tweet application",
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Configuration file to use.")
}

func loadConfig(command *cobra.Command) (*model.Config, error) {
	path, err := getConfigPath(command)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration\n%w\n", err)
	}

	fc := config.NewFileConfig(path)

	conf, err := fc.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration\n%w\n", err)
	}

	return conf, nil
}

func getConfigPath(command *cobra.Command) (string, error) {
	configPath, _ := command.Flags().GetString("config")
	if configPath != "" {
		return configPath, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	configPath = filepath.Join(homeDir, ".config", "twee.yml")
	if _, err = os.Stat(configPath); err == nil {
		return configPath, nil
	}

	return "twee.yml", nil
}

func Run(args []string) error {
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}

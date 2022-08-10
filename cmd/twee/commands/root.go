package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "twee",
	Short: "tweet",
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "Configuration file to use.")
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

package commands

import (
	"fmt"

	"github.com/kironono/twee/twitter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate your account with Twitter",
	RunE:  authCommandF,
}

func authCommandF(command *cobra.Command, args []string) error {
	config, err := loadConfig(command)
	if err != nil {
		return err
	}

	if err := twitter.RunAuth(config); err != nil {
		return fmt.Errorf("failed to auth: %w", err)
	}

	return nil
}

package commands

import (
	"fmt"

	"github.com/kironono/twee/config"
	"github.com/kironono/twee/twitter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tweetCmd)
	rootCmd.RunE = tweetCommandF
}

var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "tweet text",
	RunE:  tweetCommandF,
}

func tweetCommandF(command *cobra.Command, args []string) error {
	path, err := getConfigPath(command)
	if err != nil {
		return fmt.Errorf("failed to load configuration\n%w\n", err)
	}

	fc := config.NewFileConfig(path)

	conf, err := fc.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration\n%w\n", err)
	}

	return twitter.SendTweet(conf)
}

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
	fc := config.NewFileConfig(getConfigPath(command))
	conf, err := fc.Load()

	if err != nil {
		return fmt.Errorf("failed to load configuration\n%v\n", err)
	}

	return twitter.SendTweet(conf)
}

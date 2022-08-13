package commands

import (
	"github.com/kironono/twee/twitter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tweetCmd)
}

var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Tweet with Twitter",
	Args:  cobra.MinimumNArgs(1),
	RunE:  tweetCommandF,
}

func tweetCommandF(command *cobra.Command, args []string) error {
	config, err := loadConfig(command)
	if err != nil {
		return err
	}

	return twitter.RunTweet(config, args[0])
}

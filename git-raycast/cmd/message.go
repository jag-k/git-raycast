package cmd

import (
	"fmt"
	"git-raycast/git-raycast/git"
	"log"
	"net/url"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Create commit message based on changes",
	Long: `Generate commit message based on not-committed changes.

Calling this Deep-link:
> raycast://ai-commands/git-commit-message?arguments={diff}

More info here: https://github.com/jag-k/git-raycast/wiki/Commands#message`,
	RunE: runMessage,
}

func init() {
	rootCmd.AddCommand(messageCmd)
}

func runMessage(cmd *cobra.Command, args []string) error {
	gitDiff, err := git.GetDiff()
	if err != nil {
		return err
	}

	if gitDiff == "" {
		return fmt.Errorf("no changes found")
	}

	result, err := buildRaycastURL("git-commit-message", gitDiff)
	if err != nil {
		return err
	}

	if verbose {
		log.Println(result)
	}

	return open.Run(result)
}

func buildRaycastURL(commandName, argument string) (string, error) {
	baseUrl, err := url.Parse("raycast://ai-commands/")
	if err != nil {
		return "", err
	}

	baseUrl.Path += commandName
	params := url.Values{}
	params.Add("arguments", argument)
	baseUrl.RawQuery = params.Encode()

	return baseUrl.String(), nil
}

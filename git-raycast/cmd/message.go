package cmd

import (
	"fmt"
	"git-raycast/git-raycast/config"
	"git-raycast/git-raycast/git"
	"git-raycast/git-raycast/utils"
	"log"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var messageCmd = &cobra.Command{
	Use:     "message [command-name]",
	Aliases: []string{"msg"},
	Short:   "Create commit message based on changes",
	Long: `Generate commit message based on not-committed changes.

The Raycast AI command name can be customized:
- By passing [command-name] argument
- By setting GIT_RAYCAST_MESSAGE_NAME environment variable
- By setting git config git-raycast.message-name
- Default: git-commit-message

Calling this Deep-link:
> raycast://ai-commands/{command-name}?arguments={diff}
> raycast-x://extensions/raycast/ai/{command-name}?arguments={"diff":"{diff}"} (with --raycast-version beta)

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

	commandName, err := config.CommandName(config.MessageCommandName, args, 0)
	if err != nil {
		return err
	}

	version, err := config.RaycastVersion(raycastVersion)
	if err != nil {
		return err
	}

	result, err := utils.BuildRaycastURL(commandName, gitDiff, version)
	if err != nil {
		return err
	}

	if verbose {
		log.Println(result)
	}

	return open.Run(result)
}

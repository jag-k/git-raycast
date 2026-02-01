package cmd

import (
	"fmt"
	"git-raycast/git-raycast/git"
	"git-raycast/git-raycast/utils"
	"log"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary [command-name]",
	Short: "Create daily summary based on changes",
	Long: `Generate a summary of changes made in the repository since yesterday.

The Raycast AI command name can be customized:
- By passing [command-name] argument
- By setting GIT_RAYCAST_SUMMARY_NAME environment variable
- Default: daily-summary

Calling this Deep-link:
> raycast://ai-commands/{command-name}?arguments={diff}

More info here: https://github.com/jag-k/git-raycast/wiki/Commands#summary`,
	RunE: runSummary,
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}

func runSummary(cmd *cobra.Command, args []string) error {
	// Получаем хеш последнего коммита до вчерашнего дня
	gitHash, err := git.ExecuteCommand("log", "-1", "--until=yesterday", "--pretty=format:%H")
	if err != nil {
		return err
	}

	if gitHash == "" {
		return fmt.Errorf("no commits found")
	}

	// Получаем изменения между тем коммитом и текущим состоянием
	gitDiff, err := git.ExecuteCommand("diff", gitHash, "HEAD")
	if err != nil {
		return err
	}

	commandName := utils.GetCommandName("GIT_RAYCAST_SUMMARY_NAME", "daily-summary", args, 0)
	result, err := utils.BuildRaycastURL(commandName, gitDiff)
	if err != nil {
		return err
	}

	if verbose {
		log.Println(result)
	}

	return open.Run(result)
}

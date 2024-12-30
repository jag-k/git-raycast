package cmd

import (
	"fmt"
	"git-raycast/git-raycast/git"
	"log"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Create daily summary based on changes",
	Long: `Generate a summary of changes made in the repository since yesterday.

Calling this Deep-link:
> raycast://ai-commands/daily-summary?arguments={diff}

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

	result, err := buildRaycastURL("daily-summary", gitDiff)
	if err != nil {
		return err
	}

	if verbose {
		log.Println(result)
	}

	return open.Run(result)
}

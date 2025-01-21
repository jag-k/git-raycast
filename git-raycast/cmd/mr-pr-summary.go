package cmd

import (
	"fmt"
	"git-raycast/git-raycast/git"
	"log"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var MRPRSummaryCmd = &cobra.Command{
	Use:     "mr [target branch]",
	Aliases: []string{"pr"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("accepts at most 1 arg, received %d", len(args))
		}
		if len(args) == 1 {
			_, err := git.ExecuteCommand("rev-parse", "--quiet", "--verify", args[0])
			if err != nil {
				return fmt.Errorf("invalid branch name %q: %w", args[0], err)
			}
		}
		return nil
	},
	Short: "Generate MR/PR Summary message",
	Long: `Generate a Merge Request (or Pull Request) summary message based by commit messages between current branch and target branch.

By default, the target branch is ` + "`origin/main`" + `

Calling this Deep-link:
> raycast://ai-commands/mr-pr-summary?arguments={diff}

More info here: https://github.com/jag-k/git-raycast/wiki/Commands#mr-pr-summary`,
	RunE: runMRPRSummary,
}

func init() {
	rootCmd.AddCommand(MRPRSummaryCmd)
}

func runMRPRSummary(cmd *cobra.Command, args []string) error {
	var targetBranch = "origin/main"
	if len(args) == 1 {
		targetBranch = args[0]
	}
	gitDiff, err := git.ExecuteCommand("log", fmt.Sprintf("%s..HEAD", targetBranch), "--no-merges", "--pretty=format:%B")
	if err != nil {
		return err
	}

	if gitDiff == "" {
		return fmt.Errorf("no commits found")
	}

	result, err := buildRaycastURL("mr-pr-summary", gitDiff)
	if err != nil {
		return err
	}

	if verbose {
		log.Println(result)
	}

	return open.Run(result)
}

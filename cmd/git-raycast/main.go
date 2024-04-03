package main

import (
	"bytes"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli/v2"
)

var version = "dev"

func executeGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	result, err := cmd.Output()
	exitCode := cmd.ProcessState.ExitCode()
	if err != nil {
		return "", cli.Exit("Have a git error:\nCommand: `git "+strings.Join(args, " ")+"`\nError: "+strings.TrimSpace(stderr.String()), exitCode)
	}
	return string(result), nil
}
func git_message() error {
	gitDiff, err := executeGitCommand("diff")
	if gitDiff == "" {
		gitDiff, err = executeGitCommand("diff", "--cached")
	}
	if gitDiff == "" {
		gitDiff, err = executeGitCommand("diff", "--staged")
	}

	if err != nil {
		return err
	}
	if gitDiff == "" {
		return cli.Exit("No changes found", 1)
	}
	url := "raycast://ai-commands/git-commit-message?arguments=" + url.QueryEscape(gitDiff)
	log.Println(url)
	return open.Run(url)
}

func git_summary() error {
	gitHash, err := executeGitCommand("log", "-1", "--until=yesterday", "--pretty=format:%H")
	if err != nil {
		return err
	}
	if gitHash == "" {
		return cli.Exit("No commits found", 1)
	}
	gitDiff, err := executeGitCommand("diff", gitHash, "HEAD")
	if err != nil {
		return err
	}

	url := "raycast://ai-commands/daily-summary?arguments=" + url.QueryEscape(gitDiff)
	log.Println(url)
	return open.Run(url)
}

func main() {
	app := &cli.App{
		Name:                 "git raycast",
		Usage:                "Automate git using Raycast AI",
		Version:              version,
		HideHelpCommand:      true,
		HideHelp:             true,
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:      "message",
				Usage:     "Create commit message based on changes",
				Action:    func(c *cli.Context) error { return git_message() },
				ArgsUsage: " ",
			},
			{
				Name:      "summary",
				Usage:     "Create daily summary based on changes",
				Action:    func(c *cli.Context) error { return git_summary() },
				ArgsUsage: " ",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

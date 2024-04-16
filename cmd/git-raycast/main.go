package main

import (
	"bytes"
	"fmt"
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
func gitMessage() (string, error) {
	gitDiff, err := executeGitCommand("diff")
	if gitDiff == "" {
		gitDiff, err = executeGitCommand("diff", "--cached")
	}
	if gitDiff == "" {
		gitDiff, err = executeGitCommand("diff", "--staged")
	}

	if err != nil {
		return "", err
	}
	if gitDiff == "" {
		return "", cli.Exit("No changes found", 1)
	}
	return "raycast://ai-commands/git-commit-message?arguments=" + url.PathEscape(gitDiff), nil
}

func gitSummary() (string, error) {
	gitHash, err := executeGitCommand("log", "-1", "--until=yesterday", "--pretty=format:%H")
	if err != nil {
		return "", err
	}
	if gitHash == "" {
		return "", cli.Exit("No commits found", 1)
	}
	gitDiff, err := executeGitCommand("diff", gitHash, "HEAD")
	if err != nil {
		return "", err
	}

	return "raycast://ai-commands/daily-summary?arguments=" + url.PathEscape(gitDiff), nil
}

func createCliAction(f func() (string, error)) func(c *cli.Context) error {
	wrap := func(c *cli.Context) error {
		result, err := f()
		if err != nil {
			return err
		}
		if c.Bool("verbose") {
			log.Println(result)
		}
		return open.Run(result)
	}

	return wrap
}

func main() {
	app := &cli.App{
		Name:                 "git raycast",
		Usage:                "Automate git using Raycast AI",
		Version:              version,
		HideHelpCommand:      true,
		HideHelp:             true,
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Usage:       "Send raycast url to output",
				Required:    false,
				Value:       false,
				Destination: nil,
				Aliases:     []string{"V"},
			},
			&cli.StringFlag{
				Name:   "create-man-page",
				Value:  "-",
				Hidden: true,
				Action: func(c *cli.Context, value string) error {
					var man, err = c.App.ToManWithSection(1)
					if err != nil {
						return err
					}
					if value != "-" {
						// Write to file `value` man page
						return os.WriteFile(value, []byte(man), 0644)
					} else {
						fmt.Print(man)
					}
					cli.OsExiter(0)
					return nil
				},
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "message",
				Action: createCliAction(gitMessage),
				Usage:  "Create commit message based on changes",
				Description: strings.Replace(`Generate commit message based on not-committed changes.

				Calling this Deep-link:
				> raycast://ai-commands/git-commit-message?arguments={diff}`, "\t", "", -1),
				ArgsUsage: " ",
			},
			{
				Name:   "summary",
				Action: createCliAction(gitSummary),
				Usage:  "Create daily summary based on changes",
				Description: strings.Replace(`Generate a summary of changes made in the repository since yesterday.

				Calling this Deep-link:
				> raycast://ai-commands/daily-summary?arguments={diff}`, "\t", "", -1),
				ArgsUsage: " ",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bytes"
	"context"
	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli/v3"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

var version = "dev"

func raycastAICommand(commandName string, argument string) (string, error) {
	baseUrl, err := url.Parse("raycast://ai-commands/")
	if err != nil {
		return "", err
	}
	baseUrl.Path += commandName

	params := url.Values{}
	params.Add("arguments", argument)

	// Add Query Parameters to the URL
	baseUrl.RawQuery = params.Encode() // Escape Query Parameters

	return baseUrl.String(), err
}

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
	result, err := raycastAICommand("git-commit-message", gitDiff)
	if err != nil {
		return "", err
	}
	return result, nil
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
	result, err := raycastAICommand("daily-summary", gitDiff)
	if err != nil {
		return "", err
	}
	return result, nil
}

func createCliAction(f func() (string, error)) func(c context.Context, command *cli.Command) error {
	wrap := func(c context.Context, command *cli.Command) error {
		result, err := f()
		if err != nil {
			return err
		}
		if command.Bool("verbose") {
			log.Println(result)
		}
		return open.Run(result)
	}

	return wrap
}

func main() {
	app := &cli.Command{
		Name:                  "git raycast",
		Usage:                 "Automate git using Raycast AI",
		Version:               version,
		Description:           "Automate git using Raycast AI",
		EnableShellCompletion: true,
		Suggest:               true,
		HideHelp:              true,
		Commands: []*cli.Command{
			{
				Name:   "message",
				Action: createCliAction(gitMessage),
				Usage:  "Create commit message based on changes",
				Description: strings.Replace(`Generate commit message based on not-committed changes.

				Calling this Deep-link:
				> raycast://ai-commands/git-commit-message?arguments={diff}

				More info here: https://github.com/jag-k/git-raycast/wiki/Commands#message`, "\t", "", -1),
				ArgsUsage: " ",
			},
			{
				Name:   "summary",
				Action: createCliAction(gitSummary),
				Usage:  "Create daily summary based on changes",
				Description: strings.Replace(`Generate a summary of changes made in the repository since yesterday.

				Calling this Deep-link:
				> raycast://ai-commands/daily-summary?arguments={diff}

				More info here: https://github.com/jag-k/git-raycast/wiki/Commands#summary`, "\t", "", -1),
				ArgsUsage: " ",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Usage:       "Send raycast url to output",
				Required:    false,
				Value:       false,
				Destination: nil,
				Aliases:     []string{"V"},
			},
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

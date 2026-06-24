package config

import (
	"os"

	"git-raycast/git-raycast/git"
)

const (
	RaycastVersionStable = "stable"
	RaycastVersionBeta   = "beta"

	GitConfigRaycastVersion = "git-raycast.raycast-version"
)

type StringSetting struct {
	EnvVar  string
	GitKey  string
	Default string
}

var (
	MessageCommandName = StringSetting{
		EnvVar:  "GIT_RAYCAST_MESSAGE_NAME",
		GitKey:  "git-raycast.message-name",
		Default: "git-commit-message",
	}
	SummaryCommandName = StringSetting{
		EnvVar:  "GIT_RAYCAST_SUMMARY_NAME",
		GitKey:  "git-raycast.summary-name",
		Default: "daily-summary",
	}
	MRPRSummaryCommandName = StringSetting{
		EnvVar:  "GIT_RAYCAST_MR_PR_SUMMARY_NAME",
		GitKey:  "git-raycast.mr-pr-summary-name",
		Default: "mr-pr-summary",
	}
)

// RaycastVersion returns the configured Raycast version.
// Priority: 1. flag value, 2. environment variable, 3. git config, 4. stable.
func RaycastVersion(flagValue string) (string, error) {
	gitConfigValue, err := git.GetConfig(GitConfigRaycastVersion)
	if err != nil {
		return "", err
	}

	return firstNonEmpty(flagValue, os.Getenv("GIT_RAYCAST_VERSION"), gitConfigValue, RaycastVersionStable), nil
}

// CommandName returns the configured Raycast command name.
// Priority: 1. positional argument, 2. environment variable, 3. git config, 4. default.
func CommandName(setting StringSetting, args []string, argIndex int) (string, error) {
	if len(args) > argIndex {
		return args[argIndex], nil
	}

	gitConfigValue, err := git.GetConfig(setting.GitKey)
	if err != nil {
		return "", err
	}

	return firstNonEmpty(os.Getenv(setting.EnvVar), gitConfigValue, setting.Default), nil
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}

	return ""
}

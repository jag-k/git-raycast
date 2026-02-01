package utils

import (
	"net/url"
	"os"
)

// GetCommandName returns the command name for Raycast AI command.
// Priority: 1. args[argIndex], 2. environment variable, 3. defaultName
func GetCommandName(envVar, defaultName string, args []string, argIndex int) string {
	// Check if argument is provided
	if len(args) > argIndex {
		return args[argIndex]
	}
	// Check environment variable
	if envValue := os.Getenv(envVar); envValue != "" {
		return envValue
	}
	// Return default value
	return defaultName
}

// BuildRaycastURL constructs a Raycast AI command URL with the given command name and arguments
func BuildRaycastURL(commandName, argument string) (string, error) {
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

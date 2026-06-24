package utils

import (
	"fmt"
	"net/url"

	"git-raycast/git-raycast/config"
)

// BuildRaycastURL constructs a Raycast AI command URL with the given command name and arguments
func BuildRaycastURL(commandName, argument, raycastVersion string) (string, error) {
	baseURL, err := raycastBaseURL(raycastVersion)
	if err != nil {
		return "", err
	}

	baseURL.Path += commandName
	params := url.Values{}
	params.Add("arguments", argument)
	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}

func raycastBaseURL(raycastVersion string) (*url.URL, error) {
	switch raycastVersion {
	case config.RaycastVersionStable:
		return url.Parse("raycast://ai-commands/")
	case config.RaycastVersionBeta:
		return url.Parse("raycast-x://extensions/raycast/ai/")
	default:
		return nil, fmt.Errorf("unsupported Raycast version %q, expected %q or %q", raycastVersion, config.RaycastVersionStable, config.RaycastVersionBeta)
	}
}

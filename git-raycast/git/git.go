package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ExecuteCommand выполняет git команду и возвращает результат
func ExecuteCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	result, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git error:\nCommand: `git %s`\nError: %s",
			strings.Join(args, " "),
			strings.TrimSpace(stderr.String()))
	}

	return string(result), nil
}

// GetConfig returns a git config value using Git's standard lookup order.
func GetConfig(key string) (string, error) {
	cmd := exec.Command("git", "config", "--get", key)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	result, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			return "", nil
		}

		return "", fmt.Errorf("git error:\nCommand: `git config --get %s`\nError: %s",
			key,
			strings.TrimSpace(stderr.String()))
	}

	return strings.TrimSpace(string(result)), nil
}

// GetDiff получает изменения в репозитории
func GetDiff() (string, error) {
	gitDiff, err := ExecuteCommand("diff")
	if err != nil {
		return "", err
	}

	if gitDiff == "" {
		gitDiff, err = ExecuteCommand("diff", "--cached")
		if err != nil {
			return "", err
		}
	}

	if gitDiff == "" {
		gitDiff, err = ExecuteCommand("diff", "--staged")
		if err != nil {
			return "", err
		}
	}

	return gitDiff, nil
}

// GetDiffSince получает изменения с определенного коммита до HEAD
func GetDiffSince(hash string) (string, error) {
	return ExecuteCommand("diff", hash, "HEAD")
}

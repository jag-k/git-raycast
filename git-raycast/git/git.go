package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type DiffMode string

const (
	DiffModeAuto    DiffMode = "auto"
	DiffModeStage   DiffMode = "stage"
	DiffModeUnstage DiffMode = "unstage"
	DiffModeAll     DiffMode = "all"
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

// GetDiff returns repository changes selected by mode. Auto prioritizes staged
// changes because they are what the next commit will contain, and falls back to
// unstaged changes when the index is clean.
func GetDiff(mode DiffMode) (string, error) {
	switch mode {
	case DiffModeAuto:
		staged, err := ExecuteCommand("diff", "--cached")
		if err != nil {
			return "", err
		}
		if staged != "" {
			return staged, nil
		}
		return ExecuteCommand("diff")
	case DiffModeStage:
		return ExecuteCommand("diff", "--cached")
	case DiffModeUnstage:
		return ExecuteCommand("diff")
	case DiffModeAll:
		staged, err := ExecuteCommand("diff", "--cached")
		if err != nil {
			return "", err
		}
		unstaged, err := ExecuteCommand("diff")
		if err != nil {
			return "", err
		}
		return combineDiffs(staged, unstaged), nil
	default:
		return "", fmt.Errorf("invalid changes mode %q: must be one of auto, stage, unstage, all", mode)
	}
}

func combineDiffs(staged, unstaged string) string {
	var sections []string
	if staged != "" {
		sections = append(sections, "STAGED CHANGES:\n"+staged)
	}
	if unstaged != "" {
		sections = append(sections, "UNSTAGED CHANGES:\n"+unstaged)
	}
	return strings.Join(sections, "\n")
}

// GetDiffSince получает изменения с определенного коммита до HEAD
func GetDiffSince(hash string) (string, error) {
	return ExecuteCommand("diff", hash, "HEAD")
}

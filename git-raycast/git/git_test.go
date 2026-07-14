package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetConfig(t *testing.T) {
	isolateGitConfig(t)
	t.Setenv("GIT_CONFIG_COUNT", "1")
	t.Setenv("GIT_CONFIG_KEY_0", "git-raycast.message-name")
	t.Setenv("GIT_CONFIG_VALUE_0", "custom-message")

	result, err := GetConfig("git-raycast.message-name")
	if err != nil {
		t.Fatalf("GetConfig returned error: %v", err)
	}

	if result != "custom-message" {
		t.Fatalf("GetConfig() = %q, want %q", result, "custom-message")
	}
}

func TestGetConfigMissingValue(t *testing.T) {
	isolateGitConfig(t)

	result, err := GetConfig("git-raycast.missing")
	if err != nil {
		t.Fatalf("GetConfig returned error: %v", err)
	}

	if result != "" {
		t.Fatalf("GetConfig() = %q, want empty string", result)
	}
}

func isolateGitConfig(t *testing.T) {
	t.Helper()
	t.Setenv("GIT_CONFIG_NOSYSTEM", "1")
	t.Setenv("GIT_CONFIG_GLOBAL", "/dev/null")
	t.Setenv("HOME", t.TempDir())
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
}

func TestGetDiff(t *testing.T) {
	t.Run("auto prioritizes staged changes", func(t *testing.T) {
		setupDiffRepository(t)

		diff, err := GetDiff(DiffModeAuto)
		if err != nil {
			t.Fatalf("GetDiff() returned error: %v", err)
		}
		if !strings.Contains(diff, "+staged change") {
			t.Fatalf("GetDiff() did not include staged change:\n%s", diff)
		}
		if strings.Contains(diff, "+unstaged change") {
			t.Fatalf("GetDiff() unexpectedly included unstaged change:\n%s", diff)
		}
	})

	t.Run("auto falls back to unstaged changes", func(t *testing.T) {
		setupDiffRepository(t)
		runGit(t, "restore", "--staged", "staged.txt")

		diff, err := GetDiff(DiffModeAuto)
		if err != nil {
			t.Fatalf("GetDiff() returned error: %v", err)
		}
		if !strings.Contains(diff, "+staged change") || !strings.Contains(diff, "+unstaged change") {
			t.Fatalf("GetDiff() did not include all unstaged changes:\n%s", diff)
		}
	})

	t.Run("stage includes only staged changes", func(t *testing.T) {
		setupDiffRepository(t)

		diff, err := GetDiff(DiffModeStage)
		if err != nil {
			t.Fatalf("GetDiff() returned error: %v", err)
		}
		if !strings.Contains(diff, "+staged change") || strings.Contains(diff, "+unstaged change") {
			t.Fatalf("GetDiff() returned unexpected changes:\n%s", diff)
		}
	})

	t.Run("unstage includes only unstaged changes", func(t *testing.T) {
		setupDiffRepository(t)

		diff, err := GetDiff(DiffModeUnstage)
		if err != nil {
			t.Fatalf("GetDiff() returned error: %v", err)
		}
		if !strings.Contains(diff, "+unstaged change") || strings.Contains(diff, "+staged change") {
			t.Fatalf("GetDiff() returned unexpected changes:\n%s", diff)
		}
	})

	t.Run("all labels and includes both kinds of changes", func(t *testing.T) {
		setupDiffRepository(t)

		diff, err := GetDiff(DiffModeAll)
		if err != nil {
			t.Fatalf("GetDiff() returned error: %v", err)
		}
		for _, want := range []string{"STAGED CHANGES:", "+staged change", "UNSTAGED CHANGES:", "+unstaged change"} {
			if !strings.Contains(diff, want) {
				t.Fatalf("GetDiff() did not include %q:\n%s", want, diff)
			}
		}
	})

	t.Run("rejects an invalid mode", func(t *testing.T) {
		setupDiffRepository(t)

		_, err := GetDiff("invalid")
		if err == nil || !strings.Contains(err.Error(), "must be one of auto, stage, unstage, all") {
			t.Fatalf("GetDiff() error = %v, want invalid mode error", err)
		}
	})
}

func setupDiffRepository(t *testing.T) {
	t.Helper()
	repository := t.TempDir()
	workingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}
	if err := os.Chdir(repository); err != nil {
		t.Fatalf("change working directory: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(workingDirectory); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})

	runGit(t, "init", "--quiet")
	runGit(t, "config", "user.email", "test@example.com")
	runGit(t, "config", "user.name", "Test User")
	runGit(t, "config", "commit.gpgsign", "false")
	writeFile(t, repository, "staged.txt", "original\n")
	writeFile(t, repository, "unstaged.txt", "original\n")
	runGit(t, "add", ".")
	runGit(t, "commit", "--quiet", "-m", "initial")

	writeFile(t, repository, "staged.txt", "original\nstaged change\n")
	runGit(t, "add", "staged.txt")
	writeFile(t, repository, "unstaged.txt", "original\nunstaged change\n")
}

func writeFile(t *testing.T, directory, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(directory, name), []byte(content), 0o600); err != nil {
		t.Fatalf("write %s: %v", name, err)
	}
}

func runGit(t *testing.T, args ...string) {
	t.Helper()
	command := exec.Command("git", args...)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, output)
	}
}

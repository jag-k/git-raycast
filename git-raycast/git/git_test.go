package git

import "testing"

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

package config

import "testing"

func TestRaycastVersionDefault(t *testing.T) {
	isolateGitConfig(t)

	result, err := RaycastVersion("")
	if err != nil {
		t.Fatalf("RaycastVersion returned error: %v", err)
	}

	if result != RaycastVersionStable {
		t.Fatalf("RaycastVersion() = %q, want %q", result, RaycastVersionStable)
	}
}

func TestRaycastVersionFromGitConfig(t *testing.T) {
	isolateGitConfig(t)
	setGitConfig(t, GitConfigRaycastVersion, RaycastVersionBeta)

	result, err := RaycastVersion("")
	if err != nil {
		t.Fatalf("RaycastVersion returned error: %v", err)
	}

	if result != RaycastVersionBeta {
		t.Fatalf("RaycastVersion() = %q, want %q", result, RaycastVersionBeta)
	}
}

func TestRaycastVersionEnvTakesPriority(t *testing.T) {
	isolateGitConfig(t)
	setGitConfig(t, GitConfigRaycastVersion, RaycastVersionStable)
	t.Setenv("GIT_RAYCAST_VERSION", RaycastVersionBeta)

	result, err := RaycastVersion("")
	if err != nil {
		t.Fatalf("RaycastVersion returned error: %v", err)
	}

	if result != RaycastVersionBeta {
		t.Fatalf("RaycastVersion() = %q, want %q", result, RaycastVersionBeta)
	}
}

func TestRaycastVersionFlagTakesPriority(t *testing.T) {
	isolateGitConfig(t)
	setGitConfig(t, GitConfigRaycastVersion, RaycastVersionBeta)
	t.Setenv("GIT_RAYCAST_VERSION", RaycastVersionBeta)

	result, err := RaycastVersion(RaycastVersionStable)
	if err != nil {
		t.Fatalf("RaycastVersion returned error: %v", err)
	}

	if result != RaycastVersionStable {
		t.Fatalf("RaycastVersion() = %q, want %q", result, RaycastVersionStable)
	}
}

func TestCommandNameDefault(t *testing.T) {
	isolateGitConfig(t)

	result, err := CommandName(MessageCommandName, nil, 0)
	if err != nil {
		t.Fatalf("CommandName returned error: %v", err)
	}

	if result != MessageCommandName.Default {
		t.Fatalf("CommandName() = %q, want %q", result, MessageCommandName.Default)
	}
}

func TestCommandNameFromGitConfig(t *testing.T) {
	isolateGitConfig(t)
	setGitConfig(t, MessageCommandName.GitKey, "custom-message")

	result, err := CommandName(MessageCommandName, nil, 0)
	if err != nil {
		t.Fatalf("CommandName returned error: %v", err)
	}

	if result != "custom-message" {
		t.Fatalf("CommandName() = %q, want %q", result, "custom-message")
	}
}

func TestCommandNameEnvTakesPriority(t *testing.T) {
	isolateGitConfig(t)
	setGitConfig(t, MessageCommandName.GitKey, "custom-message")
	t.Setenv(MessageCommandName.EnvVar, "env-message")

	result, err := CommandName(MessageCommandName, nil, 0)
	if err != nil {
		t.Fatalf("CommandName returned error: %v", err)
	}

	if result != "env-message" {
		t.Fatalf("CommandName() = %q, want %q", result, "env-message")
	}
}

func TestCommandNameArgTakesPriority(t *testing.T) {
	isolateGitConfig(t)
	setGitConfig(t, MessageCommandName.GitKey, "custom-message")
	t.Setenv(MessageCommandName.EnvVar, "env-message")

	result, err := CommandName(MessageCommandName, []string{"arg-message"}, 0)
	if err != nil {
		t.Fatalf("CommandName returned error: %v", err)
	}

	if result != "arg-message" {
		t.Fatalf("CommandName() = %q, want %q", result, "arg-message")
	}
}

func setGitConfig(t *testing.T, key, value string) {
	t.Helper()
	t.Setenv("GIT_CONFIG_COUNT", "1")
	t.Setenv("GIT_CONFIG_KEY_0", key)
	t.Setenv("GIT_CONFIG_VALUE_0", value)
}

func isolateGitConfig(t *testing.T) {
	t.Helper()
	t.Setenv("GIT_CONFIG_NOSYSTEM", "1")
	t.Setenv("GIT_CONFIG_GLOBAL", "/dev/null")
	t.Setenv("HOME", t.TempDir())
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
}

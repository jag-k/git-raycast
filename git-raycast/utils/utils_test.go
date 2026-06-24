package utils

import (
	"testing"

	"git-raycast/git-raycast/config"
)

func TestBuildRaycastURLStable(t *testing.T) {
	result, err := BuildRaycastURL("git-commit-message", "hello world", config.RaycastVersionStable)
	if err != nil {
		t.Fatalf("BuildRaycastURL returned error: %v", err)
	}

	expected := "raycast://ai-commands/git-commit-message?arguments=hello+world"
	if result != expected {
		t.Fatalf("BuildRaycastURL() = %q, want %q", result, expected)
	}
}

func TestBuildRaycastURLBeta(t *testing.T) {
	result, err := BuildRaycastURL("git-commit-message", "hello world", config.RaycastVersionBeta)
	if err != nil {
		t.Fatalf("BuildRaycastURL returned error: %v", err)
	}

	expected := "raycast-x://extensions/raycast/ai/git-commit-message?arguments=hello+world"
	if result != expected {
		t.Fatalf("BuildRaycastURL() = %q, want %q", result, expected)
	}
}

func TestBuildRaycastURLUnknownVersion(t *testing.T) {
	_, err := BuildRaycastURL("git-commit-message", "hello world", "nightly")
	if err == nil {
		t.Fatal("BuildRaycastURL returned nil error for unknown Raycast version")
	}
}

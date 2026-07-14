package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestCompletionIncludesGitAdapters(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		marker string
	}{
		{name: "bash", args: []string{"completion", "bash"}, marker: "_git_raycast()"},
		{name: "zsh", args: []string{"completion", "zsh"}, marker: "_git_raycast()"},
		{name: "fish", args: []string{"completion", "fish"}, marker: "complete -c git"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer
			rootCmd.SetArgs(test.args)
			rootCmd.SetOut(&output)
			rootCmd.SetErr(&output)

			if err := rootCmd.Execute(); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(output.String(), test.marker) {
				t.Fatalf("generated completion does not contain %q", test.marker)
			}
		})
	}
}

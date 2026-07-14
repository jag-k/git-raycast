package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestGitCompletion(t *testing.T) {
	tests := []struct {
		shell  string
		marker string
	}{
		{shell: "bash", marker: "_git_raycast()"},
		{shell: "zsh", marker: "#autoload"},
	}

	for _, test := range tests {
		t.Run(test.shell, func(t *testing.T) {
			var output bytes.Buffer
			gitCompletionCmd.SetOut(&output)
			defer gitCompletionCmd.SetOut(nil)

			if err := gitCompletionCmd.RunE(gitCompletionCmd, []string{test.shell}); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(output.String(), test.marker) {
				t.Fatalf("generated completion does not contain %q", test.marker)
			}
		})
	}
}

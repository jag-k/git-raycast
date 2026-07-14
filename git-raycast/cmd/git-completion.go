package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

const bashGitCompletion = `# Completion adapter for invoking git-raycast as "git raycast".
_git_raycast()
{
    if ! declare -F __git-raycast_get_completion_results >/dev/null 2>&1; then
        _completion_loader git-raycast
    fi

    local -a git_raycast_words=("${words[@]}")
    local git_raycast_cword=$cword
    local cur prev words cword split

    words=("git-raycast" "${git_raycast_words[@]:$((__git_cmd_idx + 1))}")
    cword=$((git_raycast_cword - __git_cmd_idx))
    cur="${words[cword]}"
    if ((cword > 0)); then
        prev="${words[cword - 1]}"
    fi

    local out directive
    __git-raycast_get_completion_results
    __git-raycast_process_completion_results
}
`

const zshGitCompletion = `#autoload

# Git's zsh completion dispatches external commands to _git_<command>.
_git_raycast()
{
    emulate -L zsh

    local -a git_raycast_words=("${words[@]}")
    local git_raycast_current=$CURRENT
    local command_index

    for ((command_index = 1; command_index <= ${#git_raycast_words}; command_index++)); do
        [[ "${git_raycast_words[command_index]}" == "raycast" ]] && break
    done
    ((command_index <= ${#git_raycast_words})) || return 1

    words=("git-raycast" "${git_raycast_words[@]:${command_index}}")
    CURRENT=$((git_raycast_current - command_index + 1))

    autoload -Uz _git-raycast
    _git-raycast
}

_git_raycast "$@"
`

var gitCompletionCmd = &cobra.Command{
	Use:       "git-completion [bash|zsh]",
	Short:     "Generate completion adapters for the git raycast command",
	Hidden:    true,
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash", "zsh"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var completion string
		switch args[0] {
		case "bash":
			completion = bashGitCompletion
		case "zsh":
			completion = zshGitCompletion
		default:
			return fmt.Errorf("unsupported shell %q", args[0])
		}

		_, err := io.WriteString(cmd.OutOrStdout(), completion)
		return err
	},
}

func init() {
	rootCmd.AddCommand(gitCompletionCmd)
}

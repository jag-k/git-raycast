package cmd

import (
	"io"

	"github.com/spf13/cobra"
)

const bashGitCompletionAdapter = `

# Complete git-raycast when it is invoked as the external Git command
# "git raycast". Git's completion dispatcher calls this function by name.
_git_raycast()
{
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

const zshGitCompletionAdapter = `

# Homebrew's Git completion dispatcher uses underscores for external
# commands. The generated Cobra function above keeps its hyphenated name.
_git_raycast()
{
    emulate -L zsh

    if [[ "${words[1]}" == "git-raycast" ]]; then
        _git-raycast
        return
    fi

    local -a git_raycast_words=("${words[@]}")
    local git_raycast_current=$CURRENT
    local command_index

    for ((command_index = 1; command_index <= ${#git_raycast_words}; command_index++)); do
        [[ "${git_raycast_words[command_index]}" == "raycast" ]] && break
    done
    ((command_index <= ${#git_raycast_words})) || return 1

    words=("git-raycast" "${git_raycast_words[@]:${command_index}}")
    CURRENT=$((git_raycast_current - command_index + 1))
    _git-raycast
}

# When this file is autoloaded through the _git_raycast compatibility symlink,
# zsh loads the definition above and needs an explicit first invocation.
if [[ "${funcstack[1]}" == "_git_raycast" ]]; then
    _git_raycast
fi
`

const fishGitCompletionAdapter = `

function __git_raycast_using_git_subcommand
    set -l args (commandline -opc)
    contains -- raycast $args
end

function __git_raycast_perform_git_completion
    set -l args (commandline -opc)
    set -l command_index (contains -i -- raycast $args)
    test -n "$command_index"; or return 1

    set -l completion_args $args[(math $command_index + 1)..-1]
    set -l last_arg (string escape -- (commandline -ct))
    set -l request_comp "GIT_RAYCAST_ACTIVE_HELP=0 git-raycast __complete $completion_args $last_arg"
    set -l results (eval $request_comp 2>/dev/null)

    for line in $results[-1..1]
        if test (string trim -- $line) = ""
            set results $results[1..-2]
        else
            break
        end
    end

    test (count $results) -ge 2; or return 1
    printf "%s\n" $results[1..-2]
end

complete -c git -n __git_raycast_using_git_subcommand -f -a '(__git_raycast_perform_git_completion)'
`

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate the autocompletion script for the specified shell",
	Args:  cobra.NoArgs,
}

func completionShellCommand(shell string, generate func(io.Writer, bool) error, adapter string) *cobra.Command {
	var noDescriptions bool

	command := &cobra.Command{
		Use:               shell,
		Short:             "Generate the autocompletion script for " + shell,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, _ []string) error {
			out := cmd.OutOrStdout()
			if err := generate(out, !noDescriptions); err != nil {
				return err
			}
			_, err := io.WriteString(out, adapter)
			return err
		},
	}
	command.Flags().BoolVar(&noDescriptions, "no-descriptions", false, "disable completion descriptions")
	return command
}

func init() {
	completionCmd.AddCommand(
		completionShellCommand("bash", rootCmd.GenBashCompletionV2, bashGitCompletionAdapter),
		completionShellCommand("zsh", func(out io.Writer, descriptions bool) error {
			if descriptions {
				return rootCmd.GenZshCompletion(out)
			}
			return rootCmd.GenZshCompletionNoDesc(out)
		}, zshGitCompletionAdapter),
		completionShellCommand("fish", rootCmd.GenFishCompletion, fishGitCompletionAdapter),
	)
	rootCmd.AddCommand(completionCmd)
}

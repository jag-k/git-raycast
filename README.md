# Git with Raycast AI

Automate git using Raycast AI!

## Installation

### Homebrew

```shell
brew install jag-k/tap/git-raycast
```

### Manual

Download latest build from [Releases](https://github.com/jag-k/git-raycast/releases) and install it in any directory in
your PATH.+

## Usage

```shell
git-raycast
# OR
git raycast
```

### Supported commands

- [`message` - Create commit message based on changes](https://github.com/jag-k/git-raycast/wiki/Commands#message)
- [`mr`/`pr` - Generate MR/PR Summary message](https://github.com/jag-k/git-raycast/wiki/Commands#mr-pr-sumary)
- [`summary` - Create daily summary based on changes since yesterday](https://github.com/jag-k/git-raycast/wiki/Commands#summary)

### Raycast version

By default, `git-raycast` opens commands in stable Raycast:

```shell
git-raycast message
```

For Raycast beta, use:

```shell
git-raycast --raycast-version beta message
# OR
GIT_RAYCAST_VERSION=beta git-raycast message
# OR
git config --global git-raycast.raycast-version beta
```

Settings priority is: CLI argument or flag, environment variable, Git config, default value.

Command names can also be configured with Git config:

```shell
git config --global git-raycast.message-name git-commit-message
git config --global git-raycast.summary-name daily-summary
git config --global git-raycast.mr-pr-summary-name mr-pr-summary
```

Omit `--global` to set a value only for the current repository.

## Development

Requirements:

- [Go](https://golang.org/dl/) version 1.22 or higher

Install dependencies:

```shell
go mod download
```

Build:

```shell
go build -o ./bin/git-raycast ./git-raycast
```

Run:

```shell
./bin/git-raycast
# or
go run ./git-raycast
```

## License

[MIT](LICENSE)

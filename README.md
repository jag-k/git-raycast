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
git raycast
```

### Supported commands

- [`message` - Create commit message based on changes](https://github.com/jag-k/git-raycast/wiki/Commands#message)
- [`summary` - Create daily summary based on changes since yesterday](https://github.com/jag-k/git-raycast/wiki/Commands#summary)

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

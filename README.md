# Git with Raycast AI

Automate git using Raycast AI!

## Installation

### Homebrew

```shell
brew install jag-k/tap/git-raycast
```

### Manual

Download latest build from [Releases](https://github.com/jag-k/git-raycast/releases) and install it in any directory in your PATH.

## Usage

```shell
git raycast
```

## Supported commands

### `message` - Create commit message based on changes

Generate commit message based on not-committed changes.

```shell
git raycast message
```

Calling this Deep-link:

```
raycast://ai-commands/git-commit-message?arguments={diff}
```

[Example AI Command](https://prompts.ray.so/shared?prompts=%7B%22model%22:%22openai-gpt-4%22,%22prompt%22:%22Here%20are%20my%20code%20changes:%5Cn%60%60%60diff%5Cn%7Bargument%20name%3D%5C%22diff%5C%22%7D%5Cn%60%60%60%5Cn%5CnPlease%20generate%20a%20git%20commit%20message%20with%20the%20following%20structure:%5Cn%60%60%60%5Cn%7BCommit%20message%7D%5Cn%5Cn%7BDescription%7D%5Cn%60%60%60%5Cn%5CnFor%20example:%5Cn%60%60%60%5CnAdd%20search%20functionality%20to%20notifications%20%5Cn%5Cn-%20Implement%20search%20by%20issue%20identifiers,%20issue%20titles,%20or%20usernames%20in%20notifications%5Cn-%20Remove%20console.log%20from%20getNotificationTitle%20function%5Cn-%20Enhance%20notification%20item%20display%20with%20additional%20keywords%20for%20search%5Cn%60%60%60%22,%22highlightEdits%22:false,%22icon%22:%22stars%22,%22title%22:%22Git%20Commit%20Message%22,%22creativity%22:%22medium%22%7D)

Bash analog:

```shell
DIFF=$(git diff || git diff --cached || git diff --staged)
if [ -z "$DIFF" ]; then
  echo "No changes to commit"
  exit 1
fi

url_encoded=$(echo -n "$DIFF" | perl -MURI::Escape -ne 'print uri_escape($_)')
open "raycast://ai-commands/git-commit-message?arguments=$url_encoded"
```


### `summary` - Create daily summary based on changes since yesterday

Generate a summary of changes made in the repository since yesterday.

```shell
git raycast summary
```

Calling this Deep-link:

```
raycast://ai-commands/git-daily-summary?arguments={diff}
```

[Example AI Command (ru)](https://prompts.ray.so/shared?prompts=%7B%22model%22:%22openai-gpt-4%22,%22prompt%22:%22%D0%9D%D0%B0%D0%BF%D0%B8%D1%88%D0%B8%20%D0%BD%D0%B5%D0%B1%D0%BE%D0%BB%D1%8C%D1%88%D0%BE%D0%B5%20%D0%BE%D0%B1%D0%BE%D0%B1%D1%89%D0%B5%D0%BD%D0%B8%D0%B5%20%D0%B2%D1%81%D0%B5%D0%B3%D0%BE%20%D1%82%D0%BE%D0%B3%D0%BE,%20%D1%87%D1%82%D0%BE%20%D0%B1%D1%8B%D0%BB%D0%BE%20%D0%B7%D0%B0%20%D0%B4%D0%B5%D0%BD%D1%8C%20%D0%BE%D1%81%D0%BD%D0%BE%D0%B2%D1%8B%D0%B2%D0%B0%D1%8F%D1%81%D1%8C%20%D0%BD%D0%B0%20diff%20%D0%B7%D0%B0%20%D0%B2%D0%B5%D1%81%D1%8C%20%D0%B4%D0%B5%D0%BD%D1%8C.%5Cn%5Cn%D0%9D%D0%95%20%D0%9D%D0%A3%D0%96%D0%9D%D0%9E%20%D0%BE%D0%BF%D0%B8%D1%81%D1%8B%D0%B2%D0%B0%D1%82%D1%8C%20%D0%B8%20%D1%80%D0%B0%D1%81%D1%81%D0%BA%D0%B0%D0%B7%D1%8B%D0%B2%D0%B0%D1%82%D1%8C%20%D0%BF%D1%80%D0%BE%20%D0%BA%D0%B0%D0%B6%D0%B4%D1%8B%D0%B9%20%D1%84%D0%B0%D0%B9%D0%BB%20%D0%BE%D1%82%D0%B4%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE.%20%D0%A1%D0%B4%D0%B5%D0%BB%D0%B0%D0%B9%20%D1%81%D0%B2%D0%BE%D0%B5%D0%BE%D0%B1%D1%80%D0%B0%D0%B7%D0%BD%D1%83%D1%8E%20%D0%BA%D1%80%D0%B0%D1%82%D0%BA%D1%83%D1%8E%20%D0%B2%D1%8B%D0%B6%D0%B8%D0%BC%D0%BA%D1%83%20%D0%B7%D0%B0%20%D0%B2%D0%B5%D1%81%D1%8C%20%D0%B4%D0%B5%D0%BD%D1%8C%20%D0%B2%20%D0%BD%D0%B5%D1%81%D0%BA%D0%BE%D0%BB%D1%8C%D0%BA%D0%BE%20%D0%BF%D1%83%D0%BD%D0%BA%D1%82%D0%BE%D0%B2.%20%D0%9F%D1%80%D0%BE%D1%81%D1%82%D0%BE%20%D1%80%D0%B5%D0%B7%D1%83%D0%BB%D1%8C%D1%82%D0%B0%D1%82%20%D0%BA%D0%BE%D1%82%D0%BE%D1%80%D1%8B%D0%B9%20%D0%B1%D1%8B%D0%BB%20%D0%B4%D0%BE%D1%81%D1%82%D0%B8%D0%B3%D0%BD%D1%83%D1%82.%5Cn%5Cn%D0%9F%D0%BE%D1%81%D1%82%D0%B0%D1%80%D0%B0%D0%B9%D1%81%D1%8F%20%D1%83%D0%BB%D0%BE%D0%B6%D0%B8%D1%82%D1%81%D1%8F%20%D0%9C%D0%90%D0%9A%D0%A1%D0%98%D0%9C%D0%A3%D0%9C%20%D0%B2%204-5%20%D0%BF%D1%83%D0%BD%D0%BA%D1%82%D0%BE%D0%B2.%5Cn%5Cn%D0%9F%D1%80%D0%B8%20%D0%BD%D0%B5%D0%BE%D0%B1%D1%85%D0%BE%D0%B4%D0%B8%D0%BC%D0%BE%D1%81%D1%82%D0%B8%20%D0%BF%D0%B5%D1%80%D0%B5%D0%B2%D0%B5%D0%B4%D0%B8%20%D0%BD%D0%B0%20%D1%80%D1%83%D1%81%D1%81%D0%BA%D0%B8%D0%B9%20%D1%8F%D0%B7%D1%8B%D0%BA.%20%D0%A2%D0%B0%D0%BA%20%D0%B6%D0%B5%20%D0%B8%D1%81%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D1%83%D0%B9%20%D0%B3%D0%BB%D0%B0%D0%B3%D0%BE%D0%BB%D1%8B%20%D0%BF%D1%80%D0%BE%D1%88%D0%B5%D0%B4%D1%88%D0%B5%D0%B3%D0%BE%20%D0%B2%D1%80%D0%B5%D0%BC%D0%B5%D0%BD%D0%B8%20%D0%B2%20%D0%BC%D1%83%D0%B6%D1%81%D0%BA%D0%BE%D0%BC%20%D1%80%D0%BE%D0%B4%D0%B5%20%D0%B4%D0%BB%D1%8F%20%D1%83%D0%BA%D0%B0%D0%B7%D0%B0%D0%BD%D0%B8%D0%B9%20%D1%87%D1%82%D0%BE%20%D0%B1%D1%8B%D0%BB%D0%BE%20%D1%81%D0%B4%D0%B5%D0%BB%D0%B0%D0%BD%D0%BE.%20%D0%9D%D0%B0%D0%BF%D1%80%D0%B8%D0%BC%D0%B5%D1%80:%20%5C%22%D0%9E%D0%B1%D0%BD%D0%BE%D0%B2%D0%B8%D0%BB%20%D0%B7%D0%B0%D0%B2%D0%B8%D1%81%D0%B8%D0%BC%D0%BE%D1%81%D1%82%D0%B8%20%D0%B2%20%D0%BF%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D0%B5%5C%22%5Cn%5Cn%D0%AD%D1%82%D0%BE%D1%82%20%D0%BE%D1%82%D1%87%D1%91%D1%82%20%D1%8F%20%D0%BF%D0%BE%D0%BA%D0%B0%D0%B7%D1%8B%D0%B2%D0%B0%D1%8E%20%D0%BA%D0%BE%D0%BB%D0%BB%D0%B5%D0%B3%D0%B0%D0%BC,%20%D0%BA%D0%BE%D1%82%D0%BE%D1%80%D1%8B%D0%B5%20%D0%BD%D0%B5%20%D0%B7%D0%BD%D0%B0%D1%8E%D1%82%20%D1%81%D0%B0%D0%BC%D1%83%20%D1%81%D1%82%D1%80%D1%83%D0%BA%D1%82%D1%83%D1%80%D1%83%20%D0%BF%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D0%B0%20%D0%B8%20%D0%BD%D0%B5%20%D0%BF%D0%BE%D0%BD%D0%B8%D0%BC%D0%B0%D1%8E%D1%82%20%D0%B7%D0%B0%20%D1%87%D1%82%D0%BE%20%D0%BE%D1%82%D0%B2%D0%B5%D1%87%D0%B0%D0%B5%D1%82%20%D1%82%D0%BE%D1%82%20%D0%B8%D0%BB%D0%B8%20%D0%B8%D0%BD%D0%BE%D0%B9%20%D1%84%D0%B0%D0%B9%D0%BB.%5Cn%5Cn%D0%9E%D1%82%D0%B4%D0%B0%D0%B9%20%D1%80%D0%B5%D0%B7%D1%83%D0%BB%D1%8C%D1%82%D0%B0%D1%82%20%D0%B2%20%D1%82%D0%B0%D0%BA%D0%BE%D0%BC%20%D0%B2%D0%B8%D0%B4%D0%B5:%5Cn%5Cn%60%60%60md%5Cn%D0%9E%D1%82%D1%87%D1%91%D1%82%20%D0%B7%D0%B0%20%D1%81%D0%B5%D0%B3%D0%BE%D0%B4%D0%BD%D1%8F:%5Cn%5Cn-%20%7B%D0%B7%D0%B4%D0%B5%D1%81%D1%8C%20%D0%B2%D0%B5%D1%81%D1%8C%20%D1%81%D0%BF%D0%B8%D1%81%D0%BE%D0%BA%7D%5Cn%5Cn%D0%9F%D0%BB%D0%B0%D0%BD%D1%8B%20%D0%BD%D0%B0%20%D0%B7%D0%B0%D0%B2%D1%82%D1%80%D0%B0:%5Cn%5Cn-%20%7B%D0%BF%D0%BB%D0%B0%D0%BD%D1%8B%20%D0%BD%D0%B0%20%D0%B7%D0%B0%D0%B2%D1%82%D1%80%D0%B0%7D%5Cn%60%60%60%5Cn%5Cn%D0%92%D0%BE%D1%82%20%D0%B8%D1%81%D1%82%D0%BE%D1%80%D0%B8%D1%8F%20%D0%B8%20diff:%5Cn%60%60%60diff%5Cn%7Bargument%20name%3D%5C%22diff%5C%22%7D%5Cn%60%60%60%22,%22highlightEdits%22:false,%22icon%22:%22bullet-points%22,%22title%22:%22Git%20Daily%20Summary%22,%22creativity%22:%22medium%22%7D)

Bash analog:

```shell
GIT_HASH=$(git log -1 --until=yesterday --pretty=format:"%H")
GIT_LOGS=$(git log "$GIT_HASH" HEAD)
GIT_DIFF=$(git diff "$GIT_HASH" HEAD)

url_encoded=$(echo -n "$GIT_DIFF" | perl -MURI::Escape -ne 'print uri_escape($_)')
open "raycast://ai-commands/git-daily-summary?arguments=$url_encoded"
```

## License

[MIT](LICENSE)

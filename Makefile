completions:
	@mkdir -p completions.build
	@go run ./git-raycast/main.go completion bash > ./completions.build/git-raycast.bash
	@go run ./git-raycast/main.go completion zsh > ./completions.build/git-raycast.zsh
	@go run ./git-raycast/main.go completion fish > ./completions.build/git-raycast.fish

# Build package?
1. Create a go module.
    - `go mod init github.com/<group>/<project>`
    - remember to fill the full path of git-project, if wrong, you will get wrong when you use `go mod download`.
2. Write your PKG content.
	- Don't forget use `go test`.
3. git push & release it and use version as v1.0.0

# How to use it
1. Create a repo and `import "your repo"` in main.go
2. Use `go mod download` to download the package.
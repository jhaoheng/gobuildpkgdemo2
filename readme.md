# Build package?
1. Create a git.
2. The Git name does not force has to the same the PKG, but for clarity, we make it the same.
    - for example, this pkg call "lib".
3. Write your PKG content.
	- Don't forget use `go test`.
4. git push & release it and use version as v1.0.0

# How to use it
1. Create  repo, `go mod init <repo>`
2. `import "your repo"`
3. Then use it
    - for example : `lib.listProfiles()`
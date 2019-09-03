# readme

This is a sample about how to build a go pkg

# how to build a pkg
1. Env, if you not use `go module`, you could ignore it.
    - `go mod init github.com/<group>/<project>`
    - remember to fill the full path of git-project, if wrong, you will get wrong when you use `go mod download`.
2. Write your PKG content & test it `go test`.
3. git push & release it and use version as v1.0.0

# try it
1. build a go.mod and add this pkg to it.
2. `go mod download` & run the pkg for test.
3. see the go.mod version have some diff from `https://github.com/jhaoheng/gobuildpkgdemo`
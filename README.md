# all-about-go

## Go CLI

`go run main.go` - compiles and executes one or two files

`go build` - compiles a bunch of go source code files

`go fmt` - formats all the code in each file in the current directory

`go install` - compiles and "installs" a package

`go get` - downloads the raw source code of someone else's package

`go test` - runs any tests associated with the current project

---

## Go Packages

There are two types of packages in Go

1. **Executable package**

   `package main` defines an executable package. **Must have a func called 'main'**.

2. **Reusable package**

    Any package other than the `package main` is a reusable package that can used as a dependency (a library or a helper code).

    eg: `package calculator`, `package uploader`
---
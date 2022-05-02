# all-about-go

* Go CLI
* Go Packages
* Variables in Go

## Go CLI

`go run main.go` - compiles and executes one or two files

`go build` - compiles a bunch of go source code files

`go fmt` - formats all the code in each file in the current directory

`go install` - compiles and "installs" a package

`go get` - downloads the raw source code of someone else's package

`go test` - runs any tests associated with the current project

## Go Packages

There are two types of packages in Go

1. **Executable package**

   `package main` defines an executable package. **Must have a func called 'main'**.

2. **Reusable package**

   Any package other than the `package main` is a reusable package that can used as a dependency (a library or a helper code).

   eg: `package calculator`, `package uploader`

## Variables in Go

Go is a **statically typed language** meaning it needs the type information before the compilation.

General form of variable declaration in Go is:

```go
var name type = expression
```

The above declaration creates a variable of a particular type, attaches a name to it, and sets its initial value. **Either the `type` or the `= expression` can be omitted, but not both.**

For example:

```go
var card string = "Ace of Spades"
var x int = 1
var a int
var b, c, d = 3.14, "apple", true
```

But there is an alternative syntax as well.
`:=` is called **short variable declaration** or **declare-and-initialize construct** which takes form:

```go
name := expression
```

```go
card := "Ace of Spades"
```
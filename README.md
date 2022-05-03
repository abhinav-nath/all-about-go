# all-about-go

- [Go CLI](#go-cli)
- [Go Packages](#go-packages)
- [Variables in Go](#variables-in-go)
- [Functions and Return Types](#functions-and-return-types)
- [Arrays and Slices](#arrays-and-slices)
- [Custom Type Declarations](#custom-type-declarations)
- [Receiver Functions](#receiver-functions)

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

## Functions and Return Types

```go
func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
```

## Arrays and Slices

- **Array** - Fixed length list of things
- **Slice** - A dynamically-sized, flexible view into the elements of an array

### Slice

The type `[]T` is a slice with elements of type `T`.

```go
cards := []string{"Ace of Spades", "Ace of Diamonds"}
```

Add a new item to the slice:

```go
cards = append(cards, "Queen of Hearts")
```

> `append` does not modify the original slice. It creates a new slice instead.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```go
a[low : high]
```

The following expression creates a slice which includes elements 1 through 3 of `a`:

```go
a[1:4]
```

> **Note:** Always remember when you create a slice, it first creates an array and after that return a slice reference to it.

### Iterating Slices

```go
	cards := []string{"Ace of Spades", "Ace of Diamonds"}
	cards = append(cards, "Queen of Hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}
```

Output

```
0 Ace of Spades
1 Ace of Diamonds
2 Queen of Hearts
```

## Custom Type Declarations

Create a new type of `deck` which is a slice of strings

```go
type deck []string
```

We can use it like so

```go
cards := deck{"Ace of Spades", "Ace of Diamonds"}
```

## Receiver Functions

```go
func (d deck) print() {}
```

`d` - The reference to the actual deck variable

`deck` - Every variable of type `deck` can call this function on itself

deck.go

```go
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

main.go

```go
func main() {
	cards := deck{"Ace of Spades", "Ace of Diamonds"}
	cards = append(cards, "Queen of Hearts")
	cards.print()
}
```

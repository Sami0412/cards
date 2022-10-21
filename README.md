# The Project
- Create an application that simulates a deck of cards, with the following functionality:
    - newdeck: create list of playing cards (array of strings)
    - print: log out contents of deck of cards
    - shuffle: shuffle all cards in deck
    - deal: createa "hand" of cards
    - newdeckFronFile: load a list of cards from local machine

# Notes

To run: go run <filename>

Go is not an object oriented programming language - there is no concept of classes.

## Go CLI
go build -> compiles bunch of go source code files
- Just compiles program without executing
- e.g. go build main.go > creates main.exe file

go run -> compiles & executes one or two files
- e.g. go run main.go > runs program

go fmt -> formats all the code in each file in current directory
- Short for format

go install -> compiles and installs a package

go get -> downloads raw source code of someone else's package

go test -> run any tests associated with current project

---

## Packages
- Package == Project == Workspace
- A package is a collection of common source code files
- E.g. Package Main contains main.go, support.go, helper.go etc
- The first line of each file must declare the package to which it belongs - like a namespace in C#
    - E.g. package main
- 2 types of packages: Executable and Reusable
- Executable:
    - Generates a file we can run - .exe
- Reuseable:
    - Code used as "helpers" - code dependencies/libraries
    - Good place to put reusable logic
- The name of the package determines which is made - the word "main" is used to make an executable type package
    - Create package named main > run <go build> > main.exe created
    - Create package named anything else > run <go build> > nothing is created
- Only use main as package name when you want to create a runnable file
- This executable main package MUST have a function inside called "main"

- fmt is a standard library package included with Golang
    - Used to print out info to the terminal, mainly for debugging
- Import statement is like Using statements in C#

## Variable Declaration
var card string = "test" == card := "test"
- Go is statically typed but will infer type in some instances
- := tells Go to create a new variable and infer the type from whatever is on the right
- := can only be used when defining a new variable - not for reassigning
- := can also only be using INSIDE a function - not for initialising global variables
    - You can declare a global variable, and then initialise it inside a function
- Global variables can be declared and initialised outside of a function like: var test = "testing"
    - This is also valid inside a function

## Functions
- Must declare a return type after the function name e.g. func newCard() *string* {}

## Arrays & Slices
- Go has 2 basic data structures for handling lists of records
    - Array: fixed length list, very basic
    - Slice: an array that can change length, flexible, more features
- Both must be defined with a single data type - can't have an array or slice with multiple data types inside
- append: built0in function that adds elements to the end of a slice
    - Original slice is not changed - append returns a new slice

## Custom Type Declarations
- In an OO language, we would probably made a Deck class, and create instances of this class.
- Instead of classes, in Go we can extend the functionality of existing base types
- For this project we can create an extension of an slice of strings as our custom deck type
    - type deck []string
    - We can also attach customised functionality to this type by creating functions with a receiver - ie a method, a function that belongs to an instance
- Declare the new type, then any associated functions need to have a receiver in their declaration
- type deck []string
    func (d deck) print() {...}
    - (d deck) is the receiver
    - print is the function name
- This syntax means that every variable of type deck has access to this print method
- The variable d references the instance of the deck variable that we are working with - so in this card d = cards array
    - The receiver variable is similar to the keyword 'this' in other languages
    - In Go, convention is to name the receiver variable with a 1 or 2 letter abbreviation that matches the type of the receiver


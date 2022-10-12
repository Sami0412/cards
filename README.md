To run: go run <filename>

# Go CLI
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

# Packages
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

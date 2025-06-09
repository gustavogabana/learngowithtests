# LEARN GO WITH TESTS
https://quii.gitbook.io/learn-go-with-tests

## Go Basics
### How to Run

To run the program, execute go run file_name.go

In Go, the program must have the main package and within it. Both are required for the program to run as an executable.

Go uses the concept of modules for code organization. To create a module, `run go mod init someurl.com/module-name` in the terminal. This will generate a `go.mod` file, wich defines the module path and tracks the dependencies and their values used in the project.

To execute a test, create a file with the suffix _test.go, and run `go test` in the terminal.

Writing tests in Go is just like writing regular functions, but with a few requirements:
- The file name needs to end with the _test.go suffix;
- The function name must start with Test;
- The test function must take one argument: `t *testing.T`.

### Constansts
Constants are defined with the keyword `const variableName = "Hi"`. If you need to define multiple constants, you can group them together using parenthesis:
```
const (
    Hello = "Hello"
    World = "World"
)
```

### Subtests
A subtest is a way to group multiple related test scenarios within the same Test function. It allows you to isolate isolate and run each scenario independently. To create a subtest, use `t.Run("description", function{...})`.

### Switch
If your code contains many `if` statements, you can refactor it using a `switch` statement. This makes the code more readable and easier to extend if you need to add new conditions.

### Named return values
In Go, you can create a function with a named return value by declaring the varibale and its type inside parenthesis, right after the function parameters. For example:
```
func greetingPrefix(language string) (prefix string) {...}
```
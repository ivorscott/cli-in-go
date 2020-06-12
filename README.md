# go-cli-training

## Word Counter

The program counts the number of words or lines sent as input.

[word counter](https://github.com/ivorscott/go-cli-training/tree/master/wc)

### Explanation

The program reads from standard input (STDIN) and returns the number of words or lines. It leverages a function called `count` which takes in the `io.Reader` interface and returns an `int`. Any concrete type with a Read method implements the io.Reader interface and is allowed to be passed to count as an argument. `os.Stdin` implements the io.Reader interface so the input data can be read with `bufio.NewScanner(r)`. By default the scanner reads line by line so we need to instruct the function to read each word. The function then processes the input by iterating over each token with `bufio.Scanner.Scan()`. A command line flag is used to scan each line instead of words.

### Demo

```
❯ cd wc
echo "My first command line tool with Go" | ./wc
7

❯ cat main.go | wc -l
38
```

### Cross compilation

Cross compilation is a process of compiling programs for different platforms than the one you are currently on. By default, the current operating system and architecture is detected when using the `go build` tool. We can prefix the command with the `GOOS` environment variable (which stands for Go Operating System) to change the target operating system.

```bash
GOOS=darwin go build
GOOS=linux go build
GOOS=windows go build
```

Run `go env` to reveal all your environment variable settings.

The result of the go build command is a static binary. All the dependencies are packaged inside. So no other runtime dependencies are required. Platforms with a static program binary don't even need Go installed.

On Windows, the word counter static binary could be used like so:

`echo "Testing wc command on Windows" | wc.exe`

## Todo APP

[todo](https://github.com/ivorscott/go-cli-training/tree/master/todo)

### Demo

```
❯ cd todo/cmd/todo

❯ go build .

❯ ./todo -task "Add this item"
❯ ./todo -task "Add this item too"

❯ ./todo -list
  1: Test env vars design
  2: Add this item
  3: Add this item too

❯ ./todo -complete 3

❯ ./todo -list
  1: Test env vars design
  2: Add this item
X 3: Add this item too
```

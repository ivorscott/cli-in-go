# go-cli-training

## Word Counter

The program counts the number of words or lines sent as input.

[word counter](https://github.com/ivorscott/go-cli-training/tree/master/wc)

### Explanation

The program reads from standard input (STDIN) and returns the number of words or lines. It leverages a function called `count` which takes in the `io.Reader` interface and returns an `int`. Any concrete type with a Read method implements the io.Reader interface and is allowed to be passed to count as an argument. `os.Stdin` implements the io.Reader interface so the input data can be read with `bufio.NewScanner(r)`. By default the scanner reads line by line so we need to instruct the function to read each word. The function then processes the input by iterating over each token with `Buffer.Scan()`. A command line flag is used to scan each line instead of words.

### Demo

```
❯ cd wc
echo "My first command line tool with Go" | ./wc
7

❯ cat main.go | wc -l
38
```

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header string = `<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <title>Markdown Preview Tool</title>
  </head>
  <body>
`
	footer string = `
  </body>
</html>
`
)

func run(filename string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	htmlData := parseContent(input)
	// Create temporary file and check for errors
	temp, err := ioutil.TempFile("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}

	return preview(outName)
}

func main() {
	// Parse flags
	filename := flag.String("file", "", "Mardown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	flag.Parse()

	// If user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func parseContent(input []byte) []byte {
	//Parse the mardown file through blackfriday and blue monday
	// to generate a valid and safe HTML
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// Write html to bytes buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(outFname string, data []byte) error {
	// Write the bytes to the file
	return ioutil.WriteFile(outFname, data, 0644)
}

func preview(fname string) error {
	const GOOS string = runtime.GOOS

	switch {
	case GOOS == "darwin":
		if err := exec.Command("open", fname).Start(); err != nil {
			return err
		}

	default:
		// Locate the firefox browser in the PATH
		browserPath, err := exec.LookPath("firefox")
		if err != nil {
			return err
		}

		// Open the file in the browser
		if err := exec.Command(browserPath, fname).Start(); err != nil {
			return err
		}
	}

	return nil
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Stdin.Stat() returns an os.FileInfo
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	// os.FileInfo.Mode() returns a FileMode flag, which represents a file's mode and permission bits.
	// The .Mode() looks like this: `$ Dcrw--w----`
	// The flag we are looking for is os.ModeNamedPipe. When this flag is on it means that we have a pipe.
	// This way we can know when our command is receiving stdout from another process.
	if info.Mode() & os.ModeNamedPipe == 0 {
		fmt.Println("This command is intended to work with pipes.")
	} else {
		// fortune is a CLI tool that prints a random quotation.
		fmt.Println("Usage: fortune | ./gocowsay")
	}

	// At this point, you can query `$ fortune | ./gocowsay` in your CLI

	// The bufio package provides a buffered I/O in GoLang. 
	// NewReader returns a new Reader whose buffer has the default size.
	// We're passing the NewReader the result of the stdin.
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	// ReadRune reads a single UTF-8 encoded Unicode character and returns the rune and its size in bytes. If the encoded rune is invalid, it consumes one byte and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
	// EOF is the error returned by Read when no more input is available.Functions should return EOF only to signal a graceful end of input. If the EOF occurs unexpectedly in a structured data stream, the appropriate error is either ErrUnexpectedEOF or some other error giving more detail.
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}
}
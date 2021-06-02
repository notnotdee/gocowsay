package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// tabsToSpaces converts all tabs found in the lines slice to 4 spaces ([]string), to prevent misalignments in counting the runes
func tabsToSpaces(lines []string) []string {
	var format []string
	for _, line := range lines {
		line = strings.Replace(line, "\t", "    ", -1)
		format = append(format, line)
	}
	return format
}

// calculateMaxWidth takes a slice of strings and returns the length of the string with the longest length
func calculateMaxWidth(lines []string) int {
	width := 0
	for _, line := range lines {
		len := utf8.RuneCountInString(line)
		if len > width {
			width = len
		}
	}

	return width
}

// normalizeStringLength takes a slice of strings and appends to each one a number of spaces required to achieve an equal number of runes
func normalizeStringLength(lines []string, width int) []string {
	var format []string
	for _, line := range lines {
		str := line + strings.Repeat(" ", width-utf8.RuneCountInString(line))
		format = append(format, str)
	}

	return format
}

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
	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("This command is intended to work with pipes.")
		// fortune is a CLI tool that prints a random quotation.
		// fmt.Println("Usage: fortune | ./gocowsay")
		return
	}

	// At this point, you can query `$ fortune | ./gocowsay` in your CLI

	// The bufio package provides a buffered I/O in GoLang.
	// NewReader returns a new Reader whose buffer has the default size.
	// We're passing the NewReader the result of the stdin.
	reader := bufio.NewReader(os.Stdin)
	var lines []string

	// ReadRune reads a single UTF-8 encoded Unicode character and returns the rune and its size in bytes. If the encoded rune is invalid, it consumes one byte and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
	// EOF is the error returned by Read when no more input is available.
	for {
		line, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	for j := 0; j < len(lines); j++ {
		fmt.Printf("%c", lines[j])
	}

	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	lines = tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(lines)
	messages := normalizeStringLength(lines, maxwidth)

	// We need to...
	// Normalize all lines by appending white chars

	// Build the text bubble
	// Print the text bubble and the cow

	fmt.Println(messages)
	fmt.Println(cow)
	fmt.Println(maxwidth)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/mitchellh/go-wordwrap"
	flag "github.com/ogier/pflag"
)

var columns *int32

// readInput takes in and interprets any flags, then formats the standard input and returns a text output slice
func readInput(args []string) []string {
	// Initialize a temp string variable
	var tmps []string

	// If the length of args is 0, there are no flags, so we interpret input as text to be printed
	if len(args) == 0 {
		// Initialize a new scanner to read from os.Stdin
		s := bufio.NewScanner(os.Stdin)

		// Scan will advance the scanner to the next token for as long as that is a valid operation
		for s.Scan() {
			tmps = append(tmps, s.Text())
		}

		// If there is an error, display it
		if s.Err() != nil {
			log.Printf("Err reading stdin: %s\n", s.Err().Error())
			os.Exit(1)
		}

		// If there is no stdin, display an error
		if len(tmps) == 0 {
			fmt.Println("Err: no input from stdin")
			os.Exit(1)
		}
	} else {
		// If the length of args is anything but 0, set tmps equal to args
		tmps = args
	}

	// Initialize a messages variable
	var msgs []string
	for i := 0; i < len(tmps); i++ {
		// Replace tabs with spaces to prep the string for formatting with WrapString
		expand := strings.Replace(tmps[i], "\t", "    ", -1)

		// WrapString wraps the given string within lim width in characters; wrapping only happens at whitespace
		tmp := wordwrap.WrapString(expand, uint(*columns))

		// Use a single append to concatenate two slices
		msgs = append(msgs, strings.Split(tmp, "\n")...)
	}

	return msgs
}

// maxWidth takes in a slice of strings and returns the max width
func maxWidth(msgs []string) int {
	max := 0
	// Loop through the range of msgs and update the max width every time we encounter a larger value
	for _, msg := range msgs {
		len := utf8.RuneCountInString(msg)
		if len > max {
			max = len
		}
	}

	return max
}

// setPadding takes a slice of strings and appends to each one a number of spaces required to achieve an equal number of runes per line
func setPadding(lines []string, width int) []string {
	var format []string
	for _, line := range lines {
		str := line + strings.Repeat(" ", width-utf8.RuneCountInString(line))
		format = append(format, str)
	}

	return format
}

// buildBubble takes a slice of strings of a normalized width, prepends/appends borders on first and last line and at the start/end of each line, then returns a string with the contents of the entire formatted text bubble
func buildBubble(lines []string, width int) string {
	var borders []string
	count := len(lines)
	var format []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", width+2)
	bottom := " " + strings.Repeat("-", width+2)

	format = append(format, top)
	if count == 1 {
		str := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		format = append(format, str)
	} else {
		str := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		format = append(format, str)
		i := 1
		for ; i < count-1; i++ {
			str = fmt.Sprintf("%s %s %s", borders[4], lines[i], borders[4])
			format = append(format, str)
		}
		str = fmt.Sprintf("%s %s %s", borders[2], lines[i], borders[3])
		format = append(format, str)
	}

	format = append(format, bottom)
	return strings.Join(format, "\n")

}

func main() {
	columns = flag.Int32P("columns", "W", 40, "columns")
	// flags here

	inputs := readInput(flag.Args())

	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	width := maxWidth(inputs)
	messages := setPadding(inputs, width)
	bubble := buildBubble(messages, width)

	fmt.Println(bubble)
	fmt.Println(cow)
}

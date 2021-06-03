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

// calculateMaxWidth takes a slice of strings and returns the length of the string with max length
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

// normalizeStringLength takes a slice of strings and appends to each one a number of spaces required to achieve an equal number of runes per line
func normalizeStringLength(lines []string, width int) []string {
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
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("This command is intended to work with pipes.")
		return
	}

	var lines []string
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
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
	bubble := buildBubble(messages, maxwidth)

	fmt.Println(bubble)
	fmt.Println(cow)
}

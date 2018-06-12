package style

import (
	"fmt"
	"strconv"
	"strings"
)

// Config represents the desired configuration for the styling engine.
type Config struct {
	// SeqDH is the Double Header line character sequence.
	SeqDH string
	// SeqHL is the Header Line character sequence.
	SeqHL string
	// SeqLI is the List Item character sequence.
	SeqLI string
	// SeqLL is the Last List item character sequence.
	SeqLL string

	// Header Colour function
	HC func(string) string
	// List Colour function
	LC func(string) string
	// Bold Colour function
	BC func(string) string
	// Italic Colour function
	IC func(string) string
	// Error Colour function
	EC func(string) string

	// SeqStart is the starting sequence for tag matching.
	SeqStart []rune
	// SeqEnd is the ending sequence for tag matching.
	SeqEnd []rune
}

// TagSequence allows you to set the tag characters for the given Style.
func (c *Config) TagSequence(start, end string) {
	c.SeqStart = []rune(start)
	c.SeqEnd = []rune(end)
}

// Print prints out the string with Style tags.
func (c *Config) Print(out string) {
	fmt.Print(c.Style(out))
}

// Println prints out the string with trailing newline and Style tags.
func (c *Config) Println(out string) {
	fmt.Println(c.Style(out))
}

// Printf is analogous to fmt.Printf but with Style tags.
func (c *Config) Printf(format string, a ...interface{}) {
	fmt.Print(c.Sprintf(format, a...))
}

// Printlnf is analogous to fmt.Printf but with trailing newline and Style tags.
func (c *Config) Printlnf(format string, a ...interface{}) {
	fmt.Println(c.Sprintf(format, a...))
}

// Sprintf is analogous to fmt.Sprintf but with Style tags.
func (c *Config) Sprintf(format string, a ...interface{}) string {
	return c.Style(fmt.Sprintf(format, a...))
}

// DH double header line
func (c *Config) DH(n int) string {
	return c.LC(strings.Repeat(c.SeqDH, n))
}

// HL header line
func (c *Config) HL(n int) string {
	return c.LC(strings.Repeat(c.SeqHL, n))
}

// LI list item
func (c *Config) LI(last bool) string {
	if last {
		return c.LC(c.SeqLL)
	}
	return c.LC(c.SeqLI)
}

// Style applies style tags to the given string for this config.
func (c *Config) Style(text string) string {
	if len(c.SeqStart) == 0 {
		c.SeqStart = defaultStart
	}
	if len(c.SeqEnd) == 0 {
		c.SeqEnd = defaultEnd
	}

	for {
		input := []rune(text)
		sp, ep, fn, body := findSequence(c.SeqStart, c.SeqEnd, input)
		if sp == -1 {
			break
		}

		n, _ := strconv.Atoi(body)
		switch strings.ToUpper(fn) {
		case "DH", "DHL":
			body = c.DH(n)
		case "HL", "HR":
			body = c.HL(n)
		case "LI", "LINE":
			body = c.LI(false)
		case "LL", "LLI":
			body = c.LI(true)
		case "HC", "HEAD":
			body = c.HC(body)
		case "LC":
			body = c.LC(body)
		case "B", "BC", "BOLD":
			body = c.BC(body)
		case "I", "IC", "EM", "ITALIC":
			body = c.IC(body)
		case "E", "EC", "ERR", "ERROR":
			body = c.EC(body)
		default:
			body = "<%!:" + fn + ">"
		}

		text = string(input[0:sp])
		text += body
		text += string(input[ep+len(c.SeqEnd):])
	}

	return text
}

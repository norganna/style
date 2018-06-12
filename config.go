package style

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Config represents the desired configuration for the styling engine.
type Config struct {
	ConfigSequences
	ConfigGenerators
	ConfigColours
}

// ConfigSequences define the starting and ending sequences for tag matching.
type ConfigSequences struct {
	// SeqStart is the starting sequence for tag matching.
	SeqStart []rune
	// SeqEnd is the ending sequence for tag matching.
	SeqEnd []rune
}

// ConfigGenerators define a set of manipulators for sequence generation.
type ConfigGenerators struct {
	// GenDH is the Double Header line character sequence.
	GenDH NumericStringManipulator
	// GenHL is the Header Line character sequence.
	GenHL NumericStringManipulator
	// GenLI is the List Item character sequence.
	GenLI StringManipulator
	// GenLL is the Last List item character sequence.
	GenLL StringManipulator
}

// ConfigColours define a set of manipulators for text colourisation.
type ConfigColours struct {
	// Header Colour function
	HC StringManipulator
	// List Colour function
	LC StringManipulator
	// Bold Colour function
	BC StringManipulator
	// Italic Colour function
	IC StringManipulator
	// Error Colour function
	EC StringManipulator
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

// Errorf is an analogous to fmt.Errorf with color tags removed.
func (c *Config) Errorf(format string, a ...interface{}) error {
	return errors.New(c.StripColor().Style(fmt.Sprintf(format, a...)))
}

// Error is a shortcut to errors.New with color tags removed.
func (c *Config) Error(msg string) error {
	return errors.New(c.StripColor().Style(msg))
}

// StripColor returns a copy of the config without color functionality.
func (c *Config) StripColor() *Config {
	return &Config{
		ConfigGenerators: c.ConfigGenerators,
	}
}

// DH double header line
func (c *Config) DH(n int, text string) string {
	return c.LC(c.GenDH(n, text))
}

// HL header line
func (c *Config) HL(n int, text string) string {
	return c.LC(c.GenHL(n, text))
}

// LI list item
func (c *Config) LI(text string, last bool) string {
	if last {
		return c.LC(c.GenLL(text))
	}
	return c.LC(c.GenLI(text))
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
			body = c.DH(n, body)
		case "HL", "HR":
			body = c.HL(n, body)
		case "LI", "LINE":
			body = c.LI(body, false)
		case "LL", "LLI", "LLINE":
			body = c.LI(body, true)
		case "HC", "HEAD":
			body = c.HC(body)
		case "L", "LC":
			body = c.LC(body)
		case "B", "BC", "BOLD":
			body = c.BC(body)
		case "I", "IC", "EM", "ITALIC":
			body = c.IC(body)
		case "E", "EC", "ERR", "ERROR":
			body = c.EC(body)
		default:
			body = "<%!:" + fn + ":" + body + ">"
		}

		text = string(input[0:sp])
		text += body
		text += string(input[ep+len(c.SeqEnd):])
	}

	return text
}

package style

import (
	"strings"

	"github.com/mgutz/ansi"
	"golang.org/x/crypto/ssh/terminal"
)

// StringManipulator is a type of function that takes a string and transforms it.
type StringManipulator func(string) string

// NumericStringManipulator is a type of function that takes a number and a string and transforms it.
type NumericStringManipulator func(int, string) string

var (
	defaultStart = []rune{'‹'}
	defaultEnd   = []rune{'›'}
)

func init() {
	DefaultStyle(Box(), ASCII())
}

func repeater(s string) NumericStringManipulator {
	return func(n int, _ string) string {
		return strings.Repeat(s, n)
	}
}

func literal(s string) StringManipulator {
	return func(msg string) string {
		return s + msg
	}
}

// Box implements a colourised box drawing configuration.
func Box() Config {
	return Config{
		ConfigGenerators: ConfigGenerators{
			GenDH: repeater("═"),
			GenHL: repeater("━"),
			GenLI: literal("┣╸"),
			GenLL: literal("┗╸"),
		},
		ConfigColours: ConfigColours{
			HC: ansi.ColorFunc("green+b"),
			LC: ansi.ColorFunc("green"),
			BC: ansi.ColorFunc("cyan+b"),
			IC: ansi.ColorFunc("yellow+h"),
			EC: ansi.ColorFunc("red+h"),
		},
	}
}

// Bullet implements a colourised bulletised configuration.
func Bullet() Config {
	return Config{
		ConfigGenerators: ConfigGenerators{
			GenDH: repeater("—"),
			GenHL: repeater("–"),
			GenLI: literal("• "),
			GenLL: literal("• "),
		},
		ConfigColours: ConfigColours{
			HC: ansi.ColorFunc("green+b"),
			LC: ansi.ColorFunc("green"),
			BC: ansi.ColorFunc("cyan+b"),
			IC: ansi.ColorFunc("yellow+h"),
			EC: ansi.ColorFunc("red+h"),
		},
	}
}

// ASCII returns an uncoloured ascii art driven configuration.
func ASCII() Config {
	return Config{
		ConfigGenerators: ConfigGenerators{
			GenDH: repeater("="),
			GenHL: repeater("-"),
			GenLI: literal("|-"),
			GenLL: literal("`-"),
		},
		ConfigColours: ConfigColours{
			HC: NC,
			LC: NC,
			BC: NC,
			IC: NC,
			EC: NC,
		},
	}
}

var (
	// defaultStyle represents the desired Style configuration.
	defaultStyle Config
)

// DefaultStyle set the Style used depending upon whether stdOut is a terminal or a file descriptor.
func DefaultStyle(termConf, fdConf Config) {
	if terminal.IsTerminal(1) {
		defaultStyle = termConf
	} else {
		defaultStyle = fdConf
	}
}

// Print prints out the string with Style tags.
func Print(out string) {
	defaultStyle.Print(out)
}

// Println prints out the string with trailing newline and Style tags.
func Println(out string) {
	defaultStyle.Println(out)
}

// Printf is analogous to fmt.Printf but with Style tags.
func Printf(format string, a ...interface{}) {
	defaultStyle.Printf(format, a...)
}

// Printlnf is analogous to fmt.Printf but with trailing newline and Style tags.
func Printlnf(format string, a ...interface{}) {
	defaultStyle.Printlnf(format, a...)
}

// Sprintf is analogous to fmt.Sprintf but with Style tags.
func Sprintf(format string, a ...interface{}) string {
	return defaultStyle.Sprintf(format, a...)
}

// Errorf is a analogous to fmt.Errorf with color tags removed.
func Errorf(format string, a ...interface{}) error {
	return defaultStyle.Errorf(format, a...)
}

// Error is a shortcut to errors.New with color tags removed.
func Error(msg string) error {
	return defaultStyle.Error(msg)
}

// DH double header line
func DH(n int, text string) string {
	return defaultStyle.DH(n, text)
}

// HL header line
func HL(n int, text string) string {
	return defaultStyle.HL(n, text)
}

// LI list item
func LI(text string, last bool) string {
	return defaultStyle.LI(text, last)
}

// LC line coloured
func LC(s string) string {
	return defaultStyle.LC(s)
}

// HC header coloured
func HC(s string) string {
	return defaultStyle.HC(s)
}

// BC bold coloured
func BC(s string) string {
	return defaultStyle.BC(s)
}

// IC italic coloured
func IC(s string) string {
	return defaultStyle.IC(s)
}

// EC error coloured
func EC(s string) string {
	return defaultStyle.EC(s)
}

// NC not coloured
func NC(s string) string {
	return s
}

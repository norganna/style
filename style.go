package style

import (
	"github.com/mgutz/ansi"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	defaultStart = []rune{'‹'}
	defaultEnd = []rune{'›'}
)

func init() {
	DefaultStyle(Box(), ASCII())
}

// Box implements a colourised box drawing configuration.
func Box() Config {
	return Config{
		SeqDH: "═",
		SeqHL: "━",
		SeqLI: "┣╸",
		SeqLL: "┗╸",

		HC: ansi.ColorFunc("green+b"),
		LC: ansi.ColorFunc("green"),
		BC: ansi.ColorFunc("cyan+b"),
		IC: ansi.ColorFunc("yellow+h"),
		EC: ansi.ColorFunc("red+h"),
	}
}

// Bullet implements a colourised bulletised configuration.
func Bullet() Config {
	return Config{
		SeqDH: "—",
		SeqHL: "–",
		SeqLI: "• ",
		SeqLL: "• ",

		HC: ansi.ColorFunc("green+b"),
		LC: ansi.ColorFunc("green"),
		BC: ansi.ColorFunc("cyan+b"),
		IC: ansi.ColorFunc("yellow+h"),
		EC: ansi.ColorFunc("red+h"),
	}
}

// ASCII returns an uncoloured ascii art driven configuration.
func ASCII() Config {
	return Config{
		SeqDH: "=",
		SeqHL: "-",
		SeqLI: "|-",
		SeqLL: "`-",

		HC: NC,
		LC: NC,
		BC: NC,
		IC: NC,
		EC: NC,
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

// DH double header line
func DH(n int) string {
	return defaultStyle.DH(n)
}

// HL header line
func HL(n int) string {
	return defaultStyle.HL(n)
}

// LI list item
func LI(last bool) string {
	return defaultStyle.LI(last)
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

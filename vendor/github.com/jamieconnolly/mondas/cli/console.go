package cli

import (
	"fmt"
	"io"
	"os"
)

var (
	// Exit is the function used to exit the current program with the given status code.
	Exit = os.Exit
	// Stdin is the process's standard input.
	Stdin io.Reader = os.Stdin
	// Stdout is the process's standard output.
	Stdout io.Writer = os.Stdout
	// Stderr is the process's standard error.
	Stderr io.Writer = os.Stderr
)

// Error formats using the default formats for its operands and writes to standard error.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Error(a ...interface{}) (n int, err error) {
	return fmt.Fprint(Stderr, a...)
}

// Errorf formats according to a format specifier and writes to standard error.
// It returns the number of bytes written and any write error encountered.
func Errorf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(Stderr, format, a...)
}

// Errorln formats using the default formats for its operands and writes to standard error.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Errorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(Stderr, a...)
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(Stdout, a...)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(Stdout, format, a...)
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(Stdout, a...)
}

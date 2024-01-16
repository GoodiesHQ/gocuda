package logs

import (
	"errors"
	"io"
)

// Any of the various Log types
type LogEntry interface {
	String() string   // String version of the log (could be the original or pretty-formatted line for example)
	CSV() string      // the values of the data columns converted to strings and joined by a comma
	Fields() []string // function that returns a static list of column names
}

type Parser interface {
	ParseLine(line string) (*LogEntry, *ParseFailure)              // parse a single line (all functions ultimately use this)
	ParseLines(lines []string) ([]LogEntry, []ParseFailure)        // parse from a slice of lines
	ParseFile(filename string) ([]LogEntry, []ParseFailure, error) // parse by a filename. return logs, failures, and error for OS call
	Parse(reader io.Reader) ([]LogEntry, []ParseFailure)           // parse from a reader. Creates a new scanner and reads lines from reader
}

// Parse Failure Operations
type ParseFailure struct {
	Line string // the original erroneous line
	Err  error  // the error associated with the failure
}

// Create a ParseFailure from the line and an error directly
func NewParseFailure(line string, err error) *ParseFailure {
	return &ParseFailure{Line: line, Err: err}
}

// Create a ParseFailure from the line and an error message as a string
func NewParseFailureMessage(line, errorMessage string) *ParseFailure {
	return NewParseFailure(line, errors.New(errorMessage))
}

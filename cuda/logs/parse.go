package logs

import (
	"bufio"
	"io"
	"os"
)

func ParseLine[P Parser](parser P, line string) (LogEntry, *ParseFailure) {
	return parser.ParseLine(line)
}

func Parse[P Parser](parser P, reader io.Reader) ([]LogEntry, []ParseFailure) {
	scanner := bufio.NewScanner(reader)

	entries := make([]LogEntry, 0)
	failures := make([]ParseFailure, 0)

	for scanner.Scan() {
		entry, failure := parser.ParseLine(scanner.Text())
		if entry != nil {
			entries = append(entries, entry)
		}

		if failure != nil {
			failures = append(failures, *failure)
		}
	}

	return entries, failures
}

func ParseFile[P Parser](parser P, filename string) ([]LogEntry, []ParseFailure, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	entries, failures := Parse(parser, file)
	return entries, failures, nil
}

func ParseLines[P Parser](parser P, lines []string) ([]LogEntry, []ParseFailure) {
	entries := make([]LogEntry, 0)
	failures := make([]ParseFailure, 0)

	for _, line := range lines {
		entry, fail := parser.ParseLine(line)
		if entry != nil {
			entries = append(entries, entry)
		}

		if fail != nil {
			failures = append(failures, *fail)
		}
	}

	return entries, failures
}

package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path"

	"github.com/goodieshq/gocuda/cuda/logs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	// "github.com/joho/godotenv"
)

func main2() {
	// simple console logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Log directory
	var logDir string

	flag.StringVar(&logDir, "directory", "", "Directory containing activity logs")
	flag.Parse()

	if logDir == "" {
		flag.Usage()
		return
	}

	// Read the list of files
	files, err := os.ReadDir(logDir)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	count := len(files)
	dfmt := "%0" + fmt.Sprintf("%d", int(math.Floor(math.Log10(float64(count))))+1) + "d"
	msgfmt := "Parsing file " + dfmt + "/" + dfmt + "..."

	if count == 0 {
		log.Fatal().Msg("No Files Found")
	}
	log.Info().Msgf("Found %d files", count)

	// Create parser
	parser := &logs.ActivityLogParser{}
	failuresAll := make([]logs.ParseFailure, 0)

	csv, err := os.Create("logs.csv")
	if err != nil {
		log.Fatal().Msg("Failed to create CSV file")
	}
	defer csv.Close()

	fmt.Fprintln(csv, parser.FieldsCSV())
	for i, file := range files {
		log.Info().Msgf(msgfmt, i+1, count)
		entries, failures, err := logs.ParseFile(parser, path.Join(logDir, file.Name()))
		log.Info().Msgf("Parsed %d lines with %d error(s)", len(entries), len(failures))
		if err != nil {
			log.Fatal().Err(err).Send()
		}
		if len(failures) > 0 {
			failuresAll = append(failuresAll, failures...)
		}
		for _, entry := range entries {
			fmt.Fprintln(csv, entry.CSV())
		}
		if i == 2 {
			break
		}
	}

	if len(failuresAll) == 0 {
		log.Info().Msg("No Errors")
		return
	}
	log.Warn().Msgf("Found %d errors", len(failuresAll))
	failFile, err := os.Create("failures.txt")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer failFile.Close()

	for _, failure := range failuresAll {
		fmt.Fprintln(failFile, failure.Line)
		fmt.Fprintln(failFile, failure.Err.Error())
		fmt.Fprintln(failFile, "")
	}
}

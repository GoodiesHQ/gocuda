package main

import (
	"context"
	"os"
	"strings"

	"github.com/goodieshq/gocuda/api"
	"github.com/goodieshq/gocuda/cuda"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	godotenv.Load()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// load the URL and API key from the environment variables
	url := strings.TrimSuffix(os.Getenv("CUDA_API_URL"), "/")
	key := os.Getenv("CUDA_API_KEY")

	// Create the API client
	client := api.NewClientCC(url, key)
	ctx := context.Background()

	log.Debug().Msgf("API URL: %s Key: %s...", url, key[:5])

	obj := cuda.NetworkObject{
		Name:   "Test1",
		Type:   cuda.NetworkObjectTypeGeneric,
		Shared: false,
		Included: []cuda.NetworkObjectIncludedEntry{
			{
				Entry: cuda.NetworkObjectEntry{
					IP:      "1.2.3.4",
					Comment: "Made from the API",
				},
			},
			{
				Entry: cuda.NetworkObjectEntry{
					IP:      "4.3.2.1",
					Comment: "Made from the API",
				},
			},
		},
		Excluded: []cuda.NetworkObjectExcludedEntry{
			{
				Entry: cuda.NetworkObjectEntry{
					IP:      "4.4.4.4",
					Comment: "Made from the API",
				},
			},
		},
		// Geo: &cuda.NetworkObjectGeoEntry{Included: []string{}, Excluded: []string{}},
	}
	client.ReplaceNetworkObject(ctx, &obj)

	myObj, err := client.GetNetworkObject(ctx, "Any")
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	log.Info().Any("Object", myObj).Send()
}

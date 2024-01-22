package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/goodieshq/gocuda/api"
	"github.com/goodieshq/gocuda/cuda"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func createHandler(event *zerolog.Event, ignoreStatusCodes []int) func(err error) {
	return func(err error) {
		if err == nil {
			return
		}

		var apierr *api.ApiErr
		if errors.As(err, &apierr) {
			if slices.Contains(ignoreStatusCodes, apierr.Code) {
				log.Info().Msgf("error handler ignored status code %d (%s)", apierr.Code, apierr.Message)
				return
			}
			event.Int("status_code", apierr.Code).Msg(apierr.Message)
		}
		event.Err(err).Send()
	}
}

type handler func(err error)

func main() {
	var err error
	var data []byte
	var handle handler
	// this handler will panic on error.

	defer func() {
		// the error handler will panic on error. recover
		if r := recover(); r != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}()

	godotenv.Load()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	tea.ClearScreen()

	// load the URL and API key from the environment variables
	url := strings.TrimSuffix(os.Getenv("CUDA_API_URL"), "/")
	key := os.Getenv("CUDA_API_KEY")

	rid := os.Getenv("CUDA_RANGE_ID")
	cid := os.Getenv("CUDA_CLUSTER_ID")
	bid := os.Getenv("CUDA_BOX_ID")
	sid := os.Getenv("CUDA_SERVICE_ID")

	// Create the API client
	client := api.NewClientCC(url, key)
	ctx := context.Background()

	// provide a context to interact with a specific endpoint
	ctx = client.ContextCC(ctx, api.ContextInfoCC{
		Range:      rid,
		Cluster:    cid,
		Box:        bid,
		BoxService: sid,
	})

	objName := "Example"

	// ignore errors when deeleting object that doesn't exist
	handle = createHandler(log.Panic(), []int{http.StatusNotFound})

	// delete the object (if it exists)
	err = client.DeleteNetworkObject(ctx, objName)
	handle(err)

	// ignore trying to create an object that already exists
	handle = createHandler(log.Panic(), []int{http.StatusConflict})

	// create a new network object within the remote device's config
	err = client.MakeNetworkObject(ctx, &cuda.NetworkObject{
		Name:    objName,
		Type:    cuda.NetworkObjectTypeGeneric,
		Comment: "Just an example",
		Included: []cuda.NetworkObjectIncludedEntry{
			{
				Entry: &cuda.NetworkObjectEntry{
					IP: "7.7.7.0/24",
				},
			},
			{
				References: "Any",
			},
		},
		Excluded: []cuda.NetworkObjectExcludedEntry{
			{
				Entry: &cuda.NetworkObjectEntry{
					IP: "7.7.7.7",
				},
			},
		},
	})
	handle(err)

	// don't ignore any status codes
	handle = createHandler(log.Panic(), nil)

	err = client.DeleteNetworkObjectIncludedEntry(ctx, objName, "Any")
	handle(err)

	err = client.AddNetworkObjectIncludedEntry(ctx, objName, cuda.NetworkObjectIncludedEntry{
		References: "Any",
	})
	handle(err)

	obj, err := client.GetNetworkObject(ctx, objName)
	handle(err)

	fmt.Println()

	data, err = json.MarshalIndent(obj, "", "  ")
	handle(err)
	log.Info().Msg(string(data))

	err = client.ChangeNetworkObjectIncluded(ctx, objName, cuda.NetworkObjectIncludedChange{
		Add: []cuda.NetworkObjectIncludedEntry{
			{
				Entry: &cuda.NetworkObjectEntry{
					IP: "1.2.3.4",
				},
			},
		},
		Remove: []string{"Any"},
	})
	handle(err)

	err = client.ChangeNetworkObjectExcluded(ctx, objName, cuda.NetworkObjectExcludedChange{
		Add: []cuda.NetworkObjectExcludedEntry{
			{
				Entry: &cuda.NetworkObjectEntry{
					IP: "7.7.7.3",
				},
			},
		},
		Remove: []string{"7.7.7.7"},
	})
	handle(err)

	obj, err = client.GetNetworkObject(ctx, objName)
	handle(err)

	data, err = json.MarshalIndent(obj, "", "  ")
	handle(err)
	log.Info().Msg(string(data))

	err = client.DeleteNetworkObject(ctx, objName)
	handle(err)
}

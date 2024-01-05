package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/rs/zerolog/log"
)

type ApiErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (apiErr *ApiErr) String() string {
	return fmt.Sprintf("(%d) %s", apiErr.Code, apiErr.Message)
}

type Api struct {
	client *http.Client
	url    string
	key    string
	isCC   bool
}

func (api *Api) strBy(fwString, ccString string) string {
	if api.isCC {
		return ccString
	}
	return fwString
}

func (api *Api) addHeaders(req *http.Request) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Api-Token", api.key)
}

func (api *Api) send(req *http.Request, target interface{}, codes []int) error {
	api.addHeaders(req)
	res, err := api.client.Do(req)
	if err != nil {
		return err
	}

	if (codes == nil && res.StatusCode/100 != 2) || (codes != nil && !slices.Contains(codes, res.StatusCode)) {
		log.Warn().Msgf("Unexpected Status Code: %d", res.StatusCode)
		apiErr := &ApiErr{}
		if err := json.NewDecoder(res.Body).Decode(apiErr); err != nil {
			return err
		}
		return fmt.Errorf(apiErr.String())
	}

	if target != nil {
		if err := json.NewDecoder(res.Body).Decode(target); err != nil {
			return err
		}
	}

	return nil
}

func (api *Api) Delete(ctx context.Context, endpoint string) error {
	url := api.url + endpoint
	log.Debug().Msgf("DELETE URL: %s", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return api.send(req, nil, nil)
}

func (api *Api) Post(ctx context.Context, endpoint string, body interface{}, target interface{}) error {
	url := api.url + endpoint
	log.Debug().Msgf("POST URL: %s", url)
	var bodyBuf io.Reader = io.Reader(nil)

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyBuf = bytes.NewBuffer(data)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyBuf)
	if err != nil {
		return err
	}

	return api.send(req, target, nil)
}

func (api *Api) Patch(ctx context.Context, endpoint string, body interface{}, target interface{}) error {
	url := api.url + endpoint
	log.Debug().Msgf("PATCH URL: %s", url)
	var bodyBuf io.Reader = io.Reader(nil)

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyBuf = bytes.NewBuffer(data)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bodyBuf)
	if err != nil {
		return err
	}

	return api.send(req, target, nil)
}

func (api *Api) Put(ctx context.Context, endpoint string, body interface{}, target interface{}) error {
	url := api.url + endpoint
	log.Debug().Msgf("PUT URL: %s", url)
	var bodyBuf io.Reader = io.Reader(nil)

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyBuf = bytes.NewBuffer(data)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bodyBuf)
	if err != nil {
		return err
	}

	return api.send(req, target, nil)
}

func (api *Api) Get(ctx context.Context, endpoint string, target interface{}) error {
	url := api.url + endpoint
	log.Debug().Msgf("GET URL: %s", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	return api.send(req, target, nil)
}

func NewClient(url, key string) *Api {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Api{
		client: &http.Client{
			Transport: transport,
		},
		url: url,
		key: key,
	}
}

func NewClientCC(url, key string) *Api {
	client := NewClient(url, key)
	client.isCC = true

	return client
}

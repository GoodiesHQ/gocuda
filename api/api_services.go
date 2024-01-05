package api

import (
	"context"
	"fmt"

	"github.com/goodieshq/gocuda/cuda"
)

func (api *Api) ListServices(ctx context.Context) ([]cuda.Service, error) {
	services := &cuda.Services{}
	if err := api.Get(ctx, "/rest/control/v1/box/services?expand=true", services); err != nil {
		return nil, err
	}

	return services.Services, nil
}

func (api *Api) ListServicesSimple(ctx context.Context) ([]string, error) {
	services := &cuda.ServicesSimple{}
	if err := api.Get(ctx, "/rest/control/v1/box/services?expand=false", services); err != nil {
		return nil, err
	}

	return services.Services, nil
}

func (api *Api) GetService(ctx context.Context, service string) (*cuda.Service, error) {
	s := &cuda.Service{}
	endpoint := fmt.Sprintf("/rest/control/v1/box/services/%s", service)
	if err := api.Get(ctx, endpoint, s); err != nil {
		return nil, err
	}
	return s, nil
}

func (api *Api) RestartService(ctx context.Context, service string) error {
	endpoint := fmt.Sprintf("/rest/control/v1/box/services/%s/restart", service)
	if err := api.Post(ctx, endpoint, nil, nil); err != nil {
		return err
	}
	return nil
}

func (api *Api) StartService(ctx context.Context, service string) error {
	endpoint := fmt.Sprintf("/rest/control/v1/box/services/%s/start", service)
	if err := api.Post(ctx, endpoint, nil, nil); err != nil {
		return err
	}
	return nil
}

func (api *Api) StopService(ctx context.Context, service string) error {
	endpoint := fmt.Sprintf("/rest/control/v1/box/services/%s/stop", service)
	if err := api.Post(ctx, endpoint, nil, nil); err != nil {
		return err
	}
	return nil
}

func (api *Api) BlockService(ctx context.Context, service string) error {
	endpoint := fmt.Sprintf("/rest/control/v1/box/services/%s/stop", service)
	if err := api.Post(ctx, endpoint, nil, nil); err != nil {
		return err
	}
	return nil
}

package api

import (
	"context"

	"github.com/goodieshq/gocuda/cuda"
)

func (api *Api) ListAdmins(ctx context.Context) ([]cuda.Admin, error) {
	admins := &cuda.Admins{}
	if err := api.Get(ctx, "/rest/config/v1/box/admins?expand=true", admins); err != nil {
		return nil, err
	}
	return admins.Admins, nil
}

func (api *Api) ListAdminsSimple(ctx context.Context) ([]string, error) {
	admins := &cuda.AdminsSimple{}
	if err := api.Get(ctx, "/rest/config/v1/box/admins?expand=false", admins); err != nil {
		return nil, err
	}
	return admins.Admins, nil
}

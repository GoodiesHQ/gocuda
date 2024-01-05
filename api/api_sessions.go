package api

import (
	"context"

	"github.com/goodieshq/gocuda/cuda"
)

func (api *Api) ListSessions(ctx context.Context) ([]cuda.Session, error) {
	sessions := &cuda.Sessions{}
	if err := api.Get(ctx, "/rest/control/v1/box/sessions", sessions); err != nil {
		return nil, err
	}

	return sessions.Sessions, nil
}

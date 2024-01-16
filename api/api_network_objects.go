package api

import (
	"context"
	"fmt"

	"github.com/goodieshq/gocuda/cuda"
)

func (api *Api) EndpointRestConfigFirewallNetworkObject(ctx context.Context, objectName string) string {
	return joinPath(api.EndpointRestConfigFirewallObjectsNetworks(ctx), objectName)
}

func (api *Api) MakeNetworkObject(ctx context.Context, object *cuda.NetworkObject) error {
	endpoint := api.EndpointRestConfigFirewallObjectsNetworks(ctx)
	if err := api.Post(ctx, endpoint, object, nil); err != nil {
		return err
	}
	return nil
}

func (api *Api) ReplaceNetworkObject(ctx context.Context, object *cuda.NetworkObject) error {
	endpoint := fmt.Sprintf("%s/%s", api.EndpointNetworkObject(ctx), object.Name)
	if err := api.Put(ctx, endpoint, object, nil); err != nil {
		return err
	}
	return nil
}

func (api *Api) GetNetworkObject(ctx context.Context, name string) (*cuda.NetworkObject, error) {
	object := &cuda.NetworkObject{}
	endpoint := api.EndpointNetworkObject(ctx)
	endpoint := fmt.Sprintf(api.strBy(
		"/rest/config/v1/box/firewall/objects/networks/%s",
		"/rest/cc/v1/config/global/firewall/objects/networks/%s",
	), name)
	if err := api.Get(ctx, endpoint, object); err != nil {
		return nil, err
	}

	return object, nil
}

func (api *Api) UpdateNetworkObject(ctx context.Context, object cuda.NetworkObjectUpdate) error {
	endpoint := fmt.Sprintf(api.strBy(
		"/rest/config/v1/box/firewall/objects/networks/%s",
		"/rest/cc/v1/config/global/firewall/objects/networks/%s",
	), object.Name)
	if err := api.Patch(ctx, endpoint, &object, nil); err != nil {
		return err
	}
	return nil
}

func (api *Api) ListNetworkObjects(ctx context.Context) ([]cuda.NetworkObject, error) {
	endpoint := api.strBy(
		"/rest/config/v1/box/firewall/objects/networks?expand=true",
		"/rest/cc/v1/config/global/firewall/objects/networks?expand=true",
	)
	objects := &cuda.NetworkObjects{}
	if err := api.Get(ctx, endpoint, objects); err != nil {
		return nil, err
	}

	return objects.NetworkObjects, nil
}

func (api *Api) ListNetworkObjectsSimple(ctx context.Context) ([]string, error) {
	endpoint := api.strBy(
		"/rest/config/v1/box/firewall/objects/networks?expand=false",
		"/rest/cc/v1/config/global/firewall/objects/networks?expand=false",
	)
	objects := &cuda.NetworkObjectsSimple{}
	if err := api.Get(ctx, endpoint, objects); err != nil {
		return nil, err
	}

	return objects.NetworkObjects, nil
}

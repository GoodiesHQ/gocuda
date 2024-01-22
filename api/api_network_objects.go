package api

import (
	"context"

	"github.com/goodieshq/gocuda/cuda"
)

/* endpoint for networking objects */

func (api *Api) EndpointRestConfigFirewallObjectsNetworks(ctx context.Context) (string, error) {
	path, err := api.EndpointRestConfigFirewallObjects(ctx)
	if err != nil {
		return "", err
	}
	return joinPath(path, "networks"), nil
}

func (api *Api) EndpointRestConfigFirewallObjectsNetworksByName(ctx context.Context, objectName string) (string, error) {
	path, err := api.EndpointRestConfigFirewallObjectsNetworks(ctx)
	if err != nil {
		return "", err
	}

	return joinPath(path, objectName), nil
}

/* Implementations */

// list all network objects
func (api *Api) ListNetworkObjects(ctx context.Context) ([]cuda.NetworkObject, error) {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworks(ctx)
	if err != nil {
		return nil, err
	}

	objects := &cuda.NetworkObjects{}
	if err := api.Get(ctx, endpoint+"?expand=true", objects); err != nil {
		return nil, err
	}

	return objects.NetworkObjects, nil
}

// let all names of network objects
func (api *Api) ListNetworkObjectsSimple(ctx context.Context) ([]string, error) {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworks(ctx)
	if err != nil {
		return nil, err
	}

	objects := &cuda.NetworkObjectsSimple{}
	if err := api.Get(ctx, endpoint+"?expand=false", objects); err != nil {
		return nil, err
	}

	return objects.NetworkObjects, nil
}

// create a new network object
func (api *Api) MakeNetworkObject(ctx context.Context, object *cuda.NetworkObject) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworks(ctx)
	if err != nil {
		return err
	}

	if err := api.Post(ctx, endpoint, object, nil); err != nil {
		return err
	}

	return nil
}

// delete a network object by name
func (api *Api) DeleteNetworkObject(ctx context.Context, objectName string) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, objectName)
	if err != nil {
		return err
	}

	if err := api.Delete(ctx, endpoint); err != nil {
		return err
	}

	return nil
}

// get a network object by name
func (api *Api) GetNetworkObject(ctx context.Context, objectName string) (*cuda.NetworkObject, error) {
	object := &cuda.NetworkObject{}
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, objectName)
	if err != nil {
		return nil, err
	}

	if err := api.Get(ctx, endpoint, object); err != nil {
		return nil, err
	}

	return object, nil
}

// update a network object with a differential value change
func (api *Api) UpdateNetworkObject(ctx context.Context, object cuda.NetworkObjectUpdate) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, object.Name)
	if err != nil {
		return err
	}

	if err := api.Patch(ctx, endpoint, &object, nil); err != nil {
		return err
	}

	return nil
}

// replace a network object with a completely new value
func (api *Api) ReplaceNetworkObject(ctx context.Context, object *cuda.NetworkObject) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, object.Name)
	if err != nil {
		return err
	}

	if err := api.Put(ctx, endpoint, object, nil); err != nil {
		return err
	}

	return nil
}

func (api *Api) ChangeNetworkObjectExcluded(ctx context.Context, objectName string, object cuda.NetworkObjectExcludedChange) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, objectName)
	endpoint = joinPath(endpoint, "excluded")
	if err != nil {
		return err
	}

	if err := api.Patch(ctx, endpoint, &object, nil); err != nil {
		return err
	}

	return nil
}

func (api *Api) AddNetworkObjectExcludedEntry(ctx context.Context, objectName string, object cuda.NetworkObjectExcludedEntry) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, objectName)
	if err != nil {
		return err
	}
	endpoint = joinPath(endpoint, "excluded")

	if err := api.Post(ctx, endpoint, &object, nil); err != nil {
		return err
	}

	return nil
}

func (api *Api) DeleteNetworkObjectExcludedEntry(ctx context.Context, objectName string, entry string) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, objectName)
	if err != nil {
		return err
	}
	endpoint = joinPath(endpoint, "excluded", entry)
	if err := api.Delete(ctx, endpoint); err != nil {
		return err
	}

	return nil
}

func (api *Api) ChangeNetworkObjectIncluded(ctx context.Context, objectName string, object cuda.NetworkObjectIncludedChange) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, objectName)
	endpoint = joinPath(endpoint, "included")
	if err != nil {
		return err
	}

	if err := api.Patch(ctx, endpoint, &object, nil); err != nil {
		return err
	}

	return nil
}

func (api *Api) AddNetworkObjectIncludedEntry(ctx context.Context, objectName string, object cuda.NetworkObjectIncludedEntry) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, objectName)
	if err != nil {
		return err
	}
	endpoint = joinPath(endpoint, "included")

	if err := api.Post(ctx, endpoint, &object, nil); err != nil {
		return err
	}

	return nil
}

func (api *Api) DeleteNetworkObjectIncludedEntry(ctx context.Context, objectName string, entry string) error {
	endpoint, err := api.EndpointRestConfigFirewallObjectsNetworksByName(ctx, objectName)
	if err != nil {
		return err
	}
	endpoint = joinPath(endpoint, "included", entry)
	if err := api.Delete(ctx, endpoint); err != nil {
		return err
	}

	return nil
}

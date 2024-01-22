package api

import (
	"context"
	"fmt"
)

// Base of the REST API
func (api *Api) EndpointRestBase() string {
	if api.isCC {
		return joinPath("rest", "cc", "v1")
	} else {
		return joinPath("rest", "v1")
	}
}

func (api *Api) EndpointRestConfigBase() string {
	if api.isCC {
		return joinPath(api.EndpointRestBase(), "config")
	} else {
		return joinPath("rest", "config", "v1")
	}
}

func (api *Api) EndpointRestControlBase() string {
	if api.isCC {
		return joinPath(api.EndpointRestBase(), "control")
	} else {
		return joinPath("rest", "control", "v1")
	}
}

func (api *Api) EndpointRestConfigCtx(ctx context.Context) (string, error) {
	if !api.isCC {
		// standalone box config
		// TODO: implement all the other endpoints for standalone boxes
		return joinPath(api.EndpointRestConfigBase(), "box"), nil
	}

	var path string = api.EndpointRestConfigBase()

	rangeID := ctx.Value(CTX_CC_RANGE)
	if rangeID == nil {
		// range is empty, use global endpoint
		return joinPath(path, "global"), nil
	}

	path = joinPath(path, "ranges", rangeID.(string))

	clusterID := ctx.Value(CTX_CC_CLUSTER)
	if clusterID == nil {
		// cluster is empty, use range endpoint
		return path, nil
	}

	// from here on out, all paths will contain the cluster
	path = joinPath(path, "clusters", clusterID.(string))

	// check if box and/or boxService are provided
	if boxID := ctx.Value(CTX_CC_BOX); boxID != nil {
		path = joinPath(path, "boxes", boxID.(string))
		if boxService := ctx.Value(CTX_CC_BOX_SERVICE); boxService != nil {
			return joinPath(path, "service-container", boxService.(string)), nil
		}
		return path, nil
	}

	if serverID := ctx.Value(CTX_CC_SERVER); serverID != nil {
		if serverService := ctx.Value(CTX_CC_SERVER_SERVICE); serverService != nil {
			path = joinPath(path, "servers", serverID.(string), "services", serverService.(string))
			return path, nil
		}
		// if server ID is provided, service must be provided
		return "", fmt.Errorf("no service provided with server '%s'", serverID.(string))
	}

	clusterService := ctx.Value(CTX_CC_CLUSTER_SERVICE)
	if clusterService != nil {
		return joinPath(path, "services", clusterService.(string)), nil
	}

	sharedService := ctx.Value(CTX_CC_SHARED_SERVICE)
	if sharedService != nil {
		return joinPath(path, "shared-services", clusterService.(string)), nil
	}

	return path, nil
}

func (api *Api) EndpointRestConfigFirewall(ctx context.Context) (string, error) {
	path, err := api.EndpointRestConfigCtx(ctx)
	if err != nil {
		return "", err
	}

	if fwShared := ctx.Value(CTX_FW_SHARED); fwShared == nil || !fwShared.(bool) {
		return joinPath(path, "firewall"), nil
	}

	path = joinPath(path, "shared-firewall")

	if fwSharedType := ctx.Value(CTX_FW_SHARED_TYPE); fwSharedType != nil {
		clusterService := ctx.Value(CTX_CC_CLUSTER_SERVICE)
		sharedService := ctx.Value(CTX_CC_SHARED_SERVICE)

		if clusterService != nil || sharedService != nil {
			return "", fmt.Errorf("firewall SharedType is not permitted on cluster services or cluster shared services")
		}

		return joinPath(path, fwSharedType.(string)), nil
	}

	return path, nil
}

func (api *Api) EndpointRestConfigFirewallObjects(ctx context.Context) (string, error) {
	path, err := api.EndpointRestConfigFirewall(ctx)
	if err != nil {
		return "", err
	}
	return joinPath(path, "objects"), nil
}

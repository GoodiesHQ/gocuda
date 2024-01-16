package api

import "context"

func (api *Api) EndpointRestConfig() string {
	if api.isCC {
		return joinPath("rest", "cc", "v1", "config")
	} else {
		return joinPath("rest", "config", "v1")
	}
}

func (api *Api) EndpointRestConfigCtx(ctx context.Context) string {
	if !api.isCC {
		// standalone box config
		return joinPath(api.EndpointRestConfig(), "box")
	}

	rangeID := ctx.Value("range")
	if rangeID == nil {
		// range is empty, use global
		return joinPath(api.EndpointRestConfig(), "global")
	}

	clusterID := ctx.Value("cluster")
	if clusterID == nil {
		// cluster is empty, use range
		return joinPath(api.EndpointRestConfig(), "ranges", rangeID.(string))
	}

	boxID := ctx.Value("box")
	if boxID == nil {
		// box is empty, use cluster
		return joinPath(api.EndpointRestConfig(), "ranges", rangeID.(string), "clusters", clusterID.(string))
	} else {
		switch boxID.(string) {
		case "test":
			return ""
		}
	}

	// use a specific range/cluster/box
	return joinPath(api.EndpointRestConfig(), "ranges", rangeID.(string), "clusters", clusterID.(string), "boxes", boxID.(string))
}

func (api *Api) EndpointRestConfigFirewall(ctx context.Context) string {
	var fw string
	fwShared := ctx.Value("fw-shared")
	if fwShared == nil {
		fw = "firewall"
	} else {
		fw = "shared-firewall"
	}
	return joinPath(api.EndpointRestConfigCtx(ctx), fw)
}

func (api *Api) EndpointRestConfigFirewallObjects(ctx context.Context) string {
	return joinPath(api.EndpointRestConfigFirewall(ctx), "objects")
}

func (api *Api) EndpointRestConfigFirewallObjectsNetworks(ctx context.Context) string {
	return joinPath(api.EndpointRestConfigFirewallObjects(ctx), "networks")
}

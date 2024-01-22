package api

import (
	"context"
)

const (
	CTX_CC_RANGE           = "cc-range"
	CTX_CC_CLUSTER         = "cc-cluster"
	CTX_CC_BOX             = "cc-box"
	CTX_CC_BOX_SERVICE     = "cc-box-service"
	CTX_CC_SERVER          = "cc-server"
	CTX_CC_SERVER_SERVICE  = "cc-server-service"
	CTX_CC_CLUSTER_SERVICE = "cc-cluster-service"
	CTX_CC_SHARED_SERVICE  = "cc-shared-service"
)

const (
	CTX_FW_SHARED      = "fw-shared"
	CTX_FW_SHARED_TYPE = "fw-shared-type"
)

type FirewallSharedType string

const (
	FirewallSharedLocal   FirewallSharedType = "local"
	FirewallSharedSpecial FirewallSharedType = "special"
)

// Context information passed
type ContextInfoCC struct {
	// if empty, is global
	Range   string
	Cluster string
	// either:
	Box        string
	BoxService string
	// or...
	Server        string
	ServerService string
	// or...
	ClusterService string
	// or...
	SharedService string
}

// Create a context for control center. Global is default.
func (api *Api) ContextCC(ctx context.Context, ctxInfo ContextInfoCC) context.Context {
	if api.isCC {
		// range must be defined for anything
		if ctxInfo.Range == "" {
			return ctx
		}
		ctx = context.WithValue(ctx, CTX_CC_RANGE, ctxInfo.Range)

		// cluster not defined, uses range config
		if ctxInfo.Cluster == "" {
			return ctx
		}
		ctx = context.WithValue(ctx, CTX_CC_CLUSTER, ctxInfo.Cluster)

		// if Box is specified, only BoxService needs to be added to the context
		if ctxInfo.Box != "" {
			ctx = context.WithValue(ctx, CTX_CC_BOX, ctxInfo.Box)
			if ctxInfo.BoxService != "" {
				ctx = context.WithValue(ctx, CTX_CC_BOX_SERVICE, ctxInfo.BoxService)
			}
			return ctx
		}

		// if Server is specified, a ServerService must be provided
		if ctxInfo.Server != "" {
			ctx = context.WithValue(ctx, CTX_CC_SERVER, ctxInfo.Server)
			ctx = context.WithValue(ctx, CTX_CC_SERVER_SERVICE, ctxInfo.ServerService)
			return ctx
		}

		// cluster shared-service
		if ctxInfo.SharedService != "" {
			ctx = context.WithValue(ctx, CTX_CC_SHARED_SERVICE, ctxInfo.SharedService)
			return ctx
		}

		// cluster service
		if ctxInfo.ClusterService != "" {
			ctx = context.WithValue(ctx, CTX_CC_SHARED_SERVICE, ctxInfo.SharedService)
		}
	}

	return ctx
}

type ContextInfoFW struct {
	Shared     bool
	SharedType FirewallSharedType
}

func (api *Api) ContextFW(ctx context.Context, ctxInfo ContextInfoFW) context.Context {
	if ctxInfo.Shared {
		ctx = context.WithValue(ctx, CTX_FW_SHARED, true)
		ctx = context.WithValue(ctx, CTX_FW_SHARED_TYPE, ctxInfo.SharedType)
	}
	return ctx
}

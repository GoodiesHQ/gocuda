package api

import (
	"context"

	"github.com/goodieshq/gocuda/cuda"
)

func (api *Api) ListRanges(ctx context.Context) ([]cuda.Range, error) {
	ranges := &cuda.Ranges{}
	endpoint := joinPath(api.EndpointRestBase(), "ranges?expand=true")
	if err := api.Get(ctx, endpoint, ranges); err != nil {
		return nil, err
	}
	return ranges.Ranges, nil
}

func (api *Api) ListRangesSimple(ctx context.Context) ([]string, error) {
	ranges := &cuda.RangesSimple{}
	endpoint := joinPath(api.EndpointRestBase(), "ranges?expand=false")
	if err := api.Get(ctx, endpoint, ranges); err != nil {
		return nil, err
	}
	return ranges.Ranges, nil
}

func (api *Api) ListClusters(ctx context.Context, rangeID string) ([]cuda.Cluster, error) {
	clusters := &cuda.Clusters{}
	endpoint := joinPath(api.EndpointRestBase(), "ranges", rangeID, "clusters?expand=true")
	if err := api.Get(ctx, endpoint, clusters); err != nil {
		return nil, err
	}
	return clusters.Clusters, nil
}

func (api *Api) ListClustersSimple(ctx context.Context, rangeID string) ([]string, error) {
	clusters := &cuda.ClustersSimple{}
	endpoint := joinPath(api.EndpointRestBase(), "ranges", rangeID, "clusters?expand=false")
	if err := api.Get(ctx, endpoint, clusters); err != nil {
		return nil, err
	}
	return clusters.Clusters, nil
}

func (api *Api) ListBoxes(ctx context.Context, rangeID string, clusterName string) ([]cuda.Box, error) {
	boxes := &cuda.Boxes{}
	endpoint := joinPath(api.EndpointRestBase(), "ranges", rangeID, "clusters", clusterName, "boxes?expand=true")
	if err := api.Get(ctx, endpoint, boxes); err != nil {
		return nil, err
	}
	return boxes.Boxes, nil
}

func (api *Api) ListBoxesSimple(ctx context.Context, rangeID string, clusterName string) ([]string, error) {
	boxes := &cuda.BoxesSimple{}
	endpoint := joinPath(api.EndpointRestBase(), "ranges", rangeID, "clusters", clusterName, "boxes?expand=false")
	if err := api.Get(ctx, endpoint, boxes); err != nil {
		return nil, err
	}
	return boxes.Boxes, nil
}

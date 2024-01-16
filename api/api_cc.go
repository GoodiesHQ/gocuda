package api

import (
	"context"
	"fmt"

	"github.com/goodieshq/gocuda/cuda"
)

func (api *Api) ListRanges(ctx context.Context) ([]cuda.Range, error) {
	ranges := &cuda.Ranges{}
	endpoint := "/rest/cc/v1/ranges?expand=true"
	if err := api.Get(ctx, endpoint, ranges); err != nil {
		return nil, err
	}

	return ranges.Ranges, nil
}

func (api *Api) ListRangesSimple(ctx context.Context) ([]string, error) {
	ranges := &cuda.RangesSimple{}
	endpoint := "/rest/cc/v1/ranges?expand=false"
	if err := api.Get(ctx, endpoint, ranges); err != nil {
		return nil, err
	}

	return ranges.Ranges, nil
}

func (api *Api) ListClusters(ctx context.Context, rangeID string) ([]cuda.Cluster, error) {
	clusters := &cuda.Clusters{}
	endpoint := fmt.Sprintf("/rest/cc/v1/ranges/%s/clusters?expand=true", rangeID)
	if err := api.Get(ctx, endpoint, clusters); err != nil {
		return nil, err
	}
	return clusters.Clusters, nil
}

func (api *Api) ListClustersSimple(ctx context.Context, rangeID string) ([]string, error) {
	clusters := &cuda.ClustersSimple{}
	endpoint := fmt.Sprintf("/rest/cc/v1/ranges/%s/clusters?expand=false", rangeID)
	if err := api.Get(ctx, endpoint, clusters); err != nil {
		return nil, err
	}
	return clusters.Clusters, nil
}

func (api *Api) ListBoxes(ctx context.Context, rangeID string, clusterName string) {

}

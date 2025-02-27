package pools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/onchaindex/pools/types"
)

// TrendingPools This endpoint allows you to query the trending pools based on the provided network.
// Cache/Update Frequency: Every 60 seconds for all the API plans
func (c *PoolsClient) TrendingPools(ctx context.Context, network string, duration string, total int) (*types.TrendingPools, error) {
	totalPages := total/20 + 1
	var pools types.TrendingPools
	for i := 1; i <= totalPages; i++ {
		rUrl := fmt.Sprintf("%s/%s/%s?include=base_token&duration=%s&page=%d", c.poolsUrl(), network, "trending_pools", duration, i)
		resp, err := c.MakeReq(ctx, rUrl)
		if err != nil {
			return nil, err
		}

		var data types.TrendingPools
		err = json.Unmarshal(resp, &data)
		if err != nil {
			return nil, err
		}
		pools.Data = append(pools.Data, data.Data...)
		pools.Included = append(pools.Included, data.Included...)
		if len(data.Data) < 20 {
			break
		}
	}

	return &pools, nil
}

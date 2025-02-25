package coins

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/coins/types"
)

// NewCoins This endpoint allows you to query the latest 200 coins that recently listed on CoinGecko
//
// 👍 Tips
// CoinGecko equivalent page: https://www.coingecko.com/en/new-cryptocurrencies
// 📘 Notes
// There is no pagination required for this endpoint
// Cache/Update Frequency: Every 30 seconds
func (c *CoinsClient) NewCoins(ctx context.Context) ([]*types.NewCoinInfo, error) {
	rUrl := fmt.Sprintf("%s/%s", c.coinsUrl(), "list/new")
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []*types.NewCoinInfo
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

package tokens

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api/onchaindex/tokens/types"
)

// This endpoint allows you to query specific token info such as name, symbol, CoinGecko ID etc. based on provided token contract address on a network.
// Cache/Update frequency: every 60 seconds
func (c *TokensClient) TokenMetadata(ctx context.Context, network, tokenAddress string) (*types.TokenMetadata, error) {
	rUrl := fmt.Sprintf("%s/%s/tokens/%s/info", c.tokensUrl(), network, tokenAddress)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data types.TokenMetadataResp
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

// This endpoint allows you to query specific token data based on the provided token contract address on a network.
// Cache/Update frequency: every 30 seconds.
func (c *TokensClient) TokensMarketData(ctx context.Context, network, tokenAddress string) (*types.TokenMarketData, error) {
	rUrl := fmt.Sprintf("%s/%s/tokens/%s", c.tokensUrl(), network, tokenAddress)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data types.TokenMarketDataResp
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

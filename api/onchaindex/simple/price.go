package simple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/onchaindex/simple/types"
)

// SimplePrice  allows you to query the prices of one or more coins by using their unique Coin API IDs.
//
//	 👍 Tips
//
//	    You may obtain the coin id (api id) via several ways:
//	        refers to respective coin page and find ‘api id’
//	        refers to /coins/list endpoint
//	        refers to google sheets here
//	    You may flag to include more data such as market cap, 24hr volume, 24hr change, last updated time etc.
//	    To verify if a price is stale, you may flag include_last_updated_at=true in your request to obtain the latest updated time. Alternatively, you may flag include_24hr_change=true to determine if it returns a 'null' value.
//
//	📘Notes
//
//	    You may cross-check the price on CoinGecko and learn more about our price methodology here
//	    Cache/Update Frequency: every 30 seconds for Pro API (Analyst, Lite, Pro, Enterprise)
func (c *TokenSimpleClient) SimplePrice(ctx context.Context, network, address string, includeMarketCap bool, options ...priceOption) (*types.TokenPriceResp, error) {
	params := url.Values{}

	params.Add("include_market_cap", strconv.FormatBool(includeMarketCap))

	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/networks/%s/%s/%s?%s", c.simpleUrl(), network, "token_price", address, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data types.TokenPriceResp
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

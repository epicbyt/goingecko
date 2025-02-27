package types

import "time"

type PoolData struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		BaseTokenPriceUsd             string    `json:"base_token_price_usd"`
		BaseTokenPriceNativeCurrency  string    `json:"base_token_price_native_currency"`
		QuoteTokenPriceUsd            string    `json:"quote_token_price_usd"`
		QuoteTokenPriceNativeCurrency string    `json:"quote_token_price_native_currency"`
		BaseTokenPriceQuoteToken      string    `json:"base_token_price_quote_token"`
		QuoteTokenPriceBaseToken      string    `json:"quote_token_price_base_token"`
		Address                       string    `json:"address"`
		Name                          string    `json:"name"`
		PoolCreatedAt                 time.Time `json:"pool_created_at"`
		FdvUsd                        string    `json:"fdv_usd"`
		MarketCapUsd                  string    `json:"market_cap_usd"`
		PriceChangePercentage         struct {
			M5  string `json:"m5"`
			H1  string `json:"h1"`
			H6  string `json:"h6"`
			H24 string `json:"h24"`
		} `json:"price_change_percentage"`
		VolumeUsd struct {
			M5  string `json:"m5"`
			H1  string `json:"h1"`
			H6  string `json:"h6"`
			H24 string `json:"h24"`
		} `json:"volume_usd"`
		ReserveInUsd string `json:"reserve_in_usd"`
	} `json:"attributes"`
	Relationships struct {
		BaseToken struct {
			Data struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"base_token"`
	} `json:"relationships"`
}

type PoolTokenAttribute struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Address         string `json:"address"`
		Name            string `json:"name"`
		Symbol          string `json:"symbol"`
		ImageUrl        string `json:"image_url"`
		CoingeckoCoinId string `json:"coingecko_coin_id"`
		Decimals        int    `json:"decimals"`
	} `json:"attributes"`
}

type TrendingPools struct {
	Data     []PoolData           `json:"data"`
	Included []PoolTokenAttribute `json:"included"`
}

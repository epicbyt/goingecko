package types

type TokenMetadata struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Address         string   `json:"address"`
		Name            string   `json:"name"`
		Symbol          string   `json:"symbol"`
		ImageUrl        string   `json:"image_url"`
		CoingeckoCoinId string   `json:"coingecko_coin_id"`
		Websites        []string `json:"websites"`
		Description     string   `json:"description"`
		GtScore         float64  `json:"gt_score"`
		DiscordUrl      string   `json:"discord_url"`
		TelegramHandle  string   `json:"telegram_handle"`
		TwitterHandle   string   `json:"twitter_handle"`
		Categories      []string `json:"categories"`
	} `json:"attributes"`
}

type TokenMetadataResp struct {
	Data TokenMetadata `json:"data"`
}

type TokenMarketDataResp struct {
	Data TokenMarketData `json:"data"`
}

type TokenMarketData struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Address           string `json:"address"`
		Name              string `json:"name"`
		Symbol            string `json:"symbol"`
		ImageUrl          string `json:"image_url"`
		CoingeckoCoinId   string `json:"coingecko_coin_id"`
		Decimals          int    `json:"decimals"`
		TotalSupply       string `json:"total_supply"`
		PriceUsd          string `json:"price_usd"`
		FdvUsd            string `json:"fdv_usd"`
		TotalReserveInUsd string `json:"total_reserve_in_usd"`
		VolumeUsd         struct {
			H24 string `json:"h24"`
		} `json:"volume_usd"`
		MarketCapUsd string `json:"market_cap_usd"`
	} `json:"attributes"`
}

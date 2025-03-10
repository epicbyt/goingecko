package types

type (
	Price       map[string]PriceValues
	PriceValues map[string]float64
)

type TokenPriceResp struct {
	Data Data `json:"data"`
}

type Data struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	TokenPrices              map[string]string `json:"token_prices"`
	MarketCapUsd             map[string]string `json:"market_cap_usd"`
	H24VolumeUsd             map[string]string `json:"h24_volume_usd"`
	H24PriceChangePercentage map[string]string `json:"h24_price_change_percentage"`
	TotalReserveInUsd        map[string]string `json:"total_reserve_in_usd"`
}

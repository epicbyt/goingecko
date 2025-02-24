package types

type AssetPlatforms []Asset

type Asset struct {
	ID              string `json:"id"`
	ChainIdentifier int64  `json:"chain_identifier"`
	Name            string `json:"name"`
	ShortName       string `json:"shortName"`
	NativeCoinId    string `json:"native_coin_id"`
	Image           struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
}

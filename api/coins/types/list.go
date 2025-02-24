package types

type CoinInfo struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms"`
}

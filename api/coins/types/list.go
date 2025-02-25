package types

type CoinInfo struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms"`
}

type NewCoinInfo struct {
	Id          string `json:"id"`
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	ActivatedAt int64  `json:"activated_at"`
}

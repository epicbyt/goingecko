package api

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_CoinsList(t *testing.T) {
	c := NewDefaultClient()

	got, err := c.CoinsList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("nil response")
	}
}

func TestCoinsClient_CoinsMarket(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewProApiClient("xxx", &http.Client{})
			got, err := c.CoinsClient.CoinsMarket(context.Background(), "usd", []string{"bitcoin", "ethereum"})
			if err != nil {
				t.Errorf("CoinsList() error = %v", err)
			}
			if got == nil {
				t.Errorf("CoinsList() got = nil, want not nil")
			}
		})
	}
}

func TestCoinsClient_NewCoins(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewProApiClient("xxx", &http.Client{})
			got, err := c.CoinsClient.NewCoins(context.Background())
			if err != nil {
				t.Errorf("CoinsList() error = %v", err)
			}
			if got == nil {
				t.Errorf("CoinsList() got = nil, want not nil")
			}
		})
	}
}

func TestCoinsClient_TrendingPools(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewProApiClient("xxx", &http.Client{})

			pools, err := c.PoolsClient.TrendingPools(context.Background(), "bsc", "6h", 70)
			if err != nil {
				t.Errorf("TrendingPools() error = %v", err)
			}
			if pools == nil {
				t.Errorf("TrendingPools() got = nil, want not nil")
			}
		})
	}
}

func TestCoinsClient_TokenMetadata(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewProApiClient("xxx", &http.Client{})

			tokenMetadata, err := c.TokensClient.TokenMetadata(context.Background(), "bsc", "0xf2c88757f8d03634671208935974b60a2a28bdb3")
			if err != nil {
				t.Errorf("TrendingPools() error = %v", err)
			}
			if tokenMetadata == nil {
				t.Errorf("TrendingPools() got = nil, want not nil")
			}
		})
	}
}

func TestCoinsClient_TokenMarketData(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewProApiClient("xxx", &http.Client{})

			tokenMetadata, err := c.TokensClient.TokensMarketData(context.Background(), "bsc", "0xf2c88757f8d03634671208935974b60a2a28bdb3")
			if err != nil {
				t.Errorf("TrendingPools() error = %v", err)
			}
			if tokenMetadata == nil {
				t.Errorf("TrendingPools() got = nil, want not nil")
			}
		})
	}
}

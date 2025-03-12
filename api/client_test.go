package api

import (
	"context"
	"net/http"
	"testing"

	"github.com/JulianToledano/goingecko/v3/api/onchaindex/simple"
	sim "github.com/JulianToledano/goingecko/v3/api/simple"
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

func TestCoinsClient_SimpleTokenPrice(t *testing.T) {
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

			tokenMetadata, err := c.TokenSimpleClient.SimplePrice(context.Background(), "bsc", "0x072a55c520a28947336ab5fc080abee38bc541ee",
				true, simple.WithIncludeDayChangeOption(true),
				simple.WithIncludeDayVolumeOption(true))
			if err != nil {
				t.Errorf("TrendingPools() error = %v", err)
			}
			if tokenMetadata == nil {
				t.Errorf("TrendingPools() got = nil, want not nil")
			}
		})
	}
}

func TestCoinsClient_CoinPriceByTokenAddresses(t *testing.T) {
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
			platform := "binance-smart-chain"
			tokenMetadata, err := c.SimpleClient.SimpleTokenPrice(context.Background(), platform, "0xf8a0bf9cf54bb92f17374d9e9a321e6a111a51bd,0x25d887ce7a35172c62febfd67a1856f20faebb00,0xf2c88757f8d03634671208935974b60a2a28bdb3",
				"usd", sim.WithIncludeDayChangeOption(true))
			if err != nil {
				t.Errorf("TrendingPools() error = %v", err)
			}
			if tokenMetadata == nil {
				t.Errorf("TrendingPools() got = nil, want not nil")
			}
		})
	}
}

func TestCoinsClient_CoinPriceByContractAddress(t *testing.T) {
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
			platform := "binance-smart-chain"
			tokenMetadata, err := c.ContractClient.ContractInfo(context.Background(), platform, "0xfb6115445bff7b52feb98650c87f44907e58f802")
			if err != nil {
				t.Errorf("TrendingPools() error = %v", err)
			}
			if tokenMetadata == nil {
				t.Errorf("TrendingPools() got = nil, want not nil")
			}
		})
	}
}

package pools

import (
	"context"
	"net/http"
	"testing"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

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
			c := &PoolsClient{
				internal.NewClient(
					geckohttp.NewClient(geckohttp.WithHttpClient(http.DefaultClient)),
					internal.BaseURL,
				),
			}

			pools, err := c.TrendingPools(context.Background(), "bsc", "6h", 80)
			if err != nil {
				t.Errorf("TrendingPools() error = %v", err)
			}
			if pools == nil {
				t.Errorf("TrendingPools() got = nil, want not nil")
			}
		})
	}
}

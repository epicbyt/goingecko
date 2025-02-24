package assetPlatforms

import (
	"context"
	"net/http"
	"testing"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

func TestAssetPlatformsClient_AssetPlatforms(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AssetPlatformsClient{
				internal.NewClient(
					geckohttp.NewClient(geckohttp.WithHttpClient(http.DefaultClient)),
					internal.BaseURL,
				),
			}
			got, err := c.AssetPlatforms(context.Background())
			if err != nil {
				t.Errorf("AssetPlatforms() error = %v", err)
			}
			if got == nil {
				t.Errorf("AssetPlatforms() got = nil, want not nil")
			}
		})
	}
}

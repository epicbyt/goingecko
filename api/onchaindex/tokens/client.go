package tokens

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type TokensClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *TokensClient {
	return &TokensClient{
		internal.NewClient(c, url),
	}
}

func (c *TokensClient) tokensUrl() string {
	return c.URL + "/onchain/networks"
}

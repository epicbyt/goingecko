package simple

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type TokenSimpleClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *TokenSimpleClient {
	return &TokenSimpleClient{
		internal.NewClient(c, url),
	}
}

func (c *TokenSimpleClient) simpleUrl() string {
	return c.URL + "/onchain/simple"
}

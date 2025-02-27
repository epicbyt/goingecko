package pools

import (
	"github.com/JulianToledano/goingecko/v3/api/internal"
	geckohttp "github.com/JulianToledano/goingecko/v3/http"
)

type PoolsClient struct {
	*internal.Client
}

func NewClient(c *geckohttp.Client, url string) *PoolsClient {
	return &PoolsClient{
		internal.NewClient(c, url),
	}
}

func (c *PoolsClient) poolsUrl() string {
	return c.URL + "/onchain/networks"
}

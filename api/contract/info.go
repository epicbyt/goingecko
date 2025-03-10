package contract

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/JulianToledano/goingecko/v3/api/contract/types"
)

func (c *ContractClient) ContractInfo(ctx context.Context, id, contractAddress string) (*types.ContractAddressInfo, error) {
	rUrl := fmt.Sprintf("%s/%s/%s/%s", c.contractUrl(), id, "contract", contractAddress)
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}
	//{"error":"coin not found"}
	if strings.Contains(string(resp), "coin not found") {
		return nil, errors.New("coin not found")
	}
	var data *types.ContractAddressInfo
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

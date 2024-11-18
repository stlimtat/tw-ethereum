package ethereum

import (
	"context"
	"crypto/rand"
	"encoding/json"

	"github.com/rs/zerolog"
	"github.com/stlimtat/tw-ethereum/internal/utils"
)

const (
	METHOD_GETBALANCE = "eth_getBalance"
)

type GetBalanceRequest struct {
	EthereumRequest
	Params []string `json:"params"`
}

type GetBalanceResp struct {
	EthereumResponse
	Result string `json:"result"`
}

func (p *DefaultParser) GetBalance(ctx context.Context, address string) (int64, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str(METHOD, METHOD_GETBALANCE).
		Logger()
	b := make([]byte, 16)
	randInt, _ := rand.Read(b)

	gbReq := GetBalanceRequest{
		EthereumRequest: EthereumRequest{
			Id:      randInt,
			JSONRPC: DEFAULT_JSON_RPC,
			Method:  METHOD_GETBALANCE,
		},
		Params: []string{
			address,
			"latest",
		},
	}
	body, err := utils.DoPost(ctx, gbReq, ETHEREUM_URL)
	if err != nil {
		return 0, err
	}

	var gbResp GetBalanceResp
	err = json.Unmarshal(body, &gbResp)
	if err != nil {
		logger.Error().Err(err).Msg("json.Unmarshal")
		return -1, err
	}
	result := utils.HexToInt64(ctx, gbResp.Result)
	return result, nil
}

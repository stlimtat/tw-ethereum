package ethereum

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog"
	"github.com/stlimtat/tw-ethereum/internal/utils"
)

const (
	METHOD_GETLOG = "eth_getLogs"
)

type GetTransactionsRequest struct {
	EthereumRequest
	Params []GetLogRequestParam `json:"params"`
}

type GetLogRequestParam struct {
	Address string `json:"address"`
}

type GetTransactionsResponse struct {
	EthereumResponse
	Result []Transaction `json:"result"`
}

func (p *DefaultParser) GetTransactions(ctx context.Context, address string) ([]Transaction, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str(METHOD, METHOD_GETLOG).
		Str("address", address).
		Logger()

	gtReq := GetTransactionsRequest{
		JSONRPC: DEFAULT_JSON_RPC,
		Method:  METHOD_GETLOG,
		Params: []GetLogRequestParam{
			{
				Address: address,
			},
		},
	}
	body, err := utils.DoPost(ctx, gtReq, ETHEREUM_URL)
	if err != nil {
		return nil, err
	}
	var gtResp GetTransactionsResponse
	err = json.Unmarshal(body, &gtResp)
	if err != nil {
		logger.Error().Err(err).Msg("json.Unmarshal")
		return nil, err
	}
	return gtResp.Result, nil
}

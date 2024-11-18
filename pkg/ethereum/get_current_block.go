package ethereum

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog"
	"github.com/stlimtat/tw-ethereum/internal/utils"
)

const (
	METHOD_BLOCKNUMBER = "eth_blockNumber"
)

type GetCurrentBlockRequest struct {
	EthereumRequest
	Params []string `json:"params"`
}

type GetCurrentBlockResp struct {
	EthereumResponse
	Result string `json:"result"`
}

func (p *DefaultParser) GetCurrentBlock(ctx context.Context) int {
	logger := zerolog.Ctx(ctx).
		With().
		Str(METHOD, METHOD_BLOCKNUMBER).
		Logger()

	gcbReq := GetCurrentBlockRequest{
		JSONRPC: DEFAULT_JSON_RPC,
		Method:  METHOD_BLOCKNUMBER,
		Params:  []string{},
	}
	body, err := utils.DoPost(ctx, gcbReq, ETHEREUM_URL)
	if err != nil {
		return -1
	}

	var gcbResp GetCurrentBlockResp
	err = json.Unmarshal(body, &gcbResp)
	if err != nil {
		logger.Error().Err(err).Msg("json.Unmarshal")
		return -1
	}
	result := utils.HexToInt(gcbResp.Result)
	return result
}

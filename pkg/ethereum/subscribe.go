package ethereum

import (
	"context"
	"crypto/rand"
	"encoding/json"

	"github.com/rs/zerolog"
	"github.com/stlimtat/tw-ethereum/internal/utils"
)

const (
	METHOD_GETFILTERLOGS = "eth_getFilterLogs"
	METHOD_NEWFILTER     = "eth_newFilter"
)

type SubscribeRequest struct {
	EthereumRequest
	Params []SubscribeRequestParam `json:"params"`
}

type SubscribeRequestParam struct {
	Address string `json:"address"`
}

type SubscribeResponse struct {
	EthereumResponse
	Result string `json:"result"`
}

func (p *DefaultParser) Subscribe(ctx context.Context, address string) (bool, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str("address", address).
		Logger()

	b := make([]byte, 16)
	randInt, _ := rand.Read(b)

	// 1. if address is in p.subscribedAddr, then run eth_getFilterChanges
	// 2. else run eth_getLogs

	sReq := SubscribeRequest{
		EthereumRequest: EthereumRequest{
			Id:      randInt,
			JSONRPC: DEFAULT_JSON_RPC,
			Method:  METHOD_NEWFILTER,
		},
		Params: []SubscribeRequestParam{
			{
				Address: address,
			},
		},
	}
	body, err := utils.DoPost(ctx, sReq, ETHEREUM_URL)
	if err != nil {
		return false, err
	}
	var sResp SubscribeResponse
	err = json.Unmarshal(body, &sResp)
	if err != nil {
		logger.Error().Err(err).Msg("json.Unmarshal")
		return false, err
	}

	// 2. Add address to DefaultParser.subscribedAddr
	p.subscribedAddr[address] = sResp.Result
	return true, nil
}

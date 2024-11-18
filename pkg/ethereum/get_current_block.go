package ethereum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

	getCurrentBlockReq := GetCurrentBlockRequest{
		JSONRPC: DEFAULT_JSON_RPC,
		Method:  METHOD_BLOCKNUMBER,
		Params:  []string{},
	}
	jsonBytes, err := json.Marshal(getCurrentBlockReq)
	byteReader := bytes.NewReader(jsonBytes)

	req, err := http.NewRequest(
		http.MethodPost,
		ETHEREUM_URL,
		byteReader,
	)
	if err != nil {
		logger.Error().Err(err).Msg("NewRequest")
		return -1
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error().Err(err).Msg("DefaultClient.Do")
		return -1
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		err = fmt.Errorf("resp.StatusCode not 200")
		logger.Error().
			Err(err).
			Int("resp_status", resp.StatusCode).
			Msg("DefaultClient.Do")
		return -1
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error().Err(err).Msg("io.ReadAll")
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

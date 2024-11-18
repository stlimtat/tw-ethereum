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

	gbReq := GetBalanceRequest{
		JSONRPC: DEFAULT_JSON_RPC,
		Method:  METHOD_GETBALANCE,
		Params: []string{
			address,
			"latest",
		},
	}
	jsonBytes, err := json.Marshal(gbReq)
	byteReader := bytes.NewReader(jsonBytes)

	req, err := http.NewRequest(
		http.MethodPost,
		ETHEREUM_URL,
		byteReader,
	)
	if err != nil {
		logger.Error().Err(err).Msg("NewRequest")
		return -1, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error().Err(err).Msg("DefaultClient.Do")
		return -1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		err = fmt.Errorf("resp.StatusCode not 200")
		logger.Error().
			Err(err).
			Int("resp_status", resp.StatusCode).
			Msg("DefaultClient.Do")
		return -1, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error().Err(err).Msg("io.ReadAll")
		return -1, err
	}
	var gbResp GetBalanceResp
	err = json.Unmarshal(body, &gbResp)
	if err != nil {
		logger.Error().Err(err).Msg("json.Unmarshal")
		return -1, err
	}
	result := utils.HexToInt64(gbResp.Result)
	return result, nil
}

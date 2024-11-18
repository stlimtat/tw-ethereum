package ethereum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"
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
		Logger()

	gtReq := GetTransactionsRequest{
		JSONRPC: DEFAULT_JSON_RPC,
		Method:  METHOD_GETLOG,
	}
	jsonBytes, err := json.Marshal(gtReq)
	byteReader := bytes.NewReader(jsonBytes)

	req, err := http.NewRequest(
		http.MethodPost,
		ETHEREUM_URL,
		byteReader,
	)
	if err != nil {
		logger.Error().Err(err).Msg("NewRequest")
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error().Err(err).Msg("DefaultClient.Do")
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		err = fmt.Errorf("resp.StatusCode not 200")
		logger.Error().
			Err(err).
			Int("resp_status", resp.StatusCode).
			Msg("DefaultClient.Do")
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error().Err(err).Msg("io.ReadAll")
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

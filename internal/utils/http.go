package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

func DoPost(
	ctx context.Context,
	reqStruct any,
	url string,
) ([]byte, error) {
	logger := zerolog.Ctx(ctx).With().Interface("req", reqStruct).Logger()

	jsonBytes, err := json.Marshal(reqStruct)
	if err != nil {
		logger.Error().Err(err).Msg("json.Marshal")
		return nil, err
	}
	logger.Info().Msg("jsonMarshal")
	byteReader := bytes.NewReader(jsonBytes)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		byteReader,
	)
	if err != nil {
		logger.Error().Err(err).Msg("NewRequest")
		return nil, err
	}
	logger.Info().Msg("NewRequestWithContext")
	client := http.DefaultClient
	logger.Info().Msg("defaultClient")
	resp, err := client.Do(req)
	logger.Info().Msg("client.do")
	if err != nil {
		logger.Error().Err(err).Msg("DefaultClient.Do")
		return nil, err
	}
	logger.Info().Int("resp_status", resp.StatusCode).Msg("http.DefaultClient.Do")
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		err = fmt.Errorf("resp.StatusCode not 2xx")
		logger.Error().
			Err(err).
			Int("resp_status", resp.StatusCode).
			Msg("DefaultClient.Do")
		return nil, err
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error().Err(err).Msg("io.ReadAll")
		return nil, err
	}
	logger.Info().Bytes("result", result).Msg("result.body")
	return result, nil
}

package utils

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/rs/zerolog"
)

func HexToInt(ctx context.Context, inHex string) int {
	result64 := HexToInt64(ctx, inHex)
	result := int(result64)
	return result
}

func HexToInt64(ctx context.Context, inHex string) int64 {
	logger := zerolog.Ctx(ctx).With().Str("inHex", inHex).Logger()
	if strings.HasPrefix(inHex, "0x") {
		inHex = inHex[2:]
	}
	bi := new(big.Int)
	bi, ok := bi.SetString(inHex, 16)
	if !ok {
		err := fmt.Errorf("bi.SetString error")
		logger.Error().
			Err(err).
			Msg("bi.SetString")
	}
	result := bi.Int64()
	return result
}

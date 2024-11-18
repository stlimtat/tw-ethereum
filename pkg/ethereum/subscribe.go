package ethereum

import (
	"context"

	"github.com/rs/zerolog"
)

func (p *DefaultParser) Subscribe(ctx context.Context, address string) bool {
	logger := zerolog.Ctx(ctx).
		With().
		Str("address", address).
		Logger()
	// 1. verify that this is valid address using GetBalance
	_, err := p.GetBalance(ctx, address)
	if err != nil {
		logger.Error().Err(err).Msg("GetBalance")
		return false
	}
	// 2. Add address to DefaultParser.subscribedAddr
	p.subscribedAddr = append(p.subscribedAddr, address)
	return true
}

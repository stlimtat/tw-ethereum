package ethereum

import (
	"context"

	"github.com/rs/zerolog"
)

const (
	DEFAULT_JSON_RPC string = "2.0"
	ETHEREUM_URL     string = "https://ethereum-rpc.publicnode.com"
	METHOD                  = "method"
)

//go:generate mockgen -destination=parser_mock.go -package=ethereum . Parser
type Parser interface {
	// last parsed block
	GetCurrentBlock(ctx context.Context) int
	// add address to observer
	Subscribe(ctx context.Context, address string) bool
	// gets balance of an address
	GetBalance(ctx context.Context, address string) (int64, error)
	// list of inbound or outbound transactions for an address
	GetTransactions(ctx context.Context, address string) []Transaction
}

type EthereumRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
}

type EthereumResponse struct {
	JSONRPC string `json:"jsonrpc"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

type DefaultParser struct {
	subscribedAddr []string
}

func NewDefaultParser(
	ctx context.Context,
) *DefaultParser {
	logger := zerolog.Ctx(ctx)
	result := &DefaultParser{
		subscribedAddr: []string{},
	}
	logger.Debug().Msg("NewDefaultParser")
	return result
}

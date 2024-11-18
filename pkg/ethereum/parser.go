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

//go:generate mockgen -destination=parser_mock.go -package=ethereum . IParser
type IParser interface {
	// last parsed block
	GetCurrentBlock(ctx context.Context) int
	// add address to observer
	Subscribe(ctx context.Context, address string) bool
	// gets balance of an address
	GetBalance(ctx context.Context, address string) (int64, error)
	// list of inbound or outbound transactions for an address
	GetTransactions(ctx context.Context, address string) ([]Transaction, error)
}

type EthereumRequest struct {
	Id      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
}

type EthereumResponse struct {
	Id      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from,omitempty"`
	Gas              string `json:"gas,omitempty"`
	GasPrice         string `json:"gasPrice,omitempty"`
	Hash             string `json:"hash"`
	Input            string `json:"input,omitempty"`
	Nonce            string `json:"nonce,omitempty"`
	To               string `json:"to,omitempty"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value,omitempty"`
	V                string `json:"v,omitempty"`
	R                string `json:"r,omitempty"`
	S                string `json:"s,omitempty"`
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

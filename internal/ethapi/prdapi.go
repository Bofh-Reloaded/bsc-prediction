package ethapi

import (
	"context"
	"github.com/ethereum/go-ethereum/eth"
)

func (s *PublicBlockChainAPI) PredictBlock(ctx context.Context) (interface{}, error) {
	casted := s.b.(*eth.EthAPIBackend)
	block, _, err := casted.Miner().PredictBlock()
	if block != nil && err == nil {
		return s.rpcMarshalBlock(ctx, block, true, true)
	}
	return nil, err
}

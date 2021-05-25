package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
)

func (s *PublicEthereumAPI) PredictBlock(ctx context.Context) (interface{}, error) {
	block, _, err := s.e.Miner().PredictBlock()
	if block != nil && err == nil {
		return s.rpcMarshalBlock(ctx, block, true, true)
	}
	return nil, err
}

func (s *PublicEthereumAPI) rpcMarshalBlock(ctx context.Context, b *types.Block, inclTx bool, fullTx bool) (map[string]interface{}, error) {
	fields, err := ethapi.RPCMarshalBlock(b, inclTx, fullTx)
	if err != nil {
		return nil, err
	}
	if inclTx {
		fields["totalDifficulty"] = (*hexutil.Big)(s.e.blockchain.GetTdByHash(b.Hash()))
	}
	return fields, err
}

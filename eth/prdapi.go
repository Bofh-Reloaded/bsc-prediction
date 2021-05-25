package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
)

func (s *PublicEthereumAPI) PredictBlock(ctx context.Context) (interface{}, interface{}, error) {
	block, receipts, err := s.e.Miner().PredictBlock()
	if block != nil && err == nil {
		var logs []*types.Log
		for _, receipt := range receipts {
			for _, log := range receipt.Logs {
				log.BlockHash = block.Hash()
			}
			logs = append(logs, receipt.Logs...)
		}

		b, berr := s.rpcMarshalBlock(ctx, block, true, true)
		return b, logs, berr
	}
	return nil, nil, err
}

/*func (s *PublicEthereumAPI) PredictApplyBlock(ctx context.Context) (interface{}, error) {
	block, _, err := s.e.Miner().PredictBlock()
	if err != nil {
		return nil, err
	}

	// Apply simulating the block came from the "canonical" chain
	blocks := make([]*types.Block, 1)
	blocks[0] = block
	_, berr := s.e.blockchain.InsertChain(blocks)
	if berr != nil {
		return nil, berr
	}
	return s.rpcMarshalBlock(ctx, block, true, true)
}*/

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

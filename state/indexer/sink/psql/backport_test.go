package psql

import (
	"github.com/leftecnu/tendermint/state/indexer"
	"github.com/leftecnu/tendermint/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)

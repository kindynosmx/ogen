package primitives_test

import (
	"github.com/kindynosmx/ogen/pkg/primitives"
	"github.com/magiconair/properties/assert"
	"testing"
	"time"

	"github.com/kindynosmx/ogen/pkg/chainhash"
)

func TestGenesisBlock(t *testing.T) {
	eb := primitives.GetGenesisBlock()

	b := primitives.Block{
		Header: &primitives.BlockHeader{
			Version:                    0,
			TxsMerkleRoot:              chainhash.Hash{},
			VoteMerkleRoot:             chainhash.Hash{},
			DepositMerkleRoot:          chainhash.Hash{},
			ExitMerkleRoot:             chainhash.Hash{},
			VoteSlashingMerkleRoot:     chainhash.Hash{},
			RANDAOSlashingMerkleRoot:   chainhash.Hash{},
			ProposerSlashingMerkleRoot: chainhash.Hash{},
			GovernanceVotesMerkleRoot:  chainhash.Hash{},
			PrevBlockHash:              chainhash.Hash{},
			Timestamp:                  uint64(time.Unix(0x0, 0).Unix()),
			Slot:                       0,
			FeeAddress:                 [20]byte{},
		},
		Txs: []*primitives.Tx{},
	}

	assert.Equal(t, b, eb)
}

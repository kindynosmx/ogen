package chain

import (
	"github.com/kindynosmx/ogen/internal/chainindex"
	"github.com/kindynosmx/ogen/internal/state"
	"github.com/kindynosmx/ogen/pkg/primitives"
)

// BlockchainNotifee is a type that is notified when something changes with the blockchain.
type BlockchainNotifee interface {
	// NewTip notifies of a new tip added to the blockchain. Do not mutate state.
	NewTip(*chainindex.BlockRow, *primitives.Block, state.State, []*primitives.EpochReceipt)
	ProposerSlashingConditionViolated(slashing *primitives.ProposerSlashing)
}

// Notify registers a notifee to be notified.
func (ch *blockchain) Notify(n BlockchainNotifee) {
	ch.notifeeLock.Lock()
	defer ch.notifeeLock.Unlock()

	ch.notifees[n] = struct{}{}
}

// Unnotify unregisters a notifee to be notified.
func (ch *blockchain) Unnotify(n BlockchainNotifee) {
	ch.notifeeLock.Lock()
	defer ch.notifeeLock.Unlock()

	delete(ch.notifees, n)
}

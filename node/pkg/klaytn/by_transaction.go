package klaytn

import (
	"context"
	"fmt"
	"github.com/certusone/wormhole/node/pkg/common"
	"github.com/certusone/wormhole/node/pkg/klaytn/kabi"
	"github.com/certusone/wormhole/node/pkg/vaa"
	klayClient "github.com/klaytn/klaytn/client"
	klay_common "github.com/klaytn/klaytn/common"
	"time"
)

var (
	// SECURITY: Hardcoded ABI identifier for the LogMessagePublished topic. When using the watcher, we don't need this
	// since the node will only hand us pre-filtered events. In this case, we need to manually verify it
	// since ParseLogMessagePublished will only verify whether it parses.
	logMessagePublishedTopic = klay_common.HexToHash("0x6eb224fb001ed210e379b335e35efe88672a8ce935d981a6896b27ffdf52a3b2")
)

// MessageEventsForTransaction returns the lockup events for a given transaction.
// Returns the block number and a list of MessagePublication events.
func MessageEventsForTransaction(
	ctx context.Context,
	c *klayClient.Client,
	contract klay_common.Address,
	chainId vaa.ChainID,
	tx klay_common.Hash) ([]*common.MessagePublication, error) {

	f, err := kabi.NewKabiFilterer(contract, c)
	if err != nil {
		return nil, fmt.Errorf("failed to create ABI filterer: %w", err)
	}

	// Get transactions logs from transaction
	receipt, err := c.TransactionReceipt(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction receipt: %w", err)
	}

	// Get block
	//block, err := c.BlockByHash(ctx, receipt)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get block: %w", err)
	//}

	msgs := make([]*common.MessagePublication, 0, len(receipt.Logs))

	// Extract logs
	for _, l := range receipt.Logs {
		// SECURITY: Skip logs not produced by our contract.
		if l.Address != contract {
			continue
		}

		if l == nil {
			continue
		}

		if l.Topics[0] != logMessagePublishedTopic {
			continue
		}

		ev, err := f.ParseLogMessagePublished(*l)
		if err != nil {
			return nil, fmt.Errorf("failed to parse log: %w", err)
		}
		hash, err := KlayHashToEthHash(ev.Raw.TxHash)
		if err != nil {
			continue
		}
		message := &common.MessagePublication{
			TxHash:           hash,
			Timestamp:        time.Unix(0, 0),
			Nonce:            ev.Nonce,
			Sequence:         ev.Sequence,
			EmitterChain:     chainId,
			EmitterAddress:   PadAddress(ev.Sender),
			Payload:          ev.Payload,
			ConsistencyLevel: ev.ConsistencyLevel,
		}

		msgs = append(msgs, message)
	}
	return msgs, nil
}

package klaytn

import (
	"context"
	"fmt"
	"github.com/certusone/wormhole/node/pkg/common"
	"github.com/certusone/wormhole/node/pkg/klaytn/kabi"
	"github.com/certusone/wormhole/node/pkg/p2p"
	gossipv1 "github.com/certusone/wormhole/node/pkg/proto/gossip/v1"
	"github.com/certusone/wormhole/node/pkg/readiness"
	"github.com/certusone/wormhole/node/pkg/supervisor"
	"github.com/certusone/wormhole/node/pkg/vaa"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	klayClient "github.com/klaytn/klaytn/client"
	klay_common "github.com/klaytn/klaytn/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
	"math/big"
	"sync/atomic"
	"time"
)

var (
	klaytnConnectionErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "wormhole_klaytn_connection_errors_total",
			Help: "Total number of Klaytn connection errors (either during initial connection or while watching)",
		}, []string{"klaytn_network", "reason"})

	klaytnMessagesObserved = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "wormhole_klaytn_messages_observed_total",
			Help: "Total number of Klaytn messages observed (pre-confirmation)",
		}, []string{"klaytn_network"})
	klaytnMessagesConfirmed = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "wormhole_klaytn_messages_confirmed_total",
			Help: "Total number of Klaytn messages verified (post-confirmation)",
		}, []string{"klaytn_network"})
	currentKlaytnHeight = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "wormhole_klaytn_current_height",
			Help: "Current klaytn block height",
		}, []string{"klaytn_network"})
	queryLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "wormhole_klaytn_query_latency",
			Help: "Latency histogram for Klaytn calls (note that most interactions are streaming queries, NOT calls, and we cannot measure latency for those",
		}, []string{"klaytn_network", "operation"})
)

type (
	Watcher struct {
		// Klaytn RPC url
		url string
		// Address of the Klay contract contract
		contract klay_common.Address
		// Human-readable name of the klay network, for logging and monitoring.
		networkName string
		// Readiness component
		readiness readiness.Component
		// VAA ChainID of the network we're connecting to.
		chainID vaa.ChainID

		// Channel to send new messages to.
		msgChan chan *common.MessagePublication
		setChan chan *common.GuardianSet

		// Incoming re-observation requests from the network. Pre-filtered to only
		// include requests for our chainID.
		obsvReqC chan *gossipv1.ObservationRequest
		// 0 is a valid guardian set, so we need a nil value here
		currentGuardianSet *uint32
	}
)

func NewKlaytnWatcher(
	url string,
	contract klay_common.Address,
	networkName string,
	readiness readiness.Component,
	chainID vaa.ChainID,
	messageEvents chan *common.MessagePublication,
	setEvents chan *common.GuardianSet,
	obsvReqC chan *gossipv1.ObservationRequest,
) *Watcher {
	return &Watcher{
		url:         url,
		contract:    contract,
		networkName: networkName,
		readiness:   readiness,
		chainID:     chainID,
		msgChan:     messageEvents,
		obsvReqC:    obsvReqC,
		setChan:     setEvents}
}

func (e *Watcher) Run(ctx context.Context) error {
	logger := supervisor.Logger(ctx)

	// Initialize gossip metrics (we want to broadcast the address even if we're not yet syncing)
	p2p.DefaultRegistry.SetNetworkStats(e.chainID, &gossipv1.Heartbeat_Network{
		ContractAddress: e.contract.Hex(),
	})
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c, err := klayClient.DialContext(timeout, e.url)

	if err != nil {
		klaytnConnectionErrors.WithLabelValues(e.networkName, "dial_error").Inc()
		p2p.DefaultRegistry.AddErrorCount(e.chainID, 1)
		return fmt.Errorf("dialing klay client failed: %w", err)
	}
	f, err := kabi.NewKabiFilterer(e.contract, c)
	if err != nil {
		return fmt.Errorf("could not create wormhole contract filter: %w", err)
	}
	timeout, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	messageC := make(chan *kabi.KabiLogMessagePublished)
	messageSub, err := f.WatchLogMessagePublished(&bind.WatchOpts{Context: timeout}, messageC, nil)
	if err != nil {
		klaytnConnectionErrors.WithLabelValues(e.networkName, "subscribe_error").Inc()
		p2p.DefaultRegistry.AddErrorCount(e.chainID, 1)
		return fmt.Errorf("failed to subscribe to message publication events: %w", err)
	}

	errC := make(chan error)
	var currentBlockNumber uint64

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case r := <-e.obsvReqC:
				// This can't happen unless there is a programming error - the caller
				// is expected to send us only requests for our chainID.
				if vaa.ChainID(r.ChainId) != e.chainID {
					panic("invalid chain ID")
				}

				tx := klay_common.BytesToHash(r.TxHash)
				logger.Info("received observation request",
					zap.String("eth_network", e.networkName),
					zap.String("tx_hash", tx.Hex()))

				// SECURITY: Load the block number before requesting the transaction to avoid a
				// race condition where requesting the tx succeeds and is then dropped due to a fork,
				// but blockNumberU had already advanced beyond the required threshold.
				//
				// In the primary watcher flow, this is of no concern since we assume the node
				// always sends the head before it sends the logs (implicit synchronization
				// by relying on the same websocket connection).
				blockNumberU := atomic.LoadUint64(&currentBlockNumber)
				if blockNumberU == 0 {
					logger.Error("no block number available, ignoring observation request",
						zap.String("kalytn_network", e.networkName))
					continue
				}
				timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
				msgs, err := MessageEventsForTransaction(timeout, c, e.contract, e.chainID, tx)
				cancel()

				if err != nil {
					logger.Error("failed to process observation request",
						zap.Error(err), zap.String("eth_network", e.networkName))
					continue
				}

				for _, msg := range msgs {
					logger.Info("re-observed message publication transaction",
						zap.Stringer("tx", msg.TxHash),
						zap.Stringer("emitter_address", msg.EmitterAddress),
						zap.Uint64("sequence", msg.Sequence),
						zap.Uint64("current_block", blockNumberU),
						zap.Uint64("observed_block", 0),
						zap.String("eth_network", e.networkName),
					)
					e.msgChan <- msg
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-messageSub.Err():
				klaytnConnectionErrors.WithLabelValues(e.networkName, "subscription_error").Inc()
				errC <- fmt.Errorf("error while processing message publication subscription: %w", err)
				p2p.DefaultRegistry.AddErrorCount(e.chainID, 1)
				return
			case ev := <-messageC:
				// Request timestamp for block
				msm := time.Now()

				timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
				b, err := c.BlockByNumber(timeout, big.NewInt(int64(ev.Raw.BlockNumber)))
				cancel()
				queryLatency.WithLabelValues(e.networkName, "block_by_number").Observe(time.Since(msm).Seconds())

				if err != nil {
					klaytnConnectionErrors.WithLabelValues(e.networkName, "block_by_number_error").Inc()
					p2p.DefaultRegistry.AddErrorCount(e.chainID, 1)
					errC <- fmt.Errorf("failed to request timestamp for block %d: %w", ev.Raw.BlockNumber, err)
					return
				}
				logger.Info("Hashhhhhhhhhhhhhh", zap.String("Klaytn", ev.Raw.TxHash.String()))
				hash, err := KlayHashToEthHash(ev.Raw.TxHash)
				if err != nil {
					logger.Info("Conversion from klay hash to eth hash failed", zap.Error(err))
					continue
				}
				message := &common.MessagePublication{
					TxHash:           hash,
					Timestamp:        time.Unix(b.Time().Int64(), 0),
					Nonce:            ev.Nonce,
					Sequence:         ev.Sequence,
					EmitterChain:     e.chainID,
					EmitterAddress:   PadAddress(ev.Sender),
					Payload:          ev.Payload,
					ConsistencyLevel: ev.ConsistencyLevel,
				}

				logger.Info("found new message publication transaction", zap.Stringer("tx", ev.Raw.TxHash),
					zap.Uint64("block", ev.Raw.BlockNumber), zap.String("klay_network", e.networkName))

				klaytnMessagesObserved.WithLabelValues(e.networkName).Inc()
				e.msgChan <- message
			}
		}
	}()

	// Watch headers
	headSink := make(chan *types.Header, 2)
	headerSubscription, err := c.SubscribeNewHead(ctx, headSink)
	if err != nil {
		klaytnConnectionErrors.WithLabelValues(e.networkName, "header_subscribe_error").Inc()
		p2p.DefaultRegistry.AddErrorCount(e.chainID, 1)
		return fmt.Errorf("failed to subscribe to header events: %w", err)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-headerSubscription.Err():
				klaytnConnectionErrors.WithLabelValues(e.networkName, "header_subscription_error").Inc()
				errC <- fmt.Errorf("error while processing header subscription: %w", err)
				p2p.DefaultRegistry.AddErrorCount(e.chainID, 1)
				return
			case ev := <-headSink:
				//start := time.Now()
				blockNumberU := ev.Number.Uint64()
				atomic.StoreUint64(&currentBlockNumber, blockNumberU)
				currentHash := ev.Hash()
				logger.Debug("processing new header",
					zap.Stringer("current_block", ev.Number),
					zap.Stringer("current_blockhash", currentHash),
					zap.String("klay_network", e.networkName))
				currentKlaytnHeight.WithLabelValues(e.networkName).Set(float64(ev.Number.Int64()))
				readiness.SetReady(e.readiness)
				p2p.DefaultRegistry.SetNetworkStats(e.chainID, &gossipv1.Heartbeat_Network{
					Height:          ev.Number.Int64(),
					ContractAddress: e.contract.Hex(),
				})
				//logger.Info("processed new header",
				//	zap.Stringer("current_block", ev.Number),
				//	zap.Stringer("current_blockhash", currentHash),
				//	zap.Duration("took", time.Since(start)),
				//	zap.String("klay_network", e.networkName))
			}
		}
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errC:
		return err
	}
}

func KlayHashToEthHash(klayHash klay_common.Hash) (eth_common.Hash, error) {
	var ethHash eth_common.Hash
	res, err := klayHash.MarshalText()
	if err != nil {
		return ethHash, err
	}
	copy(ethHash[:], res)
	return ethHash, nil
}

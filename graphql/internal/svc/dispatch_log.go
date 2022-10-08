// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// logDispatcher implements dispatcher of new log events in the blockchain.
type logDispatcher struct {
	service
	inLog       chan *types.LogRecord
	knownTopics map[common.Hash]func(*types.LogRecord)
}

// name returns the name of the service used by orchestrator.
func (lgd *logDispatcher) name() string {
	return "log dispatcher"
}

// init prepares the log dispatcher to perform its function.
func (lgd *logDispatcher) init() {
	lgd.sigStop = make(chan struct{})
	lgd.knownTopics = map[common.Hash]func(*types.LogRecord){
		/* ---------------- ERC20 and ERC721 contracts related event hooks below this line ---------------- */

		/* ERC20::Approval(address indexed owner, address indexed spender, uint256 value) */
		common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"): handleErcTokenApproval,

		/* ERC20::Transfer(address indexed from, address indexed to, uint256 value) */
		common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"): handleErcTokenTransfer,

		/* ERC1155::TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value) */
		common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"): handleErc1155TransferSingle,

		/* ERC1155::TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values) */
		common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"): handleErc1155TransferBatch,
	}
}

// run starts the transaction logs dispatcher job
func (lgd *logDispatcher) run() {
	// make sure we are orchestrated
	if lgd.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", lgd.name()))
	}

	// signal orchestrator we started and go
	lgd.mgr.started(lgd)
	go lgd.execute()
}

// execute implements the dispatcher reader and router routine.
func (lgd *logDispatcher) execute() {
	// don't forget to sign off after we are done
	defer func() {
		lgd.mgr.finished(lgd)
	}()

	// wait for logs and process them
	for {
		// try to read next transaction
		select {
		case <-lgd.sigStop:
			return
		case lr, ok := <-lgd.inLog:
			// is the channel even available for reading
			if !ok {
				log.Notice("logs channel closed, terminating %s", lgd.name())
				return
			}

			// try to find the topic handler
			if nil != lr && nil != lr.Topics && 0 < len(lr.Topics) {
				handler, ok := lgd.knownTopics[lr.Topics[0]]
				if ok && lr.Block != nil && lr.Trx != nil {
					log.Debugf("known topic %s found, processing", lr.Topics[0].String())
					handler(lr)
				}
			}

			// mark the processing of this log record as finished
			lr.WatchDog.Done()
		}
	}
}

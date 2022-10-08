/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/repository/p2p"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"math/big"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etc "github.com/ethereum/go-ethereum/core/types"
)

// Repository interface defines functions the underlying implementation provides to API resolvers.
type Repository interface {
	// Account returns account at Opera blockchain for an address, nil if not found.
	Account(*common.Address) (*types.Account, error)

	// AccountBalance returns the current balance of an account at Opera blockchain.
	AccountBalance(*common.Address) (*hexutil.Big, error)

	// AccountNonce returns the current number of sent transactions of an account at Opera blockchain.
	AccountNonce(*common.Address) (*hexutil.Uint64, error)

	// AccountTransactions returns list of transaction hashes for account at Opera blockchain.
	//
	// String cursor represents cursor based on which the list is loaded. If null,
	// it loads either from top, or bottom of the list, based on the value
	// of the integer count. The integer represents the number of transaction loaded at most.
	//
	// For positive number, the list starts right after the cursor
	// (or on top without one) and loads at most defined number of transactions older than that.
	//
	// For negative number, the list starts right before the cursor
	// (or at the bottom without one) and loads at most defined number
	// of transactions newer than that.
	//
	// Transactions are always sorted from newer to older.
	AccountTransactions(*common.Address, *common.Address, *string, int32) (*types.TransactionList, error)

	// AccountsActive total number of accounts known to repository.
	AccountsActive() (hexutil.Uint64, error)

	// AccountIsKnown checks if the account of the given address is known to the API server.
	AccountIsKnown(*common.Address) bool

	// StoreAccount adds specified account detail into the repository.
	StoreAccount(*types.Account) error

	// AccountMarkActivity marks the latest account activity in the repository.
	AccountMarkActivity(*common.Address, uint64) error

	// BlockHeight returns the current height of the Opera blockchain in blocks.
	BlockHeight() (*hexutil.Big, error)

	// LastKnownBlock returns number of the last block known to the repository.
	LastKnownBlock() (uint64, error)

	// UpdateLastKnownBlock update record about last known block.
	UpdateLastKnownBlock(blockNo *hexutil.Uint64) error

	// ObservedHeaders provides a channel fed with new headers observed
	// by the connected blockchain node.
	ObservedHeaders() chan *etc.Header

	// BlockByNumber returns a block at Opera blockchain represented by a number.
	// Top block is returned if the number is not provided.
	// If the block is not found, ErrBlockNotFound error is returned.
	BlockByNumber(*hexutil.Uint64) (*types.Block, error)

	// BlockByHash returns a block at Opera blockchain represented by a hash.
	// Top block is returned if the hash is not provided.
	// If the block is not found, ErrBlockNotFound error is returned.
	BlockByHash(*common.Hash) (*types.Block, error)

	// Blocks pulls list of blocks starting on the specified block number
	// and going up, or down based on count number.
	Blocks(*uint64, int32) (*types.BlockList, error)

	// CacheBlock puts a block to the internal block ring cache.
	CacheBlock(blk *types.Block)

	// Contract extract a smart contract information by address if available.
	Contract(*common.Address) (*types.Contract, error)

	// Contracts returns list of smart contracts at Opera blockchain.
	Contracts(bool, *string, int32) (*types.ContractList, error)

	// ValidateContract tries to validate contract byte code using
	// provided source code. If successful, the contract information
	// is updated the the repository.
	ValidateContract(*types.Contract) error

	// StoreContract updates the contract in repository.
	StoreContract(*types.Contract) error

	// StoreTransaction adds a new incoming transaction from blockchain to the repository.
	StoreTransaction(*types.Block, *types.Transaction) error

	// LoadTransaction returns a transaction at Opera blockchain
	// by a hash loaded directly from the node.
	LoadTransaction(hash *common.Hash) (*types.Transaction, error)

	// Transaction returns a transaction at Opera blockchain by a hash, nil if not found.
	Transaction(*common.Hash) (*types.Transaction, error)

	// Transactions returns list of transaction hashes at Opera blockchain.
	Transactions(*string, int32) (*types.TransactionList, error)

	// TransactionsCount returns total number of transactions in the block chain.
	TransactionsCount() (uint64, error)

	// EstimateTransactionsCount returns an approximate amount of transactions on the network.
	EstimateTransactionsCount() (hexutil.Uint64, error)

	// IncTrxCountEstimate bumps the value of transaction counter estimator.
	IncTrxCountEstimate(diff uint64)

	// UpdateTrxCountEstimate updates the value of transaction counter estimator.
	UpdateTrxCountEstimate(val uint64)

	// CacheTransaction puts a transaction to the internal ring cache.
	CacheTransaction(trx *types.Transaction)

	// SendTransaction sends raw signed and RLP encoded transaction to the block chain.
	SendTransaction(hexutil.Bytes) (*types.Transaction, error)

	// Price returns a price information for the given target symbol.
	Price(sym string) (types.Price, error)

	// GasPrice provides the raw suggested value for the gas price.
	GasPrice() (hexutil.Big, error)

	// GasPriceExtended provides extended gas price information.
	GasPriceExtended() (*types.GasPrice, error)

	// StoreGasPricePeriod stores gas price period data into the persistent storage.
	StoreGasPricePeriod(*types.GasPricePeriod) error

	// GasEstimate calculates the estimated amount of Gas required to perform
	// transaction described by the input params.
	GasEstimate(*struct {
		From  *common.Address
		To    *common.Address
		Value *hexutil.Big
		Data  *string
	}) (*hexutil.Uint64, error)

	//NativeTokenAddress returns address of the native token wrapper, if available.
	//NativeTokenAddress() (*common.Address, error)

	// TokenTransactions provides list of ERC20/ERC721/ERC1155 transactions based on given filters.
	TokenTransactions(tokenType string, token *common.Address, tokenId *big.Int, acc *common.Address, txType []int32, cursor *string, count int32) (*types.TokenTransactionList, error)

	// TokenTransactionsByCall provides a list of token transaction made inside a specific
	// transaction call (blockchain transaction).
	TokenTransactionsByCall(*common.Hash) ([]*types.TokenTransaction, error)

	// Erc20Token returns an ERC20 token for the given address, if available.
	Erc20Token(*common.Address) (*types.Erc20Token, error)

	// Erc20TokensList returns a list of known ERC20 tokens ordered by their activity.
	Erc20TokensList(int32) ([]common.Address, error)

	// Erc20Assets provides list of ERC20 tokens involved with the given owner.
	Erc20Assets(common.Address, int32) ([]common.Address, error)

	// Erc20BalanceOf load the current available balance of and ERC20 token identified by the token
	// contract address for an identified owner address.
	Erc20BalanceOf(*common.Address, *common.Address) (hexutil.Big, error)

	// Erc20Allowance loads the current amount of ERC20 tokens unlocked for DeFi
	// contract by the token owner.
	Erc20Allowance(*common.Address, *common.Address, *common.Address) (hexutil.Big, error)

	// Erc20TotalSupply provides information about all available tokens
	Erc20TotalSupply(*common.Address) (hexutil.Big, error)

	// Erc20Name provides information about the name of the ERC20 token.
	Erc20Name(*common.Address) (string, error)

	// Erc20Symbol provides information about the symbol of the ERC20 token.
	Erc20Symbol(*common.Address) (string, error)

	// Erc20Decimals provides information about the decimals of the ERC20 token.
	Erc20Decimals(*common.Address) (int32, error)

	// Erc20LogoURL provides URL address of a logo of the ERC20 token.
	Erc20LogoURL(*common.Address) string

	// StoreTokenTransaction stores ERC20/ERC721/ERC1155 transaction into the repository.
	StoreTokenTransaction(*types.TokenTransaction) error

	// Erc165SupportsInterface provides information about support of the interface by the contract.
	Erc165SupportsInterface(contract *common.Address, interfaceID [4]byte) (bool, error)

	// Erc721Contract returns an ERC721 token for the given address, if available.
	Erc721Contract(*common.Address) (*types.Erc721Contract, error)

	// Erc721ContractsList returns a list of known ERC721 tokens ordered by their activity.
	Erc721ContractsList(int32) ([]common.Address, error)

	// Erc721Name provides information about the name of the ERC721 token.
	Erc721Name(*common.Address) (string, error)

	// Erc721Symbol provides information about the symbol of the ERC721 token.
	Erc721Symbol(*common.Address) (string, error)

	// Erc721TotalSupply provides information about all available tokens.
	Erc721TotalSupply(token *common.Address) (hexutil.Big, error)

	// Erc721BalanceOf provides amount of NFT tokens owned by given owner in given ERC721 contract.
	Erc721BalanceOf(token *common.Address, owner *common.Address) (hexutil.Big, error)

	// Erc721TokenURI provides URI of Metadata JSON Schema of the ERC721 token.
	Erc721TokenURI(token *common.Address, tokenId *big.Int) (string, error)

	// Erc721OwnerOf provides information about NFT token ownership.
	Erc721OwnerOf(token *common.Address, tokenId *big.Int) (common.Address, error)

	// Erc721GetApproved provides information about operator approved to manipulate with the NFT token.
	Erc721GetApproved(token *common.Address, tokenId *big.Int) (common.Address, error)

	// Erc721IsApprovedForAll provides information about operator approved to manipulate with NFT tokens of given owner.
	Erc721IsApprovedForAll(token *common.Address, owner *common.Address, operator *common.Address) (bool, error)

	// Erc1155ContractsList returns a list of known ERC1155 contracts ordered by their activity.
	Erc1155ContractsList(int32) ([]common.Address, error)

	// Erc1155Uri provides URI of Metadata JSON Schema of the token.
	Erc1155Uri(token *common.Address, tokenId *big.Int) (string, error)

	// Erc1155BalanceOf provides amount of NFT tokens owned by given owner.
	Erc1155BalanceOf(token *common.Address, owner *common.Address, tokenId *big.Int) (*big.Int, error)

	// Erc1155BalanceOfBatch provides amount of NFT tokens owned by given owner.
	Erc1155BalanceOfBatch(token *common.Address, owners *[]common.Address, tokenIds []*big.Int) ([]*big.Int, error)

	// Erc1155IsApprovedForAll provides information about operator approved to manipulate with NFT tokens of given owner.
	Erc1155IsApprovedForAll(token *common.Address, owner *common.Address, operator *common.Address) (bool, error)

	// TrxFlowVolume resolves the list of daily trx flow aggregations.
	TrxFlowVolume(from *time.Time, to *time.Time) ([]*types.DailyTrxVolume, error)

	// TrxGasSpeed provides speed of gas consumption per second by transactions.
	TrxGasSpeed(from *time.Time, to *time.Time) (float64, error)

	// GasPriceTicks provides a list of gas price ticks for the given time period.
	GasPriceTicks(from *time.Time, to *time.Time) ([]types.GasPricePeriod, error)

	// TrxFlowUpdate executes the trx flow update in the database.
	TrxFlowUpdate()

	// TrxFlowSpeed provides speed of transaction per second for the last <sec> seconds.
	TrxFlowSpeed(sec int32) (float64, error)

	// StoreFtmBurn stores the given native FTM burn per block record into the persistent storage.
	StoreFtmBurn(burn *types.FtmBurn) error

	// FtmBurnTotal provides the total amount of burned native FTM.
	FtmBurnTotal() (int64, error)

	// FtmBurnList provides list of per-block burned native FTM tokens.
	FtmBurnList(count int64) ([]types.FtmBurn, error)

	// NetworkNode returns instance of Opera network node record by its ID.
	NetworkNode(nid enode.ID) (*types.OperaNode, error)

	// StoreNetworkNode stores the given Opera node record in the persistent database.
	StoreNetworkNode(node *types.OperaNode) error

	// IsNetworkNodeKnown checks if the given network node is already registered in the persistent database.
	IsNetworkNodeKnown(id enode.ID) bool

	// NetworkNodeConfirmCheck confirms successful check of the given Opera network node.
	NetworkNodeConfirmCheck(node *enode.Node, bhp p2p.BlockHeightProvider) (bool, error)

	// NetworkNodeFailCheck registers failed check of the given Opera network node.
	NetworkNodeFailCheck(node *enode.Node) error

	// PeerInformation returns detailed information of the given peer, if it can be obtained.
	PeerInformation(node *enode.Node, bhp p2p.BlockHeightProvider) (*types.OperaNodeInformation, error)

	// NetworkNodeUpdateBatch provides a list of Opera network node addresses most suitable for status update
	// based on the registered time of the latest check.
	NetworkNodeUpdateBatch() ([]*enode.Node, error)

	// NetworkNodeBootstrapSet provides a set of known nodes to be co-used to bootstrap new search.
	NetworkNodeBootstrapSet() []*enode.Node

	// GeoLocation provides geographic location information for the given IP address using GeoIP bridge.
	GeoLocation(net.IP) (types.GeoLocation, error)

	// NetworkNodesGeoAggregated provides a list of aggregated opera nodes based on given location detail level.
	NetworkNodesGeoAggregated(level int) ([]*types.OperaNodeLocationAggregate, error)

	// Close and cleanup the repository.
	Close()
}

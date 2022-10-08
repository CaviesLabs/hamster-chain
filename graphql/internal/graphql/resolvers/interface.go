// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"context"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ApiResolver represents the API interface expected to handle API access points
type ApiResolver interface {
	// Version resolves current version of the API server.
	Version() string

	// Account resolves blockchain account by address.
	Account(struct{ Address common.Address }) (*Account, error)

	// Contracts resolves list of blockchain smart contracts encapsulated in a listable structure.
	Contracts(*struct {
		ValidatedOnly bool
		Cursor        *Cursor
		Count         int32
	}) (*ContractList, error)

	// ValidateContract resolves smart contract source code vs. deployed byte code and marks
	// the contract as validated if the match is found. Peer API points are ringed on success
	// to notify them about the change.
	ValidateContract(*struct{ Contract ContractValidationInput }) (*Contract, error)

	// Block resolves blockchain block by number or by hash. If neither is provided, the most recent block is given.
	Block(*struct {
		Number *hexutil.Uint64
		Hash   *common.Hash
	}) (*Block, error)

	// Blocks resolves list of blockchain blocks encapsulated in a listable structure.
	Blocks(*struct {
		Cursor *Cursor
		Count  int32
	}) (*BlockList, error)

	// Transaction resolves blockchain transaction by hash.
	Transaction(*struct{ Hash common.Hash }) (*Transaction, error)

	// Transactions resolves list of blockchain transactions encapsulated in a listable structure.
	Transactions(*struct {
		Cursor *Cursor
		Count  int32
	}) (*TransactionList, error)

	// OnBlock resolves subscription to new blocks' event broadcast.
	OnBlock(ctx context.Context) <-chan *Block

	// OnTransaction resolves subscription to new transactions' event broadcast.
	OnTransaction(ctx context.Context) <-chan *Transaction

	// Price resolves price details of the Opera blockchain token for the given target symbols.
	Price(*struct{ To string }) (types.Price, error)

	// GasPrice resolves the current amount of WEI for single Gas.
	GasPrice() (hexutil.Uint64, error)

	// EstimateGas resolves the estimated amount of Gas required to perform
	// transaction described by the input params.
	EstimateGas(struct {
		From  *common.Address
		To    *common.Address
		Value *hexutil.Big
		Data  *string
	}) (*hexutil.Uint64, error)

	// SendTransaction sends raw signed and RLP encoded transaction to the blockchain.
	SendTransaction(*struct{ Tx hexutil.Bytes }) (*Transaction, error)

	// Erc20Token resolves an instance of ERC20 token if available.
	Erc20Token(*struct{ Token common.Address }) *ERC20Token

	// Erc20TokenList resolves a list of instances of ERC20 tokens.
	Erc20TokenList(struct{ Count int32 }) ([]*ERC20Token, error)

	// Erc20Assets resolves a list of instances of ERC20 tokens for the given owner.
	Erc20Assets(struct {
		Owner common.Address
		Count int32
	}) ([]*ERC20Token, error)

	// ErcTokenBalance resolves the current available balance of the specified token
	// for the specified owner.
	ErcTokenBalance(args *struct {
		Owner common.Address
		Token common.Address
	}) hexutil.Big

	// ErcTotalSupply resolves the current total supply of the specified token.
	ErcTotalSupply(args *struct{ Token common.Address }) hexutil.Big

	// ErcTokenAllowance resolves the current amount of ERC20 tokens unlocked
	// by the token owner for the spender to be manipulated with.
	ErcTokenAllowance(args *struct {
		Token   common.Address
		Owner   common.Address
		Spender common.Address
	}) hexutil.Big

	// TrxVolume resolves list of daily aggregations
	// of the network transaction flow.
	TrxVolume(args struct {
		From *string
		To   *string
	}) ([]*DailyTrxVolume, error)

	// TrxSpeed resolves the recent speed of the network in transactions processed per second.
	TrxSpeed(args struct {
		Range int32
	}) (float64, error)

	// TrxGasSpeed resolves the gas consumption speed
	// of the network in transactions processed per second.
	TrxGasSpeed(args struct {
		Range int32
		To    *string
	}) (float64, error)

	// Close terminates resolver broadcast management.
	Close()
}

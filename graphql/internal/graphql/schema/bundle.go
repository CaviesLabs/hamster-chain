package gqlschema

// Auto generated GraphQL schema bundle
const schema = `
# DailyTrxVolume represents a view of an aggregated flow
# of transactions on the network on specific day.
type DailyTrxVolume {
    # day represents the day of the aggregation in format YYYY-MM-DD
    # i.e. 2021-01-23 for January 23rd, 2021
    day: String!

    # volume represent the number of transactions originated / mined
    # by the network on the day.
    volume: Int!

    # amount represents the total value of native tokens transferred
    # by the network on the day. Please note this includes only direct
    # token transfers.
    amount: BigInt!

    # gas represents the total amount of gas consumed by transactions
    # on the network on the day.
    gas: BigInt!
}

# TokenTransactionType represents a type of ERC-20/ERC-721/ERC-1155 transaction.
enum TokenTransactionType {
    TRANSFER
    MINT
    BURN
    APPROVAL
    APPROVAL_FOR_ALL
    OTHER
}

# ERC20Transaction represents a transaction on an ERC20 token.
type ERC20Transaction {
    # trxHash represents a hash of the transaction
    # executing the ERC20 call.
    trxHash: Bytes32!

    # transaction represents the transaction
    # executing the ERC20 call.
    transaction: Transaction!

    # trxIndex represents the index
    # of the ERC20 call in the transaction logs.
    trxIndex: Long!

    # tokenAddress represents the address
    # of the ERC20 token contract.
    tokenAddress: Address!

    # token represents the token detail involved.
    token: ERC20Token!

    # trxType is the type of the transaction.
    trxType: TokenTransactionType!

    # sender represents the address of the token owner
    # sending the tokens, e.g. the sender.
    sender: Address!

    # recipient represents the address of the token recipient.
    recipient: Address!

    # amount represents the amount of tokens involved
    # in the transaction; please make sure to interpret the amount
    # with the correct number of decimals from the ERC20 token detail.
    amount: BigInt!

    # timeStamp represents the Unix epoch time stamp
    # of the ERC20 transaction processing.
    timeStamp: Long!
}
# Price represents price information of core Opera token
type Price {
    "Source unit symbol."
    fromSymbol: String!

    "Target unit symbol."
    toSymbol: String!

    "Price of the source symbol unit in target symbol unit."
    price: Float!

    "Price change in last 24h."
    change24: Float!

    "Price change in percent in last 24h."
    changePct24: Float!

    "Open 24h price."
    open24: Float!

    "Highest 24h price."
    high24: Float!

    "Lowest 24h price."
    low24: Float!

    "Volume exchanged in last 24h price."
    volume24: Float!

    "Market cap of the source unit."
    marketCap: Float!

    "Timestamp of the last update of this price value."
    lastUpdate: Long!
}

# ERC1155Contract represents a generic ERC1155 multi-token contract.
type ERC1155Contract {
    # address of the token is used as the token's unique identifier.
    address: Address!

    # uri provides URI of Metadata JSON Schema for given token.
    uri(tokenId: BigInt!): String

    # balanceOf represents amount of tokens on the account.
    balanceOf(owner: Address!, tokenId: BigInt!): BigInt!

    # balanceOf represents amount of tokens on the account.
    balanceOfBatch(owners: [Address!]!, tokenIds: [BigInt!]!): [BigInt!]!

    # isApprovedForAll queries the approval status of an operator for a given owner.
    isApprovedForAll(owner: Address!, operator: Address!): Boolean
}

# TransactionList is a list of transaction edges provided by sequential access request.
type TransactionList {
    # Edges contains provided edges of the sequential list.
    edges: [TransactionListEdge!]!

    # TotalCount is the maximum number of transactions available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of transaction edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of transactions.
type TransactionListEdge {
    cursor: Cursor!
    transaction: Transaction!
}


# BlockList is a list of block edges provided by sequential access request.
type BlockList {
    # Edges contains provided edges of the sequential list.
    edges: [BlockListEdge!]!

    # TotalCount is the maximum number of blocks available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of block edges.
    pageInfo: ListPageInfo!
}

# BlockListEdge is a single edge in a sequential list of blocks.
type BlockListEdge {
    cursor: Cursor!
    block: Block!
}

# GasPriceTick represents a collected gas price tick.
type GasPriceTick {
    # fromTime is the time of the tick measurement start
    fromTime: Time!

    # toTime is the time of the tick measurement end
    toTime: Time!

    # openPrice is the opening gas price in the tick
    openPrice: Long!

    # closePrice is the closing gas price in the tick
    closePrice: Long!

    # minPrice is the lowest reached price in the tick
    minPrice: Long!

    # maxPrice is the highest reached price in the tick
    maxPrice: Long!

    # avgPrice is the average reached price in the tick
    avgPrice: Long!
}

# ListPageInfo contains information about a sequential access list page.
type ListPageInfo {
    # First is the cursor of the first edge of the edges list. null for empty list.
    first: Cursor

    # Last if the cursor of the last edge of the edges list. null for empty list.
    last: Cursor

    # HasNext specifies if there is another edge after the last one.
    hasNext: Boolean!

    # HasNext specifies if there is another edge before the first one.
    hasPrevious: Boolean!
}
# Transaction is an Opera block chain transaction.
type Transaction {
    # Hash is the unique hash of this transaction.
    hash: Bytes32!

    # Nonce is the number of transactions sent by the account prior to this transaction.
    nonce: Long!

    # Index is the index of this transaction in the block. This will
    # be null if the transaction is in a pending pool.
    index: Long

    # From is the address of the account that sent this transaction
    from: Address!

    # Sender is the account that sent this transaction
    sender: Account!

    # To is the account the transaction was sent to.
    # This is null for contract creating transactions.
    to: Address

    # contractAddress represents the address of smart contract
    # deployed by this transaction;
    # null if the transaction is not contract creation
    contractAddress: Address

    # Recipient is the account that received this transaction.
    # Null for contract creating transaction.
    recipient: Account

    # Value is the value sent along with this transaction in WEI.
    value: BigInt!

    # GasPrice is the price of gas per unit in WEI.
    gasPrice: BigInt!

    # Gas represents gas provided by the sender.
    gas: Long!

    # GasUsed is the amount of gas that was used on processing this transaction.
    # If the transaction is pending, this field will be null.
    gasUsed: Long

    # InputData is the data supplied to the target of the transaction.
    # Contains smart contract byte code if this is contract creation.
    # Contains encoded contract state mutating function call if recipient
    # is a contract address.
    inputData: Bytes!

    # BlockHash is the hash of the block this transaction was assigned to.
    # Null if the transaction is pending.
    blockHash: Bytes32

    # BlockHash is the hash of the block this transaction was assigned to.
    # Null if the transaction is pending.
    blockNumber: Long

    # Block is the block this transaction was assigned to. This will be null if
    # the transaction is pending.
    block: Block

    # Status is the return status of the transaction. This will be 1 if the
    # transaction succeeded, or 0 if it failed (due to a revert, or due to
    # running out of gas). If the transaction has not yet been processed, this
    # field will be null.
    status: Long

    # tokenTransactions represents a list of generic token transactions executed in the scope
    # of the transaction call; token type and transaction type is provided.
    tokenTransactions: [TokenTransaction!]!

    # erc20Transactions provides list of ERC-20 token transactions executed in the scope
    # of this blockchain transaction call.
    erc20Transactions: [ERC20Transaction!]!

    # erc721Transactions provides list of ERC-721 NFT transactions executed in the scope
    # of this blockchain transaction call.
    erc721Transactions: [ERC721Transaction!]!

    # erc1155Transactions provides list of ERC-1155 NFT transactions executed in the scope
    # of this blockchain transaction call.
    erc1155Transactions: [ERC1155Transaction!]!
}

# NetworkNodeGroupLevel represents the detail of network node count aggregation.
enum NetworkNodeGroupLevel {
    CONTINENT
    COUNTRY
    STATE
}

# NetworkNodeGroup represents an aggregated group of Opera network nodes.
type NetworkNodeGroup {
    # topRegion represents the name of the top level location of the aggregation group.
    topRegion: String!

    # region represents the name of the location of the aggregation group
    # based on selected detail level.
    region: String!

    # count represents the number of nodes in the aggregation group.
    count: Int!

    # latitude represents an average geographic coordinate
    # that specifies the north–south position of the group on the Earth's surface.
    latitude: Float!

    # longitude represents an average geographic coordinate
    # that specifies the east–west position of the group on the Earth's surface.
    longitude: Float!

    # pct represents the percentage share of the aggregation group
    # compared to the number of all known active nodes. The number is provided
    # as fixed point integer with 1 decimal precision (i.e. 258 = 25.8%, 1000 = 100%)
    pct: Int!
}

# NetworkNodeGroupList represents a list of network node groups with a specified detail level.
type NetworkNodeGroupList {
    # level represents the level of detail of the aggregation group list.
    level: NetworkNodeGroupLevel!

    # totalCount represents the total number of nodes in the list.
    totalCount: Int!

    # groups represents an array of groups in the list.
    groups: [NetworkNodeGroup!]!
}

# Block is an Opera block chain block.
type Block {
    # Number is the number of this block, starting at 0 for the genesis block.
    number: Long!

    # Hash is the unique block hash of this block.
    hash: Bytes32!

    # Parent is the parent block of this block.
    parent: Block

    # TransactionCount is the number of transactions in this block.
    transactionCount: Int

    # Timestamp is the unix timestamp at which this block was mined.
    timestamp: Long!

    # GasLimit represents the maximum gas allowed in this block.
    gasLimit: Long!

    # GasUsed represents the actual total used gas by all transactions in this block.
    gasUsed: Long!

    # txHashList is the list of unique hash values of transaction
    # assigned to the block.
    txHashList: [Bytes32!]!

    # txList is a list of transactions assigned to the block.
    txList: [Transaction!]!
}

# ERC721Contract represents a generic ERC721 non-fungible tokens (NFT) contract.
type ERC721Contract {
    # address of the token is used as the token's unique identifier.
    address: Address!

    # name of the token.
    name: String!

    # symbol used as an abbreviation for the token.
    symbol: String!

    # totalSupply represents total amount of tokens across all accounts
    totalSupply: BigInt

    # balanceOf represents amount of tokens on the account.
    balanceOf(owner: Address!): BigInt!

    # tokenURI provides URI of Metadata JSON Schema of the token.
    tokenURI(tokenId: BigInt!): String

    # ownerOf provides the owner of NFT identified by tokenId
    ownerOf(tokenId: BigInt!): Address

    # getApproved provides the operator approved by owner
    getApproved(tokenId: BigInt!): Address

    # isApprovedForAll queries the approval status of an operator for a given owner.
    isApprovedForAll(owner: Address!, operator: Address!): Boolean
}

# ERC20Token represents a generic ERC20 token.
type ERC20Token {
    # address of the token is used as the token's unique identifier.
    address: Address!

    # name of the token.
    name: String!

    # symbol used as an abbreviation for the token.
    symbol: String!

    # decimals is the number of decimals the token supports.
    # The most common value is 18 to mimic the ETH to WEI relationship.
    decimals: Int!

    # totalSupply represents total amount of tokens across all accounts
    totalSupply: BigInt!

    # logoURL represents a URL address of a logo of the token. It's always
    # provided, but unknown tokens have this set to a generic logo file.
    logoURL: String!

    # balanceOf represents the total available balance of the token
    # on the account regardless of the DeFi usage of the token.
    # It's effectively the amount available held by the ERC20 token
    # on the account behalf.
    balanceOf(owner: Address!): BigInt!

    # allowance represents the amount of ERC20 tokens unlocked
    # by the owner / token holder to be accessible for the given spender.
    allowance(owner: Address!, spender: Address!): BigInt!
}

# ERC721TransactionList is a list of ERC721 transaction edges provided by sequential access request.
type ERC721TransactionList {
    # Edges contains provided edges of the sequential list.
    edges: [ERC721TransactionListEdge!]!

    # TotalCount is the maximum number of ERC721 transactions available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of ERC721 transaction edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of ERC721 transactions.
type ERC721TransactionListEdge {
    cursor: Cursor!
    trx: ERC721Transaction!
}

# Contract defines block-chain smart contract information container
type Contract {
    "Address represents the contract address."
    address: Address!

    "DeployedBy represents the smart contract deployment transaction reference."
    deployedBy: Transaction!

    "transactionHash represents the smart contract deployment transaction hash."
    transactionHash: Bytes32!

    "Smart contract name. Empty if not available."
    name: String!

    "Smart contract version identifier. Empty if not available."
    version: String!

    """
    License specifies an open source license the contract was published with.
    Empty if not specified.
    """
    license: String!

    "Smart contract author contact. Empty if not available."
    supportContact: String!

    "Smart contract compiler identifier. Empty if not available."
    compiler: String!

    "Smart contract source code. Empty if not available."
    sourceCode: String!

    "Smart contract ABI definition. Empty if not available."
    abi: String!

    """
    Validated is the unix timestamp at which the source code was validated
    against the deployed byte code. Null if not validated yet.
    """
    validated: Long

    "Timestamp is the unix timestamp at which this smart contract was deployed."
    timestamp: Long!
}

# ContractValidationInput represents a set of data sent from client
# to validate deployed contract with the provided source code.
input ContractValidationInput {
    "Address of the contract being validated."
    address: Address!

    "Optional smart contract name. Maximum allowed length is 64 characters."
    name: String

    "Optional smart contract version identifier. Maximum allowed length is 14 characters."
    version: String

    "Optional smart contract author contact. Maximum allowed length is 64 characters."
    supportContact: String

    """
    License specifies an open source license the contract was published with.
    Empty if not specified.
    """
    license: String

    "Optimized specifies if the compiler was set to optimize the byte code."
    optimized: Boolean = true

    """
    OptimizeRuns specifies number of optimization runs the compiler was set
    to execute during the byte code optimizing.
    """
    optimizeRuns: Int = 200

    "Smart contract source code."
    sourceCode: String!
}

# ContractList is a list of smart contract edges provided by sequential access request.
type ContractList {
    # Edges contains provided edges of the sequential list.
    edges: [ContractListEdge!]!

    # TotalCount is the maximum number of contracts available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of contract edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of transactions.
type ContractListEdge {
    cursor: Cursor!
    contract: Contract!
}

# ERC20TransactionList is a list of ERC20 transaction edges provided by sequential access request.
type ERC20TransactionList {
    # Edges contains provided edges of the sequential list.
    edges: [ERC20TransactionListEdge!]!

    # TotalCount is the maximum number of ERC20 transactions available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of ERC20 transaction edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of ERC20 transactions.
type ERC20TransactionListEdge {
    cursor: Cursor!
    trx: ERC20Transaction!
}

# Bytes32 is a 32 byte binary string, represented by 0x prefixed hexadecimal hash.
scalar Bytes32

# Address is a 20 byte Opera address, represented as 0x prefixed hexadecimal number.
scalar Address

# BigInt is a large integer value. Input is accepted as either a JSON number,
# or a hexadecimal string alternatively prefixed with 0x. Output is 0x prefixed hexadecimal.
scalar BigInt

# Long is a 64 bit unsigned integer value.
scalar Long

# Bytes is an arbitrary length binary string, represented as 0x-prefixed hexadecimal.
# An empty byte string is represented as '0x'.
scalar Bytes

# Cursor is a string representing position in a sequential list of edges.
scalar Cursor

# Time represents date and time including time zone information in RFC3339 format.
scalar Time

# ERC1155TransactionList is a list of ERC1155 transaction edges provided by sequential access request.
type ERC1155TransactionList {
    # Edges contains provided edges of the sequential list.
    edges: [ERC1155TransactionListEdge!]!

    # TotalCount is the maximum number of ERC1155 transactions available for sequential access.
    totalCount: BigInt!

    # PageInfo is an information about the current page of ERC1155 transaction edges.
    pageInfo: ListPageInfo!
}

# TransactionListEdge is a single edge in a sequential list of ERC1155 transactions.
type ERC1155TransactionListEdge {
    cursor: Cursor!
    trx: ERC1155Transaction!
}

# ERC1155Transaction represents a transaction on an ERC1155 NFT token.
type ERC1155Transaction {
    # trxHash represents a hash of the transaction
    # executing the ERC1155 call.
    trxHash: Bytes32!

    # transaction represents the transaction
    # executing the ERC1155 call.
    transaction: Transaction!

    # trxIndex represents the index
    # of the ERC1155 call in the transaction logs.
    trxIndex: Long!

    # tokenAddress represents the address
    # of the ERC1155 token contract.
    tokenAddress: Address!

    # token represents the ERC1155 contract detail involved.
    token: ERC1155Contract!

    # tokenId represents the NFT token - one ERC1155 contract can handle multiple NFTs.
    tokenId: BigInt!

    # trxType is the type of the transaction.
    trxType: TokenTransactionType!

    # sender represents the address of the token owner
    # sending the tokens, e.g. the sender.
    sender: Address!

    # recipient represents the address of the token recipient.
    recipient: Address!

    # amount represents the amount of tokens involved in the transaction;
    # please make sure to interpret the amount with the correct number of decimals
    # from the token Metadata JSON Schema.
    amount: BigInt!

    # timeStamp represents the Unix epoch time stamp
    # of the ERC1155 transaction processing.
    timeStamp: Long!
}
# ERC721Transaction represents a transaction on an ERC721 NFT token.
type ERC721Transaction {
    # trxHash represents a hash of the transaction
    # executing the ERC721 call.
    trxHash: Bytes32!

    # transaction represents the transaction
    # executing the ERC721 call.
    transaction: Transaction!

    # trxIndex represents the index
    # of the ERC721 call in the transaction logs.
    trxIndex: Long!

    # tokenAddress represents the address
    # of the ERC721 token contract.
    tokenAddress: Address!

    # token represents the ERC721 contract detail involved.
    token: ERC721Contract!

    # tokenId represents the NFT token - one ERC721 contract can handle multiple NFTs.
    tokenId: BigInt!

    # trxType is the type of the transaction.
    trxType: TokenTransactionType!

    # sender represents the address of the token owner
    # sending the tokens, e.g. the sender.
    sender: Address!

    # recipient represents the address of the token recipient.
    recipient: Address!

    # amount represents the amount of tokens involved
    # in the transaction; please make sure to interpret the amount
    # with the correct number of decimals from the ERC721 token detail.
    amount: BigInt!

    # timeStamp represents the Unix epoch time stamp
    # of the ERC721 transaction processing.
    timeStamp: Long!
}
# TokenTransaction represents a generic token transaction
# of a supported type of token.
type TokenTransaction {
    # Hash is the hash of the executed transaction call.
    hash: Bytes32!

    # trxIndex is the index of the transaction call in a block.
    trxIndex: Long!

    # blockNumber represents the number of the block
    # the transaction was executed in.
    blockNumber: Long!

    # tokenAddress represents the address of the token involved.
    tokenAddress: Address!

    # tokenName represents the name of the token contract.
    # Is empty, if not provided for the given token.
    tokenName: String!

    # tokenSymbol represents the symbol of the token contract.
    # Is empty, if not provided for the given token.
    tokenSymbol: String!

    # tokenType represents the type of the token (i.e. ERC20/ERC721/ERC1155).
    tokenType: String!

    # tokenDecimals is the number of decimals the token supports.
    # The most common value is 18 to mimic the ETH to WEI relationship.
    tokenDecimals: Int!

    # type represents the type of the transaction executed (i.e. Transfer/Mint/Approval).
    type: String!

    # sender of the transaction.
    sender: Address!

    # recipient of the transaction.
    recipient: Address!

    # amount of tokens involved in the transaction.
    amount: BigInt!

    # multi-token contracts (ERC-721/ERC-1155) token ID involved in the transaction.
    tokenId: BigInt!

    # time stamp of the block processing.
    timeStamp: Long!
}

# Account defines block-chain account information container
type Account {
    # Address is the address of the account.
    address: Address!

    # Balance is the current balance of the Account in WEI.
    balance: BigInt!

    # TotalValue is the current total value of the account in WEI.
    # It includes available balance, delegated amount and pending rewards.
    # NOTE: This values is slow to calculate.
    totalValue: BigInt!

    # txCount represents number of transaction sent from the account (Nonce).
    txCount: Long!

    # txList represents list of transactions of the account in form of TransactionList.
    txList(recipient: Address, cursor:Cursor, count:Int!): TransactionList!

    # erc20TxList represents list of ERC20 transactions of the account.
    erc20TxList(cursor:Cursor, count:Int = 25, token: Address, txType: [TokenTransactionType!]): ERC20TransactionList!

    # erc721TxList represents list of ERC721 transactions of the account.
    erc721TxList(cursor:Cursor, count:Int = 25, token: Address, tokenId: BigInt, txType: [TokenTransactionType!]): ERC721TransactionList!

    # erc1155TxList represents list of ERC1155 transactions of the account.
    erc1155TxList(cursor:Cursor, count:Int = 25, token: Address, tokenId: BigInt, txType: [TokenTransactionType!]): ERC1155TransactionList!

    # Details about smart contract, if the account is a smart contract.
    contract: Contract
}

# Root schema definition
schema {
    query: Query
    mutation: Mutation
    subscription: Subscription
}

# Entry points for querying the API
type Query {
    # version represents the API server version responding to your requests.
    version: String!

    # Total number of accounts active on the Opera blockchain.
    accountsActive:Long!

    # Get an Account information by hash address.
    account(address:Address!):Account!

    # Get list of Contracts with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    # ValidatedOnly specifies if the list should contain all the Contracts,
    # or just contracts with validated byte code and available source/ABI.
    contracts(validatedOnly: Boolean = false, cursor:Cursor, count:Int!):ContractList!

    # Get block information by number or by hash.
    # If neither is provided, the most recent block is given.
    block(number:Long, hash: Bytes32):Block

    # Get list of Blocks with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    blocks(cursor:Cursor, count:Int!):BlockList!

    # Get transaction information for given transaction hash.
    transaction(hash:Bytes32!):Transaction

    # Get list of Transactions with at most <count> edges.
    # If <count> is positive, return edges after the cursor,
    # if negative, return edges before the cursor.
    # For undefined cursor, positive <count> starts the list from top,
    # negative <count> starts the list from bottom.
    transactions(cursor:Cursor, count:Int!):TransactionList!

    # Get filtered list of ERC20 Transactions.
    erc20Transactions(cursor:Cursor, count:Int = 25, token: Address, account: Address, txType: [TokenTransactionType!]): ERC20TransactionList!

    # Get filtered list of ERC721 Transactions.
    erc721Transactions(cursor:Cursor, count:Int = 25, token: Address, tokenId: BigInt, account: Address, txType: [TokenTransactionType!]): ERC721TransactionList!

    # Get filtered list of ERC1155 Transactions.
    erc1155Transactions(cursor:Cursor, count:Int = 25, token: Address, tokenId: BigInt, account: Address, txType: [TokenTransactionType!]): ERC1155TransactionList!

    # Returns the current price per gas in WEI units.
    gasPrice: Long!

    # estimateGas returns the estimated amount of gas required
    # for the transaction described by the parameters of the call.
    estimateGas(from: Address, to: Address, value: BigInt, data: String): Long

    # Get price details of the Opera blockchain token for the given target symbols.
    price(to:String!):Price!

    # erc20Token provides the information about an ERC20 token specified by it's
    # address, if available. The resolver returns NULL if the token does not exist.
    erc20Token(token: Address!):ERC20Token

    # erc20TokenList provides list of the most active ERC20 tokens
    # deployed on the block chain.
    erc20TokenList(count: Int = 50):[ERC20Token!]!

    # erc20Assets provides list of tokens owned by the given
    # account address.
    erc20Assets(owner: Address!, count: Int = 50):[ERC20Token!]!

    # ercTotalSupply provides the current total supply amount of a specified ERC20 token
    # identified by it's ERC20 contract address.
    ercTotalSupply(token: Address!):BigInt!

    # ercTokenBalance provides the current available balance of a specified ERC20 token
    # identified by it's ERC20 contract address.
    ercTokenBalance(owner: Address!, token: Address!):BigInt!

    # ercTokenAllowance provides the current amount of ERC20 tokens unlocked
    # by the token owner for the spender to be manipulated with.
    ercTokenAllowance(token: Address!, owner: Address!, spender: Address!):BigInt!

    # erc721Contract provides the information about ERC721 non-fungible token (NFT) by it's address.
    erc721Contract(token: Address!):ERC721Contract

    # erc721ContractList provides list of the most active ERC721 non-fungible tokens (NFT) on the block chain.
    erc721ContractList(count: Int = 50):[ERC721Contract!]!

    # erc1155Token provides the information about ERC1155 multi-token contract by it's address.
    erc1155Contract(address: Address!):ERC1155Contract

    # erc1155ContractList provides list of the most active ERC1155 multi-token contract on the block chain.
    erc1155ContractList(count: Int = 50):[ERC1155Contract!]!

    # trxVolume provides a list of daily aggregations of the network transaction flow.
    # If boundaries are not defined, last 90 days of aggregated trx flow is provided.
    # Boundaries are defined in format YYYY-MM-DD, i.e. 2021-01-23 for January 23rd, 2021.
    trxVolume(from:String, to:String):[DailyTrxVolume!]!

    # trxSpeed provides the recent speed of the network
    # as number of transactions processed per second
    # calculated for the given range denominated in secods. I.e. range:300 means last 5 minutes.
    # Minimal range is 60 seconds, any range below this value will be adjusted to 60 seconds.
    trxSpeed(range: Int = 1200): Float!

    # trxGasSpeed provides average gas consumed by transactions, either base or cumulative,
    # per second in the given date/time period. Please specify the ending date and time
    # as RFC3339 time stamp, i.e. 2021-05-14T00:00:00.000Z. The current time is used if not defined.
    # The range represents the number of seconds prior the end time stamp
    # we use to calculate the average gas consumption.
    trxGasSpeed(range: Int = 1200, to: String): Float!

    # gasPriceList provides a list of gas price ticks for the given date/time span.
    # If the end time is not specified, the list is provided up to the current date/time.
    # The maximal date/time span of the list is 30 days.
    gasPriceList(from: Time! to: Time): [GasPriceTick!]!

    # networkNodesAggregated provides an aggregated list of network nodes on the Opera network.
    networkNodesAggregated(level: NetworkNodeGroupLevel = COUNTRY): NetworkNodeGroupList!
}

# Mutation endpoints for modifying the data
type Mutation {
    # SendTransaction submits a raw signed transaction into the block chain.
    # The tx parameter represents raw signed and RLP encoded transaction data.
    sendTransaction(tx: Bytes!):Transaction

    # Validate a deployed contract byte code with the provided source code
    # so potential users can check the contract source code, access contract ABI
    # to be able to interact with the contract and get the right metadata.
    # Returns updated contract information. If the contract can not be validated,
    # it raises a GraphQL error.
    validateContract(contract: ContractValidationInput!): Contract!
}

# Subscriptions to live events broadcasting
type Subscription {
    # Subscribe to receive information about new blocks in the blockchain.
    onBlock: Block!

    # Subscribe to receive information about new transactions in the blockchain.
    onTransaction: Transaction!
}

`

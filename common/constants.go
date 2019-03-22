package common

const (
	DefaultPageSize        = 100
	DefaultOffset          = 0
	DefaultOrder           = "desc"
	DefaultHydrantCount    = 100
	DefaultHydrantInterval = 3600
	DefaultGasLimit        = 90000
	DefaultGasPrice        = 500000000000
	DefaultNodeStatus      = -1
)

const (
	Rpc_BlockNumber         = "core_blockNumber"
	Rpc_GetBlockByNumber    = "core_getBlockByNumber"
	Rpc_GetTxByHash         = "core_getTransactionByHash"
	Rpc_GetTxReceipt        = "core_getTransactionReceipt"
	Rpc_GetBlance           = "core_getBalance"
	Rpc_Call                = "core_call"
	Rpc_GetAllCandidates    = "core_getAllCandidates"
	Rpc_SendRawTransaction  = "core_sendRawTransaction"
	Rpc_GetTransactionCount = "core_getTransactionCount"
)

const (
	H_ContentType = "application/json; charset=utf-8"
)

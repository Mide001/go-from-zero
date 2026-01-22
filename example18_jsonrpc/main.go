package main

import (
	"encoding/json"
	"fmt"
)

type JSONRPCRequest struct {
	Jsonrpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      int             `json:"id"`
}

type JSONRPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      int         `json:"id"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func handleRPCRequest(reqJSON string) string {
	var req JSONRPCRequest
	json.Unmarshal([]byte(reqJSON), &req)

	var result interface{}
	var rpcErr *RPCError

	switch req.Method {
	case "eth_blockNumber":
		result = "0x10d4f"

	case "eth_getBalance":
		var params []string
		json.Unmarshal(req.Params, &params)
		result = "0xde0b6b3a7640000"

	default:
		rpcErr = &RPCError{
			Code:    -32601,
			Message: "Method not found",
		}
	}

	resp := JSONRPCResponse{
		Jsonrpc: "2.0",
		Result:  result,
		Error:   rpcErr,
		ID:      req.ID,
	}

	respJSON, _ := json.MarshalIndent(resp, "", "  ")
	return string(respJSON)
}

func example18_jsonrpc() {
	fmt.Println("=== Example 18: JSON-RPC Pattern ===")

	request1 := `{
	"jsonrpc": "2.0",
	"method": "eth_blockNumber",
	"params": [],
	"id": 1
	}`

	request2 := `{
	"jsonrpc": "2.0",
	"method": "eth_getBalance",
	"params": ["0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", "latest"],
	"id": 2
	}`

	fmt.Println("Request 1 (block number):")
	fmt.Println(handleRPCRequest(request1))

	fmt.Println("Request 2 (get balance):")
	fmt.Println(handleRPCRequest(request2))

	fmt.Println()
}

func main() {
	example18_jsonrpc()
}

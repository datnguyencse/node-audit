package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/multierr"
)

// json rpc version 2.0
// https://www.jsonrpc.org/specification

const (
	ParseErrorCode            = -32700
	InvalidRequest            = -32600
	MethodNotFound            = -32601
	InvalidParam              = -32602
	InternalError             = -32603
	RequestTimeout            = -32608
	ServerErrorInGeneral      = -32000
	SmartGatewayForwardNeeded = -32001
	JSONRPCVersion            = "2.0"

	ETHChainId                                = "eth_chainId"
	ETHBlockNumber                            = "eth_blockNumber"
	ETHGetBlockByNumber                       = "eth_getBlockByNumber"
	ETHGetBalance                             = "eth_getBalance"
	ETHGetBlockByHash                         = "eth_getBlockByHash"
	ETHGetTransactionByHash                   = "eth_getTransactionByHash"
	ETHGetTransactionByBlockHashAndIndex      = "eth_getTransactionByBlockHashAndIndex"
	ETHGetTransactionByBlockNumberAndIndex    = "eth_getTransactionByBlockNumberAndIndex"
	ETHGetBlockTransactionCountByHash         = "eth_getBlockTransactionCountByHash"
	ETHGetBlockTransactionCountByNumber       = "eth_getBlockTransactionCountByNumber"
	ETHGetTransactionCount                    = "eth_getTransactionCount"
	ETHGetTransactionReceipt                  = "eth_getTransactionReceipt"
	ETHGetLogs                                = "eth_getLogs"
	DebugTraceInternalsAndAccountsByBlockHash = "debug_traceInternalsAndAccountsByBlockHash"
)

var InternalErrorObject = ErrorObject{Code: InternalError, Message: "Internal Error"}

func init() {
}

type ServerRequest struct {
	Version string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params"`
	ID      *json.RawMessage `json:"id"`
}

type ServerResponse[T any] struct {
	Version string           `json:"jsonrpc"`
	ID      *json.RawMessage `json:"id,omitempty"`
	Result  T                `json:"result,omitempty"`
	Error   *ErrorObject     `json:"error,omitempty"`
}

type BatchServerResponse[T any] []ServerResponse[T]

func (b BatchServerResponse[T]) Result() []T {
	results := make([]T, len(b))
	for index, response := range b {
		results[index] = response.Result
	}

	return results
}

func (b BatchServerResponse[T]) Errors() []ErrorObject {
	if len(b) == 0 {
		return []ErrorObject{}
	}
	errs := make([]ErrorObject, 0)
	for _, result := range b {
		if result.Error != nil {
			errs = append(errs, *result.Error)
		}
	}
	return errs
}

func (b BatchServerResponse[T]) ToError() error {
	errs := b.Errors()
	if len(errs) == 0 {
		return nil
	}
	multiErrs := make([]error, len(errs))
	for i, err := range errs {
		multiErrs[i] = err.ToError()
	}
	return multierr.Combine(multiErrs...)
}

func newServerResponse(id *json.RawMessage, result interface{}, error *ErrorObject) *ServerResponse[any] {
	return &ServerResponse[any]{
		Version: JSONRPCVersion,
		ID:      id,
		Result:  result,
		Error:   error,
	}
}

func cannotParse(detail string) ServerResponse[any] {
	return ServerResponse[any]{
		Version: JSONRPCVersion,
		Error: &ErrorObject{
			Code:    ParseErrorCode,
			Message: fmt.Sprintf("cannot parse the body, please review again: %s", detail),
		},
	}
}

func notSupportedVersion(id *json.RawMessage) *ServerResponse[any] {
	return &ServerResponse[any]{
		Version: JSONRPCVersion,
		ID:      id,
		Error: &ErrorObject{
			Code:    InvalidRequest,
			Message: "supported only version: " + JSONRPCVersion,
		},
	}
}

func invalidRequest(id *json.RawMessage, err string) ServerResponse[any] {
	return ServerResponse[any]{
		Version: JSONRPCVersion,
		ID:      id,
		Error: &ErrorObject{
			Code:    InvalidRequest,
			Message: fmt.Sprintf("invalid request: %s", err),
		},
	}
}

func internalError(id *json.RawMessage) *ServerResponse[any] {
	return &ServerResponse[any]{
		Version: JSONRPCVersion,
		ID:      id,
		Error:   &InternalErrorObject,
	}
}

type ErrorObject struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorObjectForward() *ErrorObject {
	return &ErrorObject{
		Code: SmartGatewayForwardNeeded,
	}
}

func (err *ErrorObject) needForward() bool {
	return err.Code == SmartGatewayForwardNeeded
}

func (err *ErrorObject) ToError() error {
	bytes, e := json.Marshal(err)
	if e != nil {
		return multierr.Append(errors.New("cannot marshal ErrorObject to json"), e)
	}
	return errors.New(string(bytes[:]))
}

type TimeoutHandlerFunc func(chan HandlerFuncResult)
type HandlerFuncResult struct {
	returnValue interface{}
	errorObject *ErrorObject
}

func HandlingTimeout(timeout time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func withTimeoutHandle(ctx context.Context, invoker TimeoutHandlerFunc) (returnValue interface{}, errorObject *ErrorObject) {
	resultChan := make(chan HandlerFuncResult, 1)
	go invoker(resultChan)

	select {
	case <-ctx.Done():
		return nil, &ErrorObject{Code: RequestTimeout, Message: "request exceeds timeout"}
	case result := <-resultChan:
		return result.returnValue, result.errorObject
	}
}

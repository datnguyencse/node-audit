package rpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/multierr"
)

const DefaultClientTimeout = 10 * time.Second

type JsonRpcUrl string

type JsonRPCClient struct {
	*fiber.Client
	jsonRpcUrl JsonRpcUrl
}

func NewRPCClient(url JsonRpcUrl) *JsonRPCClient {
	client := fiber.AcquireClient()
	return &JsonRPCClient{
		Client:     client,
		jsonRpcUrl: url,
	}
}

func (client *JsonRPCClient) Forward(requestBody ServerRequest) (int, []byte, []error) {
	statusCode, body, errs := client.Post(string(client.jsonRpcUrl)).
		Timeout(DefaultClientTimeout).
		JSON(requestBody).
		Bytes()
	return statusCode, body, errs
}

func (client *JsonRPCClient) GetLatestBlock() (*BlockResponse, error) {
	var response ServerResponse[BlockResponse]
	err := send(client, blockByNumberServerRequest("latest", false), &response)
	if err != nil {
		return nil, err
	}
	return &response.Result, nil
}

func send[R any](client *JsonRPCClient, request ServerRequest, response *ServerResponse[R]) error {
	return sendWithUrl(client, string(client.jsonRpcUrl), request, response)
}

func sendWithUrl[R any](client *JsonRPCClient, url string, request ServerRequest, response *ServerResponse[R]) error {
	statusCode, _, errs := client.Post(url).Timeout(DefaultClientTimeout).JSON(request).Struct(&response)
	if len(errs) > 0 {
		return fmt.Errorf("RPC client go errors: %s", squashErrors(errs))
	}
	if statusCode != http.StatusOK {
		return fmt.Errorf("RPC server return status code: %d", statusCode)
	}
	if response.Error != nil {
		return response.Error.ToError()
	}
	return nil
}

func sendBatch[R any](client *JsonRPCClient, request []ServerRequest, response *BatchServerResponse[R]) error {
	statusCode, _, errs := client.Post(string(client.jsonRpcUrl)).Timeout(DefaultClientTimeout).JSON(request).Struct(&response)
	if len(errs) > 0 {
		return errors.New(fmt.Sprintf("RPC client go errors: %s", squashErrors(errs)))
	}
	if statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("RPC server return status code: %d", statusCode))
	}
	if err := response.ToError(); err != nil {
		return multierr.Append(errors.New("RPC server return error response"), err)
	}
	return nil
}

func traceInternalTxsAndAccountsByBlockHashRequest(blockHash common.Hash) ServerRequest {
	params := json.RawMessage(fmt.Sprintf(`["%s", {"tracer": "callTracer2"}]`, blockHash.Hex()))
	id := jsonUUID()
	return ServerRequest{
		Version: JSONRPCVersion,
		Method:  DebugTraceInternalsAndAccountsByBlockHash,
		Params:  &params,
		ID:      &id,
	}
}

func blockByNumberServerRequest(number string, fullTxs bool) ServerRequest {
	params := json.RawMessage(fmt.Sprintf(`["%s", %t]`, number, fullTxs))
	id := jsonUUID()
	return ServerRequest{
		Version: JSONRPCVersion,
		Method:  ETHGetBlockByNumber,
		Params:  &params,
		ID:      &id,
	}
}

func logsByBlockHash(blockHash common.Hash) ServerRequest {
	params := json.RawMessage(fmt.Sprintf(`[{"blockHash": "%s"}]`, blockHash.Hex()))
	id := jsonUUID()
	return ServerRequest{
		Version: JSONRPCVersion,
		Method:  ETHGetLogs,
		Params:  &params,
		ID:      &id,
	}
}

func transactionReceiptRequest(txHash common.Hash) ServerRequest {
	params := json.RawMessage(fmt.Sprintf(`["%s"]`, txHash))
	id := jsonUUID()
	return ServerRequest{
		Version: JSONRPCVersion,
		Method:  ETHGetTransactionReceipt,
		Params:  &params,
		ID:      &id,
	}
}

func transactionRequest(txHash common.Hash) ServerRequest {
	params := json.RawMessage(fmt.Sprintf(`["%s"]`, txHash))
	id := jsonUUID()
	return ServerRequest{
		Version: JSONRPCVersion,
		Method:  ETHGetTransactionByHash,
		Params:  &params,
		ID:      &id,
	}
}

func squashErrors(errs []error) string {
	length := len(errs)
	errorStrings := make([]string, length)
	for i := 0; i < len(errs); i++ {
		errorStrings[i] = errs[i].Error()
	}
	return strings.Join(errorStrings, ",")
}

func jsonUUID() json.RawMessage {
	return json.RawMessage(`"` + uuid.NewString() + `"`)
}

// Code generated — DO NOT EDIT.

package vault_match

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
	_ = rpc.API{}
	_ = json.Unmarshal
	_ = reflect.Bool
)

var VaultMatchMetaData = &bind.MetaData{
	ABI: `[{"type":"function","name":"lockMatch","inputs":[],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"resolveMatch","inputs":[{"name":"_winner","type":"uint8"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"status","inputs":[],"outputs":[{"name":"","type":"uint8"}],"stateMutability":"view"},{"type":"function","name":"winner","inputs":[],"outputs":[{"name":"","type":"uint8"}],"stateMutability":"view"},{"type":"function","name":"totalTeamA","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},{"type":"function","name":"totalTeamB","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getTotalDeposits","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},{"type":"function","name":"MATCH_ID","inputs":[],"outputs":[{"name":"","type":"uint256"}],"stateMutability":"view"},{"type":"event","name":"MatchLocked","inputs":[{"name":"matchId","type":"uint256","indexed":false}],"anonymous":false},{"type":"event","name":"MatchResolved","inputs":[{"name":"matchId","type":"uint256","indexed":false},{"name":"winner","type":"uint8","indexed":false}],"anonymous":false}]`,
}

// Contract Method Inputs
type ResolveMatchInput struct {
	Winner uint8
}

// Events
type MatchLockedTopics struct{}
type MatchLockedDecoded struct {
	MatchId *big.Int
}

type MatchResolvedTopics struct{}
type MatchResolvedDecoded struct {
	MatchId *big.Int
	Winner  uint8
}

// Main Binding Type for VaultMatch
type VaultMatch struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   VaultMatchCodec
}

type VaultMatchCodec interface {
	EncodeLockMatchMethodCall() ([]byte, error)
	EncodeResolveMatchMethodCall(in ResolveMatchInput) ([]byte, error)
	EncodeStatusMethodCall() ([]byte, error)
	DecodeStatusMethodOutput(data []byte) (uint8, error)
	EncodeWinnerMethodCall() ([]byte, error)
	DecodeWinnerMethodOutput(data []byte) (uint8, error)
	EncodeTotalTeamAMethodCall() ([]byte, error)
	DecodeTotalTeamAMethodOutput(data []byte) (*big.Int, error)
	EncodeTotalTeamBMethodCall() ([]byte, error)
	DecodeTotalTeamBMethodOutput(data []byte) (*big.Int, error)
	EncodeGetTotalDepositsMethodCall() ([]byte, error)
	DecodeGetTotalDepositsMethodOutput(data []byte) (*big.Int, error)
	EncodeMatchIDMethodCall() ([]byte, error)
	DecodeMatchIDMethodOutput(data []byte) (*big.Int, error)
	MatchLockedLogHash() []byte
	MatchResolvedLogHash() []byte
	DecodeMatchLocked(log *evm.Log) (*MatchLockedDecoded, error)
	DecodeMatchResolved(log *evm.Log) (*MatchResolvedDecoded, error)
}

func NewVaultMatch(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*VaultMatch, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultMatchMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &VaultMatch{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type Codec struct {
	abi *abi.ABI
}

func NewCodec() (VaultMatchCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultMatchMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

// ── Encode Methods ──────────────────────────────────────────────────────

func (c *Codec) EncodeLockMatchMethodCall() ([]byte, error) {
	return c.abi.Pack("lockMatch")
}

func (c *Codec) EncodeResolveMatchMethodCall(in ResolveMatchInput) ([]byte, error) {
	return c.abi.Pack("resolveMatch", in.Winner)
}

func (c *Codec) EncodeStatusMethodCall() ([]byte, error) {
	return c.abi.Pack("status")
}

func (c *Codec) EncodeWinnerMethodCall() ([]byte, error) {
	return c.abi.Pack("winner")
}

func (c *Codec) EncodeTotalTeamAMethodCall() ([]byte, error) {
	return c.abi.Pack("totalTeamA")
}

func (c *Codec) EncodeTotalTeamBMethodCall() ([]byte, error) {
	return c.abi.Pack("totalTeamB")
}

func (c *Codec) EncodeGetTotalDepositsMethodCall() ([]byte, error) {
	return c.abi.Pack("getTotalDeposits")
}

func (c *Codec) EncodeMatchIDMethodCall() ([]byte, error) {
	return c.abi.Pack("MATCH_ID")
}

// ── Decode Methods ──────────────────────────────────────────────────────

func (c *Codec) DecodeStatusMethodOutput(data []byte) (uint8, error) {
	vals, err := c.abi.Methods["status"].Outputs.Unpack(data)
	if err != nil {
		return 0, err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return 0, fmt.Errorf("failed to marshal ABI result: %w", err)
	}
	var result uint8
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}
	return result, nil
}

func (c *Codec) DecodeWinnerMethodOutput(data []byte) (uint8, error) {
	vals, err := c.abi.Methods["winner"].Outputs.Unpack(data)
	if err != nil {
		return 0, err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return 0, fmt.Errorf("failed to marshal ABI result: %w", err)
	}
	var result uint8
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}
	return result, nil
}

func decodeBigIntOutput(a *abi.ABI, method string, data []byte) (*big.Int, error) {
	vals, err := a.Methods[method].Outputs.Unpack(data)
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ABI result: %w", err)
	}
	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	return result, nil
}

func (c *Codec) DecodeTotalTeamAMethodOutput(data []byte) (*big.Int, error) {
	return decodeBigIntOutput(c.abi, "totalTeamA", data)
}

func (c *Codec) DecodeTotalTeamBMethodOutput(data []byte) (*big.Int, error) {
	return decodeBigIntOutput(c.abi, "totalTeamB", data)
}

func (c *Codec) DecodeGetTotalDepositsMethodOutput(data []byte) (*big.Int, error) {
	return decodeBigIntOutput(c.abi, "getTotalDeposits", data)
}

func (c *Codec) DecodeMatchIDMethodOutput(data []byte) (*big.Int, error) {
	return decodeBigIntOutput(c.abi, "MATCH_ID", data)
}

// ── Event Log Hashes ────────────────────────────────────────────────────

func (c *Codec) MatchLockedLogHash() []byte {
	return c.abi.Events["MatchLocked"].ID.Bytes()
}

func (c *Codec) MatchResolvedLogHash() []byte {
	return c.abi.Events["MatchResolved"].ID.Bytes()
}

func (c *Codec) DecodeMatchLocked(log *evm.Log) (*MatchLockedDecoded, error) {
	event := new(MatchLockedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MatchLocked", log.Data); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) DecodeMatchResolved(log *evm.Log) (*MatchResolvedDecoded, error) {
	event := new(MatchResolvedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MatchResolved", log.Data); err != nil {
		return nil, err
	}
	return event, nil
}

// ── Read Contract Methods ───────────────────────────────────────────────

func (c VaultMatch) Status(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[uint8] {
	calldata, err := c.Codec.EncodeStatusMethodCall()
	if err != nil {
		return cre.PromiseFromResult[uint8](0, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})
		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (uint8, error) {
		return c.Codec.DecodeStatusMethodOutput(response.Data)
	})
}

// ── Write Contract Methods (via WriteReport/onReport) ───────────────────

func (c VaultMatch) WriteReportFromLockMatch(
	runtime cre.Runtime,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeLockMatchMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c VaultMatch) WriteReportFromResolveMatch(
	runtime cre.Runtime,
	input ResolveMatchInput,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeResolveMatchMethodCall(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c *VaultMatch) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	default:
		return nil, errors.New("unknown error selector")
	}
}

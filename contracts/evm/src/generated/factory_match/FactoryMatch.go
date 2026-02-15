// Code generated — DO NOT EDIT.

package factory_match

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

var FactoryMatchMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"name\":\"\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"simpleType\":\"address\"},{\"type\":\"address\",\"name\":\"_oracle\",\"simpleType\":\"address\"}],\"id\":\"635e81e0-6f5c-4aa8-b11d-1eefa7bfc9cd\"},{\"type\":\"function\",\"name\":\"oracle\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"id\":\"0x7dc0d1d0\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"id\":\"0x715018a6\"},{\"type\":\"function\",\"name\":\"createMatch\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_teamA\",\"simpleType\":\"string\"},{\"type\":\"string\",\"name\":\"_teamB\",\"simpleType\":\"string\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"matchId\",\"simpleType\":\"uint\"},{\"type\":\"address\",\"name\":\"vault\",\"simpleType\":\"address\"}],\"id\":\"0x57447733\"},{\"type\":\"function\",\"name\":\"setOracle\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newOracle\",\"simpleType\":\"address\"}],\"id\":\"0x7adbf973\"},{\"type\":\"function\",\"name\":\"totalMatches\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x2a5b1451\"},{\"type\":\"function\",\"name\":\"getVault\",\"stateMutability\":\"view\",\"constant\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"_matchId\",\"simpleType\":\"uint\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"id\":\"0x9403b634\"},{\"type\":\"function\",\"name\":\"vaults\",\"stateMutability\":\"view\",\"constant\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"id\":\"0x8c64ea4a\"},{\"type\":\"function\",\"name\":\"nextMatchId\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0xc5adf7c9\"},{\"type\":\"function\",\"name\":\"owner\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"id\":\"0x8da5cb5b\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"simpleType\":\"address\"}],\"id\":\"0xf2fde38b\"},{\"type\":\"function\",\"name\":\"allVaults\",\"stateMutability\":\"view\",\"constant\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"id\":\"0x9094a91e\"},{\"type\":\"event\",\"name\":\"OracleUpdated\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"oldOracle\",\"simpleType\":\"address\"},{\"type\":\"address\",\"name\":\"newOracle\",\"simpleType\":\"address\"}],\"id\":\"0x078c3b417dadf69374a59793b829c52001247130433427049317bde56607b1b7\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"simpleType\":\"address\"},{\"type\":\"address\",\"name\":\"newOwner\",\"simpleType\":\"address\"}],\"id\":\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\"},{\"type\":\"event\",\"name\":\"MatchCreated\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"matchId\",\"simpleType\":\"uint\"},{\"type\":\"address\",\"name\":\"vault\",\"simpleType\":\"address\"},{\"type\":\"string\",\"name\":\"teamA\",\"simpleType\":\"string\"},{\"type\":\"string\",\"name\":\"teamB\",\"simpleType\":\"string\"}],\"id\":\"0x8fcc2ddfd2c264d34bb98b27183fb6ef55ed843dac114ab6b200cc1e8bc64324\"}]",
}

// Structs

// Contract Method Inputs
type AllVaultsInput struct {
	Arg0 *big.Int
}

type CreateMatchInput struct {
	TeamA string
	TeamB string
}

type GetVaultInput struct {
	MatchId *big.Int
}

type SetOracleInput struct {
	NewOracle common.Address
}

type TransferOwnershipInput struct {
	NewOwner common.Address
}

type VaultsInput struct {
	Arg0 *big.Int
}

// Contract Method Outputs
type CreateMatchOutput struct {
	MatchId *big.Int
	Vault   common.Address
}

// Errors

// Events
// The <Event>Topics struct should be used as a filter (for log triggers).
// Note: It is only possible to filter on indexed fields.
// Indexed (string and bytes) fields will be of type common.Hash.
// They need to he (crypto.Keccak256) hashed and passed in.
// Indexed (tuple/slice/array) fields can be passed in as is, the Encode<Event>Topics function will handle the hashing.
//
// The <Event>Decoded struct will be the result of calling decode (Adapt) on the log trigger result.
// Indexed dynamic type fields will be of type common.Hash.

type MatchCreatedTopics struct {
}

type MatchCreatedDecoded struct {
	MatchId *big.Int
	Vault   common.Address
	TeamA   string
	TeamB   string
}

type OracleUpdatedTopics struct {
}

type OracleUpdatedDecoded struct {
	OldOracle common.Address
	NewOracle common.Address
}

type OwnershipTransferredTopics struct {
}

type OwnershipTransferredDecoded struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

// Main Binding Type for FactoryMatch
type FactoryMatch struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   FactoryMatchCodec
}

type FactoryMatchCodec interface {
	EncodeAllVaultsMethodCall(in AllVaultsInput) ([]byte, error)
	DecodeAllVaultsMethodOutput(data []byte) (common.Address, error)
	EncodeCreateMatchMethodCall(in CreateMatchInput) ([]byte, error)
	DecodeCreateMatchMethodOutput(data []byte) (CreateMatchOutput, error)
	EncodeGetVaultMethodCall(in GetVaultInput) ([]byte, error)
	DecodeGetVaultMethodOutput(data []byte) (common.Address, error)
	EncodeNextMatchIdMethodCall() ([]byte, error)
	DecodeNextMatchIdMethodOutput(data []byte) (*big.Int, error)
	EncodeOracleMethodCall() ([]byte, error)
	DecodeOracleMethodOutput(data []byte) (common.Address, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodeRenounceOwnershipMethodCall() ([]byte, error)
	EncodeSetOracleMethodCall(in SetOracleInput) ([]byte, error)
	EncodeTotalMatchesMethodCall() ([]byte, error)
	DecodeTotalMatchesMethodOutput(data []byte) (*big.Int, error)
	EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error)
	EncodeVaultsMethodCall(in VaultsInput) ([]byte, error)
	DecodeVaultsMethodOutput(data []byte) (common.Address, error)
	MatchCreatedLogHash() []byte
	EncodeMatchCreatedTopics(evt abi.Event, values []MatchCreatedTopics) ([]*evm.TopicValues, error)
	DecodeMatchCreated(log *evm.Log) (*MatchCreatedDecoded, error)
	OracleUpdatedLogHash() []byte
	EncodeOracleUpdatedTopics(evt abi.Event, values []OracleUpdatedTopics) ([]*evm.TopicValues, error)
	DecodeOracleUpdated(log *evm.Log) (*OracleUpdatedDecoded, error)
	OwnershipTransferredLogHash() []byte
	EncodeOwnershipTransferredTopics(evt abi.Event, values []OwnershipTransferredTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error)
}

func NewFactoryMatch(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*FactoryMatch, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryMatchMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &FactoryMatch{
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

func NewCodec() (FactoryMatchCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryMatchMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeAllVaultsMethodCall(in AllVaultsInput) ([]byte, error) {
	return c.abi.Pack("allVaults", in.Arg0)
}

func (c *Codec) DecodeAllVaultsMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["allVaults"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeCreateMatchMethodCall(in CreateMatchInput) ([]byte, error) {
	return c.abi.Pack("createMatch", in.TeamA, in.TeamB)
}

func (c *Codec) DecodeCreateMatchMethodOutput(data []byte) (CreateMatchOutput, error) {
	vals, err := c.abi.Methods["createMatch"].Outputs.Unpack(data)
	if err != nil {
		return CreateMatchOutput{}, err
	}
	if len(vals) != 2 {
		return CreateMatchOutput{}, fmt.Errorf("expected 2 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return CreateMatchOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 *big.Int
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return CreateMatchOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return CreateMatchOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 common.Address
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return CreateMatchOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return CreateMatchOutput{
		MatchId: result0,
		Vault:   result1,
	}, nil
}

func (c *Codec) EncodeGetVaultMethodCall(in GetVaultInput) ([]byte, error) {
	return c.abi.Pack("getVault", in.MatchId)
}

func (c *Codec) DecodeGetVaultMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["getVault"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeNextMatchIdMethodCall() ([]byte, error) {
	return c.abi.Pack("nextMatchId")
}

func (c *Codec) DecodeNextMatchIdMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["nextMatchId"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeOracleMethodCall() ([]byte, error) {
	return c.abi.Pack("oracle")
}

func (c *Codec) DecodeOracleMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["oracle"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeOwnerMethodCall() ([]byte, error) {
	return c.abi.Pack("owner")
}

func (c *Codec) DecodeOwnerMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["owner"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeRenounceOwnershipMethodCall() ([]byte, error) {
	return c.abi.Pack("renounceOwnership")
}

func (c *Codec) EncodeSetOracleMethodCall(in SetOracleInput) ([]byte, error) {
	return c.abi.Pack("setOracle", in.NewOracle)
}

func (c *Codec) EncodeTotalMatchesMethodCall() ([]byte, error) {
	return c.abi.Pack("totalMatches")
}

func (c *Codec) DecodeTotalMatchesMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["totalMatches"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error) {
	return c.abi.Pack("transferOwnership", in.NewOwner)
}

func (c *Codec) EncodeVaultsMethodCall(in VaultsInput) ([]byte, error) {
	return c.abi.Pack("vaults", in.Arg0)
}

func (c *Codec) DecodeVaultsMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["vaults"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) MatchCreatedLogHash() []byte {
	return c.abi.Events["MatchCreated"].ID.Bytes()
}

func (c *Codec) EncodeMatchCreatedTopics(
	evt abi.Event,
	values []MatchCreatedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeMatchCreated decodes a log into a MatchCreated struct.
func (c *Codec) DecodeMatchCreated(log *evm.Log) (*MatchCreatedDecoded, error) {
	event := new(MatchCreatedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MatchCreated", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["MatchCreated"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) OracleUpdatedLogHash() []byte {
	return c.abi.Events["OracleUpdated"].ID.Bytes()
}

func (c *Codec) EncodeOracleUpdatedTopics(
	evt abi.Event,
	values []OracleUpdatedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeOracleUpdated decodes a log into a OracleUpdated struct.
func (c *Codec) DecodeOracleUpdated(log *evm.Log) (*OracleUpdatedDecoded, error) {
	event := new(OracleUpdatedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "OracleUpdated", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["OracleUpdated"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) OwnershipTransferredLogHash() []byte {
	return c.abi.Events["OwnershipTransferred"].ID.Bytes()
}

func (c *Codec) EncodeOwnershipTransferredTopics(
	evt abi.Event,
	values []OwnershipTransferredTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeOwnershipTransferred decodes a log into a OwnershipTransferred struct.
func (c *Codec) DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error) {
	event := new(OwnershipTransferredDecoded)
	if err := c.abi.UnpackIntoInterface(event, "OwnershipTransferred", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["OwnershipTransferred"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c FactoryMatch) AllVaults(
	runtime cre.Runtime,
	args AllVaultsInput,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeAllVaultsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeAllVaultsMethodOutput(response.Data)
	})

}

func (c FactoryMatch) GetVault(
	runtime cre.Runtime,
	args GetVaultInput,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeGetVaultMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeGetVaultMethodOutput(response.Data)
	})

}

func (c FactoryMatch) NextMatchId(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeNextMatchIdMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeNextMatchIdMethodOutput(response.Data)
	})

}

func (c FactoryMatch) Oracle(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOracleMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeOracleMethodOutput(response.Data)
	})

}

func (c FactoryMatch) Owner(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOwnerMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeOwnerMethodOutput(response.Data)
	})

}

func (c FactoryMatch) TotalMatches(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeTotalMatchesMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeTotalMatchesMethodOutput(response.Data)
	})

}

func (c FactoryMatch) Vaults(
	runtime cre.Runtime,
	args VaultsInput,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeVaultsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeVaultsMethodOutput(response.Data)
	})

}

func (c FactoryMatch) WriteReport(
	runtime cre.Runtime,
	report *cre.Report,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  c.Address.Bytes(),
		Report:    report,
		GasConfig: gasConfig,
	})
}

func (c *FactoryMatch) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	default:
		return nil, errors.New("unknown error selector")
	}
}

// MatchCreatedTrigger wraps the raw log trigger and provides decoded MatchCreatedDecoded data
type MatchCreatedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]               // Embed the raw trigger
	contract                        *FactoryMatch // Keep reference for decoding
}

// Adapt method that decodes the log into MatchCreated data
func (t *MatchCreatedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[MatchCreatedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeMatchCreated(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MatchCreated log: %w", err)
	}

	return &bindings.DecodedLog[MatchCreatedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *FactoryMatch) LogTriggerMatchCreatedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []MatchCreatedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[MatchCreatedDecoded]], error) {
	event := c.ABI.Events["MatchCreated"]
	topics, err := c.Codec.EncodeMatchCreatedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for MatchCreated: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &MatchCreatedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *FactoryMatch) FilterLogsMatchCreated(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.MatchCreatedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OracleUpdatedTrigger wraps the raw log trigger and provides decoded OracleUpdatedDecoded data
type OracleUpdatedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]               // Embed the raw trigger
	contract                        *FactoryMatch // Keep reference for decoding
}

// Adapt method that decodes the log into OracleUpdated data
func (t *OracleUpdatedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[OracleUpdatedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeOracleUpdated(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OracleUpdated log: %w", err)
	}

	return &bindings.DecodedLog[OracleUpdatedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *FactoryMatch) LogTriggerOracleUpdatedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OracleUpdatedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OracleUpdatedDecoded]], error) {
	event := c.ABI.Events["OracleUpdated"]
	topics, err := c.Codec.EncodeOracleUpdatedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for OracleUpdated: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &OracleUpdatedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *FactoryMatch) FilterLogsOracleUpdated(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.OracleUpdatedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OwnershipTransferredTrigger wraps the raw log trigger and provides decoded OwnershipTransferredDecoded data
type OwnershipTransferredTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]               // Embed the raw trigger
	contract                        *FactoryMatch // Keep reference for decoding
}

// Adapt method that decodes the log into OwnershipTransferred data
func (t *OwnershipTransferredTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[OwnershipTransferredDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeOwnershipTransferred(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OwnershipTransferred log: %w", err)
	}

	return &bindings.DecodedLog[OwnershipTransferredDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *FactoryMatch) LogTriggerOwnershipTransferredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipTransferredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipTransferredDecoded]], error) {
	event := c.ABI.Events["OwnershipTransferred"]
	topics, err := c.Codec.EncodeOwnershipTransferredTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for OwnershipTransferred: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &OwnershipTransferredTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *FactoryMatch) FilterLogsOwnershipTransferred(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.OwnershipTransferredLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

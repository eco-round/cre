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
	ABI: "[{\"type\":\"constructor\",\"name\":\"\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"simpleType\":\"address\"},{\"type\":\"address\",\"name\":\"_oracle\",\"simpleType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_matchId\",\"simpleType\":\"uint\"},{\"type\":\"string\",\"name\":\"_teamA\",\"simpleType\":\"string\"},{\"type\":\"string\",\"name\":\"_teamB\",\"simpleType\":\"string\"}],\"id\":\"1e99aefd-833a-4490-8302-1d9d30107caa\"},{\"type\":\"function\",\"name\":\"hasClaimed\",\"stateMutability\":\"view\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"simpleType\":\"bool\"}],\"id\":\"0x73b2e80e\"},{\"type\":\"function\",\"name\":\"deposit\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"inputs\":[{\"type\":\"uint8\",\"name\":\"team\",\"simpleType\":\"uint\"},{\"type\":\"uint256\",\"name\":\"amount\",\"simpleType\":\"uint\"}],\"id\":\"0xf4d4c9d7\"},{\"type\":\"function\",\"name\":\"resolveMatch\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"inputs\":[{\"type\":\"uint8\",\"name\":\"_winner\",\"simpleType\":\"uint\"}],\"id\":\"0x7986ad49\"},{\"type\":\"function\",\"name\":\"totalYield\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x01418205\"},{\"type\":\"function\",\"name\":\"emergencyWithdrawFromYield\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"id\":\"0x0e070cb0\"},{\"type\":\"function\",\"name\":\"claim\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"id\":\"0x4e71d92d\"},{\"type\":\"function\",\"name\":\"teamBName\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"simpleType\":\"string\"}],\"id\":\"0x11aa09b1\"},{\"type\":\"function\",\"name\":\"teamAName\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"simpleType\":\"string\"}],\"id\":\"0x4e108c19\"},{\"type\":\"function\",\"name\":\"paused\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"simpleType\":\"bool\"}],\"id\":\"0x5c975abb\"},{\"type\":\"function\",\"name\":\"lockMatch\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"id\":\"0xc49df529\"},{\"type\":\"function\",\"name\":\"emergencyRefund\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"simpleType\":\"address\"}],\"id\":\"0x045f7019\"},{\"type\":\"function\",\"name\":\"totalTeamB\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x2288fc2b\"},{\"type\":\"function\",\"name\":\"owner\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"id\":\"0x8da5cb5b\"},{\"type\":\"function\",\"name\":\"totalTeamA\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x669b96e2\"},{\"type\":\"function\",\"name\":\"userDeposits\",\"stateMutability\":\"view\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"},{\"type\":\"uint8\",\"name\":\"\",\"simpleType\":\"uint\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x42db5e35\"},{\"type\":\"function\",\"name\":\"getUserTotalDeposit\",\"stateMutability\":\"view\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"simpleType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0xa4c828dc\"},{\"type\":\"function\",\"name\":\"setOracle\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newOracle\",\"simpleType\":\"address\"}],\"id\":\"0x7adbf973\"},{\"type\":\"function\",\"name\":\"getExpectedPayout\",\"stateMutability\":\"view\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"simpleType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x98eca55f\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"id\":\"0x715018a6\"},{\"type\":\"function\",\"name\":\"oracle\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"id\":\"0x7dc0d1d0\"},{\"type\":\"function\",\"name\":\"status\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x200d2ed2\"},{\"type\":\"function\",\"name\":\"winner\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0xdfbf53ae\"},{\"type\":\"function\",\"name\":\"getYieldBalance\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x489cc20c\"},{\"type\":\"function\",\"name\":\"MORPHO_VAULT\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"id\":\"0x5bca49d2\"},{\"type\":\"function\",\"name\":\"getTotalDeposits\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x168a4822\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"simpleType\":\"address\"}],\"id\":\"0xf2fde38b\"},{\"type\":\"function\",\"name\":\"unpause\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"id\":\"0x3f4ba83a\"},{\"type\":\"function\",\"name\":\"USDC\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"simpleType\":\"address\"}],\"id\":\"0x89a30271\"},{\"type\":\"function\",\"name\":\"MATCH_ID\",\"stateMutability\":\"view\",\"constant\":false,\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"simpleType\":\"uint\"}],\"id\":\"0x75e10829\"},{\"type\":\"function\",\"name\":\"pause\",\"stateMutability\":\"nonpayable\",\"constant\":false,\"id\":\"0x8456cb59\"},{\"type\":\"event\",\"name\":\"OracleUpdated\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"oldOracle\",\"simpleType\":\"address\"},{\"type\":\"address\",\"name\":\"newOracle\",\"simpleType\":\"address\"}],\"id\":\"0x078c3b417dadf69374a59793b829c52001247130433427049317bde56607b1b7\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"simpleType\":\"address\"},{\"type\":\"address\",\"name\":\"newOwner\",\"simpleType\":\"address\"}],\"id\":\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\"},{\"type\":\"event\",\"name\":\"Unpaused\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"simpleType\":\"address\"}],\"id\":\"0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa\"},{\"type\":\"event\",\"name\":\"EmergencyRefund\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"simpleType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"simpleType\":\"uint\"}],\"id\":\"0xdf36e221948da014ebe0f9f6bb96696776424780da298e7f05e2f362dcd4289a\"},{\"type\":\"event\",\"name\":\"EmergencyYieldWithdrawn\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"simpleType\":\"uint\"}],\"id\":\"0x32f0f1630e950fbc850f76becff89beb4eac3e75d09ee1abcf4a231fa783a65a\"},{\"type\":\"event\",\"name\":\"Paused\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"simpleType\":\"address\"}],\"id\":\"0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258\"},{\"type\":\"event\",\"name\":\"Claimed\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"simpleType\":\"address\"},{\"type\":\"uint256\",\"name\":\"principal\",\"simpleType\":\"uint\"},{\"type\":\"uint256\",\"name\":\"yieldShare\",\"simpleType\":\"uint\"}],\"id\":\"0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a\"},{\"type\":\"event\",\"name\":\"MatchLocked\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"matchId\",\"simpleType\":\"uint\"}],\"id\":\"0x2f5f86d0163c8f8ac2745a053af228bf204225ec58d747c366151e05fc7a73a8\"},{\"type\":\"event\",\"name\":\"Deposited\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"simpleType\":\"address\"},{\"type\":\"uint8\",\"name\":\"team\",\"simpleType\":\"uint\"},{\"type\":\"uint256\",\"name\":\"amount\",\"simpleType\":\"uint\"}],\"id\":\"0x1d0787aee899a49ef81d0b11da9aca5455b46aefed042a41bd398d74619cab00\"},{\"type\":\"event\",\"name\":\"MatchResolved\",\"stateMutability\":\"\",\"constant\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"matchId\",\"simpleType\":\"uint\"},{\"type\":\"uint8\",\"name\":\"winner\",\"simpleType\":\"uint\"}],\"id\":\"0xde1c752173fee6bd3976fb1f165361c647fffc7a33987556b7f54b6a2de08ece\"}]",
}

// Structs

// Contract Method Inputs
type DepositInput struct {
	Team   uint8
	Amount *big.Int
}

type EmergencyRefundInput struct {
	User common.Address
}

type GetExpectedPayoutInput struct {
	User common.Address
}

type GetUserTotalDepositInput struct {
	User common.Address
}

type HasClaimedInput struct {
	Arg0 common.Address
}

type ResolveMatchInput struct {
	Winner uint8
}

type SetOracleInput struct {
	NewOracle common.Address
}

type TransferOwnershipInput struct {
	NewOwner common.Address
}

type UserDepositsInput struct {
	Arg0 common.Address
	Arg1 uint8
}

// Contract Method Outputs

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

type ClaimedTopics struct {
}

type ClaimedDecoded struct {
	User       common.Address
	Principal  *big.Int
	YieldShare *big.Int
}

type DepositedTopics struct {
}

type DepositedDecoded struct {
	User   common.Address
	Team   uint8
	Amount *big.Int
}

type EmergencyRefundTopics struct {
}

type EmergencyRefundDecoded struct {
	User   common.Address
	Amount *big.Int
}

type EmergencyYieldWithdrawnTopics struct {
}

type EmergencyYieldWithdrawnDecoded struct {
	Amount *big.Int
}

type MatchLockedTopics struct {
}

type MatchLockedDecoded struct {
	MatchId *big.Int
}

type MatchResolvedTopics struct {
}

type MatchResolvedDecoded struct {
	MatchId *big.Int
	Winner  uint8
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

type PausedTopics struct {
}

type PausedDecoded struct {
	Account common.Address
}

type UnpausedTopics struct {
}

type UnpausedDecoded struct {
	Account common.Address
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
	EncodeMATCHIDMethodCall() ([]byte, error)
	DecodeMATCHIDMethodOutput(data []byte) (*big.Int, error)
	EncodeMORPHOVAULTMethodCall() ([]byte, error)
	DecodeMORPHOVAULTMethodOutput(data []byte) (common.Address, error)
	EncodeUSDCMethodCall() ([]byte, error)
	DecodeUSDCMethodOutput(data []byte) (common.Address, error)
	EncodeClaimMethodCall() ([]byte, error)
	EncodeDepositMethodCall(in DepositInput) ([]byte, error)
	EncodeEmergencyRefundMethodCall(in EmergencyRefundInput) ([]byte, error)
	EncodeEmergencyWithdrawFromYieldMethodCall() ([]byte, error)
	EncodeGetExpectedPayoutMethodCall(in GetExpectedPayoutInput) ([]byte, error)
	DecodeGetExpectedPayoutMethodOutput(data []byte) (*big.Int, error)
	EncodeGetTotalDepositsMethodCall() ([]byte, error)
	DecodeGetTotalDepositsMethodOutput(data []byte) (*big.Int, error)
	EncodeGetUserTotalDepositMethodCall(in GetUserTotalDepositInput) ([]byte, error)
	DecodeGetUserTotalDepositMethodOutput(data []byte) (*big.Int, error)
	EncodeGetYieldBalanceMethodCall() ([]byte, error)
	DecodeGetYieldBalanceMethodOutput(data []byte) (*big.Int, error)
	EncodeHasClaimedMethodCall(in HasClaimedInput) ([]byte, error)
	DecodeHasClaimedMethodOutput(data []byte) (bool, error)
	EncodeLockMatchMethodCall() ([]byte, error)
	EncodeOracleMethodCall() ([]byte, error)
	DecodeOracleMethodOutput(data []byte) (common.Address, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodePauseMethodCall() ([]byte, error)
	EncodePausedMethodCall() ([]byte, error)
	DecodePausedMethodOutput(data []byte) (bool, error)
	EncodeRenounceOwnershipMethodCall() ([]byte, error)
	EncodeResolveMatchMethodCall(in ResolveMatchInput) ([]byte, error)
	EncodeSetOracleMethodCall(in SetOracleInput) ([]byte, error)
	EncodeStatusMethodCall() ([]byte, error)
	DecodeStatusMethodOutput(data []byte) (uint8, error)
	EncodeTeamANameMethodCall() ([]byte, error)
	DecodeTeamANameMethodOutput(data []byte) (string, error)
	EncodeTeamBNameMethodCall() ([]byte, error)
	DecodeTeamBNameMethodOutput(data []byte) (string, error)
	EncodeTotalTeamAMethodCall() ([]byte, error)
	DecodeTotalTeamAMethodOutput(data []byte) (*big.Int, error)
	EncodeTotalTeamBMethodCall() ([]byte, error)
	DecodeTotalTeamBMethodOutput(data []byte) (*big.Int, error)
	EncodeTotalYieldMethodCall() ([]byte, error)
	DecodeTotalYieldMethodOutput(data []byte) (*big.Int, error)
	EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error)
	EncodeUnpauseMethodCall() ([]byte, error)
	EncodeUserDepositsMethodCall(in UserDepositsInput) ([]byte, error)
	DecodeUserDepositsMethodOutput(data []byte) (*big.Int, error)
	EncodeWinnerMethodCall() ([]byte, error)
	DecodeWinnerMethodOutput(data []byte) (uint8, error)
	ClaimedLogHash() []byte
	EncodeClaimedTopics(evt abi.Event, values []ClaimedTopics) ([]*evm.TopicValues, error)
	DecodeClaimed(log *evm.Log) (*ClaimedDecoded, error)
	DepositedLogHash() []byte
	EncodeDepositedTopics(evt abi.Event, values []DepositedTopics) ([]*evm.TopicValues, error)
	DecodeDeposited(log *evm.Log) (*DepositedDecoded, error)
	EmergencyRefundLogHash() []byte
	EncodeEmergencyRefundTopics(evt abi.Event, values []EmergencyRefundTopics) ([]*evm.TopicValues, error)
	DecodeEmergencyRefund(log *evm.Log) (*EmergencyRefundDecoded, error)
	EmergencyYieldWithdrawnLogHash() []byte
	EncodeEmergencyYieldWithdrawnTopics(evt abi.Event, values []EmergencyYieldWithdrawnTopics) ([]*evm.TopicValues, error)
	DecodeEmergencyYieldWithdrawn(log *evm.Log) (*EmergencyYieldWithdrawnDecoded, error)
	MatchLockedLogHash() []byte
	EncodeMatchLockedTopics(evt abi.Event, values []MatchLockedTopics) ([]*evm.TopicValues, error)
	DecodeMatchLocked(log *evm.Log) (*MatchLockedDecoded, error)
	MatchResolvedLogHash() []byte
	EncodeMatchResolvedTopics(evt abi.Event, values []MatchResolvedTopics) ([]*evm.TopicValues, error)
	DecodeMatchResolved(log *evm.Log) (*MatchResolvedDecoded, error)
	OracleUpdatedLogHash() []byte
	EncodeOracleUpdatedTopics(evt abi.Event, values []OracleUpdatedTopics) ([]*evm.TopicValues, error)
	DecodeOracleUpdated(log *evm.Log) (*OracleUpdatedDecoded, error)
	OwnershipTransferredLogHash() []byte
	EncodeOwnershipTransferredTopics(evt abi.Event, values []OwnershipTransferredTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error)
	PausedLogHash() []byte
	EncodePausedTopics(evt abi.Event, values []PausedTopics) ([]*evm.TopicValues, error)
	DecodePaused(log *evm.Log) (*PausedDecoded, error)
	UnpausedLogHash() []byte
	EncodeUnpausedTopics(evt abi.Event, values []UnpausedTopics) ([]*evm.TopicValues, error)
	DecodeUnpaused(log *evm.Log) (*UnpausedDecoded, error)
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

func (c *Codec) EncodeMATCHIDMethodCall() ([]byte, error) {
	return c.abi.Pack("MATCH_ID")
}

func (c *Codec) DecodeMATCHIDMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["MATCH_ID"].Outputs.Unpack(data)
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

func (c *Codec) EncodeMORPHOVAULTMethodCall() ([]byte, error) {
	return c.abi.Pack("MORPHO_VAULT")
}

func (c *Codec) DecodeMORPHOVAULTMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["MORPHO_VAULT"].Outputs.Unpack(data)
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

func (c *Codec) EncodeUSDCMethodCall() ([]byte, error) {
	return c.abi.Pack("USDC")
}

func (c *Codec) DecodeUSDCMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["USDC"].Outputs.Unpack(data)
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

func (c *Codec) EncodeClaimMethodCall() ([]byte, error) {
	return c.abi.Pack("claim")
}

func (c *Codec) EncodeDepositMethodCall(in DepositInput) ([]byte, error) {
	return c.abi.Pack("deposit", in.Team, in.Amount)
}

func (c *Codec) EncodeEmergencyRefundMethodCall(in EmergencyRefundInput) ([]byte, error) {
	return c.abi.Pack("emergencyRefund", in.User)
}

func (c *Codec) EncodeEmergencyWithdrawFromYieldMethodCall() ([]byte, error) {
	return c.abi.Pack("emergencyWithdrawFromYield")
}

func (c *Codec) EncodeGetExpectedPayoutMethodCall(in GetExpectedPayoutInput) ([]byte, error) {
	return c.abi.Pack("getExpectedPayout", in.User)
}

func (c *Codec) DecodeGetExpectedPayoutMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getExpectedPayout"].Outputs.Unpack(data)
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

func (c *Codec) EncodeGetTotalDepositsMethodCall() ([]byte, error) {
	return c.abi.Pack("getTotalDeposits")
}

func (c *Codec) DecodeGetTotalDepositsMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getTotalDeposits"].Outputs.Unpack(data)
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

func (c *Codec) EncodeGetUserTotalDepositMethodCall(in GetUserTotalDepositInput) ([]byte, error) {
	return c.abi.Pack("getUserTotalDeposit", in.User)
}

func (c *Codec) DecodeGetUserTotalDepositMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getUserTotalDeposit"].Outputs.Unpack(data)
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

func (c *Codec) EncodeGetYieldBalanceMethodCall() ([]byte, error) {
	return c.abi.Pack("getYieldBalance")
}

func (c *Codec) DecodeGetYieldBalanceMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getYieldBalance"].Outputs.Unpack(data)
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

func (c *Codec) EncodeHasClaimedMethodCall(in HasClaimedInput) ([]byte, error) {
	return c.abi.Pack("hasClaimed", in.Arg0)
}

func (c *Codec) DecodeHasClaimedMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["hasClaimed"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeLockMatchMethodCall() ([]byte, error) {
	return c.abi.Pack("lockMatch")
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

func (c *Codec) EncodePauseMethodCall() ([]byte, error) {
	return c.abi.Pack("pause")
}

func (c *Codec) EncodePausedMethodCall() ([]byte, error) {
	return c.abi.Pack("paused")
}

func (c *Codec) DecodePausedMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["paused"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeRenounceOwnershipMethodCall() ([]byte, error) {
	return c.abi.Pack("renounceOwnership")
}

func (c *Codec) EncodeResolveMatchMethodCall(in ResolveMatchInput) ([]byte, error) {
	return c.abi.Pack("resolveMatch", in.Winner)
}

func (c *Codec) EncodeSetOracleMethodCall(in SetOracleInput) ([]byte, error) {
	return c.abi.Pack("setOracle", in.NewOracle)
}

func (c *Codec) EncodeStatusMethodCall() ([]byte, error) {
	return c.abi.Pack("status")
}

func (c *Codec) DecodeStatusMethodOutput(data []byte) (uint8, error) {
	vals, err := c.abi.Methods["status"].Outputs.Unpack(data)
	if err != nil {
		return *new(uint8), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(uint8), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result uint8
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(uint8), fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeTeamANameMethodCall() ([]byte, error) {
	return c.abi.Pack("teamAName")
}

func (c *Codec) DecodeTeamANameMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["teamAName"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(string), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result string
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(string), fmt.Errorf("failed to unmarshal to string: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeTeamBNameMethodCall() ([]byte, error) {
	return c.abi.Pack("teamBName")
}

func (c *Codec) DecodeTeamBNameMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["teamBName"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(string), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result string
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(string), fmt.Errorf("failed to unmarshal to string: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeTotalTeamAMethodCall() ([]byte, error) {
	return c.abi.Pack("totalTeamA")
}

func (c *Codec) DecodeTotalTeamAMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["totalTeamA"].Outputs.Unpack(data)
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

func (c *Codec) EncodeTotalTeamBMethodCall() ([]byte, error) {
	return c.abi.Pack("totalTeamB")
}

func (c *Codec) DecodeTotalTeamBMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["totalTeamB"].Outputs.Unpack(data)
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

func (c *Codec) EncodeTotalYieldMethodCall() ([]byte, error) {
	return c.abi.Pack("totalYield")
}

func (c *Codec) DecodeTotalYieldMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["totalYield"].Outputs.Unpack(data)
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

func (c *Codec) EncodeUnpauseMethodCall() ([]byte, error) {
	return c.abi.Pack("unpause")
}

func (c *Codec) EncodeUserDepositsMethodCall(in UserDepositsInput) ([]byte, error) {
	return c.abi.Pack("userDeposits", in.Arg0, in.Arg1)
}

func (c *Codec) DecodeUserDepositsMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["userDeposits"].Outputs.Unpack(data)
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

func (c *Codec) EncodeWinnerMethodCall() ([]byte, error) {
	return c.abi.Pack("winner")
}

func (c *Codec) DecodeWinnerMethodOutput(data []byte) (uint8, error) {
	vals, err := c.abi.Methods["winner"].Outputs.Unpack(data)
	if err != nil {
		return *new(uint8), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(uint8), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result uint8
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(uint8), fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}

	return result, nil
}

func (c *Codec) ClaimedLogHash() []byte {
	return c.abi.Events["Claimed"].ID.Bytes()
}

func (c *Codec) EncodeClaimedTopics(
	evt abi.Event,
	values []ClaimedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeClaimed decodes a log into a Claimed struct.
func (c *Codec) DecodeClaimed(log *evm.Log) (*ClaimedDecoded, error) {
	event := new(ClaimedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Claimed", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Claimed"].Inputs {
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

func (c *Codec) DepositedLogHash() []byte {
	return c.abi.Events["Deposited"].ID.Bytes()
}

func (c *Codec) EncodeDepositedTopics(
	evt abi.Event,
	values []DepositedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeDeposited decodes a log into a Deposited struct.
func (c *Codec) DecodeDeposited(log *evm.Log) (*DepositedDecoded, error) {
	event := new(DepositedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Deposited", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Deposited"].Inputs {
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

func (c *Codec) EmergencyRefundLogHash() []byte {
	return c.abi.Events["EmergencyRefund"].ID.Bytes()
}

func (c *Codec) EncodeEmergencyRefundTopics(
	evt abi.Event,
	values []EmergencyRefundTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeEmergencyRefund decodes a log into a EmergencyRefund struct.
func (c *Codec) DecodeEmergencyRefund(log *evm.Log) (*EmergencyRefundDecoded, error) {
	event := new(EmergencyRefundDecoded)
	if err := c.abi.UnpackIntoInterface(event, "EmergencyRefund", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["EmergencyRefund"].Inputs {
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

func (c *Codec) EmergencyYieldWithdrawnLogHash() []byte {
	return c.abi.Events["EmergencyYieldWithdrawn"].ID.Bytes()
}

func (c *Codec) EncodeEmergencyYieldWithdrawnTopics(
	evt abi.Event,
	values []EmergencyYieldWithdrawnTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeEmergencyYieldWithdrawn decodes a log into a EmergencyYieldWithdrawn struct.
func (c *Codec) DecodeEmergencyYieldWithdrawn(log *evm.Log) (*EmergencyYieldWithdrawnDecoded, error) {
	event := new(EmergencyYieldWithdrawnDecoded)
	if err := c.abi.UnpackIntoInterface(event, "EmergencyYieldWithdrawn", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["EmergencyYieldWithdrawn"].Inputs {
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

func (c *Codec) MatchLockedLogHash() []byte {
	return c.abi.Events["MatchLocked"].ID.Bytes()
}

func (c *Codec) EncodeMatchLockedTopics(
	evt abi.Event,
	values []MatchLockedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeMatchLocked decodes a log into a MatchLocked struct.
func (c *Codec) DecodeMatchLocked(log *evm.Log) (*MatchLockedDecoded, error) {
	event := new(MatchLockedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MatchLocked", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["MatchLocked"].Inputs {
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

func (c *Codec) MatchResolvedLogHash() []byte {
	return c.abi.Events["MatchResolved"].ID.Bytes()
}

func (c *Codec) EncodeMatchResolvedTopics(
	evt abi.Event,
	values []MatchResolvedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeMatchResolved decodes a log into a MatchResolved struct.
func (c *Codec) DecodeMatchResolved(log *evm.Log) (*MatchResolvedDecoded, error) {
	event := new(MatchResolvedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MatchResolved", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["MatchResolved"].Inputs {
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

func (c *Codec) PausedLogHash() []byte {
	return c.abi.Events["Paused"].ID.Bytes()
}

func (c *Codec) EncodePausedTopics(
	evt abi.Event,
	values []PausedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodePaused decodes a log into a Paused struct.
func (c *Codec) DecodePaused(log *evm.Log) (*PausedDecoded, error) {
	event := new(PausedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Paused", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Paused"].Inputs {
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

func (c *Codec) UnpausedLogHash() []byte {
	return c.abi.Events["Unpaused"].ID.Bytes()
}

func (c *Codec) EncodeUnpausedTopics(
	evt abi.Event,
	values []UnpausedTopics,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeUnpaused decodes a log into a Unpaused struct.
func (c *Codec) DecodeUnpaused(log *evm.Log) (*UnpausedDecoded, error) {
	event := new(UnpausedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Unpaused", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Unpaused"].Inputs {
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

func (c VaultMatch) MATCHID(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeMATCHIDMethodCall()
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
		return c.Codec.DecodeMATCHIDMethodOutput(response.Data)
	})

}

func (c VaultMatch) MORPHOVAULT(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeMORPHOVAULTMethodCall()
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
		return c.Codec.DecodeMORPHOVAULTMethodOutput(response.Data)
	})

}

func (c VaultMatch) USDC(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeUSDCMethodCall()
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
		return c.Codec.DecodeUSDCMethodOutput(response.Data)
	})

}

func (c VaultMatch) GetExpectedPayout(
	runtime cre.Runtime,
	args GetExpectedPayoutInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeGetExpectedPayoutMethodCall(args)
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
		return c.Codec.DecodeGetExpectedPayoutMethodOutput(response.Data)
	})

}

func (c VaultMatch) GetTotalDeposits(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeGetTotalDepositsMethodCall()
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
		return c.Codec.DecodeGetTotalDepositsMethodOutput(response.Data)
	})

}

func (c VaultMatch) GetUserTotalDeposit(
	runtime cre.Runtime,
	args GetUserTotalDepositInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeGetUserTotalDepositMethodCall(args)
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
		return c.Codec.DecodeGetUserTotalDepositMethodOutput(response.Data)
	})

}

func (c VaultMatch) GetYieldBalance(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeGetYieldBalanceMethodCall()
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
		return c.Codec.DecodeGetYieldBalanceMethodOutput(response.Data)
	})

}

func (c VaultMatch) HasClaimed(
	runtime cre.Runtime,
	args HasClaimedInput,
	blockNumber *big.Int,
) cre.Promise[bool] {
	calldata, err := c.Codec.EncodeHasClaimedMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[bool](*new(bool), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (bool, error) {
		return c.Codec.DecodeHasClaimedMethodOutput(response.Data)
	})

}

func (c VaultMatch) Oracle(
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

func (c VaultMatch) Owner(
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

func (c VaultMatch) Paused(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[bool] {
	calldata, err := c.Codec.EncodePausedMethodCall()
	if err != nil {
		return cre.PromiseFromResult[bool](*new(bool), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (bool, error) {
		return c.Codec.DecodePausedMethodOutput(response.Data)
	})

}

func (c VaultMatch) Status(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[uint8] {
	calldata, err := c.Codec.EncodeStatusMethodCall()
	if err != nil {
		return cre.PromiseFromResult[uint8](*new(uint8), err)
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

func (c VaultMatch) TeamAName(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[string] {
	calldata, err := c.Codec.EncodeTeamANameMethodCall()
	if err != nil {
		return cre.PromiseFromResult[string](*new(string), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (string, error) {
		return c.Codec.DecodeTeamANameMethodOutput(response.Data)
	})

}

func (c VaultMatch) TeamBName(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[string] {
	calldata, err := c.Codec.EncodeTeamBNameMethodCall()
	if err != nil {
		return cre.PromiseFromResult[string](*new(string), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (string, error) {
		return c.Codec.DecodeTeamBNameMethodOutput(response.Data)
	})

}

func (c VaultMatch) TotalTeamA(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeTotalTeamAMethodCall()
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
		return c.Codec.DecodeTotalTeamAMethodOutput(response.Data)
	})

}

func (c VaultMatch) TotalTeamB(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeTotalTeamBMethodCall()
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
		return c.Codec.DecodeTotalTeamBMethodOutput(response.Data)
	})

}

func (c VaultMatch) TotalYield(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeTotalYieldMethodCall()
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
		return c.Codec.DecodeTotalYieldMethodOutput(response.Data)
	})

}

func (c VaultMatch) UserDeposits(
	runtime cre.Runtime,
	args UserDepositsInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeUserDepositsMethodCall(args)
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
		return c.Codec.DecodeUserDepositsMethodOutput(response.Data)
	})

}

func (c VaultMatch) Winner(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[uint8] {
	calldata, err := c.Codec.EncodeWinnerMethodCall()
	if err != nil {
		return cre.PromiseFromResult[uint8](*new(uint8), err)
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
		return c.Codec.DecodeWinnerMethodOutput(response.Data)
	})

}

func (c VaultMatch) WriteReport(
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

func (c *VaultMatch) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	default:
		return nil, errors.New("unknown error selector")
	}
}

// ClaimedTrigger wraps the raw log trigger and provides decoded ClaimedDecoded data
type ClaimedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
}

// Adapt method that decodes the log into Claimed data
func (t *ClaimedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[ClaimedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeClaimed(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Claimed log: %w", err)
	}

	return &bindings.DecodedLog[ClaimedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *VaultMatch) LogTriggerClaimedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []ClaimedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[ClaimedDecoded]], error) {
	event := c.ABI.Events["Claimed"]
	topics, err := c.Codec.EncodeClaimedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Claimed: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &ClaimedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *VaultMatch) FilterLogsClaimed(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.ClaimedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// DepositedTrigger wraps the raw log trigger and provides decoded DepositedDecoded data
type DepositedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
}

// Adapt method that decodes the log into Deposited data
func (t *DepositedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[DepositedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeDeposited(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Deposited log: %w", err)
	}

	return &bindings.DecodedLog[DepositedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *VaultMatch) LogTriggerDepositedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DepositedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[DepositedDecoded]], error) {
	event := c.ABI.Events["Deposited"]
	topics, err := c.Codec.EncodeDepositedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Deposited: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &DepositedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *VaultMatch) FilterLogsDeposited(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DepositedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// EmergencyRefundTrigger wraps the raw log trigger and provides decoded EmergencyRefundDecoded data
type EmergencyRefundTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
}

// Adapt method that decodes the log into EmergencyRefund data
func (t *EmergencyRefundTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[EmergencyRefundDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeEmergencyRefund(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode EmergencyRefund log: %w", err)
	}

	return &bindings.DecodedLog[EmergencyRefundDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *VaultMatch) LogTriggerEmergencyRefundLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []EmergencyRefundTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[EmergencyRefundDecoded]], error) {
	event := c.ABI.Events["EmergencyRefund"]
	topics, err := c.Codec.EncodeEmergencyRefundTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for EmergencyRefund: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &EmergencyRefundTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *VaultMatch) FilterLogsEmergencyRefund(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.EmergencyRefundLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// EmergencyYieldWithdrawnTrigger wraps the raw log trigger and provides decoded EmergencyYieldWithdrawnDecoded data
type EmergencyYieldWithdrawnTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
}

// Adapt method that decodes the log into EmergencyYieldWithdrawn data
func (t *EmergencyYieldWithdrawnTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[EmergencyYieldWithdrawnDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeEmergencyYieldWithdrawn(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode EmergencyYieldWithdrawn log: %w", err)
	}

	return &bindings.DecodedLog[EmergencyYieldWithdrawnDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *VaultMatch) LogTriggerEmergencyYieldWithdrawnLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []EmergencyYieldWithdrawnTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[EmergencyYieldWithdrawnDecoded]], error) {
	event := c.ABI.Events["EmergencyYieldWithdrawn"]
	topics, err := c.Codec.EncodeEmergencyYieldWithdrawnTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for EmergencyYieldWithdrawn: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &EmergencyYieldWithdrawnTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *VaultMatch) FilterLogsEmergencyYieldWithdrawn(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.EmergencyYieldWithdrawnLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// MatchLockedTrigger wraps the raw log trigger and provides decoded MatchLockedDecoded data
type MatchLockedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
}

// Adapt method that decodes the log into MatchLocked data
func (t *MatchLockedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[MatchLockedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeMatchLocked(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MatchLocked log: %w", err)
	}

	return &bindings.DecodedLog[MatchLockedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *VaultMatch) LogTriggerMatchLockedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []MatchLockedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[MatchLockedDecoded]], error) {
	event := c.ABI.Events["MatchLocked"]
	topics, err := c.Codec.EncodeMatchLockedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for MatchLocked: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &MatchLockedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *VaultMatch) FilterLogsMatchLocked(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.MatchLockedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// MatchResolvedTrigger wraps the raw log trigger and provides decoded MatchResolvedDecoded data
type MatchResolvedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
}

// Adapt method that decodes the log into MatchResolved data
func (t *MatchResolvedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[MatchResolvedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeMatchResolved(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MatchResolved log: %w", err)
	}

	return &bindings.DecodedLog[MatchResolvedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *VaultMatch) LogTriggerMatchResolvedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []MatchResolvedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[MatchResolvedDecoded]], error) {
	event := c.ABI.Events["MatchResolved"]
	topics, err := c.Codec.EncodeMatchResolvedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for MatchResolved: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &MatchResolvedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *VaultMatch) FilterLogsMatchResolved(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.MatchResolvedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OracleUpdatedTrigger wraps the raw log trigger and provides decoded OracleUpdatedDecoded data
type OracleUpdatedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
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

func (c *VaultMatch) LogTriggerOracleUpdatedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OracleUpdatedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OracleUpdatedDecoded]], error) {
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

func (c *VaultMatch) FilterLogsOracleUpdated(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
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

func (c *VaultMatch) LogTriggerOwnershipTransferredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipTransferredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipTransferredDecoded]], error) {
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

func (c *VaultMatch) FilterLogsOwnershipTransferred(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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

// PausedTrigger wraps the raw log trigger and provides decoded PausedDecoded data
type PausedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
}

// Adapt method that decodes the log into Paused data
func (t *PausedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[PausedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodePaused(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Paused log: %w", err)
	}

	return &bindings.DecodedLog[PausedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *VaultMatch) LogTriggerPausedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []PausedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[PausedDecoded]], error) {
	event := c.ABI.Events["Paused"]
	topics, err := c.Codec.EncodePausedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Paused: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &PausedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *VaultMatch) FilterLogsPaused(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.PausedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// UnpausedTrigger wraps the raw log trigger and provides decoded UnpausedDecoded data
type UnpausedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]             // Embed the raw trigger
	contract                        *VaultMatch // Keep reference for decoding
}

// Adapt method that decodes the log into Unpaused data
func (t *UnpausedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[UnpausedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeUnpaused(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Unpaused log: %w", err)
	}

	return &bindings.DecodedLog[UnpausedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *VaultMatch) LogTriggerUnpausedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []UnpausedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[UnpausedDecoded]], error) {
	event := c.ABI.Events["Unpaused"]
	topics, err := c.Codec.EncodeUnpausedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Unpaused: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &UnpausedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *VaultMatch) FilterLogsUnpaused(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.UnpausedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// Code generated — DO NOT EDIT.

//go:build !wasip1

package vault_match

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
)

var (
	_ = errors.New
	_ = fmt.Errorf
	_ = big.NewInt
	_ = common.Big1
)

// VaultMatchMock is a mock implementation of VaultMatch for testing.
type VaultMatchMock struct {
	MATCHID                 func() (*big.Int, error)
	MORPHOVAULT             func() (common.Address, error)
	USDC                    func() (common.Address, error)
	GetExpectedAuthor       func() (common.Address, error)
	GetExpectedPayout       func(GetExpectedPayoutInput) (*big.Int, error)
	GetExpectedWorkflowId   func() ([32]byte, error)
	GetExpectedWorkflowName func() ([10]byte, error)
	GetForwarderAddress     func() (common.Address, error)
	GetTotalDeposits        func() (*big.Int, error)
	GetUserTotalDeposit     func(GetUserTotalDepositInput) (*big.Int, error)
	GetYieldBalance         func() (*big.Int, error)
	HasClaimed              func(HasClaimedInput) (bool, error)
	Oracle                  func() (common.Address, error)
	Owner                   func() (common.Address, error)
	Paused                  func() (bool, error)
	Status                  func() (uint8, error)
	SupportsInterface       func(SupportsInterfaceInput) (bool, error)
	TeamAName               func() (string, error)
	TeamBName               func() (string, error)
	TotalTeamA              func() (*big.Int, error)
	TotalTeamB              func() (*big.Int, error)
	TotalYield              func() (*big.Int, error)
	UserDeposits            func(UserDepositsInput) (*big.Int, error)
	Winner                  func() (uint8, error)
}

// NewVaultMatchMock creates a new VaultMatchMock for testing.
func NewVaultMatchMock(address common.Address, clientMock *evmmock.ClientCapability) *VaultMatchMock {
	mock := &VaultMatchMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["MATCH_ID"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.MATCHID == nil {
				return nil, errors.New("MATCH_ID method not mocked")
			}
			result, err := mock.MATCHID()
			if err != nil {
				return nil, err
			}
			return abi.Methods["MATCH_ID"].Outputs.Pack(result)
		},
		string(abi.Methods["MORPHO_VAULT"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.MORPHOVAULT == nil {
				return nil, errors.New("MORPHO_VAULT method not mocked")
			}
			result, err := mock.MORPHOVAULT()
			if err != nil {
				return nil, err
			}
			return abi.Methods["MORPHO_VAULT"].Outputs.Pack(result)
		},
		string(abi.Methods["USDC"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.USDC == nil {
				return nil, errors.New("USDC method not mocked")
			}
			result, err := mock.USDC()
			if err != nil {
				return nil, err
			}
			return abi.Methods["USDC"].Outputs.Pack(result)
		},
		string(abi.Methods["getExpectedAuthor"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetExpectedAuthor == nil {
				return nil, errors.New("getExpectedAuthor method not mocked")
			}
			result, err := mock.GetExpectedAuthor()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getExpectedAuthor"].Outputs.Pack(result)
		},
		string(abi.Methods["getExpectedPayout"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetExpectedPayout == nil {
				return nil, errors.New("getExpectedPayout method not mocked")
			}
			inputs := abi.Methods["getExpectedPayout"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetExpectedPayoutInput{
				User: values[0].(common.Address),
			}

			result, err := mock.GetExpectedPayout(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getExpectedPayout"].Outputs.Pack(result)
		},
		string(abi.Methods["getExpectedWorkflowId"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetExpectedWorkflowId == nil {
				return nil, errors.New("getExpectedWorkflowId method not mocked")
			}
			result, err := mock.GetExpectedWorkflowId()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getExpectedWorkflowId"].Outputs.Pack(result)
		},
		string(abi.Methods["getExpectedWorkflowName"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetExpectedWorkflowName == nil {
				return nil, errors.New("getExpectedWorkflowName method not mocked")
			}
			result, err := mock.GetExpectedWorkflowName()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getExpectedWorkflowName"].Outputs.Pack(result)
		},
		string(abi.Methods["getForwarderAddress"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetForwarderAddress == nil {
				return nil, errors.New("getForwarderAddress method not mocked")
			}
			result, err := mock.GetForwarderAddress()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getForwarderAddress"].Outputs.Pack(result)
		},
		string(abi.Methods["getTotalDeposits"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetTotalDeposits == nil {
				return nil, errors.New("getTotalDeposits method not mocked")
			}
			result, err := mock.GetTotalDeposits()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getTotalDeposits"].Outputs.Pack(result)
		},
		string(abi.Methods["getUserTotalDeposit"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetUserTotalDeposit == nil {
				return nil, errors.New("getUserTotalDeposit method not mocked")
			}
			inputs := abi.Methods["getUserTotalDeposit"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetUserTotalDepositInput{
				User: values[0].(common.Address),
			}

			result, err := mock.GetUserTotalDeposit(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getUserTotalDeposit"].Outputs.Pack(result)
		},
		string(abi.Methods["getYieldBalance"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetYieldBalance == nil {
				return nil, errors.New("getYieldBalance method not mocked")
			}
			result, err := mock.GetYieldBalance()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getYieldBalance"].Outputs.Pack(result)
		},
		string(abi.Methods["hasClaimed"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.HasClaimed == nil {
				return nil, errors.New("hasClaimed method not mocked")
			}
			inputs := abi.Methods["hasClaimed"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := HasClaimedInput{
				Arg0: values[0].(common.Address),
			}

			result, err := mock.HasClaimed(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["hasClaimed"].Outputs.Pack(result)
		},
		string(abi.Methods["oracle"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Oracle == nil {
				return nil, errors.New("oracle method not mocked")
			}
			result, err := mock.Oracle()
			if err != nil {
				return nil, err
			}
			return abi.Methods["oracle"].Outputs.Pack(result)
		},
		string(abi.Methods["owner"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Owner == nil {
				return nil, errors.New("owner method not mocked")
			}
			result, err := mock.Owner()
			if err != nil {
				return nil, err
			}
			return abi.Methods["owner"].Outputs.Pack(result)
		},
		string(abi.Methods["paused"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Paused == nil {
				return nil, errors.New("paused method not mocked")
			}
			result, err := mock.Paused()
			if err != nil {
				return nil, err
			}
			return abi.Methods["paused"].Outputs.Pack(result)
		},
		string(abi.Methods["status"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Status == nil {
				return nil, errors.New("status method not mocked")
			}
			result, err := mock.Status()
			if err != nil {
				return nil, err
			}
			return abi.Methods["status"].Outputs.Pack(result)
		},
		string(abi.Methods["supportsInterface"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.SupportsInterface == nil {
				return nil, errors.New("supportsInterface method not mocked")
			}
			inputs := abi.Methods["supportsInterface"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := SupportsInterfaceInput{
				InterfaceId: values[0].([4]byte),
			}

			result, err := mock.SupportsInterface(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["supportsInterface"].Outputs.Pack(result)
		},
		string(abi.Methods["teamAName"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TeamAName == nil {
				return nil, errors.New("teamAName method not mocked")
			}
			result, err := mock.TeamAName()
			if err != nil {
				return nil, err
			}
			return abi.Methods["teamAName"].Outputs.Pack(result)
		},
		string(abi.Methods["teamBName"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TeamBName == nil {
				return nil, errors.New("teamBName method not mocked")
			}
			result, err := mock.TeamBName()
			if err != nil {
				return nil, err
			}
			return abi.Methods["teamBName"].Outputs.Pack(result)
		},
		string(abi.Methods["totalTeamA"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TotalTeamA == nil {
				return nil, errors.New("totalTeamA method not mocked")
			}
			result, err := mock.TotalTeamA()
			if err != nil {
				return nil, err
			}
			return abi.Methods["totalTeamA"].Outputs.Pack(result)
		},
		string(abi.Methods["totalTeamB"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TotalTeamB == nil {
				return nil, errors.New("totalTeamB method not mocked")
			}
			result, err := mock.TotalTeamB()
			if err != nil {
				return nil, err
			}
			return abi.Methods["totalTeamB"].Outputs.Pack(result)
		},
		string(abi.Methods["totalYield"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TotalYield == nil {
				return nil, errors.New("totalYield method not mocked")
			}
			result, err := mock.TotalYield()
			if err != nil {
				return nil, err
			}
			return abi.Methods["totalYield"].Outputs.Pack(result)
		},
		string(abi.Methods["userDeposits"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.UserDeposits == nil {
				return nil, errors.New("userDeposits method not mocked")
			}
			inputs := abi.Methods["userDeposits"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 2 {
				return nil, errors.New("expected 2 input values")
			}

			args := UserDepositsInput{
				Arg0: values[0].(common.Address),
				Arg1: values[1].(uint8),
			}

			result, err := mock.UserDeposits(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["userDeposits"].Outputs.Pack(result)
		},
		string(abi.Methods["winner"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Winner == nil {
				return nil, errors.New("winner method not mocked")
			}
			result, err := mock.Winner()
			if err != nil {
				return nil, err
			}
			return abi.Methods["winner"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}

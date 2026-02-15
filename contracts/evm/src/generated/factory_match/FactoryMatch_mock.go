// Code generated — DO NOT EDIT.

//go:build !wasip1

package factory_match

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

// FactoryMatchMock is a mock implementation of FactoryMatch for testing.
type FactoryMatchMock struct {
	AllVaults    func(AllVaultsInput) (common.Address, error)
	GetVault     func(GetVaultInput) (common.Address, error)
	NextMatchId  func() (*big.Int, error)
	Oracle       func() (common.Address, error)
	Owner        func() (common.Address, error)
	TotalMatches func() (*big.Int, error)
	Vaults       func(VaultsInput) (common.Address, error)
}

// NewFactoryMatchMock creates a new FactoryMatchMock for testing.
func NewFactoryMatchMock(address common.Address, clientMock *evmmock.ClientCapability) *FactoryMatchMock {
	mock := &FactoryMatchMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["allVaults"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.AllVaults == nil {
				return nil, errors.New("allVaults method not mocked")
			}
			inputs := abi.Methods["allVaults"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := AllVaultsInput{
				Arg0: values[0].(*big.Int),
			}

			result, err := mock.AllVaults(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["allVaults"].Outputs.Pack(result)
		},
		string(abi.Methods["getVault"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetVault == nil {
				return nil, errors.New("getVault method not mocked")
			}
			inputs := abi.Methods["getVault"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetVaultInput{
				MatchId: values[0].(*big.Int),
			}

			result, err := mock.GetVault(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getVault"].Outputs.Pack(result)
		},
		string(abi.Methods["nextMatchId"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.NextMatchId == nil {
				return nil, errors.New("nextMatchId method not mocked")
			}
			result, err := mock.NextMatchId()
			if err != nil {
				return nil, err
			}
			return abi.Methods["nextMatchId"].Outputs.Pack(result)
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
		string(abi.Methods["totalMatches"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TotalMatches == nil {
				return nil, errors.New("totalMatches method not mocked")
			}
			result, err := mock.TotalMatches()
			if err != nil {
				return nil, err
			}
			return abi.Methods["totalMatches"].Outputs.Pack(result)
		},
		string(abi.Methods["vaults"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Vaults == nil {
				return nil, errors.New("vaults method not mocked")
			}
			inputs := abi.Methods["vaults"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := VaultsInput{
				Arg0: values[0].(*big.Int),
			}

			result, err := mock.Vaults(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["vaults"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}

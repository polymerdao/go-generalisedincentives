// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generalisedincentives

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AckPacket is an auto generated low-level Go binding around an user-defined struct.
type AckPacket struct {
	Success bool
	Data    []byte
}

// Height is an auto generated low-level Go binding around an user-defined struct.
type Height struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IMessageEscrowStructsIncentiveDescription is an auto generated low-level Go binding around an user-defined struct.
type IMessageEscrowStructsIncentiveDescription struct {
	MaxGasDelivery     *big.Int
	MaxGasAck          *big.Int
	RefundGasTo        common.Address
	PriceOfDeliveryGas *big.Int
	PriceOfAckGas      *big.Int
	TargetDelta        uint64
}

// IbcEndpoint is an auto generated low-level Go binding around an user-defined struct.
type IbcEndpoint struct {
	PortId    string
	ChannelId [32]byte
}

// IbcPacket is an auto generated low-level Go binding around an user-defined struct.
type IbcPacket struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}

// IncentivizedPolymerEscrowMetaData contains all meta data concerning the IncentivizedPolymerEscrow contract.
var IncentivizedPolymerEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"sendLostGasTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dispatcher\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"SEND_LOST_GAS_TO\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bounty\",\"inputs\":[{\"name\":\"fromApplication\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"incentive\",\"type\":\"tuple\",\"internalType\":\"structIMessageEscrowStructs.IncentiveDescription\",\"components\":[{\"name\":\"maxGasDelivery\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxGasAck\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"refundGasTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceOfDeliveryGas\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"priceOfAckGas\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"targetDelta\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"connectedChannels\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertEVMTo65\",\"inputs\":[{\"name\":\"evmAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"dispatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"estimateAdditionalCost\",\"inputs\":[],\"outputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"implementationAddress\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"implementationAddressHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseBounty\",\"inputs\":[{\"name\":\"fromApplication\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deliveryGasPriceIncrease\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"ackGasPriceIncrease\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"isVerifiedMessageHash\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"implementationIdentifier\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageDelivered\",\"inputs\":[{\"name\":\"sourceIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceImplementationIdentifier\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"hasMessageBeenExecuted\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenInit\",\"inputs\":[{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onCloseIbcChannel\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"proofValidPeriod\",\"inputs\":[{\"name\":\"destinationIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"duration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverAck\",\"inputs\":[{\"name\":\"messagingProtocolContext\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rawMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reemitAckMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRemoteImplementation\",\"inputs\":[{\"name\":\"destinationIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"implementation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitMessage\",\"inputs\":[{\"name\":\"destinationIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"incentive\",\"type\":\"tuple\",\"internalType\":\"structIMessageEscrowStructs.IncentiveDescription\",\"components\":[{\"name\":\"maxGasDelivery\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxGasAck\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"refundGasTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceOfDeliveryGas\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"priceOfAckGas\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"targetDelta\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"deadline\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"gasRefund\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportedVersions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"thisBytes65\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timeoutMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerChannelInit\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BountyClaimed\",\"inputs\":[{\"name\":\"destinationImplementation\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"gasSpentOnDestination\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasSpentOnSource\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"destinationRelayerReward\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"sourceRelayerReward\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BountyIncreased\",\"inputs\":[{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newDeliveryGasPrice\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"newAckGasPrice\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BountyPlaced\",\"inputs\":[{\"name\":\"destinationImplementation\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"incentive\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMessageEscrowStructs.IncentiveDescription\",\"components\":[{\"name\":\"maxGasDelivery\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"maxGasAck\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"refundGasTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceOfDeliveryGas\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"priceOfAckGas\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"targetDelta\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageAcked\",\"inputs\":[{\"name\":\"destinationImplementation\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageDelivered\",\"inputs\":[{\"name\":\"sourceImplementation\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageTimedOut\",\"inputs\":[{\"name\":\"destinationImplementation\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoteImplementationSet\",\"inputs\":[{\"name\":\"application\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"implementationAddressHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"implementationAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimeoutInitiated\",\"inputs\":[{\"name\":\"sourceImplementation\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"chainIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"messageIdentifier\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AckHasNotBeenExecuted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotRetryWrongMessage\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ChannelNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DeadlineInPast\",\"inputs\":[{\"name\":\"blocktimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actual\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DeadlineNotPassed\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actual\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DeadlineTooLong\",\"inputs\":[{\"name\":\"maxAllowed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actual\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"FeeRecipientIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ImplementationAddressAlreadySet\",\"inputs\":[{\"name\":\"currentImplementation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"IncorrectValueProvided\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"actual\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"InvalidBytes65Address\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidImplementationAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTimeoutPackage\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MessageAlreadyAcked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageAlreadyBountied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageAlreadyProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageAlreadySpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageHasInvalidContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoImplementationAddressSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonVerifiableMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughGasExecution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughGasProvided\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"actual\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"type\":\"error\",\"name\":\"NotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotImplementedError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RefundGasToIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouteDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SendLostGasToIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedChannelOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"notIbcDispatcher\",\"inputs\":[]}]",
}

// IncentivizedPolymerEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use IncentivizedPolymerEscrowMetaData.ABI instead.
var IncentivizedPolymerEscrowABI = IncentivizedPolymerEscrowMetaData.ABI

// IncentivizedPolymerEscrow is an auto generated Go binding around an Ethereum contract.
type IncentivizedPolymerEscrow struct {
	IncentivizedPolymerEscrowCaller     // Read-only binding to the contract
	IncentivizedPolymerEscrowTransactor // Write-only binding to the contract
	IncentivizedPolymerEscrowFilterer   // Log filterer for contract events
}

// IncentivizedPolymerEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncentivizedPolymerEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncentivizedPolymerEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncentivizedPolymerEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncentivizedPolymerEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncentivizedPolymerEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncentivizedPolymerEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncentivizedPolymerEscrowSession struct {
	Contract     *IncentivizedPolymerEscrow // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IncentivizedPolymerEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncentivizedPolymerEscrowCallerSession struct {
	Contract *IncentivizedPolymerEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IncentivizedPolymerEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncentivizedPolymerEscrowTransactorSession struct {
	Contract     *IncentivizedPolymerEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IncentivizedPolymerEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncentivizedPolymerEscrowRaw struct {
	Contract *IncentivizedPolymerEscrow // Generic contract binding to access the raw methods on
}

// IncentivizedPolymerEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncentivizedPolymerEscrowCallerRaw struct {
	Contract *IncentivizedPolymerEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// IncentivizedPolymerEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncentivizedPolymerEscrowTransactorRaw struct {
	Contract *IncentivizedPolymerEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncentivizedPolymerEscrow creates a new instance of IncentivizedPolymerEscrow, bound to a specific deployed contract.
func NewIncentivizedPolymerEscrow(address common.Address, backend bind.ContractBackend) (*IncentivizedPolymerEscrow, error) {
	contract, err := bindIncentivizedPolymerEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrow{IncentivizedPolymerEscrowCaller: IncentivizedPolymerEscrowCaller{contract: contract}, IncentivizedPolymerEscrowTransactor: IncentivizedPolymerEscrowTransactor{contract: contract}, IncentivizedPolymerEscrowFilterer: IncentivizedPolymerEscrowFilterer{contract: contract}}, nil
}

// NewIncentivizedPolymerEscrowCaller creates a new read-only instance of IncentivizedPolymerEscrow, bound to a specific deployed contract.
func NewIncentivizedPolymerEscrowCaller(address common.Address, caller bind.ContractCaller) (*IncentivizedPolymerEscrowCaller, error) {
	contract, err := bindIncentivizedPolymerEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowCaller{contract: contract}, nil
}

// NewIncentivizedPolymerEscrowTransactor creates a new write-only instance of IncentivizedPolymerEscrow, bound to a specific deployed contract.
func NewIncentivizedPolymerEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*IncentivizedPolymerEscrowTransactor, error) {
	contract, err := bindIncentivizedPolymerEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowTransactor{contract: contract}, nil
}

// NewIncentivizedPolymerEscrowFilterer creates a new log filterer instance of IncentivizedPolymerEscrow, bound to a specific deployed contract.
func NewIncentivizedPolymerEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*IncentivizedPolymerEscrowFilterer, error) {
	contract, err := bindIncentivizedPolymerEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowFilterer{contract: contract}, nil
}

// bindIncentivizedPolymerEscrow binds a generic wrapper to an already deployed contract.
func bindIncentivizedPolymerEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IncentivizedPolymerEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncentivizedPolymerEscrow.Contract.IncentivizedPolymerEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.IncentivizedPolymerEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.IncentivizedPolymerEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncentivizedPolymerEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.contract.Transact(opts, method, params...)
}

// SENDLOSTGASTO is a free data retrieval call binding the contract method 0x8dedc179.
//
// Solidity: function SEND_LOST_GAS_TO() view returns(address)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) SENDLOSTGASTO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "SEND_LOST_GAS_TO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SENDLOSTGASTO is a free data retrieval call binding the contract method 0x8dedc179.
//
// Solidity: function SEND_LOST_GAS_TO() view returns(address)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) SENDLOSTGASTO() (common.Address, error) {
	return _IncentivizedPolymerEscrow.Contract.SENDLOSTGASTO(&_IncentivizedPolymerEscrow.CallOpts)
}

// SENDLOSTGASTO is a free data retrieval call binding the contract method 0x8dedc179.
//
// Solidity: function SEND_LOST_GAS_TO() view returns(address)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) SENDLOSTGASTO() (common.Address, error) {
	return _IncentivizedPolymerEscrow.Contract.SENDLOSTGASTO(&_IncentivizedPolymerEscrow.CallOpts)
}

// Bounty is a free data retrieval call binding the contract method 0xfd4716c8.
//
// Solidity: function bounty(address fromApplication, bytes32 destinationIdentifier, bytes32 messageIdentifier) view returns((uint48,uint48,address,uint96,uint96,uint64) incentive)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) Bounty(opts *bind.CallOpts, fromApplication common.Address, destinationIdentifier [32]byte, messageIdentifier [32]byte) (IMessageEscrowStructsIncentiveDescription, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "bounty", fromApplication, destinationIdentifier, messageIdentifier)

	if err != nil {
		return *new(IMessageEscrowStructsIncentiveDescription), err
	}

	out0 := *abi.ConvertType(out[0], new(IMessageEscrowStructsIncentiveDescription)).(*IMessageEscrowStructsIncentiveDescription)

	return out0, err

}

// Bounty is a free data retrieval call binding the contract method 0xfd4716c8.
//
// Solidity: function bounty(address fromApplication, bytes32 destinationIdentifier, bytes32 messageIdentifier) view returns((uint48,uint48,address,uint96,uint96,uint64) incentive)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) Bounty(fromApplication common.Address, destinationIdentifier [32]byte, messageIdentifier [32]byte) (IMessageEscrowStructsIncentiveDescription, error) {
	return _IncentivizedPolymerEscrow.Contract.Bounty(&_IncentivizedPolymerEscrow.CallOpts, fromApplication, destinationIdentifier, messageIdentifier)
}

// Bounty is a free data retrieval call binding the contract method 0xfd4716c8.
//
// Solidity: function bounty(address fromApplication, bytes32 destinationIdentifier, bytes32 messageIdentifier) view returns((uint48,uint48,address,uint96,uint96,uint64) incentive)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) Bounty(fromApplication common.Address, destinationIdentifier [32]byte, messageIdentifier [32]byte) (IMessageEscrowStructsIncentiveDescription, error) {
	return _IncentivizedPolymerEscrow.Contract.Bounty(&_IncentivizedPolymerEscrow.CallOpts, fromApplication, destinationIdentifier, messageIdentifier)
}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) ConnectedChannels(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "connectedChannels", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) ConnectedChannels(arg0 *big.Int) ([32]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ConnectedChannels(&_IncentivizedPolymerEscrow.CallOpts, arg0)
}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) ConnectedChannels(arg0 *big.Int) ([32]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ConnectedChannels(&_IncentivizedPolymerEscrow.CallOpts, arg0)
}

// ConvertEVMTo65 is a free data retrieval call binding the contract method 0xea901f1a.
//
// Solidity: function convertEVMTo65(address evmAddress) pure returns(bytes)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) ConvertEVMTo65(opts *bind.CallOpts, evmAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "convertEVMTo65", evmAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConvertEVMTo65 is a free data retrieval call binding the contract method 0xea901f1a.
//
// Solidity: function convertEVMTo65(address evmAddress) pure returns(bytes)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) ConvertEVMTo65(evmAddress common.Address) ([]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ConvertEVMTo65(&_IncentivizedPolymerEscrow.CallOpts, evmAddress)
}

// ConvertEVMTo65 is a free data retrieval call binding the contract method 0xea901f1a.
//
// Solidity: function convertEVMTo65(address evmAddress) pure returns(bytes)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) ConvertEVMTo65(evmAddress common.Address) ([]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ConvertEVMTo65(&_IncentivizedPolymerEscrow.CallOpts, evmAddress)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) Dispatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "dispatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) Dispatcher() (common.Address, error) {
	return _IncentivizedPolymerEscrow.Contract.Dispatcher(&_IncentivizedPolymerEscrow.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) Dispatcher() (common.Address, error) {
	return _IncentivizedPolymerEscrow.Contract.Dispatcher(&_IncentivizedPolymerEscrow.CallOpts)
}

// EstimateAdditionalCost is a free data retrieval call binding the contract method 0x92006e73.
//
// Solidity: function estimateAdditionalCost() pure returns(address asset, uint256 amount)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) EstimateAdditionalCost(opts *bind.CallOpts) (struct {
	Asset  common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "estimateAdditionalCost")

	outstruct := new(struct {
		Asset  common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Asset = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateAdditionalCost is a free data retrieval call binding the contract method 0x92006e73.
//
// Solidity: function estimateAdditionalCost() pure returns(address asset, uint256 amount)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) EstimateAdditionalCost() (struct {
	Asset  common.Address
	Amount *big.Int
}, error) {
	return _IncentivizedPolymerEscrow.Contract.EstimateAdditionalCost(&_IncentivizedPolymerEscrow.CallOpts)
}

// EstimateAdditionalCost is a free data retrieval call binding the contract method 0x92006e73.
//
// Solidity: function estimateAdditionalCost() pure returns(address asset, uint256 amount)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) EstimateAdditionalCost() (struct {
	Asset  common.Address
	Amount *big.Int
}, error) {
	return _IncentivizedPolymerEscrow.Contract.EstimateAdditionalCost(&_IncentivizedPolymerEscrow.CallOpts)
}

// ImplementationAddress is a free data retrieval call binding the contract method 0x4993ef33.
//
// Solidity: function implementationAddress(address , bytes32 ) view returns(bytes)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) ImplementationAddress(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "implementationAddress", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ImplementationAddress is a free data retrieval call binding the contract method 0x4993ef33.
//
// Solidity: function implementationAddress(address , bytes32 ) view returns(bytes)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) ImplementationAddress(arg0 common.Address, arg1 [32]byte) ([]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ImplementationAddress(&_IncentivizedPolymerEscrow.CallOpts, arg0, arg1)
}

// ImplementationAddress is a free data retrieval call binding the contract method 0x4993ef33.
//
// Solidity: function implementationAddress(address , bytes32 ) view returns(bytes)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) ImplementationAddress(arg0 common.Address, arg1 [32]byte) ([]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ImplementationAddress(&_IncentivizedPolymerEscrow.CallOpts, arg0, arg1)
}

// ImplementationAddressHash is a free data retrieval call binding the contract method 0xdd6487fc.
//
// Solidity: function implementationAddressHash(address , bytes32 ) view returns(bytes32)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) ImplementationAddressHash(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "implementationAddressHash", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImplementationAddressHash is a free data retrieval call binding the contract method 0xdd6487fc.
//
// Solidity: function implementationAddressHash(address , bytes32 ) view returns(bytes32)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) ImplementationAddressHash(arg0 common.Address, arg1 [32]byte) ([32]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ImplementationAddressHash(&_IncentivizedPolymerEscrow.CallOpts, arg0, arg1)
}

// ImplementationAddressHash is a free data retrieval call binding the contract method 0xdd6487fc.
//
// Solidity: function implementationAddressHash(address , bytes32 ) view returns(bytes32)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) ImplementationAddressHash(arg0 common.Address, arg1 [32]byte) ([32]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ImplementationAddressHash(&_IncentivizedPolymerEscrow.CallOpts, arg0, arg1)
}

// IsVerifiedMessageHash is a free data retrieval call binding the contract method 0xaa576504.
//
// Solidity: function isVerifiedMessageHash(bytes32 ) view returns(bytes32 chainIdentifier, bytes implementationIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) IsVerifiedMessageHash(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ChainIdentifier          [32]byte
	ImplementationIdentifier []byte
}, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "isVerifiedMessageHash", arg0)

	outstruct := new(struct {
		ChainIdentifier          [32]byte
		ImplementationIdentifier []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainIdentifier = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ImplementationIdentifier = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// IsVerifiedMessageHash is a free data retrieval call binding the contract method 0xaa576504.
//
// Solidity: function isVerifiedMessageHash(bytes32 ) view returns(bytes32 chainIdentifier, bytes implementationIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) IsVerifiedMessageHash(arg0 [32]byte) (struct {
	ChainIdentifier          [32]byte
	ImplementationIdentifier []byte
}, error) {
	return _IncentivizedPolymerEscrow.Contract.IsVerifiedMessageHash(&_IncentivizedPolymerEscrow.CallOpts, arg0)
}

// IsVerifiedMessageHash is a free data retrieval call binding the contract method 0xaa576504.
//
// Solidity: function isVerifiedMessageHash(bytes32 ) view returns(bytes32 chainIdentifier, bytes implementationIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) IsVerifiedMessageHash(arg0 [32]byte) (struct {
	ChainIdentifier          [32]byte
	ImplementationIdentifier []byte
}, error) {
	return _IncentivizedPolymerEscrow.Contract.IsVerifiedMessageHash(&_IncentivizedPolymerEscrow.CallOpts, arg0)
}

// MessageDelivered is a free data retrieval call binding the contract method 0x3a24ef35.
//
// Solidity: function messageDelivered(bytes32 sourceIdentifier, bytes sourceImplementationIdentifier, bytes32 messageIdentifier) view returns(bytes32 hasMessageBeenExecuted)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) MessageDelivered(opts *bind.CallOpts, sourceIdentifier [32]byte, sourceImplementationIdentifier []byte, messageIdentifier [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "messageDelivered", sourceIdentifier, sourceImplementationIdentifier, messageIdentifier)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageDelivered is a free data retrieval call binding the contract method 0x3a24ef35.
//
// Solidity: function messageDelivered(bytes32 sourceIdentifier, bytes sourceImplementationIdentifier, bytes32 messageIdentifier) view returns(bytes32 hasMessageBeenExecuted)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) MessageDelivered(sourceIdentifier [32]byte, sourceImplementationIdentifier []byte, messageIdentifier [32]byte) ([32]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.MessageDelivered(&_IncentivizedPolymerEscrow.CallOpts, sourceIdentifier, sourceImplementationIdentifier, messageIdentifier)
}

// MessageDelivered is a free data retrieval call binding the contract method 0x3a24ef35.
//
// Solidity: function messageDelivered(bytes32 sourceIdentifier, bytes sourceImplementationIdentifier, bytes32 messageIdentifier) view returns(bytes32 hasMessageBeenExecuted)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) MessageDelivered(sourceIdentifier [32]byte, sourceImplementationIdentifier []byte, messageIdentifier [32]byte) ([32]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.MessageDelivered(&_IncentivizedPolymerEscrow.CallOpts, sourceIdentifier, sourceImplementationIdentifier, messageIdentifier)
}

// OnChanOpenInit is a free data retrieval call binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 ordering, string[] , string , string version) view returns(string selectedVersion)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) OnChanOpenInit(opts *bind.CallOpts, ordering uint8, arg1 []string, arg2 string, version string) (string, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "onChanOpenInit", ordering, arg1, arg2, version)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OnChanOpenInit is a free data retrieval call binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 ordering, string[] , string , string version) view returns(string selectedVersion)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) OnChanOpenInit(ordering uint8, arg1 []string, arg2 string, version string) (string, error) {
	return _IncentivizedPolymerEscrow.Contract.OnChanOpenInit(&_IncentivizedPolymerEscrow.CallOpts, ordering, arg1, arg2, version)
}

// OnChanOpenInit is a free data retrieval call binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 ordering, string[] , string , string version) view returns(string selectedVersion)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) OnChanOpenInit(ordering uint8, arg1 []string, arg2 string, version string) (string, error) {
	return _IncentivizedPolymerEscrow.Contract.OnChanOpenInit(&_IncentivizedPolymerEscrow.CallOpts, ordering, arg1, arg2, version)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) Owner() (common.Address, error) {
	return _IncentivizedPolymerEscrow.Contract.Owner(&_IncentivizedPolymerEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) Owner() (common.Address, error) {
	return _IncentivizedPolymerEscrow.Contract.Owner(&_IncentivizedPolymerEscrow.CallOpts)
}

// ProofValidPeriod is a free data retrieval call binding the contract method 0x5a640514.
//
// Solidity: function proofValidPeriod(bytes32 destinationIdentifier) view returns(uint64 duration)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) ProofValidPeriod(opts *bind.CallOpts, destinationIdentifier [32]byte) (uint64, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "proofValidPeriod", destinationIdentifier)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ProofValidPeriod is a free data retrieval call binding the contract method 0x5a640514.
//
// Solidity: function proofValidPeriod(bytes32 destinationIdentifier) view returns(uint64 duration)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) ProofValidPeriod(destinationIdentifier [32]byte) (uint64, error) {
	return _IncentivizedPolymerEscrow.Contract.ProofValidPeriod(&_IncentivizedPolymerEscrow.CallOpts, destinationIdentifier)
}

// ProofValidPeriod is a free data retrieval call binding the contract method 0x5a640514.
//
// Solidity: function proofValidPeriod(bytes32 destinationIdentifier) view returns(uint64 duration)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) ProofValidPeriod(destinationIdentifier [32]byte) (uint64, error) {
	return _IncentivizedPolymerEscrow.Contract.ProofValidPeriod(&_IncentivizedPolymerEscrow.CallOpts, destinationIdentifier)
}

// SupportedVersions is a free data retrieval call binding the contract method 0x7d622184.
//
// Solidity: function supportedVersions(uint256 ) view returns(string)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) SupportedVersions(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "supportedVersions", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SupportedVersions is a free data retrieval call binding the contract method 0x7d622184.
//
// Solidity: function supportedVersions(uint256 ) view returns(string)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) SupportedVersions(arg0 *big.Int) (string, error) {
	return _IncentivizedPolymerEscrow.Contract.SupportedVersions(&_IncentivizedPolymerEscrow.CallOpts, arg0)
}

// SupportedVersions is a free data retrieval call binding the contract method 0x7d622184.
//
// Solidity: function supportedVersions(uint256 ) view returns(string)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) SupportedVersions(arg0 *big.Int) (string, error) {
	return _IncentivizedPolymerEscrow.Contract.SupportedVersions(&_IncentivizedPolymerEscrow.CallOpts, arg0)
}

// ThisBytes65 is a free data retrieval call binding the contract method 0xdaf1e167.
//
// Solidity: function thisBytes65() view returns(bytes)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCaller) ThisBytes65(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _IncentivizedPolymerEscrow.contract.Call(opts, &out, "thisBytes65")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ThisBytes65 is a free data retrieval call binding the contract method 0xdaf1e167.
//
// Solidity: function thisBytes65() view returns(bytes)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) ThisBytes65() ([]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ThisBytes65(&_IncentivizedPolymerEscrow.CallOpts)
}

// ThisBytes65 is a free data retrieval call binding the contract method 0xdaf1e167.
//
// Solidity: function thisBytes65() view returns(bytes)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowCallerSession) ThisBytes65() ([]byte, error) {
	return _IncentivizedPolymerEscrow.Contract.ThisBytes65(&_IncentivizedPolymerEscrow.CallOpts)
}

// IncreaseBounty is a paid mutator transaction binding the contract method 0x5d758e22.
//
// Solidity: function increaseBounty(address fromApplication, bytes32 destinationIdentifier, bytes32 messageIdentifier, uint96 deliveryGasPriceIncrease, uint96 ackGasPriceIncrease) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) IncreaseBounty(opts *bind.TransactOpts, fromApplication common.Address, destinationIdentifier [32]byte, messageIdentifier [32]byte, deliveryGasPriceIncrease *big.Int, ackGasPriceIncrease *big.Int) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "increaseBounty", fromApplication, destinationIdentifier, messageIdentifier, deliveryGasPriceIncrease, ackGasPriceIncrease)
}

// IncreaseBounty is a paid mutator transaction binding the contract method 0x5d758e22.
//
// Solidity: function increaseBounty(address fromApplication, bytes32 destinationIdentifier, bytes32 messageIdentifier, uint96 deliveryGasPriceIncrease, uint96 ackGasPriceIncrease) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) IncreaseBounty(fromApplication common.Address, destinationIdentifier [32]byte, messageIdentifier [32]byte, deliveryGasPriceIncrease *big.Int, ackGasPriceIncrease *big.Int) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.IncreaseBounty(&_IncentivizedPolymerEscrow.TransactOpts, fromApplication, destinationIdentifier, messageIdentifier, deliveryGasPriceIncrease, ackGasPriceIncrease)
}

// IncreaseBounty is a paid mutator transaction binding the contract method 0x5d758e22.
//
// Solidity: function increaseBounty(address fromApplication, bytes32 destinationIdentifier, bytes32 messageIdentifier, uint96 deliveryGasPriceIncrease, uint96 ackGasPriceIncrease) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) IncreaseBounty(fromApplication common.Address, destinationIdentifier [32]byte, messageIdentifier [32]byte, deliveryGasPriceIncrease *big.Int, ackGasPriceIncrease *big.Int) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.IncreaseBounty(&_IncentivizedPolymerEscrow.TransactOpts, fromApplication, destinationIdentifier, messageIdentifier, deliveryGasPriceIncrease, ackGasPriceIncrease)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "onAcknowledgementPacket", packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnAcknowledgementPacket(&_IncentivizedPolymerEscrow.TransactOpts, packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnAcknowledgementPacket(&_IncentivizedPolymerEscrow.TransactOpts, packet, ack)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) OnChanOpenAck(opts *bind.TransactOpts, channelId [32]byte, arg1 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "onChanOpenAck", channelId, arg1, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnChanOpenAck(&_IncentivizedPolymerEscrow.TransactOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnChanOpenAck(&_IncentivizedPolymerEscrow.TransactOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "onChanOpenConfirm", channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnChanOpenConfirm(&_IncentivizedPolymerEscrow.TransactOpts, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnChanOpenConfirm(&_IncentivizedPolymerEscrow.TransactOpts, channelId)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 ordering, string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) returns(string selectedVersion)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) OnChanOpenTry(opts *bind.TransactOpts, ordering uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "onChanOpenTry", ordering, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 ordering, string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) returns(string selectedVersion)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) OnChanOpenTry(ordering uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnChanOpenTry(&_IncentivizedPolymerEscrow.TransactOpts, ordering, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 ordering, string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) returns(string selectedVersion)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) OnChanOpenTry(ordering uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnChanOpenTry(&_IncentivizedPolymerEscrow.TransactOpts, ordering, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnCloseIbcChannel is a paid mutator transaction binding the contract method 0x5fe39e0d.
//
// Solidity: function onCloseIbcChannel(bytes32 channelId, string , bytes32 ) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) OnCloseIbcChannel(opts *bind.TransactOpts, channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "onCloseIbcChannel", channelId, arg1, arg2)
}

// OnCloseIbcChannel is a paid mutator transaction binding the contract method 0x5fe39e0d.
//
// Solidity: function onCloseIbcChannel(bytes32 channelId, string , bytes32 ) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) OnCloseIbcChannel(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnCloseIbcChannel(&_IncentivizedPolymerEscrow.TransactOpts, channelId, arg1, arg2)
}

// OnCloseIbcChannel is a paid mutator transaction binding the contract method 0x5fe39e0d.
//
// Solidity: function onCloseIbcChannel(bytes32 channelId, string , bytes32 ) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) OnCloseIbcChannel(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnCloseIbcChannel(&_IncentivizedPolymerEscrow.TransactOpts, channelId, arg1, arg2)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes))
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) OnRecvPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes))
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnRecvPacket(&_IncentivizedPolymerEscrow.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes))
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnRecvPacket(&_IncentivizedPolymerEscrow.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "onTimeoutPacket", packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnTimeoutPacket(&_IncentivizedPolymerEscrow.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.OnTimeoutPacket(&_IncentivizedPolymerEscrow.TransactOpts, packet)
}

// ProcessPacket is a paid mutator transaction binding the contract method 0x53391d3f.
//
// Solidity: function processPacket(bytes , bytes , bytes32 ) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) ProcessPacket(opts *bind.TransactOpts, arg0 []byte, arg1 []byte, arg2 [32]byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "processPacket", arg0, arg1, arg2)
}

// ProcessPacket is a paid mutator transaction binding the contract method 0x53391d3f.
//
// Solidity: function processPacket(bytes , bytes , bytes32 ) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) ProcessPacket(arg0 []byte, arg1 []byte, arg2 [32]byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.ProcessPacket(&_IncentivizedPolymerEscrow.TransactOpts, arg0, arg1, arg2)
}

// ProcessPacket is a paid mutator transaction binding the contract method 0x53391d3f.
//
// Solidity: function processPacket(bytes , bytes , bytes32 ) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) ProcessPacket(arg0 []byte, arg1 []byte, arg2 [32]byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.ProcessPacket(&_IncentivizedPolymerEscrow.TransactOpts, arg0, arg1, arg2)
}

// RecoverAck is a paid mutator transaction binding the contract method 0xdfe1aef9.
//
// Solidity: function recoverAck(bytes messagingProtocolContext, bytes rawMessage) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) RecoverAck(opts *bind.TransactOpts, messagingProtocolContext []byte, rawMessage []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "recoverAck", messagingProtocolContext, rawMessage)
}

// RecoverAck is a paid mutator transaction binding the contract method 0xdfe1aef9.
//
// Solidity: function recoverAck(bytes messagingProtocolContext, bytes rawMessage) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) RecoverAck(messagingProtocolContext []byte, rawMessage []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.RecoverAck(&_IncentivizedPolymerEscrow.TransactOpts, messagingProtocolContext, rawMessage)
}

// RecoverAck is a paid mutator transaction binding the contract method 0xdfe1aef9.
//
// Solidity: function recoverAck(bytes messagingProtocolContext, bytes rawMessage) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) RecoverAck(messagingProtocolContext []byte, rawMessage []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.RecoverAck(&_IncentivizedPolymerEscrow.TransactOpts, messagingProtocolContext, rawMessage)
}

// ReemitAckMessage is a paid mutator transaction binding the contract method 0x8caa89ec.
//
// Solidity: function reemitAckMessage(bytes32 , bytes , bytes ) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) ReemitAckMessage(opts *bind.TransactOpts, arg0 [32]byte, arg1 []byte, arg2 []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "reemitAckMessage", arg0, arg1, arg2)
}

// ReemitAckMessage is a paid mutator transaction binding the contract method 0x8caa89ec.
//
// Solidity: function reemitAckMessage(bytes32 , bytes , bytes ) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) ReemitAckMessage(arg0 [32]byte, arg1 []byte, arg2 []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.ReemitAckMessage(&_IncentivizedPolymerEscrow.TransactOpts, arg0, arg1, arg2)
}

// ReemitAckMessage is a paid mutator transaction binding the contract method 0x8caa89ec.
//
// Solidity: function reemitAckMessage(bytes32 , bytes , bytes ) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) ReemitAckMessage(arg0 [32]byte, arg1 []byte, arg2 []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.ReemitAckMessage(&_IncentivizedPolymerEscrow.TransactOpts, arg0, arg1, arg2)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.RenounceOwnership(&_IncentivizedPolymerEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.RenounceOwnership(&_IncentivizedPolymerEscrow.TransactOpts)
}

// SetRemoteImplementation is a paid mutator transaction binding the contract method 0xffbeacc7.
//
// Solidity: function setRemoteImplementation(bytes32 destinationIdentifier, bytes implementation) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) SetRemoteImplementation(opts *bind.TransactOpts, destinationIdentifier [32]byte, implementation []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "setRemoteImplementation", destinationIdentifier, implementation)
}

// SetRemoteImplementation is a paid mutator transaction binding the contract method 0xffbeacc7.
//
// Solidity: function setRemoteImplementation(bytes32 destinationIdentifier, bytes implementation) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) SetRemoteImplementation(destinationIdentifier [32]byte, implementation []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.SetRemoteImplementation(&_IncentivizedPolymerEscrow.TransactOpts, destinationIdentifier, implementation)
}

// SetRemoteImplementation is a paid mutator transaction binding the contract method 0xffbeacc7.
//
// Solidity: function setRemoteImplementation(bytes32 destinationIdentifier, bytes implementation) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) SetRemoteImplementation(destinationIdentifier [32]byte, implementation []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.SetRemoteImplementation(&_IncentivizedPolymerEscrow.TransactOpts, destinationIdentifier, implementation)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x65a686de.
//
// Solidity: function submitMessage(bytes32 destinationIdentifier, bytes destinationAddress, bytes message, (uint48,uint48,address,uint96,uint96,uint64) incentive, uint64 deadline) payable returns(uint256 gasRefund, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) SubmitMessage(opts *bind.TransactOpts, destinationIdentifier [32]byte, destinationAddress []byte, message []byte, incentive IMessageEscrowStructsIncentiveDescription, deadline uint64) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "submitMessage", destinationIdentifier, destinationAddress, message, incentive, deadline)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x65a686de.
//
// Solidity: function submitMessage(bytes32 destinationIdentifier, bytes destinationAddress, bytes message, (uint48,uint48,address,uint96,uint96,uint64) incentive, uint64 deadline) payable returns(uint256 gasRefund, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) SubmitMessage(destinationIdentifier [32]byte, destinationAddress []byte, message []byte, incentive IMessageEscrowStructsIncentiveDescription, deadline uint64) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.SubmitMessage(&_IncentivizedPolymerEscrow.TransactOpts, destinationIdentifier, destinationAddress, message, incentive, deadline)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x65a686de.
//
// Solidity: function submitMessage(bytes32 destinationIdentifier, bytes destinationAddress, bytes message, (uint48,uint48,address,uint96,uint96,uint64) incentive, uint64 deadline) payable returns(uint256 gasRefund, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) SubmitMessage(destinationIdentifier [32]byte, destinationAddress []byte, message []byte, incentive IMessageEscrowStructsIncentiveDescription, deadline uint64) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.SubmitMessage(&_IncentivizedPolymerEscrow.TransactOpts, destinationIdentifier, destinationAddress, message, incentive, deadline)
}

// TimeoutMessage is a paid mutator transaction binding the contract method 0x3fccfa06.
//
// Solidity: function timeoutMessage(bytes32 , bytes , uint256 , bytes ) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) TimeoutMessage(opts *bind.TransactOpts, arg0 [32]byte, arg1 []byte, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "timeoutMessage", arg0, arg1, arg2, arg3)
}

// TimeoutMessage is a paid mutator transaction binding the contract method 0x3fccfa06.
//
// Solidity: function timeoutMessage(bytes32 , bytes , uint256 , bytes ) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) TimeoutMessage(arg0 [32]byte, arg1 []byte, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.TimeoutMessage(&_IncentivizedPolymerEscrow.TransactOpts, arg0, arg1, arg2, arg3)
}

// TimeoutMessage is a paid mutator transaction binding the contract method 0x3fccfa06.
//
// Solidity: function timeoutMessage(bytes32 , bytes , uint256 , bytes ) payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) TimeoutMessage(arg0 [32]byte, arg1 []byte, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.TimeoutMessage(&_IncentivizedPolymerEscrow.TransactOpts, arg0, arg1, arg2, arg3)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.TransferOwnership(&_IncentivizedPolymerEscrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.TransferOwnership(&_IncentivizedPolymerEscrow.TransactOpts, newOwner)
}

// TriggerChannelInit is a paid mutator transaction binding the contract method 0x7a805598.
//
// Solidity: function triggerChannelInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) TriggerChannelInit(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.Transact(opts, "triggerChannelInit", version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInit is a paid mutator transaction binding the contract method 0x7a805598.
//
// Solidity: function triggerChannelInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) TriggerChannelInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.TriggerChannelInit(&_IncentivizedPolymerEscrow.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInit is a paid mutator transaction binding the contract method 0x7a805598.
//
// Solidity: function triggerChannelInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) TriggerChannelInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.TriggerChannelInit(&_IncentivizedPolymerEscrow.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowSession) Receive() (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.Receive(&_IncentivizedPolymerEscrow.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowTransactorSession) Receive() (*types.Transaction, error) {
	return _IncentivizedPolymerEscrow.Contract.Receive(&_IncentivizedPolymerEscrow.TransactOpts)
}

// IncentivizedPolymerEscrowBountyClaimedIterator is returned from FilterBountyClaimed and is used to iterate over the raw logs and unpacked data for BountyClaimed events raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowBountyClaimedIterator struct {
	Event *IncentivizedPolymerEscrowBountyClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncentivizedPolymerEscrowBountyClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentivizedPolymerEscrowBountyClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncentivizedPolymerEscrowBountyClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncentivizedPolymerEscrowBountyClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentivizedPolymerEscrowBountyClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentivizedPolymerEscrowBountyClaimed represents a BountyClaimed event raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowBountyClaimed struct {
	DestinationImplementation common.Hash
	ChainIdentifier           [32]byte
	MessageIdentifier         [32]byte
	GasSpentOnDestination     uint64
	GasSpentOnSource          uint64
	DestinationRelayerReward  *big.Int
	SourceRelayerReward       *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterBountyClaimed is a free log retrieval operation binding the contract event 0x9f9aef021b9656e38cd2086f7d7ee0f6b6062c8078a049bfecd6356fdbeb1b3c.
//
// Solidity: event BountyClaimed(bytes indexed destinationImplementation, bytes32 chainIdentifier, bytes32 indexed messageIdentifier, uint64 gasSpentOnDestination, uint64 gasSpentOnSource, uint128 destinationRelayerReward, uint128 sourceRelayerReward)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) FilterBountyClaimed(opts *bind.FilterOpts, destinationImplementation [][]byte, messageIdentifier [][32]byte) (*IncentivizedPolymerEscrowBountyClaimedIterator, error) {

	var destinationImplementationRule []interface{}
	for _, destinationImplementationItem := range destinationImplementation {
		destinationImplementationRule = append(destinationImplementationRule, destinationImplementationItem)
	}

	var messageIdentifierRule []interface{}
	for _, messageIdentifierItem := range messageIdentifier {
		messageIdentifierRule = append(messageIdentifierRule, messageIdentifierItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.FilterLogs(opts, "BountyClaimed", destinationImplementationRule, messageIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowBountyClaimedIterator{contract: _IncentivizedPolymerEscrow.contract, event: "BountyClaimed", logs: logs, sub: sub}, nil
}

// WatchBountyClaimed is a free log subscription operation binding the contract event 0x9f9aef021b9656e38cd2086f7d7ee0f6b6062c8078a049bfecd6356fdbeb1b3c.
//
// Solidity: event BountyClaimed(bytes indexed destinationImplementation, bytes32 chainIdentifier, bytes32 indexed messageIdentifier, uint64 gasSpentOnDestination, uint64 gasSpentOnSource, uint128 destinationRelayerReward, uint128 sourceRelayerReward)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) WatchBountyClaimed(opts *bind.WatchOpts, sink chan<- *IncentivizedPolymerEscrowBountyClaimed, destinationImplementation [][]byte, messageIdentifier [][32]byte) (event.Subscription, error) {

	var destinationImplementationRule []interface{}
	for _, destinationImplementationItem := range destinationImplementation {
		destinationImplementationRule = append(destinationImplementationRule, destinationImplementationItem)
	}

	var messageIdentifierRule []interface{}
	for _, messageIdentifierItem := range messageIdentifier {
		messageIdentifierRule = append(messageIdentifierRule, messageIdentifierItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.WatchLogs(opts, "BountyClaimed", destinationImplementationRule, messageIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentivizedPolymerEscrowBountyClaimed)
				if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "BountyClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBountyClaimed is a log parse operation binding the contract event 0x9f9aef021b9656e38cd2086f7d7ee0f6b6062c8078a049bfecd6356fdbeb1b3c.
//
// Solidity: event BountyClaimed(bytes indexed destinationImplementation, bytes32 chainIdentifier, bytes32 indexed messageIdentifier, uint64 gasSpentOnDestination, uint64 gasSpentOnSource, uint128 destinationRelayerReward, uint128 sourceRelayerReward)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) ParseBountyClaimed(log types.Log) (*IncentivizedPolymerEscrowBountyClaimed, error) {
	event := new(IncentivizedPolymerEscrowBountyClaimed)
	if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "BountyClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentivizedPolymerEscrowBountyIncreasedIterator is returned from FilterBountyIncreased and is used to iterate over the raw logs and unpacked data for BountyIncreased events raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowBountyIncreasedIterator struct {
	Event *IncentivizedPolymerEscrowBountyIncreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncentivizedPolymerEscrowBountyIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentivizedPolymerEscrowBountyIncreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncentivizedPolymerEscrowBountyIncreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncentivizedPolymerEscrowBountyIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentivizedPolymerEscrowBountyIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentivizedPolymerEscrowBountyIncreased represents a BountyIncreased event raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowBountyIncreased struct {
	MessageIdentifier   [32]byte
	NewDeliveryGasPrice *big.Int
	NewAckGasPrice      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBountyIncreased is a free log retrieval operation binding the contract event 0x0f9ca48bcb4f4aea18e50df9f1df3c02df5c9c48f05e389a349c8ceb242da581.
//
// Solidity: event BountyIncreased(bytes32 indexed messageIdentifier, uint96 newDeliveryGasPrice, uint96 newAckGasPrice)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) FilterBountyIncreased(opts *bind.FilterOpts, messageIdentifier [][32]byte) (*IncentivizedPolymerEscrowBountyIncreasedIterator, error) {

	var messageIdentifierRule []interface{}
	for _, messageIdentifierItem := range messageIdentifier {
		messageIdentifierRule = append(messageIdentifierRule, messageIdentifierItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.FilterLogs(opts, "BountyIncreased", messageIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowBountyIncreasedIterator{contract: _IncentivizedPolymerEscrow.contract, event: "BountyIncreased", logs: logs, sub: sub}, nil
}

// WatchBountyIncreased is a free log subscription operation binding the contract event 0x0f9ca48bcb4f4aea18e50df9f1df3c02df5c9c48f05e389a349c8ceb242da581.
//
// Solidity: event BountyIncreased(bytes32 indexed messageIdentifier, uint96 newDeliveryGasPrice, uint96 newAckGasPrice)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) WatchBountyIncreased(opts *bind.WatchOpts, sink chan<- *IncentivizedPolymerEscrowBountyIncreased, messageIdentifier [][32]byte) (event.Subscription, error) {

	var messageIdentifierRule []interface{}
	for _, messageIdentifierItem := range messageIdentifier {
		messageIdentifierRule = append(messageIdentifierRule, messageIdentifierItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.WatchLogs(opts, "BountyIncreased", messageIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentivizedPolymerEscrowBountyIncreased)
				if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "BountyIncreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBountyIncreased is a log parse operation binding the contract event 0x0f9ca48bcb4f4aea18e50df9f1df3c02df5c9c48f05e389a349c8ceb242da581.
//
// Solidity: event BountyIncreased(bytes32 indexed messageIdentifier, uint96 newDeliveryGasPrice, uint96 newAckGasPrice)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) ParseBountyIncreased(log types.Log) (*IncentivizedPolymerEscrowBountyIncreased, error) {
	event := new(IncentivizedPolymerEscrowBountyIncreased)
	if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "BountyIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentivizedPolymerEscrowBountyPlacedIterator is returned from FilterBountyPlaced and is used to iterate over the raw logs and unpacked data for BountyPlaced events raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowBountyPlacedIterator struct {
	Event *IncentivizedPolymerEscrowBountyPlaced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncentivizedPolymerEscrowBountyPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentivizedPolymerEscrowBountyPlaced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncentivizedPolymerEscrowBountyPlaced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncentivizedPolymerEscrowBountyPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentivizedPolymerEscrowBountyPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentivizedPolymerEscrowBountyPlaced represents a BountyPlaced event raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowBountyPlaced struct {
	DestinationImplementation common.Hash
	ChainIdentifier           [32]byte
	MessageIdentifier         [32]byte
	Incentive                 IMessageEscrowStructsIncentiveDescription
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterBountyPlaced is a free log retrieval operation binding the contract event 0x1b7379b49b96ce33cfd9daf95718231650c40a2c92da579412e5886011867727.
//
// Solidity: event BountyPlaced(bytes indexed destinationImplementation, bytes32 chainIdentifier, bytes32 indexed messageIdentifier, (uint48,uint48,address,uint96,uint96,uint64) incentive)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) FilterBountyPlaced(opts *bind.FilterOpts, destinationImplementation [][]byte, messageIdentifier [][32]byte) (*IncentivizedPolymerEscrowBountyPlacedIterator, error) {

	var destinationImplementationRule []interface{}
	for _, destinationImplementationItem := range destinationImplementation {
		destinationImplementationRule = append(destinationImplementationRule, destinationImplementationItem)
	}

	var messageIdentifierRule []interface{}
	for _, messageIdentifierItem := range messageIdentifier {
		messageIdentifierRule = append(messageIdentifierRule, messageIdentifierItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.FilterLogs(opts, "BountyPlaced", destinationImplementationRule, messageIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowBountyPlacedIterator{contract: _IncentivizedPolymerEscrow.contract, event: "BountyPlaced", logs: logs, sub: sub}, nil
}

// WatchBountyPlaced is a free log subscription operation binding the contract event 0x1b7379b49b96ce33cfd9daf95718231650c40a2c92da579412e5886011867727.
//
// Solidity: event BountyPlaced(bytes indexed destinationImplementation, bytes32 chainIdentifier, bytes32 indexed messageIdentifier, (uint48,uint48,address,uint96,uint96,uint64) incentive)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) WatchBountyPlaced(opts *bind.WatchOpts, sink chan<- *IncentivizedPolymerEscrowBountyPlaced, destinationImplementation [][]byte, messageIdentifier [][32]byte) (event.Subscription, error) {

	var destinationImplementationRule []interface{}
	for _, destinationImplementationItem := range destinationImplementation {
		destinationImplementationRule = append(destinationImplementationRule, destinationImplementationItem)
	}

	var messageIdentifierRule []interface{}
	for _, messageIdentifierItem := range messageIdentifier {
		messageIdentifierRule = append(messageIdentifierRule, messageIdentifierItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.WatchLogs(opts, "BountyPlaced", destinationImplementationRule, messageIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentivizedPolymerEscrowBountyPlaced)
				if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "BountyPlaced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBountyPlaced is a log parse operation binding the contract event 0x1b7379b49b96ce33cfd9daf95718231650c40a2c92da579412e5886011867727.
//
// Solidity: event BountyPlaced(bytes indexed destinationImplementation, bytes32 chainIdentifier, bytes32 indexed messageIdentifier, (uint48,uint48,address,uint96,uint96,uint64) incentive)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) ParseBountyPlaced(log types.Log) (*IncentivizedPolymerEscrowBountyPlaced, error) {
	event := new(IncentivizedPolymerEscrowBountyPlaced)
	if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "BountyPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentivizedPolymerEscrowMessageAckedIterator is returned from FilterMessageAcked and is used to iterate over the raw logs and unpacked data for MessageAcked events raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowMessageAckedIterator struct {
	Event *IncentivizedPolymerEscrowMessageAcked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncentivizedPolymerEscrowMessageAckedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentivizedPolymerEscrowMessageAcked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncentivizedPolymerEscrowMessageAcked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncentivizedPolymerEscrowMessageAckedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentivizedPolymerEscrowMessageAckedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentivizedPolymerEscrowMessageAcked represents a MessageAcked event raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowMessageAcked struct {
	DestinationImplementation []byte
	ChainIdentifier           [32]byte
	MessageIdentifier         [32]byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterMessageAcked is a free log retrieval operation binding the contract event 0xf865f210bf9219c3ca273416db27a66557c44a930cfbc569e325afc1e809a307.
//
// Solidity: event MessageAcked(bytes destinationImplementation, bytes32 chainIdentifier, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) FilterMessageAcked(opts *bind.FilterOpts) (*IncentivizedPolymerEscrowMessageAckedIterator, error) {

	logs, sub, err := _IncentivizedPolymerEscrow.contract.FilterLogs(opts, "MessageAcked")
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowMessageAckedIterator{contract: _IncentivizedPolymerEscrow.contract, event: "MessageAcked", logs: logs, sub: sub}, nil
}

// WatchMessageAcked is a free log subscription operation binding the contract event 0xf865f210bf9219c3ca273416db27a66557c44a930cfbc569e325afc1e809a307.
//
// Solidity: event MessageAcked(bytes destinationImplementation, bytes32 chainIdentifier, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) WatchMessageAcked(opts *bind.WatchOpts, sink chan<- *IncentivizedPolymerEscrowMessageAcked) (event.Subscription, error) {

	logs, sub, err := _IncentivizedPolymerEscrow.contract.WatchLogs(opts, "MessageAcked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentivizedPolymerEscrowMessageAcked)
				if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "MessageAcked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageAcked is a log parse operation binding the contract event 0xf865f210bf9219c3ca273416db27a66557c44a930cfbc569e325afc1e809a307.
//
// Solidity: event MessageAcked(bytes destinationImplementation, bytes32 chainIdentifier, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) ParseMessageAcked(log types.Log) (*IncentivizedPolymerEscrowMessageAcked, error) {
	event := new(IncentivizedPolymerEscrowMessageAcked)
	if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "MessageAcked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentivizedPolymerEscrowMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowMessageDeliveredIterator struct {
	Event *IncentivizedPolymerEscrowMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncentivizedPolymerEscrowMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentivizedPolymerEscrowMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncentivizedPolymerEscrowMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncentivizedPolymerEscrowMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentivizedPolymerEscrowMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentivizedPolymerEscrowMessageDelivered represents a MessageDelivered event raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowMessageDelivered struct {
	SourceImplementation common.Hash
	ChainIdentifier      [32]byte
	MessageIdentifier    [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x6025efc987827001c12f1f5c8b7831a6439083c88f60686671451fe479ec6e32.
//
// Solidity: event MessageDelivered(bytes indexed sourceImplementation, bytes32 chainIdentifier, bytes32 indexed messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) FilterMessageDelivered(opts *bind.FilterOpts, sourceImplementation [][]byte, messageIdentifier [][32]byte) (*IncentivizedPolymerEscrowMessageDeliveredIterator, error) {

	var sourceImplementationRule []interface{}
	for _, sourceImplementationItem := range sourceImplementation {
		sourceImplementationRule = append(sourceImplementationRule, sourceImplementationItem)
	}

	var messageIdentifierRule []interface{}
	for _, messageIdentifierItem := range messageIdentifier {
		messageIdentifierRule = append(messageIdentifierRule, messageIdentifierItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.FilterLogs(opts, "MessageDelivered", sourceImplementationRule, messageIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowMessageDeliveredIterator{contract: _IncentivizedPolymerEscrow.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x6025efc987827001c12f1f5c8b7831a6439083c88f60686671451fe479ec6e32.
//
// Solidity: event MessageDelivered(bytes indexed sourceImplementation, bytes32 chainIdentifier, bytes32 indexed messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *IncentivizedPolymerEscrowMessageDelivered, sourceImplementation [][]byte, messageIdentifier [][32]byte) (event.Subscription, error) {

	var sourceImplementationRule []interface{}
	for _, sourceImplementationItem := range sourceImplementation {
		sourceImplementationRule = append(sourceImplementationRule, sourceImplementationItem)
	}

	var messageIdentifierRule []interface{}
	for _, messageIdentifierItem := range messageIdentifier {
		messageIdentifierRule = append(messageIdentifierRule, messageIdentifierItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.WatchLogs(opts, "MessageDelivered", sourceImplementationRule, messageIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentivizedPolymerEscrowMessageDelivered)
				if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDelivered is a log parse operation binding the contract event 0x6025efc987827001c12f1f5c8b7831a6439083c88f60686671451fe479ec6e32.
//
// Solidity: event MessageDelivered(bytes indexed sourceImplementation, bytes32 chainIdentifier, bytes32 indexed messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) ParseMessageDelivered(log types.Log) (*IncentivizedPolymerEscrowMessageDelivered, error) {
	event := new(IncentivizedPolymerEscrowMessageDelivered)
	if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentivizedPolymerEscrowMessageTimedOutIterator is returned from FilterMessageTimedOut and is used to iterate over the raw logs and unpacked data for MessageTimedOut events raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowMessageTimedOutIterator struct {
	Event *IncentivizedPolymerEscrowMessageTimedOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncentivizedPolymerEscrowMessageTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentivizedPolymerEscrowMessageTimedOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncentivizedPolymerEscrowMessageTimedOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncentivizedPolymerEscrowMessageTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentivizedPolymerEscrowMessageTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentivizedPolymerEscrowMessageTimedOut represents a MessageTimedOut event raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowMessageTimedOut struct {
	DestinationImplementation []byte
	ChainIdentifier           [32]byte
	MessageIdentifier         [32]byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterMessageTimedOut is a free log retrieval operation binding the contract event 0xb40655ff3c2bb146ef446b11cb0b22131f89ccd00240340eda5ef4e010794ea1.
//
// Solidity: event MessageTimedOut(bytes destinationImplementation, bytes32 chainIdentifier, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) FilterMessageTimedOut(opts *bind.FilterOpts) (*IncentivizedPolymerEscrowMessageTimedOutIterator, error) {

	logs, sub, err := _IncentivizedPolymerEscrow.contract.FilterLogs(opts, "MessageTimedOut")
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowMessageTimedOutIterator{contract: _IncentivizedPolymerEscrow.contract, event: "MessageTimedOut", logs: logs, sub: sub}, nil
}

// WatchMessageTimedOut is a free log subscription operation binding the contract event 0xb40655ff3c2bb146ef446b11cb0b22131f89ccd00240340eda5ef4e010794ea1.
//
// Solidity: event MessageTimedOut(bytes destinationImplementation, bytes32 chainIdentifier, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) WatchMessageTimedOut(opts *bind.WatchOpts, sink chan<- *IncentivizedPolymerEscrowMessageTimedOut) (event.Subscription, error) {

	logs, sub, err := _IncentivizedPolymerEscrow.contract.WatchLogs(opts, "MessageTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentivizedPolymerEscrowMessageTimedOut)
				if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "MessageTimedOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageTimedOut is a log parse operation binding the contract event 0xb40655ff3c2bb146ef446b11cb0b22131f89ccd00240340eda5ef4e010794ea1.
//
// Solidity: event MessageTimedOut(bytes destinationImplementation, bytes32 chainIdentifier, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) ParseMessageTimedOut(log types.Log) (*IncentivizedPolymerEscrowMessageTimedOut, error) {
	event := new(IncentivizedPolymerEscrowMessageTimedOut)
	if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "MessageTimedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentivizedPolymerEscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowOwnershipTransferredIterator struct {
	Event *IncentivizedPolymerEscrowOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncentivizedPolymerEscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentivizedPolymerEscrowOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncentivizedPolymerEscrowOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncentivizedPolymerEscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentivizedPolymerEscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentivizedPolymerEscrowOwnershipTransferred represents a OwnershipTransferred event raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IncentivizedPolymerEscrowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowOwnershipTransferredIterator{contract: _IncentivizedPolymerEscrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IncentivizedPolymerEscrowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IncentivizedPolymerEscrow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentivizedPolymerEscrowOwnershipTransferred)
				if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) ParseOwnershipTransferred(log types.Log) (*IncentivizedPolymerEscrowOwnershipTransferred, error) {
	event := new(IncentivizedPolymerEscrowOwnershipTransferred)
	if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentivizedPolymerEscrowRemoteImplementationSetIterator is returned from FilterRemoteImplementationSet and is used to iterate over the raw logs and unpacked data for RemoteImplementationSet events raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowRemoteImplementationSetIterator struct {
	Event *IncentivizedPolymerEscrowRemoteImplementationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncentivizedPolymerEscrowRemoteImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentivizedPolymerEscrowRemoteImplementationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncentivizedPolymerEscrowRemoteImplementationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncentivizedPolymerEscrowRemoteImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentivizedPolymerEscrowRemoteImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentivizedPolymerEscrowRemoteImplementationSet represents a RemoteImplementationSet event raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowRemoteImplementationSet struct {
	Application               common.Address
	ChainIdentifier           [32]byte
	ImplementationAddressHash [32]byte
	ImplementationAddress     []byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterRemoteImplementationSet is a free log retrieval operation binding the contract event 0x84cad7de2c9074e70f8d697042ecb3c00ae9e6622320cb8c48bdd4d5d709e796.
//
// Solidity: event RemoteImplementationSet(address application, bytes32 chainIdentifier, bytes32 implementationAddressHash, bytes implementationAddress)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) FilterRemoteImplementationSet(opts *bind.FilterOpts) (*IncentivizedPolymerEscrowRemoteImplementationSetIterator, error) {

	logs, sub, err := _IncentivizedPolymerEscrow.contract.FilterLogs(opts, "RemoteImplementationSet")
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowRemoteImplementationSetIterator{contract: _IncentivizedPolymerEscrow.contract, event: "RemoteImplementationSet", logs: logs, sub: sub}, nil
}

// WatchRemoteImplementationSet is a free log subscription operation binding the contract event 0x84cad7de2c9074e70f8d697042ecb3c00ae9e6622320cb8c48bdd4d5d709e796.
//
// Solidity: event RemoteImplementationSet(address application, bytes32 chainIdentifier, bytes32 implementationAddressHash, bytes implementationAddress)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) WatchRemoteImplementationSet(opts *bind.WatchOpts, sink chan<- *IncentivizedPolymerEscrowRemoteImplementationSet) (event.Subscription, error) {

	logs, sub, err := _IncentivizedPolymerEscrow.contract.WatchLogs(opts, "RemoteImplementationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentivizedPolymerEscrowRemoteImplementationSet)
				if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "RemoteImplementationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoteImplementationSet is a log parse operation binding the contract event 0x84cad7de2c9074e70f8d697042ecb3c00ae9e6622320cb8c48bdd4d5d709e796.
//
// Solidity: event RemoteImplementationSet(address application, bytes32 chainIdentifier, bytes32 implementationAddressHash, bytes implementationAddress)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) ParseRemoteImplementationSet(log types.Log) (*IncentivizedPolymerEscrowRemoteImplementationSet, error) {
	event := new(IncentivizedPolymerEscrowRemoteImplementationSet)
	if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "RemoteImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentivizedPolymerEscrowTimeoutInitiatedIterator is returned from FilterTimeoutInitiated and is used to iterate over the raw logs and unpacked data for TimeoutInitiated events raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowTimeoutInitiatedIterator struct {
	Event *IncentivizedPolymerEscrowTimeoutInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncentivizedPolymerEscrowTimeoutInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentivizedPolymerEscrowTimeoutInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncentivizedPolymerEscrowTimeoutInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncentivizedPolymerEscrowTimeoutInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentivizedPolymerEscrowTimeoutInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentivizedPolymerEscrowTimeoutInitiated represents a TimeoutInitiated event raised by the IncentivizedPolymerEscrow contract.
type IncentivizedPolymerEscrowTimeoutInitiated struct {
	SourceImplementation []byte
	ChainIdentifier      [32]byte
	MessageIdentifier    [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTimeoutInitiated is a free log retrieval operation binding the contract event 0x4b43f0920d04b8fb5f32167b55abce15859e8d8bb1bfe84ba2909268d77de8d2.
//
// Solidity: event TimeoutInitiated(bytes sourceImplementation, bytes32 chainIdentifier, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) FilterTimeoutInitiated(opts *bind.FilterOpts) (*IncentivizedPolymerEscrowTimeoutInitiatedIterator, error) {

	logs, sub, err := _IncentivizedPolymerEscrow.contract.FilterLogs(opts, "TimeoutInitiated")
	if err != nil {
		return nil, err
	}
	return &IncentivizedPolymerEscrowTimeoutInitiatedIterator{contract: _IncentivizedPolymerEscrow.contract, event: "TimeoutInitiated", logs: logs, sub: sub}, nil
}

// WatchTimeoutInitiated is a free log subscription operation binding the contract event 0x4b43f0920d04b8fb5f32167b55abce15859e8d8bb1bfe84ba2909268d77de8d2.
//
// Solidity: event TimeoutInitiated(bytes sourceImplementation, bytes32 chainIdentifier, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) WatchTimeoutInitiated(opts *bind.WatchOpts, sink chan<- *IncentivizedPolymerEscrowTimeoutInitiated) (event.Subscription, error) {

	logs, sub, err := _IncentivizedPolymerEscrow.contract.WatchLogs(opts, "TimeoutInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentivizedPolymerEscrowTimeoutInitiated)
				if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "TimeoutInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimeoutInitiated is a log parse operation binding the contract event 0x4b43f0920d04b8fb5f32167b55abce15859e8d8bb1bfe84ba2909268d77de8d2.
//
// Solidity: event TimeoutInitiated(bytes sourceImplementation, bytes32 chainIdentifier, bytes32 messageIdentifier)
func (_IncentivizedPolymerEscrow *IncentivizedPolymerEscrowFilterer) ParseTimeoutInitiated(log types.Log) (*IncentivizedPolymerEscrowTimeoutInitiated, error) {
	event := new(IncentivizedPolymerEscrowTimeoutInitiated)
	if err := _IncentivizedPolymerEscrow.contract.UnpackLog(event, "TimeoutInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

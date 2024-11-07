package api

import (
	"time"

	"github.com/lyn0904/spine-go/model"
)

/* Feature */

// This interface defines the functions being common to local and remote features
// A feature corresponds to a SPINE feature, see SPINE Introduction Chapter 2.2
type FeatureInterface interface {
	// Get the feature address
	Address() *model.FeatureAddressType
	// Get the feature type
	Type() model.FeatureTypeType
	// Get the feature role
	Role() model.RoleType
	// Get the feature operations
	Operations() map[model.FunctionType]OperationsInterface
	// Get the feature description
	Description() *model.DescriptionType
	// Set the feature description with a given type
	SetDescription(desc *model.DescriptionType)
	// Set the feature description with a given string
	SetDescriptionString(s string)
	// Return a descriptive feature summary as a string
	String() string
}

// Callback function used to verify if an incoming SPINE write message should be allowed or not
// The cb function has to be invoked within 1 minute, otherwise the stack will
// deny the write command
type WriteApprovalCallbackFunc func(msg *Message)

// This interface defines all the required functions need to implement a local feature
type FeatureLocalInterface interface {
	FeatureInterface
	// Get the associated DeviceLocalInterface implementation
	Device() DeviceLocalInterface
	// Get the associated EntityLocalInterface implementation
	Entity() EntityLocalInterface

	// Add a function type with allowed operations
	AddFunctionType(function model.FunctionType, read, write bool)
	// Add a callback function to be invoked when SPINE message comes in with a given msgCounterReference value
	//
	// Returns an error if there is already a callback for the msgCounter set
	AddResponseCallback(msgCounterReference model.MsgCounterType, function func(msg ResponseMessage)) error
	// Add a callback function to be invoked when a result message comes in for this feature
	AddResultCallback(function func(msg ResponseMessage))

	// Add a callback method for a server feature which is invoked to
	// check wether an incoming write message shall be approved or denied
	AddWriteApprovalCallback(function WriteApprovalCallbackFunc) error
	// This function needs to be invoked within (default) 10 seconds after the via
	// AddWriteApprovalCallback defined callback is being invoked.
	//
	// NOTE: To approve a write, ALL callbacks need to approve the write!
	//
	// ErrorType.ErrorNumber should be 0 if write is approved
	ApproveOrDenyWrite(msg *Message, err model.ErrorType)
	// Overwrite the default 1 minute timeout for write approvals
	SetWriteApprovalTimeout(duration time.Duration)

	// Clean all write approval caches for a remote device ski
	CleanWriteApprovalCaches(ski string)
	// Clean all remote device specific caches
	CleanRemoteDeviceCaches(remoteAddress *model.DeviceAddressType)
	// Clean all remote entity specific caches
	CleanRemoteEntityCaches(remoteAddress *model.EntityAddressType)

	// return all functions
	Functions() []model.FunctionType

	// Get a copy of the features data for a given function type
	DataCopy(function model.FunctionType) any
	// Update the features data for a given function type
	UpdateData(function model.FunctionType, data any, filterPartial *model.FilterType, filterDelete *model.FilterType) *model.ErrorType
	// Set the features data for a given function type
	SetData(function model.FunctionType, data any)

	// Trigger a read request message for a given FeatureRemoteInterface implementation
	RequestRemoteData(
		function model.FunctionType,
		selector any,
		elements any,
		destination FeatureRemoteInterface) (*model.MsgCounterType, *model.ErrorType)
	// Trigger a read request message for a remote ski and feature address
	RequestRemoteDataBySenderAddress(
		cmd model.CmdType,
		sender SenderInterface,
		destinationSki string,
		destinationAddress *model.FeatureAddressType,
		maxDelay time.Duration) (*model.MsgCounterType, *model.ErrorType)

	// Check if there already is a subscription to a given feature remote address
	HasSubscriptionToRemote(remoteAddress *model.FeatureAddressType) bool
	// Trigger a subscription request to a given feature remote address
	SubscribeToRemote(remoteAddress *model.FeatureAddressType) (*model.MsgCounterType, *model.ErrorType)
	// Trigger a subscription removal request for a given feature remote address
	RemoveRemoteSubscription(remoteAddress *model.FeatureAddressType) (*model.MsgCounterType, *model.ErrorType)
	// Trigger subscription removal requests for all subscriptions of this feature
	RemoveAllRemoteSubscriptions()

	// Check if there already is a binding to a given feature remote address
	HasBindingToRemote(remoteAddress *model.FeatureAddressType) bool
	// Trigger a binding request to a given feature remote address
	BindToRemote(remoteAddress *model.FeatureAddressType) (*model.MsgCounterType, *model.ErrorType)
	// Trigger a binding removal request for a given feature remote address
	RemoveRemoteBinding(remoteAddress *model.FeatureAddressType) (*model.MsgCounterType, *model.ErrorType)
	// Trigger binding removal requests for all subscriptions of this feature
	RemoveAllRemoteBindings()

	// Handle an incoming SPINE message for this feature
	HandleMessage(message *Message) *model.ErrorType

	// Get the SPINE data structure for NodeManagementDetailDiscoveryData messages for this feature
	Information() *model.NodeManagementDetailedDiscoveryFeatureInformationType
}

// Interface for local NodeManagement feature
type NodeManagementInterface interface {
	FeatureLocalInterface
}

// This interface defines all the required functions need to implement a remote feature
type FeatureRemoteInterface interface {
	FeatureInterface

	// Get the associated DeviceRemoteInterface implementation
	Device() DeviceRemoteInterface
	// Get the associated EntityRemoteInterface implementation
	Entity() EntityRemoteInterface

	// Get a copy of the features data for a given function type
	DataCopy(function model.FunctionType) any
	// Set the features data for a given function type
	// persist true will store the data, false will return the updated data without storing it
	UpdateData(persist bool, function model.FunctionType, data any, filterPartial *model.FilterType, filterDelete *model.FilterType) (any, *model.ErrorType)

	// Set the supported operations of the feature for a set of functions
	SetOperations(functions []model.FunctionPropertyType)

	// Define the maximum response duration
	SetMaxResponseDelay(delay *model.MaxResponseDelayType)
	// Get the maximum allowed response duration
	MaxResponseDelayDuration() time.Duration
}

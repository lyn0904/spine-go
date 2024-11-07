package spine

import (
	"fmt"

	"github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/util"
)

const NodeManagementFeatureId uint = 0

func NodeManagementAddress(deviceAddress *model.AddressDeviceType) *model.FeatureAddressType {
	return &model.FeatureAddressType{
		Entity:  []model.AddressEntityType{0},
		Feature: util.Ptr(model.AddressFeatureType(NodeManagementFeatureId)),
		Device:  deviceAddress,
	}
}

var _ api.NodeManagementInterface = (*NodeManagement)(nil)

type NodeManagement struct {
	*FeatureLocal
	entity api.EntityLocalInterface
}

func NewNodeManagement(id uint, entity api.EntityLocalInterface) *NodeManagement {
	f := &NodeManagement{
		FeatureLocal: NewFeatureLocal(
			id, entity,
			model.FeatureTypeTypeNodeManagement,
			model.RoleTypeSpecial),
		entity: entity,
	}

	f.AddFunctionType(model.FunctionTypeNodeManagementDetailedDiscoveryData, true, false)
	f.AddFunctionType(model.FunctionTypeNodeManagementUseCaseData, true, false)
	f.AddFunctionType(model.FunctionTypeNodeManagementSubscriptionData, true, false)
	f.AddFunctionType(model.FunctionTypeNodeManagementSubscriptionRequestCall, false, false)
	f.AddFunctionType(model.FunctionTypeNodeManagementSubscriptionDeleteCall, false, false)
	f.AddFunctionType(model.FunctionTypeNodeManagementBindingData, true, false)
	f.AddFunctionType(model.FunctionTypeNodeManagementBindingRequestCall, false, false)
	f.AddFunctionType(model.FunctionTypeNodeManagementBindingDeleteCall, false, false)
	if f.Device().FeatureSet() != nil && *f.Device().FeatureSet() != model.NetworkManagementFeatureSetTypeSimple {
		f.AddFunctionType(model.FunctionTypeNodeManagementDestinationListData, true, false)
	}

	return f
}

func (r *NodeManagement) Device() api.DeviceLocalInterface {
	return r.entity.Device()
}

func (r *NodeManagement) HandleMessage(message *api.Message) *model.ErrorType {
	switch {
	case message.Cmd.ResultData != nil:
		if err := r.processResult(message); err != nil {
			return err
		}

	case message.Cmd.NodeManagementDetailedDiscoveryData != nil:
		if err := r.handleMsgDetailedDiscoveryData(message, message.Cmd.NodeManagementDetailedDiscoveryData); err != nil {
			return model.NewErrorType(model.ErrorNumberTypeGeneralError, err.Error())
		}

	case message.Cmd.NodeManagementSubscriptionRequestCall != nil:
		if err := r.handleMsgSubscriptionRequestCall(message, message.Cmd.NodeManagementSubscriptionRequestCall); err != nil {
			return model.NewErrorType(model.ErrorNumberTypeGeneralError, err.Error())
		}

	case message.Cmd.NodeManagementSubscriptionDeleteCall != nil:
		if err := r.handleMsgSubscriptionDeleteCall(message, message.Cmd.NodeManagementSubscriptionDeleteCall); err != nil {
			return model.NewErrorType(model.ErrorNumberTypeGeneralError, err.Error())
		}

	case message.Cmd.NodeManagementSubscriptionData != nil:
		if err := r.handleMsgSubscriptionData(message); err != nil {
			return model.NewErrorType(model.ErrorNumberTypeGeneralError, err.Error())
		}

	case message.Cmd.NodeManagementBindingRequestCall != nil:
		if err := r.handleMsgBindingRequestCall(message, message.Cmd.NodeManagementBindingRequestCall); err != nil {
			return model.NewErrorType(model.ErrorNumberTypeGeneralError, err.Error())
		}

	case message.Cmd.NodeManagementBindingDeleteCall != nil:
		if err := r.handleMsgBindingDeleteCall(message, message.Cmd.NodeManagementBindingDeleteCall); err != nil {
			return model.NewErrorType(model.ErrorNumberTypeGeneralError, err.Error())
		}

	case message.Cmd.NodeManagementBindingData != nil:
		if err := r.handleMsgBindingData(message); err != nil {
			return model.NewErrorType(model.ErrorNumberTypeGeneralError, err.Error())
		}

	case message.Cmd.NodeManagementUseCaseData != nil:
		if err := r.handleMsgUseCaseData(message, message.Cmd.NodeManagementUseCaseData); err != nil {
			return model.NewErrorType(model.ErrorNumberTypeGeneralError, err.Error())
		}

	case message.Cmd.NodeManagementDestinationListData != nil:
		if err := r.handleMsgDestinationListData(message, message.Cmd.NodeManagementDestinationListData); err != nil {
			return model.NewErrorType(model.ErrorNumberTypeGeneralError, err.Error())
		}

	default:
		return model.NewErrorType(model.ErrorNumberTypeCommandNotSupported, fmt.Sprintf("nodemanagement.Handle: Cmd data not implemented: %s", message.Cmd.DataName()))
	}

	return nil
}

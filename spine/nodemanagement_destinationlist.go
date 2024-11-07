package spine

import (
	"errors"
	"fmt"

	"github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
)

func (r *NodeManagement) RequestDestinationListData(remoteDeviceAddress *model.AddressDeviceType, sender api.SenderInterface) (*model.MsgCounterType, *model.ErrorType) {
	return nil, model.NewErrorTypeFromString("Not implemented")
}

func (r *NodeManagement) processReadDestinationListData(featureRemote api.FeatureRemoteInterface, requestHeader *model.HeaderType) error {
	data := []model.NodeManagementDestinationDataType{
		r.Device().DestinationData(),
	}
	// add other remote devices here

	cmd := model.CmdType{
		NodeManagementDestinationListData: &model.NodeManagementDestinationListDataType{
			NodeManagementDestinationData: data,
		},
	}

	return featureRemote.Device().Sender().Reply(requestHeader, r.Address(), cmd)
}

func (r *NodeManagement) processReplyDestinationListData(_ *api.Message, _ model.NodeManagementDestinationListDataType) error {
	return errors.New("Not implemented")
}

func (r *NodeManagement) handleMsgDestinationListData(message *api.Message, data *model.NodeManagementDestinationListDataType) error {
	switch message.CmdClassifier {
	case model.CmdClassifierTypeRead:
		return r.processReadDestinationListData(message.FeatureRemote, message.RequestHeader)

	case model.CmdClassifierTypeReply:
		return r.processReplyDestinationListData(message, *data)

	case model.CmdClassifierTypeNotify:
		return r.processReplyDestinationListData(message, *data)

	default:
		return fmt.Errorf("nodemanagement.handleMsgDestinationListData: NodeManagementDestinationListDataType CmdClassifierType not implemented: %s", message.CmdClassifier)
	}
}

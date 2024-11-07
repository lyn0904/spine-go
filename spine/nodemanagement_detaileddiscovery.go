package spine

import (
	"errors"
	"fmt"
	"slices"

	"github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
)

// request detailed discovery data from a remote device
func (r *NodeManagement) RequestDetailedDiscovery(remoteDeviceSki string, remoteDeviceAddress *model.AddressDeviceType, sender api.SenderInterface) (*model.MsgCounterType, *model.ErrorType) {
	rfAddress := featureAddressType(NodeManagementFeatureId, EntityAddressType(remoteDeviceAddress, DeviceInformationAddressEntity))
	cmd := model.CmdType{
		NodeManagementDetailedDiscoveryData: &model.NodeManagementDetailedDiscoveryDataType{},
	}
	return r.RequestRemoteDataBySenderAddress(cmd, sender, remoteDeviceSki, rfAddress, defaultMaxResponseDelay)
}

// handle incoming detailed discovery read call
func (r *NodeManagement) processReadDetailedDiscoveryData(deviceRemote api.DeviceRemoteInterface, requestHeader *model.HeaderType) error {
	if deviceRemote == nil {
		return errors.New("nodemanagement.readDetailedDiscoveryData: invalid deviceRemote")
	}

	var entityInformation []model.NodeManagementDetailedDiscoveryEntityInformationType
	var featureInformation []model.NodeManagementDetailedDiscoveryFeatureInformationType

	for _, e := range r.Device().Entities() {
		entityInformation = append(entityInformation, *e.Information())

		for _, f := range e.Features() {
			featureInformation = append(featureInformation, *f.Information())
		}
	}

	cmd := model.CmdType{
		NodeManagementDetailedDiscoveryData: &model.NodeManagementDetailedDiscoveryDataType{
			SpecificationVersionList: &model.NodeManagementSpecificationVersionListType{
				SpecificationVersion: []model.SpecificationVersionDataType{model.SpecificationVersionDataType(SpecificationVersion)},
			},
			DeviceInformation:  r.Device().Information(),
			EntityInformation:  entityInformation,
			FeatureInformation: featureInformation,
		},
	}

	return deviceRemote.Sender().Reply(requestHeader, r.Address(), cmd)
}

// handle incoming detailed discovery reply data
func (r *NodeManagement) processReplyDetailedDiscoveryData(message *api.Message, data *model.NodeManagementDetailedDiscoveryDataType) error {
	remoteDevice := message.DeviceRemote

	deviceDescription := data.DeviceInformation.Description
	if deviceDescription == nil {
		return errors.New("nodemanagement.replyDetailedDiscoveryData: invalid DeviceInformation.Description")
	}

	remoteDevice.UpdateDevice(deviceDescription)
	entities, err := remoteDevice.AddEntityAndFeatures(true, data)
	if err != nil {
		return err
	}

	// publish event for remote device added
	payload := api.EventPayload{
		Ski:        remoteDevice.Ski(),
		EventType:  api.EventTypeDeviceChange,
		ChangeType: api.ElementChangeAdd,
		Device:     remoteDevice,
		Feature:    message.FeatureRemote,
		Data:       data,
	}
	Events.Publish(payload)

	// publish event for each added remote entity
	for _, entity := range entities {
		payload := api.EventPayload{
			Ski:        remoteDevice.Ski(),
			EventType:  api.EventTypeEntityChange,
			ChangeType: api.ElementChangeAdd,
			Device:     remoteDevice,
			Entity:     entity,
			Data:       data,
		}
		Events.Publish(payload)
	}

	return nil
}

// check if an AddressEntity slice exists in a list of AddressEntity slices
func (r *NodeManagement) addressEntityListContainsAddressEntity(list [][]model.AddressEntityType, entityAddress []model.AddressEntityType) bool {
	for _, entityList := range list {
		if slices.Equal(entityList, entityAddress) {
			return true
		}
	}

	return false
}

// process incoming detailed discovery notify with full data
// and return the data diff
func (r *NodeManagement) provideDetailedDiscoveryDiffForFullNotify(message *api.Message, data *model.NodeManagementDetailedDiscoveryDataType) *model.NodeManagementDetailedDiscoveryDataType {
	remoteDevice := message.FeatureRemote.Device()

	var existingEntities, addedEntities [][]model.AddressEntityType

	var updatedEntityInformation []model.NodeManagementDetailedDiscoveryEntityInformationType

	// search for new entities
	for _, entity := range data.EntityInformation {
		if entity.Description == nil || entity.Description.EntityAddress == nil {
			continue
		}

		// check if the entity already exists
		address := entity.Description.EntityAddress
		if remoteDevice.Entity(address.Entity) == nil {
			// does not exists
			added := model.NetworkManagementStateChangeTypeAdded
			entity.Description.LastStateChange = &added
			addedEntities = append(addedEntities, address.Entity)
			updatedEntityInformation = append(updatedEntityInformation, entity)
		} else {
			// exists
			existingEntities = append(existingEntities, address.Entity)
		}
	}

	// seach for removed entites
	for _, entity := range remoteDevice.Entities() {
		address := entity.Address()
		if !r.addressEntityListContainsAddressEntity(existingEntities, address.Entity) {
			// does not exists
			removed := model.NetworkManagementStateChangeTypeRemoved
			entityType := entity.EntityType()

			removedEntityDescription := model.NodeManagementDetailedDiscoveryEntityInformationType{
				Description: &model.NetworkManagementEntityDescriptionDataType{
					EntityAddress:   address,
					EntityType:      &entityType,
					LastStateChange: &removed,
				},
			}

			updatedEntityInformation = append(updatedEntityInformation, removedEntityDescription)
		}
	}

	data.EntityInformation = updatedEntityInformation

	// update the feature information
	var updatedFeatureInformation []model.NodeManagementDetailedDiscoveryFeatureInformationType
	for _, feature := range data.FeatureInformation {
		if feature.Description == nil || feature.Description.FeatureAddress == nil {
			continue
		}

		address := feature.Description.FeatureAddress
		// if the entity of the feature was added, add it
		// if the entity of the feature already existed, do not add it
		if r.addressEntityListContainsAddressEntity(addedEntities, address.Entity) {
			updatedFeatureInformation = append(updatedFeatureInformation, feature)
		}
	}

	data.FeatureInformation = updatedFeatureInformation

	return data
}

// handle incoming detailed discovery notify data
func (r *NodeManagement) processNotifyDetailedDiscoveryData(message *api.Message, data *model.NodeManagementDetailedDiscoveryDataType) error {
	// is this a partial request?
	if message.FilterPartial == nil {
		data = r.provideDetailedDiscoveryDiffForFullNotify(message, data)
	}

	if len(data.EntityInformation) == 0 {
		return errors.New("nodemanagement.notifyDetailedDiscoveryData: invalid EntityInformation")
	}

	for _, entity := range data.EntityInformation {
		if entity.Description == nil ||
			entity.Description.EntityAddress == nil ||
			entity.Description.LastStateChange == nil {
			return errors.New("nodemanagement.notifyDetailedDiscoveryData: invalid EntityInformation.Description")
		}

		lastStateChange := *entity.Description.LastStateChange
		remoteDevice := message.FeatureRemote.Device()

		// addition example:
		// {"data":[{"header":[{"protocolId":"ee1.0"}]},{"payload":{"datagram":[{"header":[{"specificationVersion":"1.1.1"},{"addressSource":[{"device":"d:_i:19667_PorscheEVSE-00016544"},{"entity":[0]},{"feature":0}]},{"addressDestination":[{"device":"EVCC_HEMS"},{"entity":[0]},{"feature":0}]},{"msgCounter":926685},{"cmdClassifier":"notify"}]},{"payload":[{"cmd":[[{"function":"nodeManagementDetailedDiscoveryData"},{"filter":[[{"cmdControl":[{"partial":[]}]}]]},{"nodeManagementDetailedDiscoveryData":[{"deviceInformation":[{"description":[{"deviceAddress":[{"device":"d:_i:19667_PorscheEVSE-00016544"}]}]}]},{"entityInformation":[[{"description":[{"entityAddress":[{"entity":[1,1]}]},{"entityType":"EV"},{"lastStateChange":"added"},{"description":"Electric Vehicle"}]}]]},{"featureInformation":[[{"description":[{"featureAddress":[{"entity":[1,1]},{"feature":1}]},{"featureType":"LoadControl"},{"role":"server"},{"supportedFunction":[[{"function":"loadControlLimitDescriptionListData"},{"possibleOperations":[{"read":[]}]}],[{"function":"loadControlLimitListData"},{"possibleOperations":[{"read":[]},{"write":[]}]}]]},{"description":"Load Control"}]}],[{"description":[{"featureAddress":[{"entity":[1,1]},{"feature":2}]},{"featureType":"ElectricalConnection"},{"role":"server"},{"supportedFunction":[[{"function":"electricalConnectionParameterDescriptionListData"},{"possibleOperations":[{"read":[]}]}],[{"function":"electricalConnectionDescriptionListData"},{"possibleOperations":[{"read":[]}]}],[{"function":"electricalConnectionPermittedValueSetListData"},{"possibleOperations":[{"read":[]}]}]]},{"description":"Electrical Connection"}]}],[{"description":[{"featureAddress":[{"entity":[1,1]},{"feature":3}]},{"featureType":"Measurement"},{"specificUsage":["Electrical"]},{"role":"server"},{"supportedFunction":[[{"function":"measurementListData"},{"possibleOperations":[{"read":[]}]}],[{"function":"measurementDescriptionListData"},{"possibleOperations":[{"read":[]}]}]]},{"description":"Measurements"}]}],[{"description":[{"featureAddress":[{"entity":[1,1]},{"feature":5}]},{"featureType":"DeviceConfiguration"},{"role":"server"},{"supportedFunction":[[{"function":"deviceConfigurationKeyValueDescriptionListData"},{"possibleOperations":[{"read":[]}]}],[{"function":"deviceConfigurationKeyValueListData"},{"possibleOperations":[{"read":[]}]}]]},{"description":"Device Configuration EV"}]}],[{"description":[{"featureAddress":[{"entity":[1,1]},{"feature":6}]},{"featureType":"DeviceClassification"},{"role":"server"},{"supportedFunction":[[{"function":"deviceClassificationManufacturerData"},{"possibleOperations":[{"read":[]}]}]]},{"description":"Device Classification for EV"}]}],[{"description":[{"featureAddress":[{"entity":[1,1]},{"feature":7}]},{"featureType":"TimeSeries"},{"role":"server"},{"supportedFunction":[[{"function":"timeSeriesConstraintsListData"},{"possibleOperations":[{"read":[]}]}],[{"function":"timeSeriesDescriptionListData"},{"possibleOperations":[{"read":[]}]}],[{"function":"timeSeriesListData"},{"possibleOperations":[{"read":[]},{"write":[]}]}]]},{"description":"Time Series"}]}],[{"description":[{"featureAddress":[{"entity":[1,1]},{"feature":8}]},{"featureType":"IncentiveTable"},{"role":"server"},{"supportedFunction":[[{"function":"incentiveTableConstraintsData"},{"possibleOperations":[{"read":[]}]}],[{"function":"incentiveTableData"},{"possibleOperations":[{"read":[]},{"write":[]}]}],[{"function":"incentiveTableDescriptionData"},{"possibleOperations":[{"read":[]},{"write":[]}]}]]},{"description":"Incentive Table"}]}],[{"description":[{"featureAddress":[{"entity":[1,1]},{"feature":9}]},{"featureType":"DeviceDiagnosis"},{"role":"server"},{"supportedFunction":[[{"function":"deviceDiagnosisStateData"},{"possibleOperations":[{"read":[]}]}]]},{"description":"Device Diagnosis EV"}]}],[{"description":[{"featureAddress":[{"entity":[1,1]},{"feature":10}]},{"featureType":"Identification"},{"role":"server"},{"supportedFunction":[[{"function":"identificationListData"},{"possibleOperations":[{"read":[]}]}]]},{"description":"Identification for EV"}]}]]}]}]]}]}]}}]}
		// {
		// 	"cmd":[[
		// 		{"function":"nodeManagementDetailedDiscoveryData"},
		// 		{"filter":[[{"cmdControl":[{"partial":[]}]}]]},
		// 		{"nodeManagementDetailedDiscoveryData":[
		// 			{"deviceInformation":[{"description":[{"deviceAddress":[{"device":"d:_i:19667_PorscheEVSE-00016544"}]}]}]},
		// 			{"entityInformation":[[
		// 				{"description":[
		// 					{"entityAddress":[{"entity":[1,1]}]},
		// 					{"entityType":"EV"},
		// 					{"lastStateChange":"added"},
		// 					{"description":"Electric Vehicle"}
		// 				]}
		// 			]]},
		// 			{"featureInformation":[
		// 				[{"description":[
		// 					{"featureAddress":[{"entity":[1,1]},{"feature":1}]},
		// 					{"featureType":"LoadControl"},
		// 					{"role":"server"},
		// 					{"supportedFunction":[
		// 						[{"function":"loadControlLimitDescriptionListData"},{"possibleOperations":[{"read":[]}]}],
		// 						[{"function":"loadControlLimitListData"},{"possibleOperations":[{"read":[]},{"write":[]}]}]
		// 					]},
		// 					{"description":"Load Control"}
		// 				]}],
		// ...

		// is this addition?
		if lastStateChange == model.NetworkManagementStateChangeTypeAdded {
			entities, err := remoteDevice.AddEntityAndFeatures(false, data)
			if err != nil {
				return err
			}

			// publish event for each added remote entity
			for _, entity := range entities {
				payload := api.EventPayload{
					Ski:        remoteDevice.Ski(),
					EventType:  api.EventTypeEntityChange,
					ChangeType: api.ElementChangeAdd,
					Device:     remoteDevice,
					Entity:     entity,
					Data:       data,
				}
				Events.Publish(payload)
			}
		}

		// removal example:
		// {"data":[{"header":[{"protocolId":"ee1.0"}]},{"payload":{"datagram":[{"header":[{"specificationVersion":"1.1.1"},{"addressSource":[{"device":"d:_i:19667_PorscheEVSE-00016544"},{"entity":[0]},{"feature":0}]},{"addressDestination":[{"device":"EVCC_HEMS"},{"entity":[0]},{"feature":0}]},{"msgCounter":4835},{"cmdClassifier":"notify"}]},{"payload":[{"cmd":[[{"function":"nodeManagementDetailedDiscoveryData"},{"filter":[[{"cmdControl":[{"partial":[]}]}]]},{"nodeManagementDetailedDiscoveryData":[{"deviceInformation":[{"description":[{"deviceAddress":[{"device":"d:_i:19667_PorscheEVSE-00016544"}]}]}]},{"entityInformation":[[{"description":[{"entityAddress":[{"entity":[1,1]}]},{"lastStateChange":"removed"}]}]]}]}]]}]}]}}]}
		// {
		// 	"cmd": [[
		// 			{"function": "nodeManagementDetailedDiscoveryData"},
		// 			{"filter": [[{"cmdControl": [{"partial": []}]}]]},
		// 			{"nodeManagementDetailedDiscoveryData": [
		// 					{"deviceInformation": [{"description": [{"deviceAddress": [{"device": "d:_i:19667_PorscheEVSE-00016544"}]}]}]},
		// 					{"entityInformation": [[
		// 							{
		// 								"description": [
		// 									{"entityAddress": [{"entity": [1,1]}]},
		// 									{"lastStateChange": "removed"}
		// ...

		// is this removal?
		if lastStateChange == model.NetworkManagementStateChangeTypeRemoved {
			for _, ei := range data.EntityInformation {
				if err := remoteDevice.CheckEntityInformation(false, ei); err != nil {
					return err
				}

				entityAddress := ei.Description.EntityAddress.Entity
				removedEntity := remoteDevice.RemoveEntityByAddress(entityAddress)

				// only continue if the entity existed
				if removedEntity == nil {
					continue
				}

				payload := api.EventPayload{
					Ski:        remoteDevice.Ski(),
					EventType:  api.EventTypeEntityChange,
					ChangeType: api.ElementChangeRemove,
					Device:     remoteDevice,
					Entity:     removedEntity,
					Data:       data,
				}
				Events.Publish(payload)

				// remove all subscriptions for this entity
				subscriptionMgr := r.Device().SubscriptionManager()
				subscriptionMgr.RemoveSubscriptionsForEntity(removedEntity)

				// remove all bindings for this entity
				bindingMgr := r.Device().BindingManager()
				bindingMgr.RemoveBindingsForEntity(removedEntity)

				// remove all feature caches for this entity
				r.Device().CleanRemoteEntityCaches(removedEntity.Address())
			}
		}
	}

	return nil
}

// func (f *NodeManagement) announceFeatureDiscovery(e Entity) error {
// 	entity := f.Entity()
// 	if entity == nil {
// 		return errors.New("announceFeatureDiscovery: entity not found")
// 	}
// 	device := entity.Device()
// 	if device == nil {
// 		return errors.New("announceFeatureDiscovery: device not found")
// 	}
// 	entities := device.Entities()
// 	if entities == nil {
// 		return errors.New("announceFeatureDiscovery: entities not found")
// 	}

// 	for _, le := range entities {
// 		for _, lf := range le.Features() {

// 			// connect client to server features
// 			for _, rf := range e.Features() {
// 				lr := lf.Role()
// 				rr := rf.Role()
// 				rolesValid := (lr == model.RoleTypeSpecial && rr == model.RoleTypeSpecial) || (lr == model.RoleTypeClient && rr == model.RoleTypeServer)
// 				if lf.Type() == rf.Type() && rolesValid {
// 					if cf, ok := lf.(ClientFeature); ok {
// 						if err := cf.ServerFound(rf); err != nil {
// 							return err
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return nil
// }

func (r *NodeManagement) handleMsgDetailedDiscoveryData(message *api.Message, data *model.NodeManagementDetailedDiscoveryDataType) error {
	switch message.CmdClassifier {
	case model.CmdClassifierTypeRead:
		return r.processReadDetailedDiscoveryData(message.DeviceRemote, message.RequestHeader)

	case model.CmdClassifierTypeReply:
		return r.processReplyDetailedDiscoveryData(message, data)

	case model.CmdClassifierTypeNotify:
		return r.processNotifyDetailedDiscoveryData(message, data)

	default:
		return fmt.Errorf("nodemanagement.handleDetailedDiscoveryData: NodeManagementDetailedDiscoveryData CmdClassifierType not implemented: %s", message.CmdClassifier)
	}
}

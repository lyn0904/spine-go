package spine

import (
	"sync"
	"time"

	"github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
)

type EntityLocal struct {
	*Entity
	device   api.DeviceLocalInterface
	features []api.FeatureLocalInterface

	heartbeatManager api.HeartbeatManagerInterface

	mux sync.Mutex
}

func NewEntityLocal(device api.DeviceLocalInterface,
	eType model.EntityTypeType,
	entityAddress []model.AddressEntityType,
	heartbeatTimeout time.Duration) *EntityLocal {
	entity := &EntityLocal{
		Entity: NewEntity(eType, device.Address(), entityAddress),
		device: device,
	}
	// only needed if the entity address is not DeviceInformationEntityId
	if len(entityAddress) > 0 && entityAddress[0] != model.AddressEntityType(DeviceInformationEntityId) {
		entity.heartbeatManager = NewHeartbeatManager(entity, heartbeatTimeout)
	}

	return entity
}

var _ api.EntityLocalInterface = (*EntityLocal)(nil)

/* EntityLocalInterface */

func (r *EntityLocal) Device() api.DeviceLocalInterface {
	return r.device
}

func (r *EntityLocal) HeartbeatManager() api.HeartbeatManagerInterface {
	return r.heartbeatManager
}

// Add a feature to the entity if it is not already added
func (r *EntityLocal) AddFeature(f api.FeatureLocalInterface) {
	r.mux.Lock()
	defer r.mux.Unlock()

	// check if this feature is already added
	for _, f2 := range r.features {
		if f2.Type() == f.Type() && f2.Role() == f.Role() {
			return
		}
	}
	r.features = append(r.features, f)
}

// either returns an existing feature or creates a new one
// for a given entity, featuretype and role
func (r *EntityLocal) GetOrAddFeature(featureType model.FeatureTypeType, role model.RoleType) api.FeatureLocalInterface {
	if f := r.FeatureOfTypeAndRole(featureType, role); f != nil {
		return f
	}

	r.mux.Lock()
	defer r.mux.Unlock()

	f := NewFeatureLocal(r.NextFeatureId(), r, featureType, role)

	description := string(featureType)
	switch role {
	case model.RoleTypeClient:
		description += " Client"
	case model.RoleTypeServer:
		description += " Server"
	}
	f.SetDescriptionString(description)
	r.features = append(r.features, f)

	return f
}

func (r *EntityLocal) FeatureOfTypeAndRole(featureType model.FeatureTypeType, role model.RoleType) api.FeatureLocalInterface {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, f := range r.features {
		if f.Type() == featureType && f.Role() == role {
			return f
		}
	}

	return nil
}

func (r *EntityLocal) FeatureOfAddress(addressFeature *model.AddressFeatureType) api.FeatureLocalInterface {
	r.mux.Lock()
	defer r.mux.Unlock()

	if addressFeature == nil {
		return nil
	}
	for _, f := range r.features {
		if f.Address().Feature != nil &&
			*f.Address().Feature == *addressFeature {
			return f
		}
	}

	return nil
}

func (r *EntityLocal) Features() []api.FeatureLocalInterface {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.features
}

// add a new usecase
func (r *EntityLocal) AddUseCaseSupport(
	actor model.UseCaseActorType,
	useCaseName model.UseCaseNameType,
	useCaseVersion model.SpecificationVersionType,
	useCaseDocumemtSubRevision string,
	useCaseAvailable bool,
	scenarios []model.UseCaseScenarioSupportType,
) {
	nodeMgmt := r.device.NodeManagement()

	data, err := LocalFeatureDataCopyOfType[*model.NodeManagementUseCaseDataType](nodeMgmt, model.FunctionTypeNodeManagementUseCaseData)
	if err != nil {
		data = &model.NodeManagementUseCaseDataType{}
	}

	address := model.FeatureAddressType{
		Device: r.address.Device,
		Entity: r.address.Entity,
	}

	data.AddUseCaseSupport(address, actor, useCaseName, useCaseVersion, useCaseDocumemtSubRevision, useCaseAvailable, scenarios)

	nodeMgmt.SetData(model.FunctionTypeNodeManagementUseCaseData, data)
}

// Check if a use case is already added
func (r *EntityLocal) HasUseCaseSupport(actor model.UseCaseActorType, useCaseName model.UseCaseNameType) bool {
	nodeMgmt := r.device.NodeManagement()

	data, err := LocalFeatureDataCopyOfType[*model.NodeManagementUseCaseDataType](nodeMgmt, model.FunctionTypeNodeManagementUseCaseData)
	if err != nil {
		return false
	}

	address := model.FeatureAddressType{
		Device: r.address.Device,
		Entity: r.address.Entity,
	}

	return data.HasUseCaseSupport(address, actor, useCaseName)
}

// Set the availability of a usecase. This may only be used for usescases
// that act as a client within the usecase!
func (r *EntityLocal) SetUseCaseAvailability(
	actor model.UseCaseActorType,
	useCaseName model.UseCaseNameType,
	available bool) {
	nodeMgmt := r.device.NodeManagement()

	data, err := LocalFeatureDataCopyOfType[*model.NodeManagementUseCaseDataType](nodeMgmt, model.FunctionTypeNodeManagementUseCaseData)
	if err != nil {
		return
	}

	address := model.FeatureAddressType{
		Device: r.address.Device,
		Entity: r.address.Entity,
	}

	data.SetAvailability(address, actor, useCaseName, available)

	nodeMgmt.SetData(model.FunctionTypeNodeManagementUseCaseData, data)
}

// Remove a usecase with a given actor ans usecase name
func (r *EntityLocal) RemoveUseCaseSupport(
	actor model.UseCaseActorType,
	useCaseName model.UseCaseNameType,
) {
	nodeMgmt := r.device.NodeManagement()

	data, err := LocalFeatureDataCopyOfType[*model.NodeManagementUseCaseDataType](nodeMgmt, model.FunctionTypeNodeManagementUseCaseData)
	if err != nil {
		return
	}

	address := model.FeatureAddressType{
		Device: r.address.Device,
		Entity: r.address.Entity,
	}

	data.RemoveUseCaseSupport(address, actor, useCaseName)

	nodeMgmt.SetData(model.FunctionTypeNodeManagementUseCaseData, data)
}

// Remove all usecases
func (r *EntityLocal) RemoveAllUseCaseSupports() {
	nodeMgmt := r.device.NodeManagement()

	data, err := LocalFeatureDataCopyOfType[*model.NodeManagementUseCaseDataType](nodeMgmt, model.FunctionTypeNodeManagementUseCaseData)
	if err != nil {
		return
	}

	address := model.FeatureAddressType{
		Device: r.address.Device,
		Entity: r.address.Entity,
	}

	data.RemoveUseCaseDataForAddress(address)

	nodeMgmt.SetData(model.FunctionTypeNodeManagementUseCaseData, data)
}

// Remove all subscriptions
func (r *EntityLocal) RemoveAllSubscriptions() {
	for _, item := range r.features {
		item.RemoveAllRemoteSubscriptions()
	}
}

// Remove all bindings
func (r *EntityLocal) RemoveAllBindings() {
	for _, item := range r.features {
		item.RemoveAllRemoteBindings()
	}
}

func (r *EntityLocal) Information() *model.NodeManagementDetailedDiscoveryEntityInformationType {
	res := &model.NodeManagementDetailedDiscoveryEntityInformationType{
		Description: &model.NetworkManagementEntityDescriptionDataType{
			EntityAddress: r.Address(),
			EntityType:    &r.eType,
		},
	}

	return res
}

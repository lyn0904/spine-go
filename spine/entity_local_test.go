package spine

import (
	"testing"
	"time"

	"github.com/lyn0904/spine-go/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestEntityLocalSuite(t *testing.T) {
	suite.Run(t, new(EntityLocalTestSuite))
}

type EntityLocalTestSuite struct {
	suite.Suite
}

func (suite *EntityLocalTestSuite) Test_Entity() {
	device := NewDeviceLocal("brand", "model", "serial", "code", "address", model.DeviceTypeTypeEnergyManagementSystem, model.NetworkManagementFeatureSetTypeSmart)
	entity := NewEntityLocal(device, model.EntityTypeTypeCEM, NewAddressEntityType([]uint{1}), time.Second*4)
	device.AddEntity(entity)

	f := NewFeatureLocal(1, entity, model.FeatureTypeTypeElectricalConnection, model.RoleTypeClient)
	entity.AddFeature(f)
	assert.Equal(suite.T(), 1, len(entity.Features()))

	entity.AddFeature(f)
	assert.Equal(suite.T(), 1, len(entity.Features()))

	f1 := entity.FeatureOfAddress(nil)
	assert.Nil(suite.T(), f1)

	f1 = entity.FeatureOfAddress(f.Address().Feature)
	assert.NotNil(suite.T(), f1)

	fakeAddress := model.AddressFeatureType(5)
	f1 = entity.FeatureOfAddress(&fakeAddress)
	assert.Nil(suite.T(), f1)

	f2 := entity.GetOrAddFeature(model.FeatureTypeTypeMeasurement, model.RoleTypeClient)
	assert.NotNil(suite.T(), f2)

	assert.Equal(suite.T(), 2, len(entity.Features()))

	f3 := entity.GetOrAddFeature(model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)
	assert.NotNil(suite.T(), f3)

	assert.Equal(suite.T(), 3, len(entity.Features()))

	f4 := entity.GetOrAddFeature(model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)
	assert.NotNil(suite.T(), f4)

	assert.Equal(suite.T(), 3, len(entity.Features()))

	entity.RemoveAllUseCaseSupports()

	hasUC := entity.HasUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
	)
	assert.Equal(suite.T(), false, hasUC)

	_, err := LocalFeatureDataCopyOfType[*model.NodeManagementUseCaseDataType](device.NodeManagement(), model.FunctionTypeNodeManagementUseCaseData)
	assert.NotNil(suite.T(), err)

	entity.AddUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
		model.SpecificationVersionType("1.0.0"),
		"",
		true,
		[]model.UseCaseScenarioSupportType{1, 2},
	)

	_, err = LocalFeatureDataCopyOfType[*model.NodeManagementUseCaseDataType](device.NodeManagement(), model.FunctionTypeNodeManagementUseCaseData)
	assert.Nil(suite.T(), err)

	hasUC = entity.HasUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
	)
	assert.Equal(suite.T(), true, hasUC)

	entity.AddUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
		model.SpecificationVersionType("1.0.0"),
		"",
		true,
		[]model.UseCaseScenarioSupportType{1, 2},
	)

	hasUC = entity.HasUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
	)
	assert.Equal(suite.T(), true, hasUC)

	entity.SetUseCaseAvailability(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
		false,
	)

	entity.RemoveUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
	)

	hasUC = entity.HasUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
	)
	assert.Equal(suite.T(), false, hasUC)

	entity.AddUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
		model.SpecificationVersionType("1.0.0"),
		"",
		true,
		[]model.UseCaseScenarioSupportType{1, 2},
	)

	hasUC = entity.HasUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
	)
	assert.Equal(suite.T(), true, hasUC)

	entity.RemoveAllUseCaseSupports()

	hasUC = entity.HasUseCaseSupport(
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVSECommissioningAndConfiguration,
	)
	assert.Equal(suite.T(), false, hasUC)

	entity.RemoveAllBindings()
	entity.RemoveAllSubscriptions()
}

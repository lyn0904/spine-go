package spine

import (
	"testing"
	"time"

	"github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestBindingManagerSuite(t *testing.T) {
	suite.Run(t, new(BindingManagerSuite))
}

type BindingManagerSuite struct {
	suite.Suite

	localDevice  api.DeviceLocalInterface
	writeHandler *WriteMessageHandler
	remoteDevice *DeviceRemote

	sut api.BindingManagerInterface
}

func (s *BindingManagerSuite) BeforeTest(suiteName, testName string) {
	s.localDevice = NewDeviceLocal("TestBrandName", "TestDeviceModel", "TestSerialNumber", "TestDeviceCode",
		"TestDeviceAddress", model.DeviceTypeTypeEnergyManagementSystem, model.NetworkManagementFeatureSetTypeSmart)
	remoteSki := "TestRemoteSki"

	s.writeHandler = &WriteMessageHandler{}

	sender := NewSender(s.writeHandler)
	s.remoteDevice = NewDeviceRemote(s.localDevice, remoteSki, sender)
	s.remoteDevice.address = util.Ptr(model.AddressDeviceType("Address"))

	s.sut = NewBindingManager(s.localDevice)
}

func (suite *BindingManagerSuite) Test_Bindings() {
	entity := NewEntityLocal(suite.localDevice, model.EntityTypeTypeCEM, []model.AddressEntityType{1}, time.Second*4)
	suite.localDevice.AddEntity(entity)

	localFeature := entity.GetOrAddFeature(model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)
	localClientFeature := entity.GetOrAddFeature(model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeClient)

	remoteEntity := NewEntityRemote(suite.remoteDevice, model.EntityTypeTypeEVSE, []model.AddressEntityType{1})

	remoteFeature := NewFeatureRemote(remoteEntity.NextFeatureId(), remoteEntity, model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeClient)
	remoteFeature.Address().Device = util.Ptr(model.AddressDeviceType("remoteDevice"))
	remoteEntity.AddFeature(remoteFeature)

	remoteServerFeature := NewFeatureRemote(remoteEntity.NextFeatureId(), remoteEntity, model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)
	remoteServerFeature.Address().Device = util.Ptr(model.AddressDeviceType("remoteDevice"))
	remoteEntity.AddFeature(remoteServerFeature)

	suite.remoteDevice.AddEntity(remoteEntity)

	bindingMgr := suite.localDevice.BindingManager()

	bindingRequest := model.BindingManagementRequestCallType{
		ClientAddress: util.Ptr(model.FeatureAddressType{
			Device:  util.Ptr(model.AddressDeviceType("dummy")),
			Entity:  []model.AddressEntityType{1000},
			Feature: util.Ptr(model.AddressFeatureType(1000)),
		}),
		ServerAddress: util.Ptr(model.FeatureAddressType{
			Device:  util.Ptr(model.AddressDeviceType("dummy")),
			Entity:  []model.AddressEntityType{1000},
			Feature: util.Ptr(model.AddressFeatureType(1000)),
		}),
		ServerFeatureType: util.Ptr(model.FeatureTypeTypeDeviceDiagnosis),
	}

	err := bindingMgr.AddBinding(suite.remoteDevice, bindingRequest)
	assert.NotNil(suite.T(), err)

	bindingRequest = model.BindingManagementRequestCallType{
		ClientAddress: util.Ptr(model.FeatureAddressType{
			Device:  util.Ptr(model.AddressDeviceType("dummy")),
			Entity:  []model.AddressEntityType{1000},
			Feature: util.Ptr(model.AddressFeatureType(1000)),
		}),
		ServerAddress: localClientFeature.Address(),
	}

	err = bindingMgr.AddBinding(suite.remoteDevice, bindingRequest)
	assert.NotNil(suite.T(), err)

	bindingRequest.ServerFeatureType = util.Ptr(model.FeatureTypeTypeDeviceDiagnosis)

	err = bindingMgr.AddBinding(suite.remoteDevice, bindingRequest)
	assert.NotNil(suite.T(), err)

	bindingRequest.ServerAddress = localFeature.Address()

	err = bindingMgr.AddBinding(suite.remoteDevice, bindingRequest)
	assert.NotNil(suite.T(), err)

	bindingRequest.ClientAddress = remoteServerFeature.Address()

	err = bindingMgr.AddBinding(suite.remoteDevice, bindingRequest)
	assert.NotNil(suite.T(), err)

	bindingRequest.ClientAddress = remoteFeature.Address()

	err = bindingMgr.AddBinding(suite.remoteDevice, bindingRequest)
	assert.Nil(suite.T(), err)

	subs := bindingMgr.Bindings(suite.remoteDevice)
	assert.Equal(suite.T(), 1, len(subs))

	err = bindingMgr.AddBinding(suite.remoteDevice, bindingRequest)
	assert.NotNil(suite.T(), err)

	subs = bindingMgr.Bindings(suite.remoteDevice)
	assert.Equal(suite.T(), 1, len(subs))

	address := model.FeatureAddressType{
		Device:  entity.Device().Address(),
		Entity:  entity.Address().Entity,
		Feature: util.Ptr(model.AddressFeatureType(10)),
	}
	entries := bindingMgr.BindingsOnFeature(address)
	assert.Equal(suite.T(), 0, len(entries))

	address.Feature = localFeature.Address().Feature
	entries = bindingMgr.BindingsOnFeature(address)
	assert.Equal(suite.T(), 1, len(entries))

	bindingDelete := model.BindingManagementDeleteCallType{
		ClientAddress: util.Ptr(model.FeatureAddressType{
			Device:  util.Ptr(model.AddressDeviceType("dummy")),
			Entity:  []model.AddressEntityType{1000},
			Feature: util.Ptr(model.AddressFeatureType(1000)),
		}),
		ServerAddress: util.Ptr(model.FeatureAddressType{
			Device:  util.Ptr(model.AddressDeviceType("dummy")),
			Entity:  []model.AddressEntityType{1000},
			Feature: util.Ptr(model.AddressFeatureType(1000)),
		}),
	}

	err = bindingMgr.RemoveBinding(bindingDelete, suite.remoteDevice)
	assert.NotNil(suite.T(), err)

	bindingDelete.ClientAddress = remoteServerFeature.Address()

	err = bindingMgr.RemoveBinding(bindingDelete, suite.remoteDevice)
	assert.NotNil(suite.T(), err)

	bindingDelete.ServerAddress = localClientFeature.Address()

	err = bindingMgr.RemoveBinding(bindingDelete, suite.remoteDevice)
	assert.NotNil(suite.T(), err)

	err = bindingMgr.RemoveBinding(bindingDelete, suite.remoteDevice)
	assert.NotNil(suite.T(), err)

	bindingDelete.ServerAddress = localFeature.Address()

	err = bindingMgr.RemoveBinding(bindingDelete, suite.remoteDevice)
	assert.NotNil(suite.T(), err)

	bindingDelete.ClientAddress = remoteFeature.Address()

	err = bindingMgr.RemoveBinding(bindingDelete, suite.remoteDevice)
	assert.Nil(suite.T(), err)

	subs = bindingMgr.Bindings(suite.remoteDevice)
	assert.Equal(suite.T(), 0, len(subs))

	err = bindingMgr.RemoveBinding(bindingDelete, suite.remoteDevice)
	assert.NotNil(suite.T(), err)

	err = bindingMgr.AddBinding(suite.remoteDevice, bindingRequest)
	assert.Nil(suite.T(), err)

	subs = bindingMgr.Bindings(suite.remoteDevice)
	assert.Equal(suite.T(), 1, len(subs))

	bindingMgr.RemoveBindingsForDevice(suite.remoteDevice)

	subs = bindingMgr.Bindings(suite.remoteDevice)
	assert.Equal(suite.T(), 0, len(subs))
}

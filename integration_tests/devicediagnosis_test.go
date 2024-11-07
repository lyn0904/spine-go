package integrationtests

import (
	"testing"

	"github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/spine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	dd_subscriptionRequestCall_recv_file_path        = "./testdata/dd_subscriptionRequestCall_recv.json"
	dd_subscriptionRequestCall_recv_result_file_path = "./testdata/ec_subscriptionRequestCall_recv_result.json"
)

func TestDeviceDiagnosisSuite(t *testing.T) {
	suite.Run(t, new(DeviceDiagnosisSuite))
}

type DeviceDiagnosisSuite struct {
	suite.Suite
	sut api.DeviceLocalInterface

	remoteSki string

	remoteDevice api.DeviceRemoteInterface
	writeHandler *WriteMessageHandler
}

func (s *DeviceDiagnosisSuite) BeforeTest(suiteName, testName string) {
	s.sut, s.remoteSki, s.remoteDevice, s.writeHandler = beforeTest(1, model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)

	// f.AddFunctionType(model.FunctionTypeDeviceDiagnosisHeartbeatData, true, false)

	initialCommunication(s.T(), s.remoteDevice, s.writeHandler)
}

func (s *DeviceDiagnosisSuite) TestHeartbeatSubscription_RecvNotify() {
	// Act
	msgCounter, _ := s.remoteDevice.HandleSpineMesssage(loadFileData(s.T(), dd_subscriptionRequestCall_recv_file_path))
	waitForAck(s.T(), msgCounter, s.writeHandler)

	// Assert
	remoteDevice := s.sut.RemoteDeviceForSki(s.remoteSki)
	assert.NotNil(s.T(), remoteDevice)

	ddFeature := remoteDevice.FeatureByEntityTypeAndRole(
		remoteDevice.Entity(spine.NewAddressEntityType([]uint{1})),
		model.FeatureTypeTypeDeviceDiagnosis,
		model.RoleTypeClient)
	assert.NotNil(s.T(), ddFeature)
}

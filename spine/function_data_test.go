package spine

import (
	"testing"

	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/util"
	"github.com/stretchr/testify/assert"
)

func TestFunctionData_UpdateData(t *testing.T) {
	newData := &model.DeviceClassificationManufacturerDataType{
		DeviceName: util.Ptr(model.DeviceClassificationStringType("device name")),
	}
	functionType := model.FunctionTypeDeviceClassificationManufacturerData
	sut := NewFunctionData[model.DeviceClassificationManufacturerDataType](functionType)
	_, _ = sut.UpdateData(false, true, newData, nil, nil)
	getData := sut.DataCopy()

	assert.Equal(t, newData.DeviceName, getData.DeviceName)
	assert.Equal(t, functionType, sut.FunctionType())

	// another update should not be reflected in the first dataset
	newData = &model.DeviceClassificationManufacturerDataType{
		DeviceName: util.Ptr(model.DeviceClassificationStringType("new device name")),
	}
	_, _ = sut.UpdateData(false, true, newData, nil, nil)
	getNewData := sut.DataCopy()

	assert.Equal(t, newData.DeviceName, getNewData.DeviceName)
	assert.NotEqual(t, getData.DeviceName, getNewData.DeviceName)
	assert.Equal(t, functionType, sut.FunctionType())

	_, _ = sut.UpdateDataAny(false, true, newData, nil, nil)
	getNewDataAny := sut.DataCopyAny()
	newDataAny := getNewDataAny.(*model.DeviceClassificationManufacturerDataType)

	assert.Equal(t, newData.DeviceName, newDataAny.DeviceName)
	assert.NotEqual(t, getData.DeviceName, newDataAny.DeviceName)
	assert.Equal(t, functionType, sut.FunctionType())
}

func TestFunctionData_UpdateDataPartial(t *testing.T) {
	newData := &model.ElectricalConnectionPermittedValueSetListDataType{
		ElectricalConnectionPermittedValueSetData: []model.ElectricalConnectionPermittedValueSetDataType{
			{
				ElectricalConnectionId: util.Ptr(model.ElectricalConnectionIdType(1)),
				ParameterId:            util.Ptr(model.ElectricalConnectionParameterIdType(1)),
				PermittedValueSet: []model.ScaledNumberSetType{
					{
						Range: []model.ScaledNumberRangeType{
							{
								Min: &model.ScaledNumberType{
									Number: util.Ptr(model.NumberType(6)),
									Scale:  util.Ptr(model.ScaleType(0)),
								},
							},
						},
					},
				},
			},
		},
	}
	functionType := model.FunctionTypeElectricalConnectionPermittedValueSetListData
	sut := NewFunctionData[model.ElectricalConnectionPermittedValueSetListDataType](functionType)

	_, err := sut.UpdateData(false, true, newData, &model.FilterType{CmdControl: &model.CmdControlType{Partial: &model.ElementTagType{}}}, nil)
	if assert.Nil(t, err) {
		getData := sut.DataCopy()
		assert.Equal(t, 1, len(getData.ElectricalConnectionPermittedValueSetData))
	}
}

func TestFunctionData_UpdateDataPartial_Supported(t *testing.T) {
	newData := &model.HvacOverrunListDataType{
		HvacOverrunData: []model.HvacOverrunDataType{
			{
				OverrunId: util.Ptr(model.HvacOverrunIdType(1)),
			},
		},
	}
	functionType := model.FunctionTypeHvacOverrunListData
	sut := NewFunctionData[model.HvacOverrunListDataType](functionType)

	ok := sut.SupportsPartialWrite()
	assert.True(t, ok)

	_, err := sut.UpdateData(false, true, newData, &model.FilterType{CmdControl: &model.CmdControlType{Partial: &model.ElementTagType{}}}, nil)
	assert.Nil(t, err)

	functionType = model.FunctionTypeNetworkManagementAddNodeCall
	sut2 := NewFunctionData[model.NetworkManagementAddNodeCallType](functionType)
	ok = sut2.SupportsPartialWrite()
	assert.False(t, ok)
}

package model

// HeatPumpPowerSwitchDescriptionListDataType
var _ Updater = (*HeatPumpPowerSwitchDescriptionListDataType)(nil)

func (r *HeatPumpPowerSwitchDescriptionListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpOperationPowerSwitchDescriptionDataType
	if newList != nil {
		newData = newList.(*HeatPumpPowerSwitchDescriptionListDataType).HeatPumpPowerSwitchDescriptionData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpPowerSwitchDescriptionData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpPowerSwitchDescriptionData = data
	}

	return data, success
}

// HeatPumpOperationPowerSwitchListDataType
var _ Updater = (*HeatPumpOperationPowerSwitchListDataType)(nil)

func (r *HeatPumpOperationPowerSwitchListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpOperationPowerSwitchDataType
	if newList != nil {
		newData = newList.(*HeatPumpOperationPowerSwitchListDataType).HeatPumpOperationPowerSwitchData
	}
	data, success := UpdateList(remoteWrite, r.HeatPumpOperationPowerSwitchData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpOperationPowerSwitchData = data
	}

	return data, success
}

//HeatPumpOperationModeDescriptionListDataType

var _ Updater = (*HeatPumpOperationModeDescriptionListDataType)(nil)

func (r HeatPumpOperationModeDescriptionListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpOperationModeDescriptionDataType
	if newList != nil {
		newData = newList.(*HeatPumpOperationModeDescriptionListDataType).HeatPumpOperationModeDescriptionData
	}
	data, success := UpdateList(remoteWrite, r.HeatPumpOperationModeDescriptionData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpOperationModeDescriptionData = data
	}

	return data, success
}

// HeatPumpOperationModeListDataType
var _ Updater = (*HeatPumpOperationModeListDataType)(nil)

func (r HeatPumpOperationModeListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpOperationModeDataType
	if newList != nil {
		newData = newList.(*HeatPumpOperationModeListDataType).HeatPumpOperationModeData
	}
	data, success := UpdateList(remoteWrite, r.HeatPumpOperationModeData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpOperationModeData = data
	}

	return data, success
}

// HeatPumpOperationWaterTemperatureDescriptionListDataType
var _ Updater = (*HeatPumpOperationWaterTemperatureDescriptionListDataType)(nil)

func (r HeatPumpOperationWaterTemperatureDescriptionListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpOperationWaterTemperatureDescriptionDataType
	if newList != nil {
		newData = newList.(*HeatPumpOperationWaterTemperatureDescriptionListDataType).HeatPumpOperationWaterTemperatureDescriptionData
	}
	data, success := UpdateList(remoteWrite, r.HeatPumpOperationWaterTemperatureDescriptionData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpOperationWaterTemperatureDescriptionData = data
	}

	return data, success
}

// HeatPumpWaterTemperatureListDataType
var _ Updater = (*HeatPumpWaterTemperatureListDataType)(nil)

func (r HeatPumpWaterTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpWaterTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpWaterTemperatureListDataType).HeatPumpWaterTemperatureData
	}
	data, success := UpdateList(remoteWrite, r.HeatPumpWaterTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpWaterTemperatureData = data
	}

	return data, success
}

// UpdateList HeatPumpRoomTemperatureDescriptionListDataType
var _ Updater = (*HeatPumpRoomTemperatureDescriptionListDataType)(nil)

func (r HeatPumpRoomTemperatureDescriptionListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpRoomTemperatureDescriptionDataType
	if newList != nil {
		newData = newList.(*HeatPumpRoomTemperatureDescriptionListDataType).HeatPumpRoomTemperatureDescriptionData
	}
	data, success := UpdateList(remoteWrite, r.HeatPumpRoomTemperatureDescriptionData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpRoomTemperatureDescriptionData = data
	}

	return data, success
}

// UpdateList HeatPumpRoomTemperatureListDataType
var _ Updater = (*HeatPumpRoomTemperatureListDataType)(nil)

func (r HeatPumpRoomTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpRoomTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpRoomTemperatureListDataType).HeatPumpRoomTemperatureData
	}
	data, success := UpdateList(remoteWrite, r.HeatPumpRoomTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpRoomTemperatureData = data
	}

	return data, success
}

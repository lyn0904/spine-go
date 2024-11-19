package model

var _ Updater = (*HeatPumpArea1WaterTempPowerSwitchListDataType)(nil)

func (r HeatPumpArea1WaterTempPowerSwitchListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpSwitchDataType
	if newList != nil {
		newData = newList.(*HeatPumpArea1WaterTempPowerSwitchListDataType).HeatPumpArea1WaterTempPowerSwitchData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpArea1WaterTempPowerSwitchData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpArea1WaterTempPowerSwitchData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpArea2WaterTempPowerSwitchListDataType)(nil)

func (r HeatPumpArea2WaterTempPowerSwitchListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpSwitchDataType
	if newList != nil {
		newData = newList.(*HeatPumpArea2WaterTempPowerSwitchListDataType).HeatPumpArea2WaterTempPowerSwitchData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpArea2WaterTempPowerSwitchData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpArea2WaterTempPowerSwitchData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpWaterTankPowerSwitchListDataType)(nil)

func (r HeatPumpWaterTankPowerSwitchListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpSwitchDataType
	if newList != nil {
		newData = newList.(*HeatPumpWaterTankPowerSwitchListDataType).HeatPumpWaterTankPowerSwitchData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpWaterTankPowerSwitchData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpWaterTankPowerSwitchData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpRoomTempPowerSwitchListDataType)(nil)

func (r HeatPumpRoomTempPowerSwitchListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpSwitchDataType
	if newList != nil {
		newData = newList.(*HeatPumpRoomTempPowerSwitchListDataType).HeatPumpRoomTempPowerSwitchData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpRoomTempPowerSwitchData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpRoomTempPowerSwitchData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpModeListDataType)(nil)

func (r HeatPumpModeListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpModeDataType
	if newList != nil {
		newData = newList.(*HeatPumpModeListDataType).HeatPumpModeData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpModeData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpModeData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpArea1WaterTemperatureListDataType)(nil)

func (r HeatPumpArea1WaterTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpArea1WaterTemperatureListDataType).HeatPumpArea1WaterTemperatureData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpArea1WaterTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpArea1WaterTemperatureData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpArea2WaterTemperatureListDataType)(nil)

func (r HeatPumpArea2WaterTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpArea2WaterTemperatureListDataType).HeatPumpArea2WaterTemperatureData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpArea2WaterTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpArea2WaterTemperatureData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpCurrentWaterTemperatureListDataType)(nil)

func (r HeatPumpCurrentWaterTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpCurrentWaterTemperatureListDataType).HeatPumpCurrentWaterTemperatureData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpCurrentWaterTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpCurrentWaterTemperatureData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpWaterTankTemperatureListDataType)(nil)

func (r HeatPumpWaterTankTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpWaterTankTemperatureListDataType).HeatPumpWaterTankTemperatureData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpWaterTankTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpWaterTankTemperatureData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpCurrentWaterTankTemperatureListDataType)(nil)

func (r HeatPumpCurrentWaterTankTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpCurrentWaterTankTemperatureListDataType).HeatPumpCurrentWaterTankTemperatureData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpCurrentWaterTankTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpCurrentWaterTankTemperatureData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpRoomTemperatureListDataType)(nil)

func (r HeatPumpRoomTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpRoomTemperatureListDataType).HeatPumpRoomTemperatureData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpRoomTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpRoomTemperatureData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpCurrentRoomTemperatureListDataType)(nil)

func (r HeatPumpCurrentRoomTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpCurrentRoomTemperatureListDataType).HeatPumpCurrentRoomTemperatureData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpCurrentRoomTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpCurrentRoomTemperatureData = data
	}

	return data, success
}

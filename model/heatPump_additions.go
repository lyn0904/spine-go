package model

var _ Updater = (*HeatPumpArea1WaterTemperaturePowerSwitchListDataType)(nil)

func (r HeatPumpArea1WaterTemperaturePowerSwitchListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpSwitchDataType
	if newList != nil {
		newData = newList.(*HeatPumpArea1WaterTemperaturePowerSwitchListDataType).HeatPumpArea1WaterTemperaturePowerSwitchData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpArea1WaterTemperaturePowerSwitchData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpArea1WaterTemperaturePowerSwitchData = data
	}

	return data, success
}

var _ Updater = (*HeatPumpArea2WaterTemperaturePowerSwitchListDataType)(nil)

func (r HeatPumpArea2WaterTemperaturePowerSwitchListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpSwitchDataType
	if newList != nil {
		newData = newList.(*HeatPumpArea2WaterTemperaturePowerSwitchListDataType).HeatPumpArea2WaterTemperaturePowerSwitchData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpArea2WaterTemperaturePowerSwitchData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpArea2WaterTemperaturePowerSwitchData = data
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

var _ Updater = (*HeatPumpRoomTemperaturePowerSwitchListDataType)(nil)

func (r HeatPumpRoomTemperaturePowerSwitchListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpSwitchDataType
	if newList != nil {
		newData = newList.(*HeatPumpRoomTemperaturePowerSwitchListDataType).HeatPumpRoomTemperaturePowerSwitchData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpRoomTemperaturePowerSwitchData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpRoomTemperaturePowerSwitchData = data
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

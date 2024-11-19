package model

var _ Updater = (*HeatPumpPowerSwitchListDataType)(nil)

func (r HeatPumpPowerSwitchListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpSwitchDataType
	if newList != nil {
		newData = newList.(*HeatPumpPowerSwitchListDataType).HeatPumpPowerSwitchData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpPowerSwitchData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpPowerSwitchData = data
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

var _ Updater = (*HeatPumpTemperatureListDataType)(nil)

func (r HeatPumpTemperatureListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpTemperatureDataType
	if newList != nil {
		newData = newList.(*HeatPumpTemperatureListDataType).HeatPumpTemperatureData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpTemperatureData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpTemperatureData = data
	}

	return data, success
}

package model

var _ Updater = (*HeatPumpOperationPowerSwitchDescriptionListDataType)(nil)

func (r HeatPumpOperationPowerSwitchDescriptionListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpOperationPowerSwitchDescriptionDataType
	if newList != nil {
		newData = newList.(*HeatPumpOperationPowerSwitchDescriptionListDataType).HeatPumpOperationPowerSwitchDescriptionData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpOperationPowerSwitchDescriptionData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpOperationPowerSwitchDescriptionData = data
	}

	return data, success
}

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

var _ Updater = (*HeatPumpOperationTemperatureDescriptionListDataType)(nil)

func (r HeatPumpOperationTemperatureDescriptionListDataType) UpdateList(remoteWrite, persist bool, newList any, filterPartial, filterDelete *FilterType) (any, bool) {
	var newData []HeatPumpOperationTemperatureDescriptionDataType
	if newList != nil {
		newData = newList.(*HeatPumpOperationTemperatureDescriptionListDataType).HeatPumpOperationTemperatureDescriptionData
	}

	data, success := UpdateList(remoteWrite, r.HeatPumpOperationTemperatureDescriptionData, newData, filterPartial, filterDelete)

	if success && persist {
		r.HeatPumpOperationTemperatureDescriptionData = data
	}

	return data, success
}

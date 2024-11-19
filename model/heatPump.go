package model

type HeatPumpIdType uint
type HeatPumpPowerSwitchType bool

type HeatPumpSwitchDataType struct {
	Id          *HeatPumpIdType          `json:"id,omitempty" eebus:"key"`
	PowerSwitch *HeatPumpPowerSwitchType `json:"powerSwitch,omitempty"`
	Description *DescriptionType         `json:"description,omitempty"`
}

type HeatPumpModeTypeType string

const (
	HeatPumpModeTypeTypeAuto    HeatPumpModeTypeType = "auto"
	HeatPumpModeTypeTypeEco     HeatPumpModeTypeType = "eco"
	HeatPumpModeTypeTypeHeating HeatPumpModeTypeType = "heating"
	HeatPumpModeTypeTypeCold    HeatPumpModeTypeType = "cold"
)

type HeatPumpModeDataType struct {
	Id          *HeatPumpIdType       `json:"id,omitempty" eebus:"key"`
	Mode        *HeatPumpModeTypeType `json:"heatPumpMode,omitempty"`
	Description *DescriptionType      `json:"description,omitempty"`
}

type HeatPumpTemperatureType uint

type HeatPumpTemperatureDataType struct {
	Id          *HeatPumpIdType
	Temperature *HeatPumpTemperatureType
	Description *DescriptionType
}

type HeatPumpArea1WaterTempPowerSwitchListDataType struct {
	HeatPumpArea1WaterTempPowerSwitchData []HeatPumpSwitchDataType `json:"heatPumpArea1WaterTempPowerSwitchData,omitempty"`
}

type HeatPumpArea2WaterTempPowerSwitchListDataType struct {
	HeatPumpArea2WaterTempPowerSwitchData []HeatPumpSwitchDataType `json:"heatPumpArea2WaterTempPowerSwitchData,omitempty"`
}

type HeatPumpWaterTankPowerSwitchListDataType struct {
	HeatPumpWaterTankPowerSwitchData []HeatPumpSwitchDataType `json:"heatPumpWaterTankPowerSwitchData,omitempty"`
}

type HeatPumpRoomTempPowerSwitchListDataType struct {
	HeatPumpRoomTempPowerSwitchData []HeatPumpSwitchDataType `json:"heatPumpRoomTempPowerSwitchData,omitempty"`
}

type HeatPumpModeListDataType struct {
	HeatPumpModeData []HeatPumpModeDataType `json:"heatPumpModeData,omitempty"`
}

type HeatPumpArea1WaterTemperatureListDataType struct {
	HeatPumpArea1WaterTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpArea1WaterTemperatureData,omitempty"`
}

type HeatPumpArea2WaterTemperatureListDataType struct {
	HeatPumpArea2WaterTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpArea2WaterTemperatureData,omitempty"`
}

type HeatPumpCurrentWaterTemperatureListDataType struct {
	HeatPumpCurrentWaterTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpCurrentWaterTemperatureData,omitempty"`
}

type HeatPumpWaterTankTemperatureListDataType struct {
	HeatPumpWaterTankTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpWaterTankTemperatureData,omitempty"`
}

type HeatPumpCurrentWaterTankTemperatureListDataType struct {
	HeatPumpCurrentWaterTankTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpCurrentWaterTankTemperatureData,omitempty"`
}

type HeatPumpRoomTemperatureListDataType struct {
	HeatPumpRoomTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpRoomTemperatureData,omitempty"`
}

type HeatPumpCurrentRoomTemperatureListDataType struct {
	HeatPumpCurrentRoomTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpCurrentRoomTemperatureData,omitempty"`
}

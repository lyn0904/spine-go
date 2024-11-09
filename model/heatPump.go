package model

type HeatPumpIdType uint
type HeatPumpPowerSwitchType bool

type HeatPumpSwitchDataType struct {
	Id          *HeatPumpIdType          `json:"id,omitempty" eebus:"key"'`
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
	Id          *HeatPumpIdType       `json:"id,omitempty" eebus:"key"'`
	Mode        *HeatPumpModeTypeType `json:"heatPumpMode,omitempty"`
	Description *DescriptionType      `json:"description,omitempty"`
}

type HeatPumpTemperatureType uint

type HeatPumpTemperatureDataType struct {
	Id          *HeatPumpIdType
	Temperature *HeatPumpTemperatureType
	Description *DescriptionType
}

// HeatPumpArea1WaterTemperaturePowerSwitchListDataType 区域1水温开关
type HeatPumpArea1WaterTemperaturePowerSwitchListDataType struct {
	HeatPumpArea1WaterTemperaturePowerSwitchData []HeatPumpSwitchDataType `json:"heatPumpArea1WaterTemperaturePowerSwitchData,omitempty"`
}

// HeatPumpArea2WaterTemperaturePowerSwitchListDataType 区域2水温开关
type HeatPumpArea2WaterTemperaturePowerSwitchListDataType struct {
	HeatPumpArea2WaterTemperaturePowerSwitchData []HeatPumpSwitchDataType `json:"heatPumpArea2WaterTemperaturePowerSwitchData,omitempty"`
}

// HeatPumpWaterTankPowerSwitchListDataType 水箱开关
type HeatPumpWaterTankPowerSwitchListDataType struct {
	HeatPumpWaterTankPowerSwitchData []HeatPumpSwitchDataType `json:"heatPumpWaterTankPowerSwitchData,omitempty"`
}

// HeatPumpRoomTemperaturePowerSwitchListDataType 室温开关
type HeatPumpRoomTemperaturePowerSwitchListDataType struct {
	HeatPumpRoomTemperaturePowerSwitchData []HeatPumpSwitchDataType `json:"heatPumpRoomTemperaturePowerSwitchData,omitempty"`
}

// HeatPumpModeListDataType 模式
type HeatPumpModeListDataType struct {
	HeatPumpModeData []HeatPumpModeDataType `json:"heatPumpModeData,omitempty"`
}

// HeatPumpArea1WaterTemperatureListDataType 区域1水温
type HeatPumpArea1WaterTemperatureListDataType struct {
	HeatPumpArea1WaterTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpArea1WaterTemperatureData,omitempty"`
}

// HeatPumpArea2WaterTemperatureListDataType 区域2水温
type HeatPumpArea2WaterTemperatureListDataType struct {
	HeatPumpArea2WaterTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpArea2WaterTemperatureData,omitempty"`
}

// HeatPumpRoomTemperatureListDataType 室温
type HeatPumpRoomTemperatureListDataType struct {
	HeatPumpRoomTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpRoomTemperatureData,omitempty"`
}

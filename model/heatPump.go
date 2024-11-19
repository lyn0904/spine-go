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

// HeatPumpPowerSwitchListDataType 开关
type HeatPumpPowerSwitchListDataType struct {
	HeatPumpPowerSwitchData []HeatPumpSwitchDataType `json:"HeatPumpPowerSwitchData,omitempty"`
}

// HeatPumpModeListDataType 模式
type HeatPumpModeListDataType struct {
	HeatPumpModeData []HeatPumpModeDataType `json:"heatPumpModeData,omitempty"`
}

// HeatPumpTemperatureListDataType 温度
type HeatPumpTemperatureListDataType struct {
	HeatPumpTemperatureData []HeatPumpTemperatureDataType `json:"heatPumpTemperatureData,omitempty"`
}

package model

type HeatPumpDescription string

type HeatPumpPowerSwitchId uint
type HeatPumpPowerSwitch bool

type HeatPumpPowerSwitchDescriptionDataType struct {
	Id          *HeatPumpPowerSwitchId
	PowerSwitch *HeatPumpPowerSwitch `json:"powerSwitch,omitempty"`
	Description *HeatPumpDescription
}

type HeatPumpPowerSwitchDescriptionListDataType struct {
	HeatPumpPowerSwitchDescriptionData []HeatPumpPowerSwitchDescriptionDataType `json:"heatPumpPowerSwitchDescriptionData,omitempty"`
}

type HeatPumpPowerSwitchDataType struct {
	Id          *HeatPumpPowerSwitchId
	PowerSwitch *HeatPumpPowerSwitch `json:"powerSwitch,omitempty"`
}

type HeatPumpModeId uint
type HeatPumpMode string

const (
	HeatPumpModeAuto    HeatPumpMode = "auto"
	HeatPumpModeEco     HeatPumpMode = "eco"
	HeatPumpModeHeating HeatPumpMode = "heating"
	HeatPumpModeCold    HeatPumpMode = "cold"
)

type HeatPumpOperationModeDescriptionDataType struct {
	ModeId      *HeatPumpModeId      `json:"modeId,omitempty"`
	Mode        *HeatPumpMode        `json:"mode,omitempty"`
	Description *HeatPumpDescription `json:"description,omitempty"`
}

type HeatPumpOperationModeDescriptionListDataType struct {
	HeatPumpOperationModeDescriptionData []HeatPumpOperationModeDescriptionDataType `json:"heatPumpOperationModeDescriptionData,omitempty"`
}

type HeatPumpOperationModeDataType struct {
	Mode *HeatPumpMode `json:"mode,omitempty"`
}

type HeatPumpWaterTemperatureId uint
type HeatPumpWaterMinTemperature uint
type HeatPumpWaterMaxTemperature uint

type HeatPumpWaterTemperatureDescriptionDataType struct {
	Id          *HeatPumpWaterTemperatureId  `json:"id,omitempty"`
	Max         *HeatPumpWaterMaxTemperature `json:"max,omitempty"`
	Min         *HeatPumpWaterMaxTemperature `json:"min,omitempty"`
	Description *HeatPumpDescription         `json:"description,omitempty"`
}

type HeatPumpWaterTemperatureDescriptionListDataType struct {
	HeatPumpWaterTemperatureDescriptionData []HeatPumpWaterTemperatureDescriptionDataType `json:"heatPumpWaterTemperatureDescriptionData,omitempty"`
}

type HeatPumpWaterTemperature uint

type HeatPumpWaterTemperatureDataType struct {
	Id    *HeatPumpWaterTemperatureId `json:"waterTemperatureId,omitempty"`
	Value *HeatPumpWaterTemperature   `json:"value,omitempty"`
}

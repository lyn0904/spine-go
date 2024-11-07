package model

type HeatPumpDescription string

type HeatPumpPowerSwitchId uint
type HeatPumpPowerSwitch bool

type HeatPumpOperationPowerSwitchDescriptionDataType struct {
	Id          *HeatPumpPowerSwitchId
	PowerSwitch *HeatPumpPowerSwitch `json:"powerSwitch,omitempty"`
	Description *HeatPumpDescription
}

type HeatPumpPowerSwitchDescriptionListDataType struct {
	HeatPumpPowerSwitchDescriptionData []HeatPumpOperationPowerSwitchDescriptionDataType `json:"heatPumpPowerSwitchDescriptionData,omitempty"`
}

type HeatPumpOperationPowerSwitchDataType struct {
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
	Id          *HeatPumpModeId      `json:"modeId,omitempty"`
	Mode        *HeatPumpMode        `json:"mode,omitempty"`
	Description *HeatPumpDescription `json:"description,omitempty"`
}

type HeatPumpOperationModeDescriptionListDataType struct {
	HeatPumpOperationModeDescriptionData []HeatPumpOperationModeDescriptionDataType `json:"heatPumpOperationModeDescriptionData,omitempty"`
}

type HeatPumpOperationModeDataType struct {
	Id   *HeatPumpModeId `json:"modeId,omitempty"`
	Mode *HeatPumpMode   `json:"mode,omitempty"`
}

type HeatPumpWaterTemperatureId uint
type HeatPumpWaterMinTemperature uint
type HeatPumpWaterMaxTemperature uint

type HeatPumpOperationWaterTemperatureDescriptionDataType struct {
	Id          *HeatPumpWaterTemperatureId  `json:"id,omitempty"`
	Max         *HeatPumpWaterMaxTemperature `json:"max,omitempty"`
	Min         *HeatPumpWaterMinTemperature `json:"min,omitempty"`
	Description *HeatPumpDescription         `json:"description,omitempty"`
}

type HeatPumpOperationWaterTemperatureDescriptionListDataType struct {
	HeatPumpOperationWaterTemperatureDescriptionData []HeatPumpOperationWaterTemperatureDescriptionDataType `json:"heatPumpWaterTemperatureDescriptionData,omitempty"`
}

type HeatPumpWaterTemperature uint

type HeatPumpOperationWaterTemperatureDataType struct {
	Id    *HeatPumpWaterTemperatureId `json:"waterTemperatureId,omitempty"`
	Value *HeatPumpWaterTemperature   `json:"value,omitempty"`
}

package model

type HeatPumpDescription string

type HeatPumpPowerSwitchId uint
type HeatPumpPowerSwitch bool

type HeatPumpOperationPowerSwitchDescriptionDataType struct {
	Id          *HeatPumpPowerSwitchId `json:"id,omitempty" eebus:"key"`
	PowerSwitch *HeatPumpPowerSwitch   `json:"powerSwitch,omitempty"`
	Description *HeatPumpDescription   `json:"description,omitempty"`
}

type HeatPumpPowerSwitchDescriptionListDataType struct {
	HeatPumpPowerSwitchDescriptionData []HeatPumpOperationPowerSwitchDescriptionDataType `json:"heatPumpPowerSwitchDescriptionData,omitempty"`
}

type HeatPumpOperationPowerSwitchDataType struct {
	Id          *HeatPumpPowerSwitchId `json:"id,omitempty" eebus:"key"`
	PowerSwitch *HeatPumpPowerSwitch   `json:"powerSwitch,omitempty"`
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
	Id          *HeatPumpModeId      `json:"id,omitempty" eebus:"key"`
	Mode        *HeatPumpMode        `json:"mode,omitempty"`
	Description *HeatPumpDescription `json:"description,omitempty"`
}

type HeatPumpOperationModeDescriptionListDataType struct {
	HeatPumpOperationModeDescriptionData []HeatPumpOperationModeDescriptionDataType `json:"heatPumpOperationModeDescriptionData,omitempty"`
}

type HeatPumpOperationModeDataType struct {
	Id   *HeatPumpModeId `json:"id,omitempty" eebus:"key"`
	Mode *HeatPumpMode   `json:"mode,omitempty"`
}

type HeatPumpWaterTemperatureId uint
type HeatPumpWaterMinTemperature uint
type HeatPumpWaterMaxTemperature uint

type HeatPumpOperationWaterTemperatureDescriptionDataType struct {
	Id          *HeatPumpWaterTemperatureId  `json:"id,omitempty" eebus:"key"`
	Max         *HeatPumpWaterMaxTemperature `json:"max,omitempty"`
	Min         *HeatPumpWaterMinTemperature `json:"min,omitempty"`
	Description *HeatPumpDescription         `json:"description,omitempty"`
}

type HeatPumpOperationWaterTemperatureDescriptionListDataType struct {
	HeatPumpOperationWaterTemperatureDescriptionData []HeatPumpOperationWaterTemperatureDescriptionDataType `json:"heatPumpWaterTemperatureDescriptionData,omitempty"`
}

type HeatPumpWaterTemperature uint

type HeatPumpWaterTemperatureDataType struct {
	Id    *HeatPumpWaterTemperatureId `json:"id,omitempty" eebus:"key"`
	Value *HeatPumpWaterTemperature   `json:"value,omitempty"`
}

type HeatPumpRoomTemperatureId uint
type HeatPumpRoomTemperature uint
type HeatPumpRoomMinTemperature uint
type HeatPumpRoomMaxTemperature uint

type HeatPumpRoomTemperatureDescriptionDataType struct {
	Id          *HeatPumpRoomTemperatureId  `json:"id,omitempty" eebus:"key"`
	Max         *HeatPumpRoomMaxTemperature `json:"max,omitempty"`
	Min         *HeatPumpRoomMinTemperature `json:"min,omitempty"`
	Description *HeatPumpDescription        `json:"description,omitempty"`
}

type HeatPumpRoomTemperatureDescriptionListDataType struct {
	HeatPumpRoomTemperatureDescriptionData []HeatPumpRoomTemperatureDescriptionDataType `json:"heatPumpRoomTemperatureDescriptionData,omitempty"`
}

type HeatPumpRoomTemperatureDataType struct {
	Id    *HeatPumpRoomTemperatureId `json:"id,omitempty" eebus:"key"`
	Value *HeatPumpRoomTemperature   `json:"value,omitempty"`
}

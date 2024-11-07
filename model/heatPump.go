package model

type HeatPumpDescriptionType string

type HeatPumpPowerSwitchIdType uint
type HeatPumpPowerSwitchType bool

type HeatPumpOperationPowerSwitchDescriptionDataType struct {
	HeatPumpPowerSwitchId *HeatPumpPowerSwitchIdType `json:"heatPumpPowerSwitchId,omitempty" eebus:"key"`
	HeatPumpPowerSwitch   *HeatPumpPowerSwitchType   `json:"heatPumpPowerSwitch,omitempty"`
	HeatPumpDescription   *HeatPumpDescriptionType   `json:"heatPumpDescription,omitempty"`
}

type HeatPumpPowerSwitchDescriptionListDataType struct {
	HeatPumpPowerSwitchDescriptionData []HeatPumpOperationPowerSwitchDescriptionDataType `json:"heatPumpPowerSwitchDescriptionData,omitempty"`
}

type HeatPumpOperationPowerSwitchDataType struct {
	HeatPumpPowerSwitchId *HeatPumpPowerSwitchIdType `json:"heatPumpPowerSwitchId,omitempty" eebus:"key"`
	HeatPumpPowerSwitch   *HeatPumpPowerSwitchType   `json:"heatPumpPowerSwitch,omitempty"`
}

type HeatPumpOperationPowerSwitchListDataType struct {
	HeatPumpOperationPowerSwitchData []HeatPumpOperationPowerSwitchDataType
}

type HeatPumpModeIdType uint
type HeatPumpModeType string

const (
	HeatPumpModeAuto    HeatPumpModeType = "auto"
	HeatPumpModeEco     HeatPumpModeType = "eco"
	HeatPumpModeHeating HeatPumpModeType = "heating"
	HeatPumpModeCold    HeatPumpModeType = "cold"
)

type HeatPumpOperationModeDescriptionDataType struct {
	HeatPumpModeId      *HeatPumpModeIdType      `json:"heatPumpModeId,omitempty" eebus:"key"`
	HeatPumpMode        *HeatPumpModeType        `json:"heatPumpMode,omitempty"`
	HeatPumpDescription *HeatPumpDescriptionType `json:"heatPumpDescription,omitempty"`
}

type HeatPumpOperationModeDescriptionListDataType struct {
	HeatPumpOperationModeDescriptionData []HeatPumpOperationModeDescriptionDataType `json:"heatPumpOperationModeDescriptionData,omitempty"`
}

type HeatPumpOperationModeDataType struct {
	HeatPumpModeId *HeatPumpModeIdType `json:"heatPumpModeId,omitempty" eebus:"key"`
	HeatPumpMode   *HeatPumpModeType   `json:"heatPumpMode,omitempty"`
}

type HeatPumpOperationModeListDataType struct {
	HeatPumpOperationModeData []HeatPumpOperationModeDataType
}

type HeatPumpWaterTemperatureIdType uint
type HeatPumpWaterMinTemperatureType uint
type HeatPumpWaterMaxTemperatureType uint

type HeatPumpOperationWaterTemperatureDescriptionDataType struct {
	HeatPumpWaterTemperatureId  *HeatPumpWaterTemperatureIdType  `json:"heatPumpWaterTemperatureId,omitempty" eebus:"key"`
	HeatPumpWaterMaxTemperature *HeatPumpWaterMaxTemperatureType `json:"heatPumpWaterMaxTemperature,omitempty"`
	HeatPumpWaterMinTemperature *HeatPumpWaterMinTemperatureType `json:"heatPumpWaterMinTemperature,omitempty"`
	HeatPumpDescription         *HeatPumpDescriptionType         `json:"heatPumpDescription,omitempty"`
}

type HeatPumpOperationWaterTemperatureDescriptionListDataType struct {
	HeatPumpOperationWaterTemperatureDescriptionData []HeatPumpOperationWaterTemperatureDescriptionDataType `json:"heatPumpWaterTemperatureDescriptionData,omitempty"`
}

type HeatPumpWaterTemperatureType uint

type HeatPumpWaterTemperatureDataType struct {
	HeatPumpWaterTemperatureId *HeatPumpWaterTemperatureIdType `json:"heatPumpWaterTemperatureId,omitempty" eebus:"key"`
	HeatPumpWaterTemperature   *HeatPumpWaterTemperatureType   `json:"heatPumpWaterTemperature,omitempty"`
}

type HeatPumpWaterTemperatureListDataType struct {
	HeatPumpWaterTemperatureData []HeatPumpWaterTemperatureDataType
}

type HeatPumpRoomTemperatureIdType uint
type HeatPumpRoomTemperatureType uint
type HeatPumpRoomMinTemperatureType uint
type HeatPumpRoomMaxTemperatureType uint

type HeatPumpRoomTemperatureDescriptionDataType struct {
	HeatPumpRoomTemperatureId  *HeatPumpRoomTemperatureIdType  `json:"heatPumpRoomTemperatureId,omitempty" eebus:"key"`
	HeatPumpRoomMaxTemperature *HeatPumpRoomMaxTemperatureType `json:"heatPumpRoomMaxTemperature,omitempty"`
	HeatPumpRoomMinTemperature *HeatPumpRoomMinTemperatureType `json:"heatPumpRoomMinTemperature,omitempty"`
	HeatPumpDescription        *HeatPumpDescriptionType        `json:"heatPumpDescription,omitempty"`
}

type HeatPumpRoomTemperatureDescriptionListDataType struct {
	HeatPumpRoomTemperatureDescriptionData []HeatPumpRoomTemperatureDescriptionDataType `json:"heatPumpRoomTemperatureDescriptionData,omitempty"`
}

type HeatPumpRoomTemperatureDataType struct {
	HeatPumpRoomTemperatureId *HeatPumpRoomTemperatureIdType `json:"heatPumpRoomTemperatureId,omitempty" eebus:"key"`
	HeatPumpRoomTemperature   *HeatPumpRoomTemperatureType   `json:"heatPumpRoomTemperature,omitempty"`
}

type HeatPumpRoomTemperatureListDataType struct {
	HeatPumpRoomTemperatureData []HeatPumpRoomTemperatureDataType
}

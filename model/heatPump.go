package model

type HeatPumpIdType uint
type HeatPumpPowerSwitchType bool
type HeatPumpModeTypeType string

const (
	HeatPumpModeTypeTypeAuto    HeatPumpModeTypeType = "auto"    //自动
	HeatPumpModeTypeTypeEco     HeatPumpModeTypeType = "eco"     //节能
	HeatPumpModeTypeTypeHeating HeatPumpModeTypeType = "heating" //制热
	HeatPumpModeTypeTypeCold    HeatPumpModeTypeType = "cold"    //制冷
)

type HeatPumpTemperatureType uint

type HeatPumpOperationPowerSwitchDescriptionDataType struct {
	Id          *HeatPumpIdType          `json:"id,omitempty" eebus:"key"`
	PowerSwitch *HeatPumpPowerSwitchType `json:"powerSwitch,omitempty"`
	Label       *LabelType               `json:"label,omitempty"`
	Description *DescriptionType         `json:"description,omitempty"`
}

type HeatPumpOperationModeDescriptionDataType struct {
	Id          *HeatPumpIdType       `json:"id,omitempty" eebus:"key"`
	Mode        *HeatPumpModeTypeType `json:"heatPumpMode,omitempty"`
	Label       *LabelType            `json:"label,omitempty"`
	Description *DescriptionType      `json:"description,omitempty"`
}

type HeatPumpOperationTemperatureDescriptionDataType struct {
	Id          *HeatPumpIdType          `json:"id,omitempty" eebus:"key"`
	Temperature *HeatPumpTemperatureType `json:"temperature,omitempty"`
	Label       *LabelType               `json:"label,omitempty"`
	Description *DescriptionType         `json:"description,omitempty"`
}

type HeatPumpOperationPowerSwitchDescriptionListDataType struct {
	HeatPumpOperationPowerSwitchDescriptionData []HeatPumpOperationPowerSwitchDescriptionDataType `json:"heatPumpOperationPowerSwitchDescriptionData,omitempty"`
}

type HeatPumpOperationModeDescriptionListDataType struct {
	HeatPumpOperationModeDescriptionData []HeatPumpOperationModeDescriptionDataType `json:"heatPumpOperationModeDescriptionData,omitempty"`
}

type HeatPumpOperationTemperatureDescriptionListDataType struct {
	HeatPumpOperationTemperatureDescriptionData []HeatPumpOperationTemperatureDescriptionDataType `json:"heatPumpOperationTemperatureDescriptionData,omitempty"`
}

type HeatPumpOperationPowerSwitchDescriptionListDataSelectorsType struct {
	Id *HeatPumpIdType `json:"id,omitempty"`
}

type HeatPumpOperationModeDescriptionListDataSelectorsType struct {
	Id *HeatPumpIdType `json:"id,omitempty"`
}

type HeatPumpOperationTemperatureDescriptionListDataSelectorsType struct {
	Id *HeatPumpIdType `json:"id,omitempty"`
}

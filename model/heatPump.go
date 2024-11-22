package model

type HeatPumpIdType uint
type HeatPumpPowerSwitchType bool
type HeatPumpModeTypeType string
type HeatPumpMinTemperature uint
type HeatPumpMaxTemperature uint
type HeatPumpTemperatureUnit string

const (
	HeatPumpModeTypeTypeAuto    HeatPumpModeTypeType = "auto"    //自动
	HeatPumpModeTypeTypeEco     HeatPumpModeTypeType = "eco"     //节能
	HeatPumpModeTypeTypeHeating HeatPumpModeTypeType = "heating" //制热
	HeatPumpModeTypeTypeCold    HeatPumpModeTypeType = "cold"    //制冷
)

const (
	HeatPumpTemperatureUnitC HeatPumpTemperatureUnit = "c"
	HeatPumpTemperatureUnitF HeatPumpTemperatureUnit = "f"
)

type HeatPumpTemperatureType uint

type HeatPumpPowerSwitchDescriptionDataType struct {
	Id          *HeatPumpIdType          `json:"id,omitempty" eebus:"key"`
	PowerSwitch *HeatPumpPowerSwitchType `json:"powerSwitch,omitempty"`
	Label       *LabelType               `json:"label,omitempty"`
	Description *DescriptionType         `json:"description,omitempty"`
}

type HeatPumpModeDescriptionDataType struct {
	Id          *HeatPumpIdType       `json:"id,omitempty" eebus:"key"`
	Mode        *HeatPumpModeTypeType `json:"heatPumpMode,omitempty"`
	Label       *LabelType            `json:"label,omitempty"`
	Description *DescriptionType      `json:"description,omitempty"`
}

type HeatPumpTemperatureDescriptionDataType struct {
	Id          *HeatPumpIdType          `json:"id,omitempty" eebus:"key"`
	Temperature *HeatPumpTemperatureType `json:"temperature,omitempty"`
	Unit        *HeatPumpTemperatureUnit `json:"unit,omitempty"`
	Min         *HeatPumpMinTemperature  `json:"min,omitempty"`
	Max         *HeatPumpMaxTemperature  `json:"Max,omitempty"`
	Label       *LabelType               `json:"label,omitempty"`
	Description *DescriptionType         `json:"description,omitempty"`
}

type HeatPumpOperationPowerSwitchDescriptionListDataType struct {
	HeatPumpOperationPowerSwitchDescriptionData []HeatPumpPowerSwitchDescriptionDataType `json:"heatPumpOperationPowerSwitchDescriptionData,omitempty"`
}

type HeatPumpOperationModeDescriptionListDataType struct {
	HeatPumpOperationModeDescriptionData []HeatPumpModeDescriptionDataType `json:"heatPumpOperationModeDescriptionData,omitempty"`
}

type HeatPumpOperationTemperatureDescriptionListDataType struct {
	HeatPumpOperationTemperatureDescriptionData []HeatPumpTemperatureDescriptionDataType `json:"heatPumpOperationTemperatureDescriptionData,omitempty"`
}

type HeatPumpCurrentTemperatureDescriptionListDataType struct {
	HeatPumpCurrentTemperatureDescriptionData []HeatPumpTemperatureDescriptionDataType `json:"heatPumpCurrentTemperatureDescriptionData,omitempty"`
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

type HeatPumpCurrentTemperatureDescriptionListDataSelectorsType struct {
	Id *HeatPumpIdType `json:"id,omitempty"`
}

type HeatPumpOperationPowerSwitchDescriptionDataElementsType struct {
	Id          *ElementTagType `json:"id,omitempty"`
	PowerSwitch *ElementTagType `json:"powerSwitch,omitempty"`
	Label       *ElementTagType `json:"label,omitempty"`
	Description *ElementTagType `json:"description,omitempty"`
}

type HeatPumpOperationModeDescriptionDataElementsType struct {
	Id          *ElementTagType `json:"id,omitempty"`
	Mode        *ElementTagType `json:"heatPumpMode,omitempty"`
	Label       *ElementTagType `json:"label,omitempty"`
	Description *ElementTagType `json:"description,omitempty"`
}

type HeatPumpOperationTemperatureDescriptionDataElementsType struct {
	Id          *ElementTagType `json:"id,omitempty"`
	Temperature *ElementTagType `json:"temperature,omitempty"`
	Unit        *ElementTagType `json:"unit,omitempty"`
	Min         *ElementTagType `json:"min,omitempty"`
	Max         *ElementTagType `json:"Max,omitempty"`
	Label       *ElementTagType `json:"label,omitempty"`
	Description *ElementTagType `json:"description,omitempty"`
}

type HeatPumpCurrentTemperatureDescriptionDataElementsType struct {
	Id          *ElementTagType `json:"id,omitempty"`
	Temperature *ElementTagType `json:"temperature,omitempty"`
	Unit        *ElementTagType `json:"unit,omitempty"`
	Min         *ElementTagType `json:"min,omitempty"`
	Max         *ElementTagType `json:"Max,omitempty"`
	Label       *ElementTagType `json:"label,omitempty"`
	Description *ElementTagType `json:"description,omitempty"`
}

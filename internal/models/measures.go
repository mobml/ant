package models

type MeasureType string

const (
	MeasureBoolean  MeasureType = "boolean"
	MeasureInteger  MeasureType = "integer"
	MeasureFloat    MeasureType = "float"
	MeasureHours    MeasureType = "hours"
	MeasureDistance MeasureType = "distance"
	MeasureCustom   MeasureType = "custom"
)

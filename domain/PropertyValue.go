package domain

type PropertyValue struct {
	Id             float64 `json:"id" gorm:"primaryKey"`
	Value          string  `json:"value" `
	ScenarioId     string  `json:"scenarioId"`
	OrganizationId string  `json:"organizationId"`
	BusinessUnitId string  `json:"businessUnitId"`
}

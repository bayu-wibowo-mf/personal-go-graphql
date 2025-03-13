package model

type EmployeeDemography struct {
	DateKey            string  `json:"date_key" pg:"date_key, pk"`
	TenantUUID         string  `json:"tenant_uuid" pg:"tenant_uuid, pk"`
	DemographyCategory string  `json:"demography_category" pg:"demography_category, pk"`
	DemographyName     string  `json:"demography_name" pg:"demography_name, pk"`
	DemographyValue    float64 `json:"demography_Value" pg:"demography_value"`
}

type Tabler interface {
	TableName() string
}

func (EmployeeDemography) TableName() string {
	return "dm_employee_demography"
}

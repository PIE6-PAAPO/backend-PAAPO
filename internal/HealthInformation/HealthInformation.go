package healthinformation

type HealthInformation struct {
	Weight                    float64 `json:"weight"`
	Height                    float64 `json:"height"`
	Smoker                    bool    `json:"smoker"`
	AlcoholConsumption        string  `json:"alcohol_consumption"`
	PhysicalActivityFrequency string  `json:"physical_activity_frequency"`
	PhysicalActivityType      string  `json:"physical_activity_type"`
	UserID                    string  `gorm:"uniqueIndex" json:"user_id"`
}

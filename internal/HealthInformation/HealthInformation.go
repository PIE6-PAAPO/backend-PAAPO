package healthinformation

import "gorm.io/gorm"

type HealthInformation struct {
	gorm.Model
	Weight                    float64 `json:"weight" gorm:"not null"`
	Height                    float64 `json:"height" gorm:"not null"`
	Smoker                    bool    `json:"smoker" gorm:"not null;default:false"`
	AlcoholConsumption        string  `json:"alcohol_consumption" gorm:"type:varchar(255)"`
	PhysicalActivityFrequency string  `json:"physical_activity_frequency" gorm:"type:varchar(255)"`
	PhysicalActivityType      string  `json:"physical_activity_type" gorm:"type:varchar(255)"`
	UserID                    string  `gorm:"uniqueIndex;not null" json:"user_id"`
}

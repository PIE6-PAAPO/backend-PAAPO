package medicaldata

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Sequela struct {
	RestrictMobity  bool `json:"restrict_mobility" gorm:"not null;default:false"`
	JointPain       bool `json:"joint_pain" gorm:"not null;default:false"`
	ShoulderArmPain bool `json:"shoulder_arm_pain" gorm:"not null;default:false"`
	Neuromas        bool `json:"neuromas" gorm:"not null;default:false"`
	CervicalPain    bool `json:"cervical_pain" gorm:"not null;default:false"`
	BackPain        bool `json:"back_pain" gorm:"not null;default:false"`
	Swelling        bool `json:"swelling" gorm:"not null;default:false"`
	Fatigue         bool `json:"fatigue" gorm:"not null;default:false"`
}

type MedicalData struct {
	gorm.Model
	UserID                    string         `gorm:"uniqueIndex;not null" json:"user_id"`
	DiagnosisDate             *time.Time     `json:"diagnosis_date"`
	SurgeryDate               *time.Time     `json:"surgery_date"`
	SurgeryType               string         `json:"surgery_type" gorm:"type:varchar(255)"`
	PhysiotherapyReferral     bool           `json:"physiotherapy_referral" gorm:"not null;default:false"`
	PhysiotherapyDuration     *time.Time     `json:"physiotherapy_duration"`
	SufficientRecovery        bool           `json:"sufficient_recovery" gorm:"not null;default:false"`
	PresentSequela            Sequela        `json:"present_sequela" gorm:"embedded;embeddedPrefix:present_"`
	RemainingSequela          Sequela        `json:"remaining_sequela" gorm:"embedded;embeddedPrefix:remaining_"`
	ImpactOnDailyLife         string         `json:"impact_on_daily_life" gorm:"type:text"`
	AffectsIndependence       bool           `json:"affects_independence" gorm:"not null;default:false"`
	PostSurgeryActivities     string         `json:"post_surgery_activities" gorm:"type:text"`
	MedicalConditions         datatypes.JSON `json:"medical_conditions" gorm:"type:jsonb"`
	OtherDiagnoses            string         `json:"other_diagnoses" gorm:"type:text"`
	Medications               string         `json:"medications" gorm:"type:text"`
}

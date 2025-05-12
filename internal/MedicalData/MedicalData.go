package medicaldata

import "time"

type Sequela struct {
	RestrictMobity  bool `json:"restrict_mobility"`
	JointPain       bool `json:"joint_pain"`
	ShoulderArmPain bool `json:"shoulder_arm_pain"`
	Neuromas        bool `json:"neuromas"`
	CervicalPain    bool `json:"cervical_pain"`
	BackPain        bool `json:"back_pain"`
	Swelling        bool `json:"swelling"`
	Fatigue         bool `json:"fatigue"`
}

type MedicalData struct {
	DiagnosisDate         *time.Time `json:"diagnosis_date"`
	SurgeryDate           *time.Time `json:"surgery_date"`
	SurgeryType           string     `json:"surgery_type"`
	PhysiotherapyReferral bool       `json:"physiotherapy_referral"`
	PhysiotherapyDuration *time.Time `json:"physiotherapy_duration"`
	SufficientRecovery    bool       `json:"sufficient_recovery"`
	PresentSequela        Sequela    `json:"present_sequela"`
	RemainingSequela      Sequela    `json:"remaining_sequela"`
	ImpactOnDailyLife     string     `json:"impact_on_daily_life"`
	AffectsIndependence   bool       `json:"affects_independence"`
	PostSurgeryActivities string     `json:"post_surgery_activities"`
	MedicalConditions     []string   `json:"medical_conditions"`
	OtherDiagnoses        string     `json:"other_diagnoses"`
	Medications           string     `json:"medications"`
}

package schemas

type Change struct {
	Commit  string `json:"commit" gorm:"type:text"`
	Summary string `json:"summary" gorm:"type:text"`
	Message string `json:"message" gorm:"type:text"`
}

var ChangeNil Change

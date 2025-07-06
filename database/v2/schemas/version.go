package schemas

import "time"

type Version struct {
	Name    string     `gorm:"type:varchar(48);size:48;index"`
	Project string     `gorm:"type:varchar(48);size:48;index"`
	Group   string     `gorm:"type:varchar(48);size:48;index"`
	Time    *time.Time // `gorm:"type:timestamp"`
}

func (v *Version) GetName() string {
	return v.Name
}

func (v *Version) GetTime() *time.Time {
	if v.Time != nil {
		return v.Time
	}
	return nil
}

type VersionFamily struct {
	Project string     `gorm:"type:varchar(48);size:48;index"`
	Name    string     `gorm:"type:varchar(48);size:48;index"`
	Time    *time.Time // `gorm:"type:datetime"`
}

func (v *VersionFamily) GetName() string {
	return v.Name
}

func (v *VersionFamily) GetTime() *time.Time {
	if v.Time != nil {
		return v.Time
	}
	return nil
}

type IVersion interface {
	GetTime() *time.Time
	GetName() string
}

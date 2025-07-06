package schemas

type Project struct {
	ID           string `gorm:"type:varchar(48);size:48;primary_key"`
	Name         string `gorm:"type:varchar(48);size:48;index"`
	FriendlyName string `gorm:"type:varchar(48);size:48;"`
}

type Channel string

const (
	ChannelDefault      Channel = "default"
	ChannelExperimental Channel = "experimental"
)

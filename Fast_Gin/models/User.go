package models

type User struct {
	Username    string `gorm:"size:128"`
	ID          uint   `gorm:"primaryKey"`
	Password    string `gorm:"size:128;not null"`
	CreatedTime int64  `gorm:"autoCreateTime"`
	UpdatedTime int64  `gorm:"autoUpdateTime:milli"`
}

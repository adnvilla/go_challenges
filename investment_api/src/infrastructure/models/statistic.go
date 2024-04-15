package models

type Statistic struct {
	Success    bool `gorm:"column:success"`
	Investment int  `gorm:"column:investment"`
}

package repo

import "gorm.io/gorm"

type User struct {
	gorm.DB
	Id    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type NewUser struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type Company struct {
	gorm.DB
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"Name"`
	Location string `json:"Location"`
}

type NewCompany struct {
	Name     string `json:"Name"`
	Location string `json:"Location"`
}

type Job struct {
	gorm.DB
	JobId       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	CompanyId   int    `json:"CompanyId"`
}

type NewJob struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	CompanyId   int    `json:"CompanyId"`
}

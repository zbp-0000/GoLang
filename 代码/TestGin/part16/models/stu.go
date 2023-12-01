package models

type Student struct {
	Id int `gorm:"primary_key"`
	Name string
	Age int
}

package demostruct


type Student struct {
	SId int `gorm:"primary_key"`
	SNo int
	Name string
	Sex string
	Age int
	//关联表：
	Course []Course `gorm:"many2many:Student2Course"`
}

type Course struct {
	CId int `gorm:"primary_key"`
	CName string
	TeacherName string
	Room string
}


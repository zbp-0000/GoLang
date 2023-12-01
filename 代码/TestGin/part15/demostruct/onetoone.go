package demostruct

type User struct{
	UserId int `gorm:"primary_key;AUTO_INCREMENT"`
	Age int
	Name string
	//指定外键：
	IID int
}

type UserInfo struct {
	InfoID int `gorm:"primary_key;AUTO_INCREMENT"`
	Pic string
	Address string
	Email string
	//关联关系
	User User `gorm:"ForeignKey:IID;AssociationForeignKey:InfoID"`
}

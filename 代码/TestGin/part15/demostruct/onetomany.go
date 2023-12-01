package demostruct


type Author struct {
	AID int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	Age int
	Sex string
	//关联关系：
	Article []Article `gorm:"ForeignKey:AuId;AssociationForeignKey:AID"`
}

type Article struct {
	ArId int `gorm:"primary_key;AUTO_INCREMENT"`
	Title string
	Content string
	Desc string
	//设置外键：
	AuId int
}
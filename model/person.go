package model

type Person struct {
	// gorm.Model, uncomment this line to turn on gorm.Model
	ID        uint   `gorm:"column:id;primaryKey"`
	Name      string `gorm:"column:name"`
	Address   string `gorm:"column:address"`
	AccountId uint   `gorm:"column:account_id;unique"`
}

// this method below prevent pluralize from gorm
func (Person) TableName() string {
	return "Person"
}

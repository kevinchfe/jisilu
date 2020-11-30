package model

// Tag 标签表
type Tag struct {
	Id         int64  `gorm:"column:id; PRIMARY_KEY; autoIncrement" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	CreatedOn  int64  `gorm:"column:created_on" json:"created_on"`
	CreatedBy  string `gorm:"column:created_by" json:"created_by"`
	ModifiedOn int64  `gorm:"column:modified_on" json:"modified_on"`
	ModifiedBy string `gorm:"column:modified_by" json:"modified_by"`
	DeletedOn  int64  `gorm:"column:deleted_on" json:"deleted_on"`
	IsDel      int8   `gorm:"column:is_del" json:"is_del"`
	State      int8   `gorm:"column:state" json:"state"`
}

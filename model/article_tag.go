package model

//`id` INT(10) UNSIGNED NOT NULL,
//`article_id` INT(11) NOT NULL COMMENT '文章ID',
//`tag_id` INT(10) UNSIGNED NOT NULL COMMENT '标签ID',
//`created_on` INT(10) UNSIGNED NULL DEFAULT NULL COMMENT '创建时间',
//`created_by` VARCHAR(100) NULL DEFAULT '' COMMENT '创建人',
//`modified_on` INT(10) UNSIGNED NULL DEFAULT NULL COMMENT '修改时间',
//`modified_by` VARCHAR(100) NULL DEFAULT '' COMMENT '修改人',
//`deleted_on` INT(10) UNSIGNED NULL DEFAULT NULL COMMENT '删除时间',
//`is_del` TINYINT(3) UNSIGNED NULL DEFAULT NULL COMMENT '是否删除 0为未删除、1为已删除',
type ArticleTag struct {
	Id         int64  `gorm:"column:id; PRIMARY_KEY; autoIncrement" json:"id"`
	ArticleId  int64  `gorm:"column:article_id" json:"article_id"`
	TagId      int64  `gorm:"column:tag_id" json:"tag_id"`
	CreatedOn  int64  `gorm:"column:created_on" json:"created_on"`
	CreatedBy  string `gorm:"column:created_by" json:"created_by"`
	ModifiedOn int64  `gorm:"column:modified_on" json:"modified_on"`
	ModifiedBy string `gorm:"column:modified_by" json:"modified_by"`
	DeletedOn  int64  `gorm:"column:deleted_on" json:"deleted_on"`
	IsDel      int8   `gorm:"column:is_del" json:"is_del"`
}

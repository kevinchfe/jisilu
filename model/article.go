package model

import (
	"qtmd-php/database"
)

//`id` INT(10) UNSIGNED NOT NULL,
//`title` VARCHAR(100) NULL DEFAULT '' COMMENT '文章标题',
//`desc` VARCHAR(255) NULL DEFAULT '' COMMENT '文章简述',
//`cover_image_url` VARCHAR(255) NULL DEFAULT '' COMMENT '封面图片地址',
//`content` LONGTEXT NULL COMMENT '文章内容',
//`created_on` INT(10) UNSIGNED NULL DEFAULT NULL COMMENT '新建时间',
//`created_by` VARCHAR(100) NULL DEFAULT '' COMMENT '创建人',
//`modified_on` INT(10) UNSIGNED NULL DEFAULT NULL COMMENT '修改时间',
//`modified_by` VARCHAR(100) NULL DEFAULT '' COMMENT '修改人',
//`deleted_on` INT(10) UNSIGNED NULL DEFAULT NULL COMMENT '删除时间',
//`is_del` TINYINT(3) UNSIGNED NULL DEFAULT NULL COMMENT '是否删除 0为未删除、1为已删除',
//`state` TINYINT(3) UNSIGNED NULL DEFAULT NULL COMMENT '状态 0为禁用、1为启用',
type Article struct {
	Id            int64  `gorm:"column:id; PRIMARY_KEY; autoIncrement" json:"id"`
	Title         string `gorm:"column:title" json:"title"`
	Desc          string `gorm:"column:desc" json:"desc"`
	CoverImageUrl string `gorm:"column:cover_image_url" json:"cover_image_url"`
	Content       string `gorm:"column:content" json:"content"`
	CreatedOn     int64  `gorm:"column:created_on" json:"created_on"`
	CreatedBy     string `gorm:"column:created_by" json:"created_by"`
	ModifiedOn    int64  `gorm:"column:modified_on" json:"modified_on"`
	ModifiedBy    string `gorm:"column:modified_by" json:"modified_by"`
	DeletedOn     int64  `gorm:"column:deleted_on" json:"deleted_on"`
	IsDel         int8   `gorm:"column:is_del" json:"is_del"`
	State         int8   `gorm:"column:state" json:"state"`
}

func (a Article) Create() (*Article, error) {
	if err := database.DB.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

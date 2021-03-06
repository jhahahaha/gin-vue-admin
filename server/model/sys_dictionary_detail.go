// 自动生成模板SysDictionaryDetail
package model

import (
	"gorm.io/gorm"
	"time"
)

// 如果含有time.Time 请自行import time包
type SysDictionaryDetail struct {
	gorm.Model
	Label           string `json:"label" form:"label" gorm:"column:label;comment:'展示值'"`
	Value           int    `json:"value" form:"value" gorm:"column:value;comment:'字典值'"`
	Status          *bool  `json:"status" form:"status" gorm:"column:status;comment:'启用状态'"`
	Sort            int    `json:"sort" form:"sort" gorm:"column:sort;comment:'排序标记'"`
	SysDictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" gorm:"column:sys_dictionary_id;comment:'关联标记'"`
}

func SysDictionaryDetailData() []SysDictionaryDetail {
	status := new(bool)
	*status = true
	return []SysDictionaryDetail{
		{gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "smallint", 1, status, 1, 2},
		{gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumint", 2, status, 2, 2},
		{gorm.Model{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "int", 3, status, 3, 2},
		{gorm.Model{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "bigint", 4, status, 4, 2},
		{gorm.Model{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "data", 0, status, 0, 3},
		{gorm.Model{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "time", 1, status, 1, 3},
		{gorm.Model{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "year", 2, status, 2, 3},
		{gorm.Model{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "datetime", 3, status, 3, 3},
		{gorm.Model{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "timestamp", 5, status, 5, 3},
		{gorm.Model{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "float", 0, status, 0, 4},
		{gorm.Model{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "double", 1, status, 1, 4},
		{gorm.Model{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "decimal", 2, status, 2, 4},
		{gorm.Model{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "char", 0, status, 0, 5},
		{gorm.Model{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "varchar", 1, status, 1, 5},
		{gorm.Model{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyblob", 2, status, 2, 5},
		{gorm.Model{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinytext", 3, status, 3, 5},
		{gorm.Model{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "text", 4, status, 4, 5},
		{gorm.Model{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "blob", 5, status, 5, 5},
		{gorm.Model{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumblob", 6, status, 6, 5},
		{gorm.Model{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "mediumtext", 7, status, 7, 5},
		{gorm.Model{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longblob", 8, status, 8, 5},
		{gorm.Model{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "longtext", 9, status, 9, 5},
		{gorm.Model{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "tinyint", 0, status, 0, 6},
	}
}
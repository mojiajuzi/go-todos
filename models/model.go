package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	// TodoModel todo模型
	TodoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}

	// TransformedTodo 格式化返回需要做的事情
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

// migrateModel 数据迁移
func migrateModel(db *gorm.DB) {
	db.AutoMigrate(&TodoModel{})
}

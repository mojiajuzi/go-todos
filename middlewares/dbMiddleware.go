package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AddDBConnect 添加数据库连接到中间件
func AddDBConnect(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConnect", db)
	}
}

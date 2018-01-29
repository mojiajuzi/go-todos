package main

import (
	"hello/controllers"
	"hello/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var dbConnect *gorm.DB

func main() {
	DbConnect()
	//注册路由以及中间件
	r := registerRoute()
	r.Run()
}

func registerRoute() *gin.Engine {

	// 添加路由，并使用基础的中间件
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.AddDBConnect(dbConnect))
	//设置路由
	todo := r.Group("/api/v1/todos")
	{
		todo.POST("/", controllers.Store)
		todo.GET("/", controllers.Index)
		todo.GET("/:id", controllers.Show)
		todo.PUT("/:id", controllers.Update)
		todo.DELETE("/:id", controllers.Destory)
	}
	return r
}

// DbConnect 数据库连接以及数据表迁移
func DbConnect() {
	db, err := gorm.Open("mysql", "root:happybird@tcp(0.0.0.0:43306)/todo?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	dbConnect = db
}

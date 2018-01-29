package controllers

import (
	"errors"
	"fmt"
	"hello/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Index 获取友情连接列表
func Index(c *gin.Context) {
	var todos []models.TodoModel
	db, err := getDB(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// Store  创建todo
func Store(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := models.TodoModel{Title: c.PostForm("title"), Completed: completed}
	db, err := getDB(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully", "resourceId": todo.ID})
}

//Show  获取详情
func Show(c *gin.Context) {
	todoID, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(todoID)
	db, err := getDB(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
	var todo models.TodoModel
	db.First(&todo, todoID)
	c.JSON(http.StatusOK, todo)
}

// Update 更新
func Update(c *gin.Context) {
	todoID, _ := strconv.Atoi(c.Param("id"))

	db, err := getDB(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
	var todo models.TodoModel
	db.First(&todo, todoID)
	todo.Title = c.PostForm("title")
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo.Completed = completed
	db.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// Destory 删除
func Destory(c *gin.Context) {
	todoID, _ := strconv.Atoi(c.Param("id"))
	db, err := getDB(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err})
	}
	var todo models.TodoModel
	db.Delete(&todo, todoID)
	c.JSON(http.StatusOK, todo)
}

func getDB(c *gin.Context) (*gorm.DB, error) {
	db, ok := c.Get("dbConnect")
	if !ok {
		return nil, errors.New("数据库获取错误")
	}
	dbConnect := db.(*gorm.DB)
	return dbConnect, nil
}

package api

import (
	"fmt"
	"golang_gorm/database"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Message  string `json:"msg" binding:"required"`
}

func GetAllListMethod(c *gin.Context) {
	var todoList []database.Todo
	database.DB.Find(&todoList)

	c.JSON(http.StatusOK, gin.H{"Result": todoList})
}

func CreateUserMethod(c *gin.Context) {
	var input CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	todoList := database.Todo{Username: input.Username, Title: input.Title, Message: input.Message}
	database.DB.Create(&todoList)

	c.JSON(http.StatusOK, gin.H{"result": todoList})
}

func GetUserMethod(c *gin.Context) {
	var todoList []database.Todo

	if err := database.DB.Where("username = ?", c.Query("username")).Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "not found user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"List": todoList})
}

func DeleteUserMethod(c *gin.Context) {
	var todoList []database.Todo

	if err := database.DB.Where("id = ?", c.Param("id")).First(&todoList).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found delete"})
		return
	}
	database.DB.Delete(&todoList)
	c.JSON(http.StatusOK, gin.H{"result": todoList})
}

func UploadFileMethod(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8000/file" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

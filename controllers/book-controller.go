package controllers

import (
	"book-rest-api/database"
	"book-rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H {
			"error":"ID has to be integer",
		})
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error

	if err != nil {
		c.JSON(400, gin.H {
			"error":"cannot find book " + err.Error(),
		})

		return 
	}

	c.JSON(200, book)

}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H {
			"error":"Cannot bind json"  + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H {
			"error":"Cannot create book"  + err.Error(),
		})
		return
	}
	c.JSON(201, book)
}

func ShowBooks (c *gin.Context){
	db := database.GetDatabase()


	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H {
			"error":"Cannot find books" + err.Error(),
		})
		return
	}

	c.JSON(200, books)

}

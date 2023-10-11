package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sample/initializers"
	model "sample/models"

	"github.com/gin-gonic/gin"
	"github.com/go-swagger/go-swagger/examples/generated/models"
)

func GetHandler(c *gin.Context) {
	url := c.Query("url")

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get the data", err)
		return
	}

	defer res.Body.Close()

	//body := map[string]interface{}{}
	var post model.Post
	if err1 := json.NewDecoder(res.Body).Decode(&post); err1 != nil {
		fmt.Println("Failed to decode", err1)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"body":   post.Body,
		"userId": post.UserID,
	})

}

func DiscountHandler(c *gin.Context) {
	var body model.Req
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var resBody model.Res
	resBody.Discount = body.Amount * 10
	resBody.Suc = "sucker"
	//result := body.Amount * 10

	c.JSON(http.StatusOK, gin.H{
		"Discount": resBody.Discount,
		"message":  resBody.Suc,
	})

}

func RegisterHandler(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	db := initializers.GetGormDB()

	if err := db.Create(&body).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})

}

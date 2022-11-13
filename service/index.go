package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_gateway/models"
	"strconv"
)

func GetIndex(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	data := make([]*models.UserBasic, 10)
	intid, _ := strconv.Atoi(id)
	data = models.GetUser(intid)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

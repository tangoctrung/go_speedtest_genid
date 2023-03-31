package controllers

import (
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

func GenUUID(c *gin.Context) {
	// dataRes := []string{}

	// for i := 0; i < 100; i++ {
	// 	id := uuid.New().String()
	// 	dataRes = append(dataRes, id)
	// }
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": nil,
		"data":    uuid.New().String(),
	})
}

func GenSnowflakeID(c *gin.Context) {
	// dataRes := []string{}

	// for i := 0; i < 100; i++ {
	// 	id := n.Generate().String()
	// 	dataRes = append(dataRes, id)
	// }

	n, _ := snowflake.NewNode(1)

	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": nil,
		"data":    n.Generate().String(),
	})
}

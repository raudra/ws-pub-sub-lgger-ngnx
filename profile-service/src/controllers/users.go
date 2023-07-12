package controllers

import (
	"fmt"
	"log"
	"net/http"
	"profile-service/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Err     interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func CreateUser(c *gin.Context) {
	args := make(map[string]interface{})

	if err := c.BindJSON(&args); err != nil {
		log.Fatalln(err)
	}

	user, err := services.CreateUser(args)

	if err != nil {
		resp := Response{
			Success: false,
			Err:     err,
		}
		c.JSON(http.StatusUnprocessableEntity, resp)
	} else {
		resp := Response{
			Success: true,
			Data:    map[string]interface{}{"user": user},
		}
		c.JSON(http.StatusOK, resp)
	}
}

func GetUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := services.GetUser(int(userId))

	if err != nil {
		fmt.Println(err)
		resp := Response{
			Success: false,
			Err:     fmt.Errorf("%v", err),
		}
		c.JSON(404, resp)
	} else {
		resp := Response{
			Success: true,
			Data:    map[string]interface{}{"user": user},
		}
		c.JSON(http.StatusOK, resp)
	}
}

func GetUserByMobile(c *gin.Context) {
	mobileNo := c.Query("mobileNo")

	user, err := services.GetUserByMobile(mobileNo)

	if err != nil {
		fmt.Println(err)
		resp := Response{
			Success: false,
			Err:     fmt.Sprintf("%s", err),
		}
		c.JSON(404, resp)
	} else {
		resp := Response{
			Success: true,
			Data:    map[string]interface{}{"user": user},
		}
		c.JSON(http.StatusOK, resp)
	}
}

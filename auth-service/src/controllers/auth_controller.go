package controllers

import (
	"auth-service/src/models"
	"auth-service/src/services"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Success bool        `json:"success"`
	Err     interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func CreateSession(c *gin.Context) {
	args := make(map[string]interface{})

	if err := c.BindJSON(&args); err != nil {
		log.Err(err).
			Msg("Error while parsing the params")
	}

	mobileNo := args["mobileNo"].(string)
	otp := int(args["otp"].(float64))

	session, err := services.CreateSession(mobileNo, otp)

	resp := new(Response)

	if err == nil {
		resp.Success = true
		resp.Data = map[string]interface{}{"message": "Session created successfully",
			"session": session,
		}
		c.JSON(http.StatusOK, resp)
	} else {
		resp.Success = false
		resp.Err = fmt.Sprintf("%s", err)
		c.JSON(http.StatusUnprocessableEntity, resp)
	}

}

func ValidateSession(c *gin.Context) {
	token := c.Request.Header["Token"]

	var err error
	var u *models.User

	if len(token) > 0 {
		u, err = services.ValidateSession(token[0])
	} else {
		err = errors.New("Missing headers")
	}

	resp := new(Response)

	if err == nil {
		resp.Success = true
		resp.Data = map[string]interface{}{
			"user": u,
		}
		c.JSON(http.StatusOK, resp)
	} else {
		resp.Err = fmt.Sprintf("%s", err)
		resp.Success = false
		c.JSON(401, resp)
	}
}

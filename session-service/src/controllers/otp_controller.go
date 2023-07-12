package controllers

import (
	"fmt"
	"net/http"
	"session-service/src/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Success bool        `json:"success"`
	Err     interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SendOtp(c *gin.Context) {
	args := make(map[string]interface{})

	if err := c.BindJSON(&args); err != nil {
		log.Err(err).
			Msg("Error while parsing the params")
	}

	mobileNo := args["mobileNo"].(string)

	status, err := services.SendOtp(mobileNo)

	resp := Response{
		Success: status,
	}

	if status {
		resp.Data = map[string]interface{}{"message": "Successfully send otp"}
		c.JSON(http.StatusOK, resp)
	} else {
		resp.Err = fmt.Sprintf("%s", err)
		c.JSON(http.StatusUnprocessableEntity, resp)
	}

}

func ValidateOtp(c *gin.Context) {
	args := make(map[string]interface{})

	if err := c.BindJSON(&args); err != nil {
		log.Err(err).
			Msg("Error while parsing the params")
	}

	mobileNo := args["mobileNo"].(string)
	otp := int(args["otp"].(float64))

	status, err := services.ValidateOtp(mobileNo, otp)

	resp := Response{
		Success: status,
	}

	if status {
		resp.Data = map[string]interface{}{"message": "Successfully validated otp"}
		c.JSON(http.StatusOK, resp)
	} else {
		resp.Err = fmt.Sprintf("%s", err)
		c.JSON(http.StatusUnprocessableEntity, resp)
	}

}

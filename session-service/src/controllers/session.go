package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Resp struct {
	Status bool        `json:"success"`
	Err    interface{} `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func CreateSession(c *gin.Context) {
	args := make(map[string]interface{})

	if err := c.BindJSON(&args); err != nil {
		log.Fatal(err)
	}

	services.CreateSession(args)

}

func DeleteSession(c *gin.Context) {

}

package server

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type MsFormServerHandler gin.HandlerFunc

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Form(c *gin.Context) {
	resAsStr := `{
		"form": {
			"fields": {
				"entries": [
					{
						"clazz": "combo",
						"uuid": "individual",
						"values": [ "person", "company" ]
					}, {
						"clazz": "text",
						"uuid": "id_company",
						"validation": ".*"
					}, {
						"clazz": "text",
						"uuid": "id_person"
					}
				]
			}
		}
	}`

	var res gin.H

	json.Unmarshal([]byte(resAsStr), &res)

	c.JSON(200, res)
}

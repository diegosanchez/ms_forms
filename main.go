package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/form", func(c *gin.Context) {
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
	})

	router.Run(":" + port)
}

package main

import (
	"database/sql"
	"fmt"

	"example.com/m/v2/config"
	"example.com/m/v2/dbo"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {

	server := gin.Default()

	server.POST("/create-block", func(c *gin.Context) {

		config := config.Loadconfig()
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			config.PG_HOST, config.PG_PORT, config.PG_USER, config.PG_PASSWORD, config.PG_DATABASE)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		body, err := c.GetRawData()
		dbo.InsertNewBlock(db, string(body))
		c.Status(200)
	})

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run()
	fmt.Println("Successfully connected!")

}

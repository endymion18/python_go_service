package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_service/pkg/database"
	"log"
	"time"
)

type Greet struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func (h *Handler) greet(c *gin.Context) {
	name := c.Query("name")
	if len([]rune(name)) == 0 {
		c.JSON(400, "Wrong query parameter")
		return
	}

	if len([]rune(name)) > 30 {
		c.JSON(400, "Name is too long")
		return
	}

	conn := database.GetDatabaseConnection()
	_, err := conn.Exec("INSERT INTO greet_info_go VALUES "+
		"(default, $1, default)", name)
	if err != nil {
		panic(err)
	}

	c.JSON(200, fmt.Sprintf("Привет, %s от Go!", name))
}

func (h *Handler) greetHistory(c *gin.Context) {
	conn := database.GetDatabaseConnection()
	rows, err := conn.Query("SELECT * FROM greet_info_go")
	if err != nil {
		log.Fatalf("database error %s", err.Error())
	}
	var gList []Greet
	for rows.Next() {
		var g Greet
		err := rows.Scan(&g.Id, &g.Name, &g.Date)
		if err != nil {
			panic(err)
		}
		gList = append(gList, g)
	}

	c.JSON(200, gList)
}

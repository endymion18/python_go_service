package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go_service/pkg/database"
	"io"
	"log"
	"net/http"
)

func (h *Handler) pythonGreet(c *gin.Context) {
	name := c.Query("name")
	if len([]rune(name)) == 0 {
		c.JSON(400, "Wrong query parameter")
		return
	}

	if len([]rune(name)) > 30 {
		c.JSON(400, "Name is too long")
		return
	}

	requestURL := fmt.Sprintf("%s/greet?name=%s", viper.GetString("server_url"), name)

	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("request error %s", err.Error())
	}

	var responseContent string
	if res.StatusCode != 200 {
		c.JSON(400, "Python server error")
		return
	}

	err = json.NewDecoder(res.Body).Decode(&responseContent)
	if err != nil {
		return
	}

	conn := database.GetDatabaseConnection()
	_, err = conn.Exec("INSERT INTO greeting VALUES "+
		"(default, $1, default)", responseContent)
	if err != nil {
		log.Fatalf("database error %s", err.Error())
	}

	c.JSON(200, responseContent)
}

func (h *Handler) pythonGreetHistory(c *gin.Context) {
	requestURL := fmt.Sprintf("%s/greet/history", viper.GetString("server_url"))

	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("request error %s", err.Error())
	}

	if res.StatusCode != 200 {
		c.JSON(400, "Python server error")
		return
	}

	body, err := io.ReadAll(res.Body)

	conn := database.GetDatabaseConnection()
	_, err = conn.Exec("INSERT INTO history VALUES"+
		"(default, $1, default)", body)
	if err != nil {
		log.Fatalf("database error %s", err.Error())
	}

	c.Data(200, "application/json", body)

}

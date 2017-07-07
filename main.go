package main

import (
	"net/http"
	"os"

	centrifuge "github.com/centrifugal/centrifuge-go"
	"github.com/centrifugal/centrifugo/libcentrifugo/auth"
	"github.com/gin-gonic/gin"
)

type centrifugo struct {
	URL       string   `json:"url"`
	User      string   `json:"user"`
	Timestamp string   `json:"timestamp"`
	Token     string   `json:"token"`
	Channels  []string `json:"channels"`
}

func main() {
	r := gin.Default()
	r.Static("/web", "./public")

	r.GET("/websocket/:user", func(c *gin.Context) {
		websocket := centrifugo{
			URL:       os.Getenv("CENTRIFUGO_URL"),
			User:      c.Param("user"),
			Timestamp: centrifuge.Timestamp(),
			Token:     auth.GenerateClientToken(os.Getenv("CENTRIFUGO_SECRET"), c.Param("user"), centrifuge.Timestamp(), ""),
			Channels:  []string{"all", c.Param("user")},
		}

		c.JSON(http.StatusOK, websocket)
	})

	r.Run()
}

package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"slack-bot/models"
	"slack-bot/repo"
	"slack-bot/service"
	"slack-bot/util"
	"strings"
)

func InitRoutes() {
	e := echo.New()

	e.GET("/", SayHello)
	e.GET("/alerts", FetchAlerts)
	e.POST("/events", ProcessEvents)
	port := os.Getenv("PORT")
	fmt.Println(port)
	e.Logger.Fatal(e.Start(":" + port))
}

func ProcessEvents(c echo.Context) error {
	botEvent := new(models.BotEvent)
	err := c.Bind(&botEvent)
	util.HError(err)
	Message := strings.Join(strings.Split(botEvent.Event.Text, " ")[1:], " ")
	log.Info("SLACK-REQ: User: ", botEvent.Event.User, " Message: ", Message)
	service.SendASingleMessage(Message)
	err = c.String(http.StatusOK, botEvent.Event.Text)
	util.HError(err)
	return nil
}

func FetchAlerts(c echo.Context) error {
	ctx := context.Background()
	db := repo.GetDBConn()
	event := new(models.Event)
	err := db.NewSelect().Model(event).Scan(ctx)
	util.HError(err)
	fmt.Println(event)
	return c.String(http.StatusOK, event.EventTitle)
}

func SayHello(c echo.Context) error {
	err := c.JSON(http.StatusOK, c.Param("challenge"))
	util.HError(err)
	return nil
}

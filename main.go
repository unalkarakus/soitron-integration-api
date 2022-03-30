package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	hostName := os.Getenv("SOURCE_API_URL")
	if hostName == "" {
		hostName = "http://localhost:8080"
	}

	e.GET("/", func(c echo.Context) error {
		client := &http.Client{}
		req, err := http.NewRequest("GET", hostName+"/ping", nil)
		resp, err := client.Do(req)
		bodyText, err := ioutil.ReadAll(resp.Body)
		resultString := string(bodyText)
		if err != nil {
			return c.HTML(http.StatusOK, "API failed :( ")
		}
		return c.HTML(http.StatusOK, "Result is here! "+resultString)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "Current API is OK!"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8090"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

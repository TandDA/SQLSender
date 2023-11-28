package main

import (
	"net/http"

	"github.com/TandDA/SQLSender/pkg/postgres"
	"github.com/labstack/echo/v4"
)

func main() {
	postgres := postgres.New()
	postgres.Execute("SELECT * FROM test_table")

	e := echo.New()
	e.File("/", "index.html")
	e.GET("/send-request", func(c echo.Context) error {
		var test []interface{} = []interface{}{ae{23, "dima"},ae{1,"geor"}}
		return c.JSON(http.StatusOK, test)
	})
	e.Logger.Fatal(e.Start(":8080"))
} 

type ae struct {
	Id   int
	Name string
}

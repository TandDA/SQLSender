package main

import (
	"net/http"

	"github.com/TandDA/SQLSender/pkg/postgres"
	"github.com/labstack/echo/v4"
)

func main() {
	client := postgres.New()
	client.Execute("SELECT * FROM test_table")

	e := echo.New()
	e.File("/", "index.html")
	e.POST("/send-request", func(c echo.Context) error {
		sql := c.FormValue("sql")
		result := client.Execute(sql)
		return c.JSON(http.StatusOK, result)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

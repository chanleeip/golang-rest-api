package main

import (
	"fmt"
	"net/http"
	"roman-numerals/mongodb"
	"roman-numerals/romanNumerals"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	value := mongodb.MongoInit()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("The Client is %v", value))
	})
	e.GET("/ping", func(c echo.Context) error {
		status := mongodb.Ping()
		return c.String(http.StatusOK, status)
	})
	e.GET("/get-collections", func(c echo.Context) error {
		collections := mongodb.GetAllCollections("whatsapp-mon")
		return c.String(http.StatusOK, strings.Join(collections, ", "))
	})
	e.GET("/test-1/:number", func(c echo.Context) error {
		num, err := strconv.Atoi(c.Param("number"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid number")
		}
		if num > 10 || num < 1 {
			return c.String(http.StatusNotFound, fmt.Sprintf("Out of range"))
		}
		value, _ := romanNumerals.GetValue(num)
		return c.String(http.StatusOK, fmt.Sprintf("The Numeral u sent is %d and Roman Numeral is %s", num, value))
	})
	e.Logger.Fatal(e.Start(":1323"))
}

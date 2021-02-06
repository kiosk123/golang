package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kiosk123/golang/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// return c.String(http.StatusOK, "Hello, World!")
		return c.File("home.html")
	})

	// e.POST("/scrape", func(c echo.Context) error {
	// 	defer os.Remove(fileName) // 콜백일 경우 Defer 실행 안됨

	// 	term := c.FormValue("term")
	// 	term = scrapper.CleanString(term)
	// 	term = strings.ToLower(term)
	// 	scrapper.Scrape(term)

	// 	return c.Attachment(fileName, fileName) //file download
	// })

	e.POST("/scrape", handlePost)

	e.Logger.Fatal(e.Start(":1323"))
	fmt.Println("-- program end --")
}

func handlePost(c echo.Context) error {
	defer os.Remove(fileName)

	term := c.FormValue("term")
	term = scrapper.CleanString(term)
	term = strings.ToLower(term)
	fmt.Println("term is :", term)
	scrapper.Scrape(term)

	return c.Attachment(fileName, fileName) //file download
}

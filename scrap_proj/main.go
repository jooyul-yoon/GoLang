package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jooyul-yoon/learngo/scrap_proj/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	fmt.Println("original:", c.FormValue("term"))
	term := strings.ToLower(cleanString(c.FormValue("term")))
	fmt.Println("term:", term)
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "jobs.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

// clean string removing whitespace
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), "+")
}

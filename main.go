package main

import (
	"log"
	"sync/atomic"

	"github.com/gofiber/fiber/v2"
)

type quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotes = []quote{
	{"Demos never go perfectly.", "Harsh Sharma"},
	{"Talk is cheap. Show me the code.", "Linus Torvalds"},
	{"Premature optimization is the root of all evil.", "Donald Knuth"},
	{"Any fool can write code that a computer can understand. Good programmers write code that humans can understand.", "Martin Fowler"},
	{"Programs must be written for people to read, and only incidentally for machines to execute.", "Harold Abelson"},
}

var next atomic.Uint64

func getQuote(c *fiber.Ctx) error {
	i := next.Add(1) - 1
	return c.JSON(quotes[i%uint64(len(quotes))])
}

func main() {
	app := fiber.New()
	app.Get("/", getQuote)
	log.Fatal(app.Listen(":8080"))
}

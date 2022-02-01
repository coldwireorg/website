package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load env values: %v\n", err)
	} else {
		log.Println("Loaded env values successfully")
	}
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "layouts/main")
	})

	listener := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")
	err := app.Listen(listener)
	if err != nil {
		log.Println("Unable to start server on [" + listener + "]")
		panic(err)
	} else {
		log.Println("Listening on [" + listener + "]")
	}
}

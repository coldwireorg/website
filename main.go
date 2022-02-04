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
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"nav": fiber.Map{
				"color": "green",
			},
			"infos": fiber.Map{},
		}, "layouts/home")
	})

	app.Get("/join", func(c *fiber.Ctx) error {
		return c.Render("pages/join", fiber.Map{
			"nav": fiber.Map{
				"color": "white",
			},
			"infos": fiber.Map{
				"title":       "Join",
				"description": "Write articles, communicate, give your voice, design or contribute to our projets!",
			},
		}, "layouts/page")
	})

	app.Get("/blog", func(c *fiber.Ctx) error {
		return c.Render("pages/blog", fiber.Map{
			"nav": fiber.Map{
				"color": "white",
			},
			"infos": fiber.Map{
				"title":       "",
				"description": "",
			},
		}, "layouts/page")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("pages/about", fiber.Map{
			"nav": fiber.Map{
				"color": "white",
			},
			"infos": fiber.Map{
				"title":       "About",
				"description": "Who are we, what do we do and why it is a legitimate question?",
			},
		}, "layouts/page")
	})

	app.Get("/donate", func(c *fiber.Ctx) error {
		return c.Render("pages/donate", fiber.Map{
			"nav": fiber.Map{
				"color": "white",
			},
			"infos": fiber.Map{
				"title":       "Donate",
				"description": "Help us building resilient and safe organizing tools!",
			},
		}, "layouts/page")
	})

	app.Get("/tos", func(c *fiber.Ctx) error {
		return c.SendFile("./static/TOS.txt")
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

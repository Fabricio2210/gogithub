package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gogit/router"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	router.DefaultRouter(app, "chatlogs", "https://github.com/Fabricio2210/chatlogs.git", "../chatlogs", "go")
	router.DefaultRouter(app, "filetransfer", "https://github.com/Fabricio2210/filetransfer.git", "../filetransfer", "go")
	router.DefaultRouter(app, "gofiber", "https://github.com/Fabricio2210/archivechatlogsserver.git", "../gofiber", "go")
	router.DefaultRouter(app, "chatpy", "https://github.com/Fabricio2210/chatpy.git", "../chatpy", "python")
	router.DefaultRouter(app, "superchatexc", "https://github.com/Fabricio2210/superchatexc.git", "../superchatexc", "python")
	fmt.Println("Server listening on port 5000")
	app.Listen(":5000")
}

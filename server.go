package main

import (
	"github.com/gofiber/fiber"
	"log"
	"strconv"
)

//Server api
type Server struct {
	ytd	ytDownloader
}

type video struct {
    URL string `json:"url" xml:"url" form:"url" query:"url"`
}

func newServer() Server {
	return Server{ytd: newYtDownloader()}
}

func (server *Server) getYturl(c *fiber.Ctx) string {
	v := new(video)

	if err := c.BodyParser(v); err != nil {
		log.Fatal(err)
	}

	log.Println(v.URL) // john
	return v.URL
}

func (server *Server) downloadVideoInfo(c *fiber.Ctx) {
	url := server.getYturl(c)
	err := server.ytd.loadVideoInfo(url)

	if err != nil {
		c.SendStatus(401) //error
	}
	//get video info in json
	dlFormats := server.ytd.getVideoInfo()
	c.Send(dlFormats)
}

func (server *Server) downloadVideobyFormat(c *fiber.Ctx) {
	formatID := c.Params("formatId")

	println(formatID)
	if !isInt(formatID) {
		c.SendStatus(401) //error
	}
	//todo get correct formatId not the index
	formatIndex, _ := strconv.Atoi(formatID)
	filename, err := server.ytd.download(downloadDirectory, uint(formatIndex))

	if err != nil {
		c.SendStatus(401) //error
	}
	filePath := downloadDirectory + filename
	println("FILE = ", filePath)
	c.Download(filePath, filename)
}

func (server *Server) start() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		// println(c.Params("url"))
		c.Send("request has been received by the API")
	})

	// GET http://localhost:3000/download?url=VideoURL
	app.Get("/download", server.downloadVideoInfo)

	// GET http://localhost:3000/download/FormatID
	app.Get("/download/:formatId", server.downloadVideobyFormat)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})


	app.Listen(3000)
  }

package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type service interface {
	Create(msg string) (string, error)
	Get(id string) (string, error)
}

type Handlers struct {
	service
}

func New(service service) *Handlers {
	return &Handlers{service}
}
func (h *Handlers) Post(c *fiber.Ctx) error {
	// Read the param noteId
	//link := c.Params("Link")
	//var shortUrl model.Links
	type shortLink struct {
		ShortUrl string `json:"shortUrl"`
	}
	var shortUrl shortLink
	err := c.BodyParser(&shortUrl)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	fmt.Println("Post - ", shortUrl.ShortUrl)
	msg, err := h.service.Create(shortUrl.ShortUrl)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	return c.SendString(msg)
	// Find the note with the given id
	// Return the note with the id
}

func (h *Handlers) Get(c *fiber.Ctx) error {
	link := c.Params("Link")
	fmt.Println("Get - ", link)
	msg, err := h.service.Get(link)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}
	return c.SendString(msg)
}

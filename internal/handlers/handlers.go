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
	link := c.Params("Link")
	fmt.Println("Post - ", link)
	msg, err := h.service.Create(link)
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

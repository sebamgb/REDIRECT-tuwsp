package handlers

import (
	"log"
)

type handler struct {
	id string
}

func NewHandler(id string) *handler {
	return &handler{
		id: id,
	}
}

func (h *handler) GetId() string {
	return h.id
}

func (h *handler) UpdateValue(value string) {
	log.Printf("%s -> %s", value, h.id)
}

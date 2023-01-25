package pattern

import (
	"log"
)

type Topic interface {
	Register(observer Observer)
	broadcast()
	Figure() bool
}

type Observer interface {
	GetId() string
	UpdateValue(string)
}

type Redirect struct {
	observers []Observer
	Status    string
	figureIn  bool
}

func NewRedirect(status string) *Redirect {
	return &Redirect{
		Status: status,
	}
}

func (r *Redirect) UpdateFigureIn(n string) {
	if r.Status == n {
		r.figureIn = true
		r.broadcast()
		return
	} else {
		log.Println("do not figure")
		r.figureIn = false
	}
}

func (r *Redirect) Register(observer Observer) {
	r.observers = append(r.observers, observer)
}

func (r *Redirect) broadcast() {
	for _, observer := range r.observers {
		observer.UpdateValue(r.Status)
	}
}

func (r *Redirect) Figure() bool {
	return r.figureIn
}

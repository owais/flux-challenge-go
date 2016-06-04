package planets

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket"
	"github.com/owais/flux-challenge/models"
)

func NewPlanetsListener() chan *models.Planet {
	channel := make(chan *models.Planet)
	ws, err := websocket.New("ws://localhost:4000/")
	if err != nil {
		panic(err)
	}
	ws.AddEventListener("message", false, func(ev *js.Object) {
		go func() {
			data := js.Global.Get("JSON").Call("parse", ev.Get("data"))
			channel <- &models.Planet{
				Object: data,
			}
		}()
	})
	return channel
}

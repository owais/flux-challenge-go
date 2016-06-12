package main

import (
	"github.com/owais/flux-challenge/api/jedis"
	"github.com/owais/flux-challenge/api/planets"
	"github.com/owais/flux-challenge/components"
	"gitlab.com/owais/rendr/dom"
)

func main() {
	app := components.NewApp()
	rootID := "#app-container"

	planetsChan := planets.NewPlanetsListener()
	jedisChan, scrollChan := jedis.NewJedisListener(planetsChan)

	app.ScrollChan = scrollChan
	dom.Render(rootID, app)

	for {
		select {

		case app.Planet = <-planetsChan:
			dom.Render(rootID, app)

		case app.Jedis = <-jedisChan:
			dom.Render(rootID, app)
		}
	}
}

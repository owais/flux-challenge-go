package main

import (
	"github.com/owais/flux-challenge/api/jedis"
	"github.com/owais/flux-challenge/api/planets"
	"github.com/owais/flux-challenge/components"
	"gitlab.com/owais/rendr/dom"
)

func main() {

	app := components.NewApp()
	rootId := "#app-container"

	planetsChan := planets.NewPlanetsListener()
	jedisChan, scrollChan := jedis.NewJedisListener(planetsChan)

	app.ScrollChan = scrollChan
	dom.Render(rootId, app)

	for {
		select {

		case app.Planet = <-planetsChan:
			dom.Render(rootId, app)

		case app.Jedis = <-jedisChan:
			dom.Render(rootId, app)
		}
	}
}

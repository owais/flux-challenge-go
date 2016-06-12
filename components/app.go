package components

import (
	"github.com/owais/flux-challenge/models"
	c "gitlab.com/owais/rendr/components"
)

type App struct {
	c.Component
	Planet *models.Planet
	Jedis  *models.Jedis

	ScrollChan chan int
}

func NewApp() App {
	return App{Jedis: &models.Jedis{}}
}

func (app App) Render() c.Renderer {
	return app.Append(
		c.Class("css-root"),
		header(app.Planet),
		c.Section(
			c.Class("css-scrollable-list"),
			jediList(app.Planet, app.Jedis),
			scroller(app.Jedis, app.ScrollChan),
		),
	)
}

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
	return app.SetAttr("class", "css-root").Append(
		header(app.Planet),
		c.Section(
			c.Attrs{"class": "css-scrollable-list"},
			jediList(app.Planet, app.Jedis),
			scroller(app.Jedis, app.ScrollChan),
		),
	)
}

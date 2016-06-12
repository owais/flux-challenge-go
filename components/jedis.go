package components

import (
	"github.com/owais/flux-challenge/models"
	c "gitlab.com/owais/rendr/components"
)

func jedi(j models.Jedi, highlight bool) c.Renderer {
	var style c.Renderer
	if highlight {
		style = c.Style("color", "red")
	}
	return c.Div(
		style,
		c.H3(nil, c.Text(j.Name)),
		c.H6(nil, c.Text("Homeworld: "+j.Home.Name)),
	)
}

func jediList(planet *models.Planet, jedis *models.Jedis) c.Renderer {

	children := []c.Renderer{
		c.Class("css-slots"),
	}

	for _, j := range jedis {
		child := c.Text("")
		if j.Object != nil {
			child = jedi(j, j.Home.Id == planet.Id)
		}
		children = append(children, c.Li(c.Class("css-slot"), child))
	}

	return c.Ul(children...)
}

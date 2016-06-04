package components

import (
	"github.com/owais/flux-challenge/models"
	c "gitlab.com/owais/rendr/components"
)

func jedi(j models.Jedi, highlight bool) c.Renderer {
	attrs := c.Attrs{}
	if highlight {
		attrs["style"] = "color: red;"
	}
	return c.Div(
		attrs,
		c.H3(nil, c.Text(j.Name)),
		c.H6(nil, c.Text("Homeworld: "+j.Home.Name)),
	)
}

func jediList(planet *models.Planet, jedis *models.Jedis) c.Renderer {

	children := []c.Renderer{}

	for _, j := range jedis {
		child := c.Text("")
		if j.Object != nil {
			child = jedi(j, j.Home.Id == planet.Id)
		}
		children = append(children, c.Li(c.Attrs{"class": "css-slot"}, child))
	}

	return c.Ul(
		c.Attrs{"class": "css-slots"},
		children...,
	)
}

package components

import (
	"github.com/owais/flux-challenge/models"
	c "gitlab.com/owais/rendr/components"
)

func header(planet *models.Planet) c.Renderer {
	name := ""
	if planet != nil && planet.Object != nil {
		name = planet.Name
	}
	return c.H1(
		c.Attrs{"class": "css-planet-monitor"},
		c.Text("Obi-wan currently on "+name),
	)
}

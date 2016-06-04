package components

import (
	"github.com/owais/flux-challenge/models"
	c "gitlab.com/owais/rendr/components"
	"honnef.co/go/js/dom"
)

func scroller(jedis *models.Jedis, scrollChan chan int) c.Renderer {
	first, last := jedis[0], jedis[4]

	scroll := func(offset int) c.Handler {
		return func(e dom.Event) {
			go func() {
				scrollChan <- offset
			}()
		}
	}

	up := c.Button(c.Attrs{"class": "css-button-up"})
	if first.Object != nil && first.MasterId == 0 {
		up.SetAttr("class", up.GetAttr("class")+" css-button-disabled")
	} else {
		up = up.On("click", scroll(2))
	}

	down := c.Button(c.Attrs{"class": "css-button-down"})

	if last.Object != nil && last.ApprenticeId == 0 {
		down.SetAttr("class", down.GetAttr("class")+" css-button-disabled")
	} else {
		down = down.On("click", scroll(-2))
	}

	return c.Div(c.Attrs{"class": "css-scroll-buttons"}, up, down)
}

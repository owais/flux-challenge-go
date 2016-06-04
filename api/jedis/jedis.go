package jedis

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gopherjs/gopherjs/js"
	"github.com/owais/flux-challenge/models"
)

type jediListener struct {
	jedis      *models.Jedis
	jedisChan  chan *models.Jedis
	scrollChan chan int
}

func NewJedisListener(planets chan *models.Planet) (chan *models.Jedis, chan int) {
	listener := jediListener{
		&models.Jedis{},
		make(chan *models.Jedis),
		make(chan int),
	}
	go listener.fetch(3616)
	go listener.watchScroll()
	return listener.jedisChan, listener.scrollChan
}

func (l *jediListener) fetch(id int) {
	if id == 0 {
		return
	}

	resp, _ := http.Get(fmt.Sprintf("http://localhost:3000/dark-jedis/%d", id))
	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)
	data := js.Global.Get("JSON").Call("parse", string(contents))

	l.addJedi(models.Jedi{Object: data})
}

func (l *jediListener) addJedi(new models.Jedi) {
	empty := 0
	position := -1
	for i, jedi := range l.jedis {
		if jedi.Object == nil {
			empty++
			continue
		}
		if jedi.ApprenticeId == new.Id {
			position = i + 1
			break
		}

		if jedi.MasterId == new.Id {
			position = i - 1
			break
		}
	}
	if empty == 5 {
		position = 2
	}

	if position >= 0 && position < 5 {
		l.jedis[position] = new
		l.jedisChan <- l.jedis

		prev := position - 1
		if prev >= 0 {
			if l.jedis[prev].Object == nil {
				go l.fetch(new.MasterId)
			}
		}

		next := position + 1
		if next < 5 {
			if l.jedis[next].Object == nil {
				go l.fetch(new.ApprenticeId)
			}
		}
	}
}

func (l *jediListener) watchScroll() {
	for {
		offset := <-l.scrollChan
		newJedis := &models.Jedis{}
		for i, jedi := range l.jedis {
			j := i + offset
			if j < 0 || j >= 5 {
				continue
			}
			newJedis[j] = jedi
		}
		l.jedis = newJedis
		l.jedisChan <- l.jedis

		for _, jedi := range l.jedis {
			if jedi.Object != nil {
				go l.fetch(jedi.MasterId)
				go l.fetch(jedi.ApprenticeId)
			}
		}
	}
}

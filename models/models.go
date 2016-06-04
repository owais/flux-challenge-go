package models

import "github.com/gopherjs/gopherjs/js"

type Planet struct {
	*js.Object
	Id   int    `js:"id"`
	Name string `js:"name"`
}

type Jedi struct {
	*js.Object
	Id           int    `js:"id"`
	Name         string `js:"name"`
	Home         Planet `js:"homeworld"`
	MasterId     int    `js:"master.id"`
	ApprenticeId int    `js:"apprentice.id"`
}

type Jedis [5]Jedi

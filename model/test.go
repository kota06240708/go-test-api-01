package model

type Test struct {
	Model

	Name     string `json:name`
	Contenxt string `json:context`
}

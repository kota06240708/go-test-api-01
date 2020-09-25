package model

type User struct {
	Model

	Name    string `json:name`
	Context string `json:context`
}

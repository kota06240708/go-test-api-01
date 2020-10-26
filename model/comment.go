package model

type Comment struct {
	Model

	Comment string `json:comment binding:"-"`
	UserId  uint   `json:user_id binding:"-"`
	User    *User  `binding:"-"`
}

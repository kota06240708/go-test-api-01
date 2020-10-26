package model

type User struct {
	Model

	Name        string     `json:name`
	Password    string     `json:password`
	Email       string     `json:email`
	Description string     `json:description`
	Comment     []*Comment `gorm:"ForeignKey:UserId; binding:"-"`
}

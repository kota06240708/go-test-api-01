package model

type Comment struct {
	Model

	Comment string `json:comment`
	UserId  int    `json:user_id`
	User    User   `gorm:"ForeignKey:userID;AssociationForeignKey:ID"`
}

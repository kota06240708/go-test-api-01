package request

type Comment struct {
	UserId  int    `json:"user_id"`
	Comment string `json:"comment" validate:"required"`
}

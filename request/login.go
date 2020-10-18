package request

type Login struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

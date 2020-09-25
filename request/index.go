package request

type User struct {
	Name    string `json:"name" validate:"required"`
	Context string `json:"context" validate:"required"`
}

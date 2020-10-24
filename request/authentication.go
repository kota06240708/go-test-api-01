package request

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

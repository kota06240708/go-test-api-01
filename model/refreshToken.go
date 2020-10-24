package model

import (
	"time"
)

// Model is basic struct
type RefreshToken struct {
	Model

	Token  string    `json:token`
	Expire time.Time `json:expire`
}

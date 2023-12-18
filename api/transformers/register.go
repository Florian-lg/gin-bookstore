package transformers

import (
	"core"
	"inputs"
	"models"
)

type RegisterTransformer struct {
	core.Transformer
}

func (t *RegisterTransformer) Transform(input inputs.UserRegister) models.User {
	return models.User{
		Username: input.Username,
		Password: input.Password,
		Email: input.Email,
	}
}
package transformers

import (
	"core"
	"github.com/google/uuid"
	"inputs"
	"models"
	"time"
)

type Author struct {
	core.Transformer
}

func (t *Author) Transform(input *inputs.Author) models.Author {
	return models.Author{
		Id:        uuid.New(),
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Timestamp: models.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

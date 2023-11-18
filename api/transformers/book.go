package transformers

import (
	"github.com/google/uuid"
	"inputs"
	"models"
	"time"
)

func Transform(input *inputs.Book) models.Book {
	return models.Book{
		Id:     uuid.New(),
		Title:  input.Title,
		Price:  input.Price,
		Author: input.Author,
		Timestamp: models.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

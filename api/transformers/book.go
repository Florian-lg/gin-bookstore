package transformers

import (
	"core"
	"github.com/google/uuid"
	"inputs"
	"models"
	"repositories"
	"time"
)

type Book struct {
	core.Transformer
}

func (t *Book) Transform(input *inputs.Book, model *models.Book) models.Book {
	_, author := new(repositories.Author).Find(input.Author)

	if model != nil {
		return models.Book{
			Id:     model.Id,
			Title:  input.Title,
			Price:  input.Price,
			Author: author,
			Timestamp: models.Timestamp{
				CreatedAt: model.Timestamp.CreatedAt,
				UpdatedAt: time.Now(),
			},
		}
	}

	return models.Book{
		Id:     uuid.New(),
		Title:  input.Title,
		Price:  input.Price,
		Author: author,
		Timestamp: models.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

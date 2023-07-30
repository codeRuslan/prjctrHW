package book

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/exp/slog"
)

type Service struct {
	DB *gorm.DB
}

type Book struct {
	gorm.Model
	BookName string
	Author   string
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAllComments() ([]Book, error) {
	slog.Debug("Getting all Books from DB")
	var books []Book
	if result := s.DB.Find(&books); result.Error != nil {
		slog.Warn("Failed to get all the books from DB")
		return books, result.Error
	}
	return books, nil
}

func (s *Service) GetComment(ID uint) (Book, error) {
	slog.Debug("Getting Book from DB" + string(ID))
	var book Book
	if result := s.DB.First(&book, ID); result.Error != nil {
		slog.Debug("Unable to retrieve book from DB")
		return Book{}, result.Error
	}

	return book, nil
}

func (s *Service) PostComment(book Book) (Book, error) {
	slog.Debug("Posting book to DB" + string(book.ID))

	if result := s.DB.Save(&book); result.Error != nil {
		return Book{}, result.Error
	}

	return book, nil

}

func (s *Service) DeleteBook(ID uint) error {
	if result := s.DB.Delete(&Book{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

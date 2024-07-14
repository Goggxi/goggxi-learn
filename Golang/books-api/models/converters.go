package models

import (
	"book-api/models/entities"
	"book-api/models/responses"
)

func ConvertUserToResponse(user entities.User) responses.UserResponse {
	return responses.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

}

func ConvertAuthorToResponse(author entities.Author) responses.AuthorResponse {
	return responses.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Title:     author.Title,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}

func ConvertBookAttrsToResponse(bookAttrs entities.BookAttrs) responses.BookAttrsResponse {
	return responses.BookAttrsResponse{
		ID:          bookAttrs.ID,
		Publisher:   bookAttrs.Publisher,
		Pages:       bookAttrs.Pages,
		Description: bookAttrs.Description,
		Status:      bookAttrs.Status,
		CreatedAt:   bookAttrs.CreatedAt,
		UpdatedAt:   bookAttrs.UpdatedAt,
	}
}

func ConvertBookToResponse(book entities.Book, author entities.Author, bookAttrs entities.BookAttrs) responses.BookResponse {
	return responses.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		UserID:    book.UserID,
		Author:    ConvertAuthorToResponse(author),
		BookAttrs: ConvertBookAttrsToResponse(bookAttrs),
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

func ConvertBooksToResponse(books []entities.Book, authors []entities.Author, bookAttrsList []entities.BookAttrs) []responses.BookResponse {
	bookResponses := make([]responses.BookResponse, len(books))
	for i, book := range books {
		bookResponses[i] = ConvertBookToResponse(book, authors[i], bookAttrsList[i])
	}
	return bookResponses
}

package models

import (
	entities2 "book-api/internal/models/entities"
	responses2 "book-api/internal/models/responses"
)

func ConvertUserToResponse(user entities2.User) responses2.UserResponse {
	return responses2.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

}

func ConvertAuthorToResponse(author entities2.Author) responses2.AuthorResponse {
	return responses2.AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Title:     author.Title,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}

func ConvertBookAttrsToResponse(bookAttrs entities2.BookAttrs) responses2.BookAttrsResponse {
	return responses2.BookAttrsResponse{
		ID:          bookAttrs.ID,
		Publisher:   bookAttrs.Publisher,
		Pages:       bookAttrs.Pages,
		Description: bookAttrs.Description,
		Status:      bookAttrs.Status,
		CreatedAt:   bookAttrs.CreatedAt,
		UpdatedAt:   bookAttrs.UpdatedAt,
	}
}

func ConvertBookToResponse(book entities2.Book, author entities2.Author, bookAttrs entities2.BookAttrs) responses2.BookResponse {
	return responses2.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		UserID:    book.UserID,
		Author:    ConvertAuthorToResponse(author),
		BookAttrs: ConvertBookAttrsToResponse(bookAttrs),
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

func ConvertBooksToResponse(books []entities2.Book, authors []entities2.Author, bookAttrsList []entities2.BookAttrs) []responses2.BookResponse {
	bookResponses := make([]responses2.BookResponse, len(books))
	for i, book := range books {
		bookResponses[i] = ConvertBookToResponse(book, authors[i], bookAttrsList[i])
	}
	return bookResponses
}

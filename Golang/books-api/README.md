# books API

## Endpoints
### `prefix: /api/v1` - for all endpoints

## Public

### `POST /auth/signup` - Create a new user
- Request Body:
```json
{
  "username": "username",
  "password": "password"
}
```
- Response `201`:
```json
{
  "message": "User created successfully",
  "data": {
    "user": {
      "id": "id",
      "username": "username",
      "created_at": "created_at",
      "updated_at": "updated_at"
    },
    "token": "token"
  }
}
```

### `POST /auth/login` - Login a user
- Request Body:
```json
{
  "username": "username",
  "password": "password"
}
```

- Response `200`:
```json
{
  "message": "User logged in successfully",
  "data": {
    "user": {
      "id": "id",
      "username": "username",
      "created_at": "created_at",
      "updated_at": "updated_at"
    },
    "token": "token"
    }
}
```

### `GET /books` - Get all books
- request body:
```json
{
  "page": "page",
  "limit": "limit"
}
```

- response `200`:
```json
{
  "message": "Books found",
  "data": {
    "books": [
      {
        "id": "id",
        "title": "title",
        "user_id": "user_id",
        "author": {
          "id": "id",
          "name": "name",
          "title": "title",
          "created_at": "created_at",
          "updated_at": "updated_at"
        },
        "book_attrs": {
          "id": "id", 
          "publisher": "publisher",
          "pages": "pages",
          "description": "description",
          "status": "status",
          "created_at": "created_at",
          "updated_at": "updated_at"
        },
        "created_at": "created_at",
        "updated_at": "updated_at"
      }
    ]
  }
}
```

### `GET /books/:id` - Get a book by id
- query parameters:
```json
{
  "id": "id"
}
```

- response `200`:
```json
{
  "message": "Book found",
  "data": {
    "book": {
      "id": "id",
      "title": "title",
      "user" : {
        "id": "id",
        "username": "username",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "author": {
        "id": "id",
        "name": "name",
        "title": "title",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "book_attrs": {
        "id": "id",
        "publisher": "publisher",
        "pages": "pages",
        "description": "description",
        "status": "status",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "created_at": "created_at",
      "updated_at": "updated_at"
    }
  }
}
```

### `GET /books/search/:query` - Search for books by title or author
- query parameters:
```json
{
  "query": "query"
}
```

- response `200`:
```json
{
  "message": "Books found",
  "data": {
    "books": [
      {
        "id": "id",
        "title": "title",
        "user_id": "user_id",
        "author": {
          "id": "id",
          "name": "name",
          "title": "title",
          "created_at": "created_at",
          "updated_at": "updated_at"
        },
        "book_attrs": {
          "id": "id",
          "publisher": "publisher",
          "pages": "pages",
          "description": "description",
          "status": "status",
          "created_at": "created_at",
          "updated_at": "updated_at"
        },
        "created_at": "created_at",
        "updated_at": "updated_at"
      }
    ]
  }
}
```


## Private
### `headers: { Authorization: Bearer token }` - for all private endpoints
- Request Headers:
```json
{
  "Authorization": "Bearer token"
}
```

### `GET /auth/logout` - Logout a user
- Response `200`:
```json
{
  "message": "User logged out successfully"
}
```

### `GET /auth/me` - Get the current user
- Response `200`:
```json
{
  "message": "User found",
  "data": {
    "user": {
      "id": "id",
      "username": "username"
    }
  }
}
```

### `POST /books` - Create a new book
- Request Body:
```json
{
  "title": "title",
  "author_id": "author_id",
  "publisher": "publisher",
  "pages": "pages",
  "description": "description",
  "status": "status"
}
```

- Response `201`:
```json
{
  "message": "Book created successfully",
  "data": {
    "book": {
      "id": "id",
      "title": "title",
      "user_id": "user_id",
      "author": {
        "id": "id",
        "name": "name",
        "title": "title",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "book_attrs": {
        "id": "id",
        "publisher": "publisher",
        "pages": "pages",
        "description": "description",
        "status": "status",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "created_at": "created_at",
      "updated_at": "updated_at"
    }
  }
}
```

### `PUT /books/:id` - Update a book by id
- Request Body:
```json
{
  "title": "title",
  "author_id": "author_id",
  "publisher": "publisher",
  "pages": "pages",
  "description": "description",
  "status": "status"
}
```

- Response `200`:
```json
{
  "message": "Book updated successfully",
  "data": {
    "book": {
      "id": "id",
      "title": "title",
      "user_id": "user_id",
      "author": {
        "id": "id",
        "name": "name",
        "title": "title",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "book_attrs": {
        "id": "id",
        "publisher": "publisher",
        "pages": "pages",
        "description": "description",
        "status": "status",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "created_at": "created_at",
      "updated_at": "updated_at"
    }
  }
}
```

### `DELETE /books/:id` - Delete a book by id
- query parameters:
```json
{
  "id": "id"
}
```

- Response `200`:
```json
{
  "message": "Book deleted successfully"
}
```


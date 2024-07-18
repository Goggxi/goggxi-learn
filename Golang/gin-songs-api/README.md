# Gin Songs API

## Description
This is a simple API for managing artists, albums and songs.

## Technologies
- Go
- Gin-Gonic
- PostgresSQL
- Swagger
- Docker for PostgresSQL

## Installation
### Docker PostgresSQL Setup
- Pull the PostgresSQL image:
```bash
docker pull postgres
```

- Create a PostgresSQL container:
```bash
docker run --name postgres-container -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```

- Access the PostgresSQL container:
```bash
docker exec -it postgres-container psql -U postgres
```

- Create a database:
```sql
CREATE DATABASE db_gin_songs_api_golang;
```

- Connect to the database:
```sql
\c db_gin_songs_api_golang
```

- Create the tables:
open the `migrations.sql` file in the `db` directory and copy the contents to the terminal.

### Project Setup
- Clone the repository
- Create a `.env` file in the root directory and add the following environment variables:
```env
# APP
SERVER_ADDRESS=:3000

# DB
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=db_gin_songs_api_golang
DB_SSLMODE=disable
```
- Run the following command to start the application:
```bash
go mod tidy
go run main.go
```

## Entity Schema
### Artist
- id: string
- name: string
- bio: string
- created_at: string
- updated_at: string

### Album
- id: string
- name: string
- genre: string
- artist_id: string (foreign key)
- release_date: string
- created_at: string
- updated_at: string

### Song
- id: string
- name: string
- album_id: string (foreign key)
- duration: float
- release_date: string
- created_at: string
- updated_at: string

## Endpoints
### `prefix: /api/v1` - for all endpoints

### `POST /artists` - Create Artist
- Request Body:
```json
{
  "name": "name",
  "bio": "bio"
}
```
- Response `201`:
```json
{
  "message": "Artist created successfully",
  "data": {
    "id": "id",
    "name": "name",
    "bio": "bio",
    "created_at": "created_at",
    "updated_at": "updated_at"
  }
}
```

### `GET /artists` - Get Artists
- Response `200`:
```json
{
  "message": "Artists retrieved successfully",
  "data": [
    {
      "id": "id",
      "name": "name",
      "bio": "bio",
      "created_at": "created_at",
      "updated_at": "updated_at"
    }
  ]
}
```

### `GET /artists/:id` - Get Artist
- Response `200`:
```json
{
  "message": "Artist retrieved successfully",
  "data": {
    "id": "id",
    "name": "name",
    "bio": "bio",
    "created_at": "created_at",
    "updated_at": "updated_at"
  }
}
```

### `PUT /artists/:id` - Update Artist
- Request Body:
```json
{
  "name": "name",
  "bio": "bio"
}
```

- Response `200`:
```json
{
  "message": "Artist updated successfully",
  "data": {
    "id": "id",
    "name": "name",
    "bio": "bio",
    "created_at": "created_at",
    "updated_at": "updated_at"
  }
}
```

### `DELETE /artists/:id` - Delete Artist
- Response `200`:
```json
{
  "message": "Artist deleted successfully"
}
```

### `POST /albums` - Create Album
- Request Body:
```json
{
  "name": "name",
  "genre": "genre",
  "artist_id": "artist_id",
  "release_date": "release_date"
}
```

- Response `201`:
```json
{
  "message": "Album created successfully",
  "data": {
    "id": "id",
    "name": "name",
    "genre": "genre",
    "artist": {
      "id": "id",
      "name": "name",
      "bio": "bio",
      "created_at": "created_at",
      "updated_at": "updated_at"
    },
    "release_date": "release_date",
    "created_at": "created_at",
    "updated_at": "updated_at"
  }
}
```

### `GET /albums` - Get Albums
- Response `200`:
```json
{
  "message": "Albums retrieved successfully",
  "data": [
    {
      "id": "id",
      "name": "name",
      "genre": "genre",
      "artist": {
        "id": "id",
        "name": "name",
        "bio": "bio",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "release_date": "release_date",
      "created_at": "created_at",
      "updated_at": "updated_at"
    }
  ]
}
```

### `GET /albums/:id` - Get Album
- Response `200`:
```json
{
  "message": "Album retrieved successfully",
  "data": {
    "id": "id",
    "name": "name",
    "genre": "genre",
    "artist": {
      "id": "id",
      "name": "name",
      "bio": "bio",
      "created_at": "created_at",
      "updated_at": "updated_at"
    },
    "release_date": "release_date",
    "created_at": "created_at",
    "updated_at": "updated_at"
  }
}
```

### `PUT /albums/:id` - Update Album
- Request Body:
```json
{
  "name": "name",
  "genre": "genre",
  "artist_id": "artist_id",
  "release_date": "release_date"
}
```

- Response `200`:
```json
{
  "message": "Album updated successfully",
  "data": {
    "id": "id",
    "name": "name",
    "genre": "genre",
    "artist": {
      "id": "id",
      "name": "name",
      "bio": "bio",
      "created_at": "created_at",
      "updated_at": "updated_at"
    },
    "release_date": "release_date",
    "created_at": "created_at",
    "updated_at": "updated_at"
  }
}
```

### `DELETE /albums/:id` - Delete Album
- Response `200`:
```json
{
  "message": "Album deleted successfully"
}
```

### `POST /songs` - Create Song
- Request Body:
```json
{
  "name": "name",
  "album_id": "album_id",
  "duration": "duration",
  "release_date": "release_date"
}
```

- Response `201`:
```json
{
  "message": "Song created successfully",
  "data": {
    "id": "id",
    "name": "name",
    "album": {
      "id": "id",
      "name": "name",
      "genre": "genre",
      "artist": {
        "id": "id",
        "name": "name",
        "bio": "bio",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "release_date": "release_date",
      "created_at": "created_at",
      "updated_at": "updated_at"
    },
    "duration": 0.0,
    "release_date": "release_date",
    "created_at": "created_at",
    "updated_at": "updated_at"
  }
}
```

### `GET /songs` - Get Songs
- Response `200`:
```json
{
  "message": "Songs retrieved successfully",
  "data": [
    {
      "id": "id",
      "name": "name",
      "album": {
        "id": "id",
        "name": "name",
        "genre": "genre",
        "artist": {
          "id": "id",
          "name": "name",
          "bio": "bio",
          "created_at": "created_at",
          "updated_at": "updated_at"
        },
        "release_date": "release_date",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "duration": 0.0,
      "release_date": "release_date",
      "created_at": "created_at",
      "updated_at": "updated_at"
    }
  ]
}
```

### `GET /songs/:id` - Get Song
- Response `200`:
```json
{
  "message": "Song retrieved successfully",
  "data": {
    "id": "id",
    "name": "name",
    "album": {
      "id": "id",
      "name": "name",
      "genre": "genre",
      "artist": {
        "id": "id",
        "name": "name",
        "bio": "bio",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "release_date": "release_date",
      "created_at": "created_at",
      "updated_at": "updated_at"
    },
    "duration": 0.0,
    "release_date": "release_date",
    "created_at": "created_at",
    "updated_at": "updated_at"
  }
}
```

### `PUT /songs/:id` - Update Song
- Request Body:
```json
{
  "name": "name",
  "album_id": "album_id",
  "duration": "duration",
  "release_date": "release_date"
}
```

- Response `200`:
```json
{
  "message": "Song updated successfully",
  "data": {
    "id": "id",
    "name": "name",
    "album": {
      "id": "id",
      "name": "name",
      "genre": "genre",
      "artist": {
        "id": "id",
        "name": "name",
        "bio": "bio",
        "created_at": "created_at",
        "updated_at": "updated_at"
      },
      "release_date": "release_date",
      "created_at": "created_at",
      "updated_at": "updated_at"
    },
    "duration": 0.0,
    "release_date": "release_date",
    "created_at": "created_at",
    "updated_at": "updated_at"
  }
}
```

### `DELETE /songs/:id` - Delete Song
- Response `200`:
```json
{
  "message": "Song deleted successfully"
}
```

## Error Responses
- Response multiple errors:
```json
{
  "message": "HTTP Status Message",
  "error": [
    "error message",
    "error message"
  ]
}
```

- Response single error:
```json
{
  "message": "HTTP Status Message",
  "error": "error message"
}
```

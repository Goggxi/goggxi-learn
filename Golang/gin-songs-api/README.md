# Gin Songs API

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
- duration: string
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
    "duration": "duration",
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
      "duration": "duration",
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
    "duration": "duration",
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
    "duration": "duration",
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

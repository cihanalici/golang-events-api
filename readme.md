# Events REST API Using Go Gin

This is a simple REST API for event system using Go Gin and Sqlite3. User can signup, login and authenticate with JWT token. User can create, update, delete events. User can register and unregister for events.

## Installation

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/golang-jwt/jwt/v5
go get -u golang.org/x/crypto
go get -u github.com/mattn/go-sqlite3

go run .
```

## Endpoints

### Signup For User

```http
POST /signup
```

### Login For User

```http
POST /login
```

### Get All Events

```http
GET /events
```

### Get Event by ID

```http
GET /events/:id
```

### Create Event

```http
POST /events (Authorization required)
```

### Update  Event

```http
PUT /events/:id (Authorization required) (user can only update their own events)
```

### Delete Booking

```http
DELETE /events/:id (Authorization required) (user can only delete their own events)
```

### Register for Event

```http
POST /events/:id/register (Authorization required)
```

### Unregister for Event

```http
POST /events/:id/unregister (Authorization required)
```

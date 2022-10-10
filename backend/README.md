# Available APIs

Running on `http://localhost:8080/`

- `POST` : `/api/login`
- `POST` : `/api/register`
- `GET` :`/api/category`
- `GET` : `/api/post/:id`
- `GET` : `/api/comments`

## Need Authentication
### Profile
- `GET, PATCH` : `/api/profil`
- `PUT` : `/api/profil/avatar`

### Forum Post
- `GET, POST, PUT` : `/api/post`
- `POST` : `/api/post/images/:id`
- `DELETE` : `/api/post/:id`

### Comments
- `GET, POST, PUT` : `/api/comments`
- `DELETE` : `/api/comments/:id`

### Post Like
- `POST, DELETE` : `/api/post/:id/likes`

### Comments Like
- `POST, DELETE` : `/api/comments/:id/likes`

### Notification
- `GET` : `/api/notifications`
- `PUT` : `/api/notifications/read`

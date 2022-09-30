# Available APIs

Running on `http://localhost:8080/`

- `POST` : `/api/login`
<!-- - `POST` : `/api/logout` -->
- `POST` : `/api/register`
- `GET` :`/api/category`

## Need Authentication
### Profile
- `GET` & `PATCH` : `/api/profil`
- `PUT` : `/api/profil/avatar`

### Forum Post
- `GET & POST` : `/api/post`
- `POST` : `/api/post/images/:id`

### Comments
- `GET & POST & PUT` : `/api/comments`
- `DELETE` : `/api/comments/:id`

### Post Like
- `POST & DELETE` : `/api/post/:id/likes`
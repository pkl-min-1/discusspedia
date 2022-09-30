package main

import (
	"database/sql"

	"github.com/pkl-min-1/discusspedia/backend/api"
	"github.com/pkl-min-1/discusspedia/backend/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:basis-app.db")
	if err != nil {
		panic(err)
	}

	commentRepo := repository.NewCommentRepository(db)
	likeRepo := repository.NewLikeRepository(db)
	notifRepo := repository.NewNotificationRepository(db)
	postsRepo := repository.NewPostRepository(db)
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	questionnaireRepo := repository.NewQuestionnaireRepository(db)

	mainAPI := api.NewAPI(*commentRepo, *likeRepo, *notifRepo, *postsRepo, *userRepo, *categoryRepo, *questionnaireRepo)
	mainAPI.Start()
}

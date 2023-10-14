package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"tutorial.sqlc.dev/app/configs"
	"tutorial.sqlc.dev/app/internal/db"

	_ "github.com/lib/pq"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	dbConn, err := sql.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New(),
		Name:        "Go",
		Description: sql.NullString{String: "Go programming language", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}
}

package main

import (
	"context"
	"database/sql"
	"fmt"

	"tutorial.sqlc.dev/app/configs"
	"tutorial.sqlc.dev/app/internal/db"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory db.CreateCategoryParams, argsCourse db.CreateCoursesParams) error {
	return c.callTx(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, argsCategory)
		if err != nil {
			return err
		}

		argsCourse.CategoryID = argsCategory.ID
		err = q.CreateCourses(ctx, argsCourse)
		if err != nil {
			return err
		}

		return nil
	})
}
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

	courseDB := NewCourseDB(dbConn)

	category := db.CreateCategoryParams{ID: uuid.New(), Name: "Programação"}
	course := db.CreateCoursesParams{ID: uuid.New(), Name: "Alura", Price: 1992.99}

	if err := courseDB.CreateCourseAndCategory(ctx, category, course); err != nil {
		panic(err)
	}

	courses, err := courseDB.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Println(course.ID, course.CategoryID, course.Name, course.Description.String, course.Price, course.CategoryName)
	}
}

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"todo/ent"
	"todo/ent/todo"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Api struct {
	client *ent.Client
}

type ListTodoParams struct {
	Page         *int
	ItemsPerPage *int
}

func (a *Api) ListTodo(ctx echo.Context, params ListTodoParams) error {
	page := 0
	if params.Page != nil {
		page = *params.Page
	}
	itemsPerPage := 5
	if params.ItemsPerPage != nil {
		itemsPerPage = *params.ItemsPerPage
	}
	ees, err := a.client.Todo.Query().Order(ent.Desc(todo.FieldID)).
		Offset(page * *params.ItemsPerPage).
		Limit(itemsPerPage).
		All(context.Background())
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	return ctx.JSON(200, ees)
}

func (a *Api) CreateTodo(ctx echo.Context) error {
	var ee ent.Todo
	err := json.NewDecoder(ctx.Request().Body).Decode(&ee)
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	e := a.client.Todo.Create().
		SetText(ee.Text).
		SetStatus(ee.Status).
		SetPriority(ee.Priority)
	if !ee.CreatedAt.IsZero() {
		e.SetCreatedAt(ee.CreatedAt)
	}
	if ee2, err := e.Save(context.Background()); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	} else {
		ee = *ee2
	}

	return ctx.JSON(200, ee)
}

func (a *Api) DeleteTodo(ctx echo.Context, id int) error {
	e := a.client.Todo.DeleteOneID(int(id))
	err := e.Exec(context.Background())
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	return nil
}

func (a *Api) ReadTodo(ctx echo.Context, id int) error {
	e, err := a.client.Todo.Get(context.Background(), int(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.ErrNotFound
		}
		log.Println(err)
		return echo.ErrBadRequest
	}
	return ctx.JSON(200, e)
}

func (a *Api) UpdateTodo(ctx echo.Context, id int) error {
	var ee ent.Todo
	err := json.NewDecoder(ctx.Request().Body).Decode(&ee)
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	e := a.client.Todo.UpdateOneID(int(id)).
		SetText(ee.Text).
		SetStatus(ee.Status).
		SetPriority(ee.Priority)

	if ee2, err := e.Save(context.Background()); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	} else {
		ee = *ee2
	}
	return ctx.JSON(200, ee)
}

func main() {
	url := "host=127.0.0.1 port=5432 sslmode=disable password=postgres user=postgres"
	client, err := ent.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	client.Schema.Create(context.Background())

	e := echo.New()
	myApi := &Api{client: client}
	RegisterHandlers(e, myApi)
	e.Static("/", "static")
	e.Logger.Fatal(e.Start(":8989"))

}

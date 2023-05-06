package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"time"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Content   string    `bun:"content,notnull"`
	Done      bool      `bun:"done"`
	Until     time.Time `bun:"until,nullzero"`
	CreatedAt time.Time
	UpdatedAt time.Time `bun:",nullzero"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}

func main() {
	url := "host=127.0.0.1 port=5432 sslmode=disable password=postgres user=postgres"
	sqldb, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()
	db := bun.NewDB(sqldb, pgdialect.New())
	defer db.Close()

	var todos []Todo
	ctx := context.Background()
	err = db.NewSelect().Model(&todos).Order("created_at").Where("until is not null").Where("done is false").Scan(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(todos) == 0 {
		return
	}
	from := mail.Address{Name: "TODO Reminder", Address: "mail from"}
	to := "to address"
	var buf bytes.Buffer
	buf.WriteString("From:" + from.String() + "\r\n")
	buf.WriteString("To:" + to + "\r\n")
	buf.WriteString("Subject: TODO Reminder")
	buf.WriteString("\r\n")
	buf.WriteString("This is your todo list\r\n")
	for _, todo := range todos {
		fmt.Fprintf(&buf, "%s %s\n", todo.Until, todo.Content)
	}
	smtpauth := smtp.PlainAuth(
		"maildomain",
		"user",
		"pass",
		"authserver",
	)

	server := "google.com"
	err = smtp.SendMail(
		server,
		smtpauth,
		from.Address,
		[]string{"to"},
		buf.Bytes())

	if err != nil {
		log.Fatal(err)
	}
}

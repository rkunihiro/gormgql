package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rkunihiro/gormgql/generated"
	"github.com/rkunihiro/gormgql/repository"
	"github.com/rkunihiro/gormgql/resolver"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conn, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/dbname?parseTime=true")
	if err != nil {
		panic(fmt.Errorf("%v", err))
	}

	db, err := gorm.Open(
		mysql.New(mysql.Config{Conn: conn, SkipInitializeWithVersion: true}),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver.NewResolver(
					userRepo,
					postRepo,
				),
			},
		),
	)
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		return gqlerror.Errorf("something went wrong")
	})

	http.Handle("/graphql", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

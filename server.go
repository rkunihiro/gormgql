package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rkunihiro/gormgql/generated"
	"github.com/rkunihiro/gormgql/repository"
	"github.com/rkunihiro/gormgql/resolver"
)

const defaultPort = "8080"

func initLogger() {
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.CallerFieldName = "caller"
	zerolog.MessageFieldName = "message"
	zerolog.ErrorFieldName = "error"
	zerolog.ErrorStackFieldName = "stack"
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z07:00"
	log.Logger = zerolog.New(os.Stdout).With().
		Timestamp().
		// Caller().
		// Stack().
		Logger()
}

func main() {
	initLogger()
	log.Info().Msg("start")
	defer func() {
		err := recover()
		if err != nil {
			log.Error().Msgf("failed %v", err)
			os.Exit(1)
		}
		log.Info().Msg("success")
		os.Exit(0)
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conn, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/dbname?parseTime=true")
	if err != nil {
		panic(fmt.Errorf("sql.Open failed: %v", err))
	}

	db, err := gorm.Open(
		mysql.New(mysql.Config{Conn: conn, SkipInitializeWithVersion: true}),
		&gorm.Config{},
	)
	if err != nil {
		panic(fmt.Errorf("failed to connect database %v", err))
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

	r := mux.NewRouter()
	r.Handle("/graphql", srv)
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	r.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		_, _ = w.Write([]byte("ok"))
	})
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/go-microfrontend/auth/internal/processes"
	"github.com/go-microfrontend/auth/internal/repository"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx := context.Background()

	client, err := client.Dial(client.Options{HostPort: os.Getenv("TEMPORAL_ADDR")})
	if err != nil {
		slog.Error("failed to connect to temporal", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer client.Close()

	conn, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer conn.Close()

	repo := repository.New(conn)
	activities := processes.New(repo)

	w := worker.New(client, os.Getenv("TASK_QUEUE"), worker.Options{})
	for _, workflow := range processes.Workflows {
		w.RegisterWorkflow(workflow)
	}
	w.RegisterActivity(activities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		slog.Error("failed to run worker", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

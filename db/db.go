// Package dbtest contains supporting code for running tests that hit the DB.
package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func UpTestingDB(t *testing.T) (string, func()) {
	ctx := context.Background()

	dbName := "test"
	dbUser := "test"
	dbPassword := "test"

	// 1. Start the postgres ctr and run any migrations on it
	postgresContainer, err := postgres.Run(
		ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
		postgres.WithSQLDriver("pgx"),
	)
	testcontainers.CleanupContainer(t, postgresContainer)
	require.NoError(t, err)

	cs, err := postgresContainer.ConnectionString(context.Background())
	require.NoError(t, err)

	return cs, func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			require.NoError(t, err)
		}
	}
}

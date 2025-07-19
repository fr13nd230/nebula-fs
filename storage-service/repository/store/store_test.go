package store

import (
	"os"
	"testing"

	"github.com/fr13nd230/nebula-fs/storage-service/config"
	"github.com/stretchr/testify/require"
)

// TestNewDb tests if *Queries is acquired and a connection
// has been succesfully established to the database.
func TestNewDb(t *testing.T) {
	if err := config.LoadEnv("../../.env"); err != nil {
		t.Error("TEST FAILED, unable to load .env file variables")
		os.Exit(1)
	}
	drv := config.GetVar("PSQL_DRIVER_PATH")
	q, err := NewDB(drv)
	require.NoError(t, err)
	require.NotEmpty(t, q)
}
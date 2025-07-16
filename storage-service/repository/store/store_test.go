package store

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNewDb tests if *Queries is acquired and a connection
// has been succesfully established to the database.
func TestNewDb(t *testing.T) {
	q, err := NewDB("")
	require.NoError(t, err)
	require.NotEmpty(t, q)
}
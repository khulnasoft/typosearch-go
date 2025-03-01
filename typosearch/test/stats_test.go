//go:build integration
// +build integration

package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatsRetrieve(t *testing.T) {
	_, err := typosearchClient.Stats().Retrieve(context.Background())
	require.NoError(t, err)
}

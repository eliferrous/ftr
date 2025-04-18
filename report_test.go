package ftr

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMTRRunnerReport(t *testing.T) {
	r := NewMTRRunner()

	_, err := r.Run(context.Background(), "example.com", 5)

	require.NoError(t, err)
}

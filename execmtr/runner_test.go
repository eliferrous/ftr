package execmtr

import (
	"context"
	"os/exec"
	"testing"

	"github.com/eliferrous/ftr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMTRRunner(t *testing.T) {

	if _, err := exec.LookPath("mtr"); err != nil {
		t.Skip("mtr not found, is it installed? – skipping")
	}

	r := ftr.NewMTRRunner()

	context := context.Background()
	report, err := r.Run(context, "google.com", 1)

	require.NoError(t, err, "failed running runner ")
	assert.NotNil(t, report, "report should not be nil")

}

func TestExtractIP(t *testing.T) {
	ip, ok := extractIP("dns.google (8.8.8.8)")
	require.True(t, ok)
	require.Equal(t, "8.8.8.8", ip)
}

func TestExtractIP6(t *testing.T) {
	ip, ok := extractIP("loopback1.NWRKNJMD-PPR02-CC.ALTER.NET (2600:802::2f)")
	require.True(t, ok)
	require.Equal(t, "2600:802::2f", ip)
}

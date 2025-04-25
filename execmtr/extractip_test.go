package execmtr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

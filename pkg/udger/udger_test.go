package udger

import (
	"fmt"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Datacenter(t *testing.T) {
	t.Skip("This test currently fails: no such table: udger_datacenter_range")
	tmp := t.TempDir()
	client, err := NewClient(path.Join(tmp, "db"))
	require.NoError(t, err)

	tests := []struct {
		addr    string
		want    int
		wantErr bool
	}{
		{
			addr:    "1.12.0.0",
			want:    1277,
			wantErr: false,
		},
		{
			addr:    "1.15.0.0",
			want:    1277,
			wantErr: false,
		},
		{
			addr:    "5.175.219.0",
			want:    485,
			wantErr: false,
		},
		{
			addr:    "2001:1be0:1000:100::",
			want:    630,
			wantErr: false,
		},
		{
			addr:    "2a01:7e00::",
			want:    777,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s -> %d", tt.addr, tt.want), func(t *testing.T) {
			did, err := client.Datacenter(tt.addr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, did)
			}
		})
	}
}

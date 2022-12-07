package compress

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestLz4Compress_Compress(t *testing.T) {
	tests := []struct {
		name    string
		lc      *Lz4Compress
		data    []byte
		want    []byte
		wantErr error
	}{
		{
			name:    "normal case:",
			lc:      &Lz4Compress{},
			data:    []byte("hello world"),
			want:    []byte("hello world"),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.lc.Compress(tt.data)
			t.Log("err:", err)
			assert.Equal(t, err, tt.wantErr)

			got, err = tt.lc.Uncompress(got)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

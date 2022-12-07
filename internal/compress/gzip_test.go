package compress

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGzipCompress_Compress(t *testing.T) {
	tests := []struct {
		name    string
		gc      *GzipCompress
		data    []byte
		want    []byte
		wantErr error
	}{
		{
			name:    "normal case:",
			gc:      &GzipCompress{},
			data:    []byte("hello world"),
			want:    []byte("hello world"),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.gc.Compress(tt.data)
			assert.Equal(t, err, tt.wantErr)
			got, err = tt.gc.Uncompress(got)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

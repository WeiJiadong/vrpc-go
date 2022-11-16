// Package codec 包含协议打解包
package codec

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestVrpcMsg_Marshal(t *testing.T) {
	type fields struct {
		MsgLen              uint32
		BodyLen             uint32
		MsgID               int64
		Vesion              byte
		Compress            byte
		Serialization       byte
		CalleeServerNameLen byte
		CalleeMethodNameLen byte
		CalleeServerName    string
		CalleeMethodName    string
		Mate                map[string]string
		Body                []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "打解包测试：",
			fields: fields{
				MsgID:            1,
				Vesion:           0,
				Compress:         0,
				Serialization:    0,
				CalleeServerName: "server",
				CalleeMethodName: "method",
				Mate: map[string]string{
					"key":  "val",
					"key1": "val1",
				},
				Body: []byte("hello world"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VrpcMsg{
				MsgLen:              0,
				BodyLen:             uint32(len(tt.fields.Body)),
				MsgID:               tt.fields.MsgID,
				Vesion:              tt.fields.Vesion,
				Compress:            tt.fields.Compress,
				Serialization:       tt.fields.Serialization,
				CalleeServerNameLen: byte(len(tt.fields.CalleeServerName)),
				CalleeMethodNameLen: byte(len(tt.fields.CalleeMethodName)),
				CalleeServerName:    tt.fields.CalleeServerName,
				CalleeMethodName:    tt.fields.CalleeMethodName,
				Mate:                tt.fields.Mate,
				Body:                tt.fields.Body,
			}

			got, err := c.Marshal()
			if err != nil {
				t.Errorf("VrpcMsg.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			c1 := &VrpcMsg{}
			err = c1.Unmarshal(got)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, c, c1)
		})
	}
}

package utils

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoPost(t *testing.T) {
	tests := []struct {
		name       string
		req        any
		statusCode int
		want       bool
		wantErr    bool
	}{
		{
			"happy",
			struct {
				A string `json:"something_long"`
				B int    `json:"another_something"`
			}{"blah", 2},
			http.StatusOK,
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			server := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					defer r.Body.Close()
					bodyBytes, err := io.ReadAll(r.Body)
					assert.NoError(t, err)
					var obj map[string]json.RawMessage
					err = json.Unmarshal(bodyBytes, &obj)
					assert.NoError(t, err)
					v := reflect.ValueOf(tt.req)
					typeOfReq := v.Type()
					for i := range v.NumField() {
						field := typeOfReq.Field(i)
						tag := string(field.Tag)
						tagSub := tag[6 : len(tag)-1]
						assert.Contains(t, obj, tagSub)
					}

					// validate the request
					w.WriteHeader(tt.statusCode)
					if tt.want {
						_, _ = w.Write([]byte(`{"value":"fixed"}`))
					}
				}))
			defer server.Close()

			got, err := DoPost(ctx, tt.req, server.URL)
			if tt.want && got == nil {
				t.Errorf("got %q, want %v", got, tt.want)
			}
			if tt.wantErr && err == nil {
				t.Errorf("error: got %q, want %v", err, tt.wantErr)
			}
		})
	}
}

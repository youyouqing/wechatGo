package util

import (
	"net/http"
	"testing"
)

func TestResultJson(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

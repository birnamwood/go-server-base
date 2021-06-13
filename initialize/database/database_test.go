package database

import (
	"github.com/birnamwood/go-server-base/initialize/config"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		env  string
	}{
		{
			name: "データベース初期化local",
			env:  "local",
		},
		{
			name: "データベース初期化development",
			env:  "development",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Init(tt.env, "environment")
			Init()
		})
	}
}

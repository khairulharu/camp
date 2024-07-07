package database

import (
	"campsite/internal/config"
	"database/sql"
	"reflect"
	"testing"
)

func TestNewMysqlConnection(t *testing.T) {
	type args struct {
		config *config.Config
	}
	tests := []struct {
		name string
		args args
		want *sql.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMysqlConnection(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMysqlConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

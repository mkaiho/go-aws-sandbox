package infrastructure

import (
	"testing"
)

func TestMySQLConfig_DSN(t *testing.T) {
	type fields struct {
		Host     string
		User     string
		Database string
		Password string
		Port     int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "return DSN value",
			fields: fields{
				Host:     "test_host",
				User:     "test_user",
				Database: "test_db",
				Password: "test_pass",
				Port:     3306,
			},
			want: "test_user:test_pass@tcp(test_host:3306)/test_db?parseTime=true&charset=utf8mb4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &MySQLConfig{
				Host:     tt.fields.Host,
				User:     tt.fields.User,
				Database: tt.fields.Database,
				Password: tt.fields.Password,
				Port:     tt.fields.Port,
			}
			if got := c.GetDSN(); got != tt.want {
				t.Errorf("MySQLConfig.DSN() = %v, want %v", got, tt.want)
			}
		})
	}
}

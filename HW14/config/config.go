package conf

import (
	"os"
)

var APIKeyHeader string
var APIKeys map[string]string

func init() {
	APIKeyHeader = os.Getenv("APIKeyHeader")
	APIKeys = map[string]string{
		"api-key-1": "user1",
		"api-key-2": "user2",
	}
}

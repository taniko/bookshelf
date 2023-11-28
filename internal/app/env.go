package app

import "os"

func IsDevelopment() bool {
	return os.Getenv("APP_ENV") == "development"
}

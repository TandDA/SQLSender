package main

import (
	"github.com/TandDA/SQLSender/pkg/postgres"
)

func main() {
	postgres := postgres.New()
	postgres.Execute("SELECT * FROM specialties")
}

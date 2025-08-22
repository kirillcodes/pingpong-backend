package main

import (
	"pingPong/internal/db"
)

func main() {
	database := db.Connect()
	_ = database
}
